package contractsAPI

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"quasar-evm/config"
	"quasar-evm/contractsAPI/oracle"
	"quasar-evm/pkg/logger"
)

// IOracleAPI is an oracle smart-contract API package
type IOracleAPI interface {
	// PushPrice is used to push given price to the oracle smart contract for given token
	PushPrice(token *Token, price *big.Int) error
	// ParseTokens is used to parse tokens from given token symbols
	ParseTokens(symbols []string) ([]*Token, error)
	// Close is used to close the API package
	Close()
}

// oracleAPI is an implementation of the IOracleAPI
type oracleAPI struct {
	conf *config.WEB3

	rpc    *ethclient.Client
	signer *bind.TransactOpts
	oracle *oracle.Oracle
}

// NewOracleAPI is used to get new IOracleAPI implementation instance
func NewOracleAPI(rpc *ethclient.Client, conf *config.WEB3) (IOracleAPI, error) {
	o := &oracleAPI{
		rpc:  rpc,
		conf: conf,
	}

	if err := o.init(); err != nil {
		return nil, fmt.Errorf("err init orcale API: %w", err)
	}

	return o, nil
}

func (t *oracleAPI) init() (err error) {
	if t.conf.SignerPK == "" {
		return fmt.Errorf("no signer pk provided")
	}

	pk, err := crypto.HexToECDSA(t.conf.SignerPK)
	if err != nil {
		return fmt.Errorf("err get PK: %w", err)
	}

	if t.signer, err = bind.NewKeyedTransactorWithChainID(pk, big.NewInt(int64(t.conf.ChainID))); err != nil {
		return fmt.Errorf("err create signer: %w", err)
	}

	if t.oracle, err = oracle.NewOracle(common.HexToAddress(t.conf.OracleAddress), t.rpc); err != nil {
		return fmt.Errorf("err init contract instance: %w", err)
	}

	return nil
}

func (t *oracleAPI) PushPrice(token *Token, price *big.Int) error {
	tx, err := t.oracle.PushPrice(t.signer, token.ID, price)
	if err != nil {
		return fmt.Errorf("err push price: %w", err)
	}

	logger.Log().WithField("layer", "ContractsAPI-PushPrice").Infof(
		"push price: %v, for: %s, tx hash: %s", price.Uint64(), token.Symbol, tx.Hash(),
	)

	return nil
}

func (t *oracleAPI) ParseTokens(symbols []string) ([]*Token, error) {
	tokens := []*Token{}

	for _, s := range symbols {
		id, isSupported, err := t.oracle.GetCurrencyID(nil, s)
		if err != nil {
			logger.Log().WithField("layer", "ContractsAPI-ParseTokens").Errorf(
				"err GetCurrencyID: %s", err.Error(),
			)
			continue
		}

		if isSupported {
			tokens = append(tokens, &Token{Symbol: s, ID: id})
		}
	}

	return tokens, nil
}

func (t *oracleAPI) Close() {

}
