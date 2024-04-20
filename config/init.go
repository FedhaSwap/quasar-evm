package config

import (
	"github.com/spf13/viper"
)

// init initialize default config params
func init() {
	// environment - could be "local", "prod", "dev"
	viper.SetDefault("env", "prod")

	viper.SetDefault("tokens", []string{
		"ADA", "ALGO", "ARB",
		"AVAX", "BNB", "BTC",
		"DOGE", "ETH", "FIL",
		"FLR", "LTC", "MATIC",
		"SOL", "USDC", "USDT",
		"XDC", "XLM", "XRP",
	})

	viper.SetDefault("ws.url", "wss://oracle.gateway.fm")
	viper.SetDefault("ws.WriteFrequencyMS", 5000)

	viper.SetDefault("web3.SendFrequencyMS", 10000)
	viper.SetDefault("web3.RpcURL", "https://sn2-stavanger-rpc.eu-north-2.gateway.fm")
	viper.SetDefault("web3.SignerPK", "")
	viper.SetDefault("web3.ChainID", 686669576)
	viper.SetDefault("web3.OracleAddress", "0x22a01749E0C916c859B584528D37e838A8a8eba7")
}
