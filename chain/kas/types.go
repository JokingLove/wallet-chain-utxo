package kas

type BalanceResp struct {
	Address string `json:"address"`
	Balance int64  `json:"balance"`
}
type Utxo struct {
	Address  string   `json:"address"` // Kaspa 地址
	Outpoint struct { // 输出点
		TransactionID string `json:"transactionId"` // 交易 ID
		Index         int    `json:"index"`         // 输出索引
	} `json:"outpoint"`
	UtxoEntry struct { // UTXO 条目
		Amount          string   `json:"amount"` // 金额，单位：sompi (字符串类型)
		ScriptPublicKey struct { // 脚本公钥
			ScriptPublicKey string `json:"scriptPublicKey"` // 脚本公钥内容
		} `json:"scriptPublicKey"`
		BlockDaaScore string `json:"blockDaaScore"` // 区块 DAA 分数 (字符串类型)
		IsCoinbase    bool   `json:"isCoinbase"`    // 是否为 Coinbase 交易
	} `json:"utxoEntry"`
}

// FeeEstimate 定义费用估计的完整结构
type FeeEstimate struct {
	PriorityBucket FeeBucket   `json:"priorityBucket"` // 高优先级桶
	NormalBuckets  []FeeBucket `json:"normalBuckets"`  // 普通桶列表
	LowBuckets     []FeeBucket `json:"lowBuckets"`     // 低优先级桶列表
}

// FeeBucket 定义单个桶的结构（内嵌使用）
type FeeBucket struct {
	Feerate          int64   `json:"feerate"`          // 费用率
	EstimatedSeconds float64 `json:"estimatedSeconds"` // 预计秒数
}

// Block 定义区块的结构，内嵌所有子结构体
type Block struct {
	Header struct { // 内嵌区块头
		Version              int64      `json:"version"`              // 版本号 (int)
		HashMerkleRoot       string     `json:"hashMerkleRoot"`       // Merkle 根哈希 (string)
		AcceptedIDMerkleRoot string     `json:"acceptedIdMerkleRoot"` // 已接受 ID 的 Merkle 根 (string)
		UtxoCommitment       string     `json:"utxoCommitment"`       // UTXO 承诺 (string)
		Timestamp            string     `json:"timestamp"`            // 时间戳 (string)
		Bits                 uint32     `json:"bits"`                 // 难度位 (uint32)
		Nonce                string     `json:"nonce"`                // Nonce (string)
		DaaScore             string     `json:"daaScore"`             // DAA 分数 (string)
		BlueWork             string     `json:"blueWork"`             // 蓝色工作量 (string)
		Parents              []struct { // 内嵌父区块列表
			ParentHashes []string `json:"parentHashes"` // 父区块哈希列表 (string array)
		} `json:"parents"`
		BlueScore    string `json:"blueScore"`    // 蓝色分数 (string)
		PruningPoint string `json:"pruningPoint"` // 剪枝点 (string)
	} `json:"header"`
	Transactions []struct { // 内嵌交易列表
		Inputs []struct { // 内嵌输入列表
			PreviousOutpoint struct { // 内嵌前一输出点
				TransactionID string `json:"transactionId"` // 交易 ID (string)
				Index         uint32 `json:"index"`         // 索引 (uint32)
			} `json:"previousOutpoint"`
			SignatureScript string `json:"signatureScript"` // 签名脚本 (string)
			SigOpCount      int    `json:"sigOpCount"`      // 签名操作计数 (int)
			Sequence        string `json:"sequence"`        // 序列号 (string)
		} `json:"inputs"`
		Outputs []struct { // 内嵌输出列表
			Amount          string   `json:"amount"` // 金额 (string)
			ScriptPublicKey struct { // 内嵌脚本公钥
				ScriptPublicKey string `json:"scriptPublicKey"` // 脚本公钥内容 (string)
				Version         int    `json:"version"`         // 版本 (int)
			} `json:"scriptPublicKey"`
			VerboseData struct { // 内嵌详细数据
				ScriptPublicKeyType    string `json:"scriptPublicKeyType"`    // 脚本公钥类型 (string)
				ScriptPublicKeyAddress string `json:"scriptPublicKeyAddress"` // 脚本公钥地址 (string)
			} `json:"verboseData"`
		} `json:"outputs"`
		SubnetworkID string   `json:"subnetworkId"` // 子网络 ID (string)
		Payload      string   `json:"payload"`      // 负载 (string, 可为空字符串)
		VerboseData  struct { // 内嵌交易详细数据
			TransactionID string `json:"transactionId"` // 交易 ID (string)
			Hash          string `json:"hash"`          // 哈希 (string)
			BlockHash     string `json:"blockHash"`     // 区块哈希 (string)
			BlockTime     string `json:"blockTime"`     // 区块时间 (string)
			ComputeMass   string `json:"computeMass"`   // 计算质量 (string)
		} `json:"verboseData"`
		Version  int    `json:"version"`  // 版本 (int)
		LockTime string `json:"lockTime"` // 锁定时间 (string)
		Gas      string `json:"gas"`      // Gas (string)
		Mass     string `json:"mass"`     // 质量 (string)
	} `json:"transactions"`
	VerboseData struct { // 内嵌区块详细数据
		Hash                string   `json:"hash"`                // 区块哈希 (string)
		Difficulty          float64  `json:"difficulty"`          // 难度值 (float)
		SelectedParentHash  string   `json:"selectedParentHash"`  // 选择的父区块哈希 (string)
		TransactionIDs      []string `json:"transactionIds"`      // 交易 ID 列表 (string array)
		BlueScore           string   `json:"blueScore"`           // 蓝色分数 (string)
		ChildrenHashes      []string `json:"childrenHashes"`      // 子区块哈希列表 (string array)
		MergeSetBluesHashes []string `json:"mergeSetBluesHashes"` // 蓝色合并集哈希列表 (string array)
		MergeSetRedsHashes  []string `json:"mergeSetRedsHashes"`  // 红色合并集哈希列表 (string array)
		IsChainBlock        bool     `json:"isChainBlock"`        // 是否为链上区块 (bool)
	} `json:"verboseData"`
	Extra struct { // 内嵌额外信息
		Color        interface{} `json:"color"`        // 颜色 (null，用 interface{})
		MinerAddress string      `json:"minerAddress"` // 矿工地址 (string)
		MinerInfo    string      `json:"minerInfo"`    // 矿工信息 (string)
	} `json:"extra"`
}

