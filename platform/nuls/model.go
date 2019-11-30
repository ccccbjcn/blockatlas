package nuls

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type JsonRpcRequest struct {
	JsonRpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      int         `json:"id"`
}

type JsonRpcResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  Result `json:"result"`
}

type Result struct {
	PageNumber int           `json:"pageNumber"`
	PageSize   int           `json:"pageSize"`
	TotalCount int           `json:"totalCount"`
	List       []interface{} `json:"list"`
}

type Tx struct {
	TxHash       string            `json:"txHash"`
	Address      string            `json:"address"`
	Type         int               `json:"type"`
	CreateTime   int64             `json:"createTime"`
	Heigth       int64             `json:"height"`
	ChainId      int               `json:"chainId"`
	AssetId      int               `json:"assetId"`
	Symbol       string            `json:"symbol"`
	Values       blockatlas.Amount `json:"values"`
	Fee          Fee               `json:"fee"`
	Balance      string            `json:"balance"`
	TransferType int               `json:"transferType"`
	Status       int               `json:"status"`
}
type Fee struct {
	ChainId int               `json:"chainId"`
	AssetId int               `json:"assetId"`
	Symbol  string            `json:"symbol"`
	Value   blockatlas.Amount `json:"value"`
}

type GetAccountTxsParam struct {
	ChainId    int    `json:"chainId"`
	PageNumber int    `json:"pageNumber"`
	PageSize   int    `json:"pageSize"`
	Address    string `json:"address"`
	TxType     int    `json:"txType"`
	IsHidden   bool   `json:"isHidden"`
}

type Validator struct {
	TxHash            string            `json:"txHash"`
	AgentId           string            `json:"agentId"`
	AgentAddress      string            `json:"agentAddress"`
	PackingAddress    string            `json:"packingAddress"`
	RewardAddress     string            `json:"rewardAddress"`
	AgentAlias        string            `json:"agentAlias"`
	Deposit           blockatlas.Amount `json:"deposit"`
	CommissionRate    int               `json:"commissionRate"`
	CreateTime        int64             `json:"createTime"`
	Status            int               `json:"status"`
	TotalDeposit      blockatlas.Amount `json:"totalDeposit"`
	DepositCount      int               `json:"depositCount"`
	CreditValue       float64           `json:"creditValue"`
	TotalPackingCount int               `json:"totalPackingCount"`
	LostRate          float64           `json:"lostRate"`
	LastRewardHeight  int64             `json:"lastRewardHeight"`
	DeleteHash        string            `json:"deleteHash"`
	BlockHeight       int64             `json:"blockHeight"`
	DeleteHeight      int64             `json:"deleteHeight"`
	TotalReward       blockatlas.Amount `json:"totalReward"`
	CommissionReward  blockatlas.Amount `json:"commissionReward"`
	AgentReward       blockatlas.Amount `json:"agentReward"`
	RoundPackingTime  int64             `json:"roundPackingTime"`
	Version           int               `json:"version"`
	Type              int               `json:"type"`
}

type GetConsensusNodesParam struct {
	ChainId    int `json:"chainId"`
	PageNumber int `json:"pageNumber"`
	PageSize   int `json:"pageSize"`
	Type       int `json:"type"`
}

type Delegation struct {
	TxHash       string            `json:"txHash"`
	Amount       blockatlas.Amount `json:"amount"`
	AgentHash    string            `json:"agentHash"`
	Address      string            `json:"address"`
	CreateTime   int64             `json:"createTime"`
	BlockHeight  int64             `json:"blockHeight"`
	DeleteHeight int64             `json:"deleteHeight"`
	Type         int               `json:type`
	Fee          Fee               `json:"fee"`
}
type GetAccountConsensusParam struct {
	ChainId    int    `json:"chainId"`
	PageNumber int    `json:"pageNumber"`
	PageSize   int    `json:"pageSize"`
	Address    string `json:"address"`
	AgentHash  string `json:"agentHash"`
}
