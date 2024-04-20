package config

// Scheme represents the application configuration scheme.
type Scheme struct {
	// Env is the application environment.
	Env string

	// Tokens that will be sent to the oracle contract
	Tokens []string
	// WS is a web socket configs
	WS *WS
	// WEB3 is a blockchain configs
	WEB3 *WEB3
}

// WEB3 is a blockchain configs
type WEB3 struct {
	// OracleAddress smart-contract address for the rollup oracle
	OracleAddress string
	// RpcURL url for rpc-provider
	RpcURL string
	// ChainID rollup chain ID
	ChainID int
	// SignerPK is a wallet private key. Should be the same as the oracle owner. Shall never be hardcoded
	SignerPK string
	// SendFrequencyMS is a send to oracle frequency in milliseconds
	SendFrequencyMS int
}

// WS is a web socket configs
type WS struct {
	// URL is a demon url address
	URL string
	// WriteFrequencyMS is a write from demon frequency in milliseconds
	WriteFrequencyMS int
}