// Transaction 定义交易的结构，内嵌所有子结构体
type Transaction struct {
	SubnetworkID            string      `json:"subnetwork_id"`              // 子网络 ID (string)
	TransactionID           string      `json:"transaction_id"`             // 交易 ID (string)
	Hash                    string      `json:"hash"`                       // 哈希 (string)
	Mass                    string      `json:"mass"`                       // 质量 (string, 可为 null 但示例中是字符串)
	Payload                 interface{} `json:"payload"`                    // 负载 (null，用 interface{})
	BlockHash               []string    `json:"block_hash"`                 // 区块哈希列表 (string array)
	BlockTime               int64       `json:"block_time"`                 // 区块时间 (int)
	IsAccepted              bool        `json:"is_accepted"`                // 是否被接受 (bool)
	AcceptingBlockHash      string      `json:"accepting_block_hash"`       // 接受区块哈希 (string)
	AcceptingBlockBlueScore int64       `json:"accepting_block_blue_score"` // 接受区块蓝色分数 (int)
	AcceptingBlockTime      int64       `json:"accepting_block_time"`       // 接受区块时间 (int)
	Inputs                  []struct {  // 内嵌输入列表
		TransactionID            string   `json:"transaction_id"`          // 交易 ID (string)
		Index                    int      `json:"index"`                   // 输入索引 (int)
		PreviousOutpointHash     string   `json:"previous_outpoint_hash"`  // 前一输出点哈希 (string)
		PreviousOutpointIndex    string   `json:"previous_outpoint_index"` // 前一输出点索引 (string)
		PreviousOutpointResolved struct { // 内嵌前一输出点解析数据
			TransactionID          string `json:"transaction_id"`            // 交易 ID (string)
			Index                  int    `json:"index"`                     // 输出索引 (int)
			Amount                 int64  `json:"amount"`                    // 金额 (int)
			ScriptPublicKey        string `json:"script_public_key"`         // 脚本公钥 (string)
			ScriptPublicKeyAddress string `json:"script_public_key_address"` // 脚本公钥地址 (string)
			ScriptPublicKeyType    string `json:"script_public_key_type"`    // 脚本公钥类型 (string)
		} `json:"previous_outpoint_resolved"`
		PreviousOutpointAddress string `json:"previous_outpoint_address"` // 前一输出点地址 (string)
		PreviousOutpointAmount  int64  `json:"previous_outpoint_amount"`  // 前一输出点金额 (int)
		SignatureScript         string `json:"signature_script"`          // 签名脚本 (string)
		SigOpCount              string `json:"sig_op_count"`              // 签名操作计数 (string)
	} `json:"inputs"`
	Outputs []struct { // 内嵌输出列表
		TransactionID          string `json:"transaction_id"`            // 交易 ID (string)
		Index                  int    `json:"index"`                     // 输出索引 (int)
		Amount                 int64  `json:"amount"`                    // 金额 (int)
		ScriptPublicKey        string `json:"script_public_key"`         // 脚本公钥 (string)
		ScriptPublicKeyAddress string `json:"script_public_key_address"` // 脚本公钥地址 (string)
		ScriptPublicKeyType    string `json:"script_public_key_type"`    // 脚本公钥类型 (string)
	} `json:"outputs"`
}
