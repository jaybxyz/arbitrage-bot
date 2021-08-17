package rpc

import (
	"context"
	"fmt"

	rpcclient "github.com/tendermint/tendermint/rpc/client"
	rpc "github.com/tendermint/tendermint/rpc/client/http"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

var (
	DefaultRPCTimeout = uint(5)
)

// Client wraps RPC client connection.
type Client struct {
	rpcclient.Client
}

// NewClient creates RPC client.
func NewClient(rpcURL string) (*Client, error) {
	rpcClient, err := rpc.NewWithTimeout(rpcURL, "/websocket", DefaultRPCTimeout)
	if err != nil {
		return &Client{}, fmt.Errorf("failed to connect RPC client: %s", err)
	}

	return &Client{rpcClient}, nil
}

// GetNetworkChainID returns network chain id.
func (c *Client) GetNetworkChainID(ctx context.Context) (string, error) {
	status, err := c.Status(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get status: %v", err)
	}

	return status.NodeInfo.Network, nil
}

// GetStatus returns the status of the blockchain network.
func (c *Client) GetStatus(ctx context.Context) (*tmctypes.ResultStatus, error) {
	return c.Status(ctx)
}
