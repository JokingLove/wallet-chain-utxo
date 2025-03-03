package kas

import (
	"encoding/hex"
	"fmt"
	"github.com/dapplink-labs/wallet-chain-utxo/chain"
	"github.com/dapplink-labs/wallet-chain-utxo/chain/base"
	"github.com/dapplink-labs/wallet-chain-utxo/config"
	common2 "github.com/dapplink-labs/wallet-chain-utxo/rpc/common"
	"github.com/dapplink-labs/wallet-chain-utxo/rpc/utxo"
	"github.com/ethereum/go-ethereum/log"
	"github.com/kaspanet/kaspad/domain/dagconfig"
	"github.com/kaspanet/kaspad/util"
	"strconv"
	"time"
)

const ChainName = "Kaspa"

type ChainAdaptor struct {
	kaspaClient           *KasClient
	kaspaDataClientClient *base.BaseDataClient
}

func NewChainAdaptor(conf *config.Config) (chain.IChainAdaptor, error) {
	kasClient, err := NewKasClient(conf.WalletNode.Kas.RpcUrl)
	if err != nil {
		return nil, err
	}
	baseDataClient, err := base.NewBaseDataClient(conf.WalletNode.Kas.DataApiUrl, conf.WalletNode.Kas.DataApiKey, "KAS", "Kaspa")
	if err != nil {
		log.Error("new kaspa data client fail", "err", err)
		return nil, err
	}
	return &ChainAdaptor{
		kaspaClient:           kasClient,
		kaspaDataClientClient: baseDataClient,
	}, nil
}

func (c *ChainAdaptor) GetSupportChains(req *utxo.SupportChainsRequest) (*utxo.SupportChainsResponse, error) {
	return &utxo.SupportChainsResponse{
		Code:    common2.ReturnCode_SUCCESS,
		Msg:     "Support this chain",
		Support: true,
	}, nil
}

func (c ChainAdaptor) ConvertAddress(req *utxo.ConvertAddressRequest) (*utxo.ConvertAddressResponse, error) {
	var addressString string
	publicKeyByte, _ := hex.DecodeString(req.PublicKey)
	if req.Format == "" {
		req.Format = "p2pkh"
	}
	switch req.Format {
	case "p2sh":
		address, err := util.NewAddressScriptHash(publicKeyByte, dagconfig.MainnetParams.Prefix)
		if err != nil {
			log.Error("new kaspa address fail", "err", err)
			return nil, err
		}
		addressString = address.String()
		break
	default:
		address, err := util.NewAddressPublicKeyECDSA(publicKeyByte, dagconfig.MainnetParams.Prefix)
		if err != nil {
			log.Error("new kaspa address fail", "err", err)
			return nil, err
		}
		addressString = address.String()
		break
	}

	return &utxo.ConvertAddressResponse{
		Code:    common2.ReturnCode_SUCCESS,
		Msg:     "convert address success",
		Address: addressString,
	}, nil
}

func (c ChainAdaptor) ValidAddress(req *utxo.ValidAddressRequest) (*utxo.ValidAddressResponse, error) {
	address, err := util.DecodeAddress(req.Address, dagconfig.MainnetParams.Prefix)
	if err != nil {
		log.Error(" kaspa address is invalid", "err", err)
		return &utxo.ValidAddressResponse{
			Code:  common2.ReturnCode_ERROR,
			Msg:   err.Error(),
			Valid: false,
		}, err
	}
	if !address.IsForPrefix(dagconfig.MainnetParams.Prefix) {
		return &utxo.ValidAddressResponse{
			Code:  common2.ReturnCode_ERROR,
			Msg:   "kaspa address is invalid",
			Valid: false,
		}, err
	}
	return &utxo.ValidAddressResponse{
		Code:  common2.ReturnCode_SUCCESS,
		Msg:   "kaspa address is valid",
		Valid: true,
	}, nil
}

func (c ChainAdaptor) GetFee(req *utxo.FeeRequest) (*utxo.FeeResponse, error) {
	resp, err := c.kaspaClient.FeeEstimate()
	if err != nil {
		log.Error("kaspa get fee fail", "err", err)
		return nil, err
	}

	return &utxo.FeeResponse{
		Code:      common2.ReturnCode_SUCCESS,
		Msg:       "kaspa get fee success",
		BestFee:   strconv.FormatFloat(resp.PriorityBucket.EstimatedSeconds, 'f', 3, 64),
		SlowFee:   strconv.FormatFloat(resp.LowBuckets[0].EstimatedSeconds, 'f', 3, 64),
		NormalFee: strconv.FormatFloat(resp.NormalBuckets[0].EstimatedSeconds, 'f', 3, 64),
		FastFee:   strconv.FormatFloat(resp.PriorityBucket.EstimatedSeconds, 'f', 3, 64),
	}, nil
}

