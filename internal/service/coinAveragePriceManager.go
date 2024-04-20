package service

import (
	"fmt"
	"math/big"
	"time"

	"golang.org/x/sync/syncmap"

	"quasar-evm/contractsAPI"
	"quasar-evm/pkg/wsClient"
)

// SendCoinAveragePrice is used to subscribe on the avg price and send results to the oracle smart-contract
func (s *service) SendCoinAveragePrice(tokens []string) {
	parsedTokens, err := s.oracle.ParseTokens(tokens)
	if err != nil {
		logErr(fmt.Sprintf("err parse tokens: %s", err.Error()), "SendCoinAveragePrice")
		return
	}

	if len(parsedTokens) == 0 {
		logErr("err parse tokens: all given tokens are invalid. All tokens should be registered in the oracle "+
			"smart-contract and have same symbol in the contract and demon", "SendCoinAveragePrice")
		return
	}

	manager := newCoinAvgPriceManager(len(s.avgPriceManagers), s.oracle, s.wsClient, parsedTokens)
	s.avgPriceManagers = append(s.avgPriceManagers, manager)

	go manager.runWriter()

	manager.ticker = time.NewTicker(s.sendFreqMS)
	logInfo(fmt.Sprintf("set ticker for sender to %v", s.sendFreqMS), "SendCoinAveragePrice")
	go manager.runSender()
}

// coinAVGPriceManager is used to manage the write-send flow
type coinAVGPriceManager struct {
	// id is a WS id
	id int

	oracle   contractsAPI.IOracleAPI
	wsClient wsClient.IWSClient

	stream      chan *wsClient.CoinAveragePriceStream
	stopWriter  chan struct{}
	stopSender  chan struct{}
	resubscribe chan struct{}

	ticker *time.Ticker

	// tokens are the tokens for each write-send flow
	tokens []*contractsAPI.Token
	// prices are the prices for next send action
	prices *syncmap.Map
}

// newCoinAvgPriceManager is used to get new coinAVGPriceManager instance
func newCoinAvgPriceManager(
	id int,
	oracle contractsAPI.IOracleAPI,
	ws wsClient.IWSClient,
	tokens []*contractsAPI.Token,
) *coinAVGPriceManager {
	return &coinAVGPriceManager{
		id:          id,
		oracle:      oracle,
		wsClient:    ws,
		stream:      make(chan *wsClient.CoinAveragePriceStream),
		stopWriter:  make(chan struct{}),
		stopSender:  make(chan struct{}),
		resubscribe: make(chan struct{}),
		tokens:      tokens,
		prices:      &syncmap.Map{},
	}
}

// parsePrices is used to get big int prices from the prices map
func (s *coinAVGPriceManager) parsePrices() (prices []*big.Int) {
	for _, t := range s.tokens {
		p := big.NewInt(0)
		pp, ok := s.prices.Load(t)
		if ok {
			pb, ok := pp.(*big.Int)
			if ok {
				p = pb
			}
		}
		prices = append(prices, p)
	}

	return prices
}

// tokenNames is used to get token names string from the tokens struct
func (s *coinAVGPriceManager) tokenNames() []string {
	tokenNames := []string{}

	for _, t := range s.tokens {
		tokenNames = append(tokenNames, t.Symbol)
	}

	return tokenNames
}

// close is used to close coin average price sender
func (s *coinAVGPriceManager) close() {
	close(s.stopWriter)
	close(s.stopSender)
}
