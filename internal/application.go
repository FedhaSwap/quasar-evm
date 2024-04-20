package internal

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	version "github.com/misnaged/annales/versioner"

	"quasar-evm/config"
	"quasar-evm/contractsAPI"
	"quasar-evm/internal/service"
	"quasar-evm/pkg/wsClient"
)

// App is main microservice application instance that
// have all necessary dependencies inside structure
type App struct {
	// application configuration
	config *config.Scheme

	version *version.Version

	rpc          *ethclient.Client
	contractsAPI contractsAPI.IOracleAPI
	ws           wsClient.IWSClient
	service      service.IService
}

// NewApplication create new App instance
func NewApplication() (app *App, err error) {
	ver, err := version.NewVersion()
	if err != nil {
		return nil, fmt.Errorf("init app version: %w", err)
	}

	return &App{
		config:  &config.Scheme{},
		version: ver,
	}, nil
}

// Init initialize application and all necessary instances
func (app *App) Init() (err error) {
	if app.rpc, err = ethclient.Dial(app.config.WEB3.RpcURL); err != nil {
		return fmt.Errorf("err dial rpc: %w", err)
	}

	if app.contractsAPI, err = contractsAPI.NewOracleAPI(app.rpc, app.config.WEB3); err != nil {
		return fmt.Errorf("err init oracle api: %w", err)
	}

	sendFrequency, err := time.ParseDuration(fmt.Sprintf("%vms", app.config.WEB3.SendFrequencyMS))
	if err != nil {
		return fmt.Errorf("err parse send freq ms duration: %w", err)
	}

	app.ws = wsClient.NewClient(app.config.WS)
	app.service = service.NewService(app.ws, app.contractsAPI, sendFrequency)

	return nil
}

// Serve start serving Application service
func (app *App) Serve() error {
	app.service.SendCoinAveragePrice(app.config.Tokens)

	// Gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-quit

	app.Stop()

	return nil
}

// Stop shutdown the application
func (app *App) Stop() {
	if app.service != nil {
		app.service.Close()
	}

	if app.ws != nil {
		app.ws.Close()
	}

	if app.contractsAPI != nil {
		app.contractsAPI.Close()
	}

	if app.rpc != nil {
		app.rpc.Close()
	}
}

// Config return App config Scheme
func (app *App) Config() *config.Scheme {
	return app.config
}

// Version return application current version
func (app *App) Version() string {
	return app.version.String()
}

// CreateAddr is create address string from host and port
func CreateAddr(host string, port int) string {
	return fmt.Sprintf("%s:%v", host, port)
}