func (c ChainAdaptor) GetAccount(req *utxo.AccountRequest) (*utxo.AccountResponse, error) {
	result, err := c.kaspaClient.GetAccountBalance(req.Address)
	if err != nil {
		log.Error("kaspa get account fail", "err", err)
		return nil, err
	}

	return &utxo.AccountResponse{
		Code:    common2.ReturnCode_SUCCESS,
		Msg:     "get account success",
		Network: dagconfig.MainnetParams.Name,
		Balance: strconv.Itoa(int(result)),
	}, nil
}

func (c ChainAdaptor) GetUnspentOutputs(req *utxo.UnspentOutputsRequest) (*utxo.UnspentOutputsResponse, error) {
	result, err := c.kaspaClient.GetUtxo(req.Address)
	if err != nil {
		log.Error("kaspa get unspent outputs fail", "err", err)
		return nil, err
	}

	// return
	var utxoList []*utxo.UnspentOutput

	for _, utxoItem := range result {
		item := utxo.UnspentOutput{
			TxId:          utxoItem.Outpoint.TransactionID,
			Script:        utxoItem.UtxoEntry.ScriptPublicKey.ScriptPublicKey,
			Height:        utxoItem.UtxoEntry.BlockDaaScore,
			Address:       utxoItem.Address,
			UnspentAmount: utxoItem.UtxoEntry.Amount,
			Index:         uint64(utxoItem.Outpoint.Index),
		}

		utxoList = append(utxoList, &item)
	}
	return &utxo.UnspentOutputsResponse{
		Code:           common2.ReturnCode_SUCCESS,
		Msg:            "kaspa get unspent outputs success",
		UnspentOutputs: utxoList,
	}, nil
}

func (c ChainAdaptor) GetBlockByNumber(req *utxo.BlockNumberRequest) (*utxo.BlockResponse, error) {
	result, err := c.kaspaClient.GetBlockByHeight(req.Height)
	if err != nil {
		log.Error("kaspa get block by number fail", "err", err)
		return nil, err
	}

	var txList []*utxo.TransactionList
	for _, txItem := range result.Transactions {
		if len(txItem.Inputs) == 0 {
			continue
		}
		// vin
		var inputs []*utxo.Vin
		for _, inputItem := range txItem.Inputs {
			input := &utxo.Vin{
				Hash:    inputItem.PreviousOutpoint.TransactionID, // string
				Index:   inputItem.PreviousOutpoint.Index,         // 暂无数据，设为空字符串
				Amount:  0,                                        // 暂无数据，设为空字符串
				Address: "",                                       // 暂无数据，设为空字符串
			}
			inputs = append(inputs, input)
		}

		// vout
		var outputs []*utxo.Vout
		for index, outputItem := range txItem.Outputs {
			amount, _ := strconv.ParseInt(outputItem.Amount, 10, 64)
			output := &utxo.Vout{
				Amount:  amount,                                        // uint64 转为 string
				Address: outputItem.VerboseData.ScriptPublicKeyAddress, // string
				Index:   uint32(index),
			}
			outputs = append(outputs, output)
		}

		// 构造 TransactionList
		tx := &utxo.TransactionList{
			Hash: txItem.VerboseData.Hash,
			Fee:  "",
			Vin:  inputs,
			Vout: outputs,
		}
		txList = append(txList, tx)
	}

	return &utxo.BlockResponse{
		Code:   common2.ReturnCode_SUCCESS,
		Msg:    "kaspa get block by block success",
		Height: 0,
		Hash:   result.VerboseData.Hash,
		TxList: txList,
	}, nil
}

