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
	GetConsensusNodes  = "getConsensusNodes"
	GetAccountConsensus = "getAccountConsensus"
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
		PageNumber: {strconv.FormatInt(1, 10)},
		PageSize: {strconv.FormatInt(blockatlas.ValidatorsPerPage, 10)},
		Address: address,
		TxType: 0,
		IsHidden: false,
	}
	InitJsonprc(GetAccountTxsMethod, params, &rpcRequest)
	err := c.Post(&rpcResponse, PubSvcPath, JsonRpcRequest)
	return rpcResponse.Result.List, err
}

func (c *Client) GetValidators() (validators []Validator, err error) {
	var rpcResponse JsonRpcResponse
	var rpcRequest JsonRpcRequest
	var params GetConsensusNodesParam = {
		ChainId: ChainId,
		PageNumber: {strconv.FormatInt(1, 10)},
		PageSize: {strconv.FormatInt(blockatlas.ValidatorsPerPage, 10)},
		Type: 0,  // all nodes
	}
	InitJsonprc(GetConsensusNodes, params, &rpcRequest)
	err := c.Post(&rpcResponse, PubSvcPath, JsonRpcRequest)
	return rpcResponse.Result.List, err
}

func (c *Client) GetDelegations(address string) (delegations []Delegation, err error) {
	var rpcResponse JsonRpcResponse
	var rpcRequest JsonRpcRequest
	var params GetConsensusNodesParam = {
		ChainId: ChainId,
		PageNumber: {strconv.FormatInt(1, 10)},
		PageSize: {strconv.FormatInt(blockatlas.ValidatorsPerPage, 10)},
		Address: address,
		AgentHash: null,
	}
	InitJsonprc(GetAccountConsensus, params, &rpcRequest)
	err := c.Post(&rpcResponse, PubSvcPath, JsonRpcRequest)
	return rpcResponse.Result.List, err
}