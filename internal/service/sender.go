package service

import (
	"fmt"
)

// runSender is used to send token prices. Will not send price if it is 0
func (s *coinAVGPriceManager) runSender() {
	for {
		select {
		case <-s.stopSender:
			logInfo("stop...", "Sender")
			return
		case <-s.ticker.C:
			logInfo(fmt.Sprintf("pushing prices"), "Sender")

			prices := s.parsePrices()

			for i, t := range s.tokens {
				if prices[i].Uint64() == 0 {
					logInfo("price is 0, push stopped", "Sender")
					continue
				}

				if err := s.oracle.PushPrice(t, prices[i]); err != nil {
					logErr(fmt.Sprintf("err push price: %s", err.Error()), "Sender")
				}
			}
		}
	}
}
