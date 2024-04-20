package service

import (
	"time"

	"quasar-evm/contractsAPI"
	"quasar-evm/pkg/wsClient"
)

// IService is a service layer interface
type IService interface {
	// SendCoinAveragePrice is used to send coin average price from the ws service to the oracle smart-contracts
	SendCoinAveragePrice(tokens []string)
	// Close is used to stop the service
	Close()
}

// service is a service-layer struct implementing IService interface
type service struct {
	oracle   contractsAPI.IOracleAPI
	wsClient wsClient.IWSClient

	sendFreqMS time.Duration

	avgPriceManagers []*coinAVGPriceManager
	resubscribe      chan struct{}
}

// NewService is used to get new service instance
func NewService(ws wsClient.IWSClient, oracle contractsAPI.IOracleAPI, sendFreqMS time.Duration) IService {
	logInfo("creating new service...", "Init")
	c := &service{
		avgPriceManagers: make([]*coinAVGPriceManager, 0),
		sendFreqMS:       sendFreqMS,
	}

	if ws != nil {
		c.wsClient = ws
		c.resubscribe = ws.Resubscribe()
	}

	if oracle != nil {
		c.oracle = oracle
	}

	go c.listenResubscribe()

	return c
}

func (s *service) listenResubscribe() {
	for {
		select {
		case <-s.resubscribe:
			for _, v := range s.avgPriceManagers {
				v.resubscribe <- struct{}{}
			}
		}
	}
}

// Close is used to close the service and all dependencies
func (s *service) Close() {
	logInfo("service closing...", "Close")
	for _, v := range s.avgPriceManagers {
		v.close()
	}
}
