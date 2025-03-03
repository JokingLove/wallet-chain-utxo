package kas

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	gresty "github.com/go-resty/resty/v2"
	"strconv"
)

var errBlockChainHTTPError = errors.New("kaspa http error")

type KasClient struct {
	client *gresty.Client
}

func NewKasClient(url string) (*KasClient, error) {
	// validate if blockchain url is provided or not
	if url == "" {
		return nil, fmt.Errorf("blockchain URL cannot be empty")
	}

	client := gresty.New()
	client.SetBaseURL(url)
	client.OnAfterResponse(func(c *gresty.Client, r *gresty.Response) error {
		statusCode := r.StatusCode()
		if statusCode >= 400 {
			method := r.Request.Method
			url := r.Request.URL
			return fmt.Errorf("%d cannot %s %s: %w", statusCode, method, url, errBlockChainHTTPError)
		}
		return nil
	})
	return &KasClient{
		client: client,
	}, nil
}

func (kc *KasClient) GetAccountBalance(address string) (int64, error) {
	// 定义响应结构体
	respData := &BalanceResp{}
	// 发送 GET 请求
	resp, err := kc.client.R().
		SetHeader("Accept", "application/json"). // 设置请求头
		SetResult(respData).                     // 自动解析 JSON 到结构体
		Get("/addresses/" + address + "/balance")

	if err != nil {
		log.Error("kaspa get account balance error", "address", address, "err", err)
		return 0, err
	}

	// 检查响应状态（OnAfterResponse 已处理，这里仅为额外安全）
	if resp.IsError() {
		return 0, fmt.Errorf("API returned error: %s", resp.String())
	}

	return respData.Balance, nil

}

func (kc *KasClient) GetUtxo(address string) ([]Utxo, error) {
	var utxos []Utxo
	response, err := kc.client.R().
		SetResult(&utxos).
		Get("/addresses/" + address + "/utxos")
	if err != nil {
		log.Error("kaspa get utxos error", "address", address, "err", err)
		return nil, err
	}

	if !response.IsSuccess() {
		return nil, fmt.Errorf("API returned error: %s", response.String())
	}
	return utxos, nil
}

func (kc *KasClient) FeeEstimate() (FeeEstimate, error) {
	var fee FeeEstimate
	response, err := kc.client.R().
		SetResult(&fee).
		Get("/info/fee-estimate")
	if err != nil {
		log.Error("get kaspa fee estimate fail", "err", err)
		return fee, err
	}

	if !response.IsSuccess() {
		return fee, fmt.Errorf("API returned error: %s", response.String())
	}
	return fee, nil
}

func (kc *KasClient) GetBlockByHash(hash string) (*Block, error) {
	var block Block
	response, err := kc.client.R().
		SetResult(&block).
		Get("/blocks/" + hash + "?includeTransactions=true&includeColor=false")
	if err != nil {
		log.Error("kaspa get block by hash error", "hash", hash, "err", err)
		return nil, err
	}
	if !response.IsSuccess() {
		return nil, fmt.Errorf("GetBlockByHash API returned error: %s", response.String())
	}
	return &block, nil
}

func (kc *KasClient) GetBlockByHeight(height int64) (*Block, error) {
	url := "/blocks-from-bluescore?includeTransactions=true"
	if height > 0 {
		url += "&blueScore=" + strconv.FormatInt(height, 10)
	}
	var block []*Block
	response, err := kc.client.R().
		SetResult(&block).
		Get(url)
	if err != nil {
		log.Error("kaspa get block by blue score error", "height", height, "err", err)
		return nil, err
	}
	if !response.IsSuccess() {
		return nil, fmt.Errorf("GetBlockByHeight API returned error: %s", response.String())
	}
	return block[0], nil
}

func (kc *KasClient) GetTxByHash(hash string) (Transaction, error) {
	var tx Transaction
	response, err := kc.client.R().
		SetResult(&tx).
		Get("transactions/" + hash + "?inputs=true&outputs=true&resolve_previous_outpoints=full")
	if err != nil {
		log.Error("kaspa get block by hash error", "hash", hash, "err", err)
		return tx, err
	}
	if !response.IsSuccess() {
		return tx, fmt.Errorf("GetTxByHash API returned error: %s", response.String())
	}
	return tx, nil
}

func (kc *KasClient) GetTxByAddress(address string, pagesize, offset int) ([]*Transaction, error) {
	var txs []*Transaction
	response, err := kc.client.R().
		SetResult(&txs).
		Get(
			fmt.Sprintf("/addresses/%s/full-transactions?limit=%d&offset=%d&resolve_previous_outpoints=full",
				address, pagesize, offset))
	if err != nil {
		log.Error("kaspa get block by address error", "address", address, "err", err)
		return nil, err
	}
	if !response.IsSuccess() {
		return nil, fmt.Errorf("GetTxByAddress API returned error: %s", response.String())
	}
	return txs, nil
}

//func (kc *KasClient) SendTx() (*Transaction, error) {
//	kc.client.R().
//		SetResult()
//}
