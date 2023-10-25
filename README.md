# Tendermint RPC WebSocket Listener

POC for an service that uses the `subscribe` WebSocket endpoint from the
[CometBFT (Tendermint) RPC][1]. This receives events that match the query that
may be used when they occur.

The process uses a query on `osmosis-1` for swap events and logs the `token_in`
value, but this is easily configurable for other use cases.

## Installation

- `go mod download`
- `go run main.go`
- `go build`

## Usage

```sh
./rpc-listener-ws
```

The process writes data to the console.

[1]: https://docs.cometbft.com/main/rpc/
