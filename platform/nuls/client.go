package nuls

import (
	"fmt"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Client struct {
	blockatlas.Request
}

const (
	JsonRpcVersion = "2.0"
	PubSvcPath     = "/"
	Id             = 8964
	GetAccountTxsMethod  = "getAccountTxs"
	ChainId        = 1
)

func InitJsonprc(method string, params interface{}, jsonrpc *JsonrpcRequest) {
	jsonrpc.JsonRpc = JsonRpcVersion
	jsonrpc.Method = method
	jsonrpc.Id = Id
	jsonrpc.Params = params
}

func (c *Client) GetTxsOfAddress(address, token string) ([]Tx, error) {
	var rpcResponse JsonRpcResponse
	var rpcRequest JsonRpcRequest
	var params GetAccountTxsParam = {
		ChainId: ChainId,
		PageNumber: 1,
		PageSize: 100,
		Address: address,
		TxType: 0,
		IsHidden: false,
	}
	InitJsonprc(GetAccountTxsMethod, params, &rpcRequest)
	err := c.Post(&txs, PubSvcPath, body)
	return rpcResponse.Result.List, err
}

func (c *Client) GetAccount(address string) (*Account, error) {
	path := fmt.Sprintf("v1/accounts/%s", address)

	var accounts Account
	err := c.Get(&accounts, path, nil)

	return &accounts, err
}

func (c *Client) GetAccountVotes(address string) (*AccountData, error) {
	var account AccountData
	err := c.Post(&account, "wallet/getaccount", VotesRequest{Address: address, Visible: true})
	return &account, err
}

func (c *Client) GetTokenInfo(id string) (*Asset, error) {
	path := fmt.Sprintf("v1/assets/%s", id)

	var asset Asset
	err := c.Get(&asset, path, nil)

	return &asset, err
}

func (c *Client) GetValidators() (validators Validators, err error) {
	err = c.Get(&validators, "wallet/listwitnesses", nil)
	if err != nil {
		return validators, err
	}
	return validators, err
}
