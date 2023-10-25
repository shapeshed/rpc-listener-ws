// Package main is the entry point, containing the main function
package main

import (
	"context"
	"fmt"
	"strconv"

	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
)

const (
	// RPCEndpoint is the CometBFT (Tendermint) RPC endpoint
	// RPCEndpoint = "https://rpc.testnet.osmosis.zone:443"
	RPCEndpoint = "https://rpc.osmosis.zone:443"

	// WsEndpoint is the websocket endpoint, usually `/websocket`
	WsEndpoint = "/websocket"

	// Subscriber is an arbitrary string that can be used to manage a subscription
	Subscriber = "gobot"

	// Query is the query to subscribe to events matching the query
	//Query = "wasm-apply_funding._contract_address = 'osmo1cnj84q49sp4sd3tsacdw9p4zvyd8y46f2248ndq2edve3fqa8krs9jds9g'"
	Query = "token_swapped.module = 'gamm'"
)

func main() {
	// Create a new Tendermint RPC client
	c, err := rpchttp.New(RPCEndpoint, WsEndpoint)
	c.Start()
	if err != nil {
		panic(err)
	}

	// Create a context for the subscription
	ctx := context.Background()

	// Subscribe to the WebSocket connection
	eventCh, err := c.Subscribe(ctx, Subscriber, Query)
	if err != nil {
		panic(err)
	}

	// Create a goroutine to print events from the channel
	go func() {
		for {
			event := <-eventCh
			txHeight, err := strconv.ParseInt(event.Events["tx.height"][0], 10, 64)
			blockInfo, err := c.Block(ctx, &txHeight)
			if err != nil {
				panic(err)
			}
			fmt.Printf("time: %+v\n", blockInfo.Block.Time)
			//fmt.Printf("funding_rate: %+v\n", event.Events["wasm-apply_funding.funding_rate"])
			fmt.Printf("height: %+v\n", txHeight)
			fmt.Printf("funding_rate: %+v\n", event.Events["token_swapped.tokens_in"])

			// Insert into database here
		}
	}()

	// Keep the main goroutine running
	select {}
}
