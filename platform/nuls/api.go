package nuls

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	services "github.com/trustwallet/blockatlas/services/assets"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client = Client{blockatlas.InitClient(viper.GetString("nuls.api"))}
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.NULS]
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	Txs, err := p.client.GetTxsOfAddress(address, "")
	if err != nil && len(Txs) == 0 {
		return nil, err
	}

	var txs []blockatlas.Tx
	for _, srcTx := range Txs {
		tx, ok := Normalize(address, &srcTx)
		if ok {
			txs = append(txs, tx)
		}
	}

	return txs, nil
}

func (p *Platform) GetValidators() (blockatlas.ValidatorPage, error) {
	results := make(blockatlas.ValidatorPage, 0)
	validators, err := p.client.GetValidators()

	if err != nil {
		return results, err
	}

	for _, v := range validators {
		if val, ok := normalizeValidator(v); ok {
			results = append(results, val)
		}
	}

	return results, nil
}

func (p *Platform) GetDetails() blockatlas.StakingDetails {
	return blockatlas.StakingDetails{
		Reward:        blockatlas.StakingReward{Annual: 12},
		MinimumAmount: blockatlas.Amount("2000"),
		LockTime:      0,
		Type:          blockatlas.DelegationTypeDelegate,
	}
}

/// Normalize converts a Nuls transaction into the generic model
func Normalize(from string, srcTx *Tx) (tx blockatlas.Tx, ok bool) {
	return blockatlas.Tx{
		ID:   srcTx.TxHash,
		Coin: coin.NULS,
		Date: srcTx.CreateTime / 1000,
		From: from,
		To:   srcTx.Address,
		Fee:  srcTx.Fee.Value,
		Meta: blockatlas.Transfer{
			Value:    srcTx.Values,
			Symbol:   coin.Coins[coin.TRX].Symbol,
			Decimals: coin.Coins[coin.TRX].Decimals,
		},
	}, true
}

func getDetails(commissionRate int) blockatlas.StakingDetails {
	return blockatlas.StakingDetails{
		Reward:        blockatlas.StakingReward{Annual: float64(commissionRate / 100.0)},
		MinimumAmount: blockatlas.Amount("2000"),
		LockTime:      0,
		Type:          blockatlas.DelegationTypeDelegate,
	}
}

func normalizeValidator(v Validator) (validator blockatlas.Validator, ok bool) {
	return blockatlas.Validator{
		Status:  true,
		ID:      v.RewardAddress,
		Details: getDetails(v.CommissionRate),
	}, true
}

func (p *Platform) GetDelegations(address string) (blockatlas.DelegationsPage, error) {
	results := make(blockatlas.DelegationsPage, 0)

	delegations, err := p.client.GetDelegations(address)
	if err != nil {
		return nil, err
	}
	validators, err := services.GetValidatorsMap(p)
	if err != nil {
		return nil, err
	}
	results = append(results, NormalizeDelegations(delegations, validators)...)
	return results, nil
}

func (p *Platform) UndelegatedBalance(address string) (string, error) {
	return "0", nil
}

func NormalizeDelegations(delegations []Delegation, validators blockatlas.ValidatorMap) []blockatlas.Delegation {
	results := make([]blockatlas.Delegation, 0)
	for _, v := range delegations {
		validator, ok := validators[v.Address]
		if !ok {
			logger.Error("Validator not found", validator)
			continue
		}
		delegation := blockatlas.Delegation{
			Delegator: validator,
			Value:     string(v.Amount),
			Status:    blockatlas.DelegationStatusActive,
		}
		results = append(results, delegation)
	}
	return results
}
