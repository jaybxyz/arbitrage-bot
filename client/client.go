package client

import (
	"github.com/kogisin/arbitrage-bot/client/grpc"
	"github.com/kogisin/arbitrage-bot/client/rpc"

	"github.com/rs/zerolog/log"
)

// Client is a wrapper for various clients.
type Client struct {
	RPC  *rpc.Client
	GRPC *grpc.Client
}

// NewClient creates a new Client with the given configuration.
func NewClient(rpcURL string, grpcURL string) (*Client, error) {
	log.Debug().Str("rpcURL", rpcURL).Msg("connecting rpc client...")
	rpcClient, err := rpc.NewClient(rpcURL)
	if err != nil {
		return &Client{}, err
	}

	log.Debug().Str("grpcURL", grpcURL).Msg("connecting grpc client...")
	grpcClient, err := grpc.NewClient(grpcURL)
	if err != nil {
		return &Client{}, err
	}

	return &Client{
		RPC:  rpcClient,
		GRPC: grpcClient,
	}, nil
}

// GetRPCClient returns RPC client.
func (c *Client) GetRPCClient() *rpc.Client {
	return c.RPC
}

// GetGRPCClient returns GRPC client.
func (c *Client) GetGRPCClient() *grpc.Client {
	return c.GRPC
}

// Stop defers the node stop execution to the RPC and GRPC clients.
func (c Client) Stop() error {
	err := c.RPC.Stop()
	if err != nil {
		return err
	}

	err = c.GRPC.Close()
	if err != nil {
		return err
	}
	return nil
}
