# Quasar-EVM

Quasar EVM is a backend part of the Quasar project (see
[Quasar-Contract repository](https://github.com/gateway-fm/quasar-contract)). Quasar-EVM service is used to listen
to the token average prices and write them to the oracle smart-contract. It provides solid stabilized write-send flow
where write and send actions are independent of each other. Service will resubscribe if needed and keep the same send
frequency even if the prices are outdated.

## Table of content

- [Getting started](#getting-started)
- [Project structure](#project-structure)
- [Env variables](#env-variables)
- [Deploy and Run](#deploy-and-run)

## Getting started

To run the service firstly you need to deploy the [Quasar-Contract](https://github.com/gateway-fm/quasar-contract) with
the same wallet address that you will use in this project deploy. After fill the [env variables](#env-variables) and
run the service local or deploy it remotely.

## Project structure

### Infrastructure

#### Config and env variables

Config scheme can be found in the [scheme.go](./config/scheme.go) file. There is 3 main config blocks - the `WEB3` for
the blockchain configs, the `WS` for the web socket connection and demon work configs and `TOKENS` for the list of
token symbols you want to write and send to the smart-contract. Read more in the
[Env variables readme block](#env-variables)

#### Makefile

Shortcuts for all main service commands such as `run`, `test`, `build` etc. are implement in the [Malefile](./Makefile).
To run commands use make syntax, for example:

```shell
make run
```

#### Dockerfile

[Docker](./Dockerfile) file can be found in the project route. When project deployed on the remote environment the
docker
file can be used in the CI/CD flows.

### External packages

#### Smart-contract

Smart-contract package implementation can be found in the [contractsAPI](./contractsAPI) directory. The
[oracle](./contractsAPI/oracle) directory contains smart-contract API implementation generated with the `abigen` CLI
util.
The ABI for the code generation is stored in the [abis](./abis) directory.

[ContractAPI.go](./contractsAPI/contractAPI.go) implements only needed methods to communicate with the oracle
smart-contract. See the `IOracleAPI` interface

#### WS Client

[WS Client](./pkg/wsClient) package is implementing methods to communicate with the currency price demon using WS
connection. Example of the wsClient usage can be fond in the [example directory](./pkg/wsClient/example). Also see the
[IWSClient](./pkg/wsClient/client.go) interface for more information.

### Internal package

[Service](./internal/service) is a main internal package. It is managing the write and send flow independently of
each other using the oracle and ws client packages and provides reconnections if needed.

## Env variables

For default values see [config/init.go](./config/init.go)

| Name                 | Example / Default value                      | Description                                                                                    | Default |
|----------------------|----------------------------------------------|------------------------------------------------------------------------------------------------|---------|
| WEB3_ORACLEADDRESS   | 0x078413b8b5a2614081813619c594fE84Ef839535   | Deployed oracle address                                                                        | ❌       |
| WEB3_RPCURL          | https://jesterfair-rpc.eu-north-2.gateway.fm | Rollup RPC-provider URL                                                                        | ❌       |
| WEB3_CHAINID         | 1413045109                                   | Rollup chain ID                                                                                | ❌       |
| WEB3_SIGNERPK        |                                              | Deployed oracle owner wallet private key                                                       | ❌       |
| TOKENS               | "SOL", "USDC", "USDT", ...                   | Tokens data for which will be send to the oracle contract                                      | ✅       |
| WS_URL               | wss://oracle.gateway.fm                      | Demon ws URL                                                                                   | ✅       |
| WS_WRITEFREQUENCYMS  | 5000                                         | Frequancy for write flow in milliseconds. Better to be less or equal to the send flow requancy | ✅       |
| WEB3_SENDFREQUENCYMS | 10000                                        | Frequancy for send flow in milliseconds. Better to be more or equal to the write flow requancy | ✅       |

## Deploy and Run

Before you deploy this service, you need to create a rollup and deploy the 
[Quasar-Contract](https://github.com/gateway-fm/quasar-contract) to the rollup with the same wallet, that you will be
using in this service as a `WEB3_SIGNERPK` variable.

To deploy the service you need to set the env variables, after that you can use the [Dockerfile](./Dockerfile), no 
additional service are needed to be deployed. Or you can run service locally using the `make run` command. 

You can also build the binary using the `make build` command and run the binary using `quasar-evm serve` command.
