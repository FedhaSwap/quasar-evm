package service

import (
	"fmt"
	"math/big"
	"time"

	"quasar-evm/pkg/wsClient"
)

// runWriter is used to run writer
func (s *coinAVGPriceManager) runWriter() {
	logInfo("start", "Writer")
	go s.listenAndSendARGPrice(s.tokenNames(), s.id, s.stream, s.stopWriter)

	if err := s.subscribeCoinAveragePrice(s.tokenNames(), s.id, s.stream); err != nil {
		s.close()
	}
}

// subscribeCoinAveragePrice is used to send subscribe message to the ws server
func (s *coinAVGPriceManager) subscribeCoinAveragePrice(tokens []string, id int, stream chan *wsClient.CoinAveragePriceStream) error {
	if err := s.wsClient.SubscribeCoinAveragePrice(tokens, id, stream); err != nil {
		time.Sleep(time.Second * 5)
		return s.subscribeCoinAveragePrice(tokens, id, stream)
	}

	return nil
}

// listenAndSendARGPrice is used to listen to the CoinAveragePriceStream chanel and write prices to the
// coinAVGPriceManager. Will resubscribe automatically if needed.
func (s *coinAVGPriceManager) listenAndSendARGPrice(tokens []string, id int, stream chan *wsClient.CoinAveragePriceStream, stopWriter chan struct{}) {
	for {
		select {
		case <-stopWriter:
			logInfo("stop...", "Writer")
			return
		case <-s.resubscribe:
			if err := s.subscribeCoinAveragePrice(tokens, id, stream); err != nil {
				logErr(fmt.Sprintln("err resubscribe:", err.Error()), "Writer")
				return
			}
		case data := <-stream:
			logInfo(fmt.Sprintf("received data on the %s coin", data.Coin), "Writer")

			for _, t := range s.tokens {
				if data.Coin == t.Symbol {
					price := big.NewFloat(data.Value)
					price = price.Mul(price, big.NewFloat(100000))
					integer, _ := price.Int64()

					s.prices.Store(t, big.NewInt(integer))
					break
				}
			}
		}
	}
}
