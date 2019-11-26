package nuls

import (
	"encoding/json"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type JsonRpcRequest struct {
	JsonRpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Id      uint        `json:"id"`
}
type GetAccountTxsParam struct {
	ChainId    uint   `json:"chainId"`
	PageNumber uint   `json:"pageNumber"`
	PageSize   uint   `json:"pageSize"`
	Address    string `json:"address"`
	TxType     unit   `json:"txType"`
	IsHidden   bool   `json:"isHidden"`
}

type JsonRpcResponse struct {
	Jsonrpc string      `json:"jsonrpc"`
	Id      uint        `json:"id"`
	Result  interface{} `json:"result"`
}

type Tx struct {
	TxHash     string `json:"txHash"`
	Address    string `json:"address"`
	Type       unit   `json:"type"`
	CreateTime int64  `json:"createTime"`
}

/*
"txHash": "a8611112f2b35385ee84f85……",		//交易hash
"address": "tNULSeBaMrbMRiFA……",			//账户地址
"type": 1,									//交易类型
"createTime": 1531152,						//交易时间，单位秒
"height": 0,								//交易被打包确定的区块高度
"chainId": 2,								//资产的链id
"assetId": 1,								//资产id
"symbol": "NULS",							//资产符号
"values": 1000000000000000,					//交易金额
"fee": { 									//bigInt	手续费
	"chainId": 100,							//手续费链id
	"assetId": 1,							//手续费资产id
	"symbol": "ATOM",						//手续费资产符号
	"value": 100000							//手续费金额
},
"balance": 1000000000000000,				//交易后账户的余额
"transferType": 1,							// -1:转出, 1:转入
"status": 1									//交易状态 0:未确认,1:已确认
}
*/

type Page struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
	Txs     []Tx   `json:"data"`
}

type Tx struct {
	ID        string `json:"txID"`
	BlockTime int64  `json:"block_timestamp"`
	Data      TxData `json:"raw_data"`
}

type TxData struct {
	Contracts []Contract `json:"contract"`
}

type Contract struct {
	Type      string      `json:"type"`
	Parameter interface{} `json:"parameter"`
}

type TransferContract struct {
	Value TransferValue `json:"value"`
}

type TransferValue struct {
	Amount       blockatlas.Amount `json:"amount"`
	OwnerAddress string            `json:"owner_address"`
	ToAddress    string            `json:"to_address"`
}

// Type for token transfer
type TransferAssetContract struct {
	Value TransferAssetValue `json:"value"`
}

type TransferAssetValue struct {
	TransferValue
	AssetName string `json:"asset_name"`
}

type Account struct {
	Data []AccountData `json:"data"`
}

type AccountData struct {
	Balance  uint      `json:"balance"`
	AssetsV2 []AssetV2 `json:"assetV2"`
	Votes    []Votes   `json:"votes"`
	Frozen   []Frozen  `json:"frozen"`
}

type AssetV2 struct {
	Key string `json:"key"`
}

type Votes struct {
	VoteAddress string `json:"vote_address"`
	VoteCount   int    `json:"vote_count"`
}

type Frozen struct {
	ExpireTime    int64       `json:"expire_time"`
	FrozenBalance interface{} `json:"frozen_balance,string"`
}

type Asset struct {
	Data []AssetInfo `json:"data"`
}

type AssetInfo struct {
	Name     string `json:"name"`
	Symbol   string `json:"abbr"`
	ID       string `json:"id"`
	Decimals uint   `json:"precision"`
}

type Validators struct {
	Witnesses []Validator `json:"witnesses"`
}

type Validator struct {
	Address string `json:"address"`
}

type VotesRequest struct {
	Address string `json:"address"`
	Visible bool   `json:"visible"`
}

func (c *Contract) UnmarshalJSON(buf []byte) error {
	var contractInternal struct {
		Type      string          `json:"type"`
		Parameter json.RawMessage `json:"parameter"`
	}
	err := json.Unmarshal(buf, &contractInternal)
	if err != nil {
		return err
	}
	switch contractInternal.Type {
	case "TransferContract":
		var transfer TransferContract
		err = json.Unmarshal(contractInternal.Parameter, &transfer)
		c.Parameter = transfer
	case "TransferAssetContract":
		var tokenTransfer TransferAssetContract
		err = json.Unmarshal(contractInternal.Parameter, &tokenTransfer)
		c.Parameter = tokenTransfer
	}
	return err
}