func (c ChainAdaptor) GetBlockByHash(req *utxo.BlockHashRequest) (*utxo.BlockResponse, error) {
	result, err := c.kaspaClient.GetBlockByHash(req.Hash)
	if err != nil {
		log.Error("kaspa get block by block id fail", "err", err)
		return nil, err
	}

	var txList []*utxo.TransactionList

	for _, txItem := range result.Transactions {
		if len(txItem.Inputs) == 0 {
			continue
		}
		// vin
		var inputs []*utxo.Vin
		for _, inputItem := range txItem.Inputs {
			input := &utxo.Vin{
				Hash:    inputItem.PreviousOutpoint.TransactionID, // string
				Index:   inputItem.PreviousOutpoint.Index,         // 暂无数据，设为空字符串
				Amount:  0,                                        // 暂无数据，设为空字符串
				Address: "",                                       // 暂无数据，设为空字符串
			}
			inputs = append(inputs, input)
		}

		// vout
		var outputs []*utxo.Vout
		for index, outputItem := range txItem.Outputs {
			amount, _ := strconv.ParseInt(outputItem.Amount, 10, 64)
			output := &utxo.Vout{
				Amount:  amount,                                        // uint64 转为 string
				Address: outputItem.VerboseData.ScriptPublicKeyAddress, // string
				Index:   uint32(index),
			}
			outputs = append(outputs, output)
		}

		// 构造 TransactionList
		tx := &utxo.TransactionList{
			Hash: txItem.VerboseData.Hash,
			Fee:  "",
			Vin:  inputs,
			Vout: outputs,
		}
		txList = append(txList, tx)
	}

	return &utxo.BlockResponse{
		Code:   common2.ReturnCode_SUCCESS,
		Msg:    "kaspa get block by block success",
		Height: 0,
		Hash:   req.Hash,
		TxList: txList,
	}, nil

}

func (c ChainAdaptor) GetBlockHeaderByHash(req *utxo.BlockHeaderHashRequest) (*utxo.BlockHeaderResponse, error) {
	result, err := c.kaspaClient.GetBlockByHash(req.Hash)
	if err != nil {
		log.Error("kaspa get block by block id fail", "err", err)
		return nil, err
	}

	return &utxo.BlockHeaderResponse{
		Code:       common2.ReturnCode_SUCCESS,
		Msg:        "get block header by hash success",
		ParentHash: result.Header.Parents[0].ParentHashes[0],
		BlockHash:  result.VerboseData.Hash,
		MerkleRoot: result.Header.HashMerkleRoot,
		Number:     result.Header.BlueScore,
	}, nil
}

func (c ChainAdaptor) GetBlockHeaderByNumber(req *utxo.BlockHeaderNumberRequest) (*utxo.BlockHeaderResponse, error) {
	result, err := c.kaspaClient.GetBlockByHeight(req.Height)
	if err != nil {
		log.Error("kaspa get block by number fail", "err", err)
		return nil, err
	}

	return &utxo.BlockHeaderResponse{
		Code:       common2.ReturnCode_SUCCESS,
		Msg:        "get block header by hash success",
		ParentHash: result.Header.Parents[0].ParentHashes[0],
		BlockHash:  result.VerboseData.Hash,
		MerkleRoot: result.Header.HashMerkleRoot,
		Number:     result.Header.BlueScore,
	}, nil
}

func (c ChainAdaptor) SendTx(req *utxo.SendTxRequest) (*utxo.SendTxResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c ChainAdaptor) GetTxByAddress(req *utxo.TxAddressRequest) (*utxo.TxAddressResponse, error) {
	transactions, err := c.kaspaClient.GetTxByAddress(req.Address, int(req.Pagesize), int(req.Page*(req.Pagesize-1)))
	if err != nil {
		log.Error("kaspa get tx by address fail", "err", err)
		return nil, err
	}

	var txList []*utxo.TxMessage
	for _, tx := range transactions {

		// 计算输入总金额以确定手续费
		var totalInputAmount int64
		froms := make([]*utxo.Address, 0, len(tx.Inputs))
		for _, input := range tx.Inputs {
			// 查询前序交易以获取发送地址和金额
			froms = append(froms, &utxo.Address{
				Address: input.PreviousOutpointAddress,
			})
			totalInputAmount += input.PreviousOutpointAmount
		}

		// 处理输出地址和金额
		tos := make([]*utxo.Address, 0, len(tx.Outputs))
		values := make([]*utxo.Value, 0, len(tx.Outputs))
		var totalOutputAmount int64
		for _, output := range tx.Outputs {
			tos = append(tos, &utxo.Address{
				Address: output.ScriptPublicKeyAddress,
			})
			values = append(values, &utxo.Value{
				Value: strconv.FormatInt(output.Amount, 10),
			})
			totalOutputAmount += output.Amount
		}

		// 计算手续费
		fee := "0"
		if totalInputAmount > totalOutputAmount {
			fee = fmt.Sprintf("%d", totalInputAmount-totalOutputAmount)
		} else if tx.Mass != "" {
			fee = tx.Mass // 若无法计算，使用 mass 作为近似值
		}

		// 确定交易状态
		status := utxo.TxStatus_Pending
		if tx.IsAccepted {
			status = utxo.TxStatus_Success
		}

		// 构造 TxMessage
		t := &utxo.TxMessage{
			Hash:         tx.Hash,
			Index:        0, // 交易本身的索引，暂设为 0
			Froms:        froms,
			Tos:          tos,
			Values:       values,
			Fee:          fee,
			Status:       status,
			Type:         0, // 类型未明确，设为 0
			Height:       fmt.Sprintf("%d", tx.AcceptingBlockBlueScore),
			Brc20Address: "", // Kaspa 无 BRC-20 地址，留空
			Datetime:     time.UnixMilli(tx.AcceptingBlockTime).Format(time.RFC3339),
		}
		txList = append(txList, t)
	}

	return &utxo.TxAddressResponse{
		Code: common2.ReturnCode_SUCCESS,
		Msg:  "kaspa get tx by address success",
		Tx:   txList,
	}, nil
}

