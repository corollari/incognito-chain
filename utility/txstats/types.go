package main

type BCInfoJsonResponse struct {
	ID     int                     `json:"Id"`
	Result GetBlockChainInfoResult `json:"Result"`
}

type BlockJsonResponse struct {
	ID     int             `json:"Id"`
	Result []BlockResponse `json:"Result"`
}

type TxJsonResponse struct {
	ID     int               `json:"Id"`
	Result TransactionDetail `json:"Result"`
}

type GetBlockChainInfoResult struct {
	ChainName    string                   `json:"ChainName"`
	BestBlocks   map[int]GetBestBlockItem `json:"BestBlocks"`
	ActiveShards int                      `json:"ActiveShards"`
}

type GetBestBlockItem struct {
	Height              uint64 `json:"Height"`
	Hash                string `json:"Hash"`
	TotalTxs            uint64 `json:"TotalTxs"`
	BlockProducer       string `json:"BlockProducer"`
	ValidationData      string `json:"ValidationData"`
	Epoch               uint64 `json:"Epoch"`
	Time                int64  `json:"Time"`
	RemainingBlockEpoch uint64 `json:"RemainingBlockEpoch"`
	EpochBlock          uint64 `json:"EpochBlock"`
}

type BlockResponse struct {
	Txs    []GetBlockTxResult `json:"Txs"`
	Height uint64             `json:"Height"`
}

type GetBlockTxResult struct {
	Hash     string `json:"Hash"`
	Locktime int64  `json:"Locktime"`
	HexData  string `json:"HexData"`
}

type TransactionDetail struct {
	BlockHash   string `json:"BlockHash"`
	BlockHeight uint64 `json:"BlockHeight"`
	TxSize      uint64 `json:"TxSize"`
	Index       uint64 `json:"Index"`
	ShardID     byte   `json:"ShardID"`
	Hash        string `json:"Hash"`
	Version     int8   `json:"Version"`
	Type        string `json:"Type"` // Transaction type
	LockTime    string `json:"LockTime"`
	IsPrivacy   bool   `json:"IsPrivacy"`
}