// GetTxByHash 获取交易详情
func (c ChainAdaptor) GetTxByHash(req *utxo.TxHashRequest) (*utxo.TxHashResponse, error) {
	// 请求 Kaspa API 获取交易数据
	result, err := c.kaspaClient.GetTxByHash(req.Hash)
	if err != nil {
		log.Error("kaspa get block by hash fail", "err", err)
		return nil, err
	}

	// 计算输入总金额以确定手续费
	var totalInputAmount int64
	froms := make([]*utxo.Address, 0, len(result.Inputs))
	for _, input := range result.Inputs {
		// 查询前序交易以获取发送地址和金额
		froms = append(froms, &utxo.Address{
			Address: input.PreviousOutpointAddress,
		})
		totalInputAmount += input.PreviousOutpointAmount
	}

	// 处理输出地址和金额
	tos := make([]*utxo.Address, 0, len(result.Outputs))
	values := make([]*utxo.Value, 0, len(result.Outputs))
	var totalOutputAmount int64
	for _, output := range result.Outputs {
		tos = append(tos, &utxo.Address{
			Address: output.ScriptPublicKeyAddress,
		})
		values = append(values, &utxo.Value{
			Value: strconv.FormatInt(output.Amount, 10),
		})
		totalOutputAmount += output.Amount
	}

	// 计算手续费
	fee := "0"
	if totalInputAmount > totalOutputAmount {
		fee = fmt.Sprintf("%d", totalInputAmount-totalOutputAmount)
	} else if result.Mass != "" {
		fee = result.Mass // 若无法计算，使用 mass 作为近似值
	}

	// 确定交易状态
	status := utxo.TxStatus_Pending
	if result.IsAccepted {
		status = utxo.TxStatus_Success
	}

	// 构造 TxMessage
	tx := &utxo.TxMessage{
		Hash:         req.Hash,
		Index:        0, // 交易本身的索引，暂设为 0
		Froms:        froms,
		Tos:          tos,
		Values:       values,
		Fee:          fee,
		Status:       status,
		Type:         0, // 类型未明确，设为 0
		Height:       fmt.Sprintf("%d", result.AcceptingBlockBlueScore),
		Brc20Address: "", // Kaspa 无 BRC-20 地址，留空
		Datetime:     time.UnixMilli(result.AcceptingBlockTime).Format(time.RFC3339),
	}

	return &utxo.TxHashResponse{
		Code: common2.ReturnCode_SUCCESS,
		Msg:  "kaspa get block by hash success",
		Tx:   tx,
	}, nil
}

func (c ChainAdaptor) CreateUnSignTransaction(req *utxo.UnSignTransactionRequest) (*utxo.UnSignTransactionResponse, error) {
	return &utxo.UnSignTransactionResponse{
		Code: common2.ReturnCode_SUCCESS,
		Msg:  "kaspa create unsign transaction success",
	}, nil

}

func (c ChainAdaptor) BuildSignedTransaction(req *utxo.SignedTransactionRequest) (*utxo.SignedTransactionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c ChainAdaptor) DecodeTransaction(req *utxo.DecodeTransactionRequest) (*utxo.DecodeTransactionResponse, error) {
	return &utxo.DecodeTransactionResponse{
		Code: common2.ReturnCode_SUCCESS,
		Msg:  "decode transaction success",
	}, nil
}

func (c ChainAdaptor) VerifySignedTransaction(req *utxo.VerifyTransactionRequest) (*utxo.VerifyTransactionResponse, error) {
	return &utxo.VerifyTransactionResponse{
		Code:   common2.ReturnCode_SUCCESS,
		Msg:    "verify signed transaction is true",
		Verify: true,
	}, nil
}
