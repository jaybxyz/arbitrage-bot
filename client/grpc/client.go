package grpc

import (
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

const (
	DefaultTLSPort     = "443"
	DefaultGRPCTimeout = 10 * time.Second // detecting conn failure issue: https://github.com/grpc/grpc-go/issues/849
)

// Client wraps GRPC client connection.
type Client struct {
	*grpc.ClientConn
}

// NewClient creates GRPC client.
func NewClient(grpcURL string) (*Client, error) {
	var grpcopts []grpc.DialOption
	var conn *grpc.ClientConn
	var err error

	urls := strings.Split(grpcURL, ":")
	if 2 < len(urls) {
		panic(fmt.Sprintf("incorrect grpc endpoint: %s", urls))
	}

	// handle TLS/SSL connection
	if urls[1] == DefaultTLSPort {
		grpcopts = []grpc.DialOption{
			grpc.WithTransportCredentials(credentials.NewTLS(nil)),
			grpc.WithTimeout(DefaultGRPCTimeout),
		}
		conn, err = grpc.Dial(grpcURL, grpcopts...)
		if err != nil {
			return &Client{}, fmt.Errorf("failed to connect GRPC client: %v", err)
		}
	} else {
		grpcopts = []grpc.DialOption{
			grpc.WithInsecure(),
			grpc.WithBlock(),
			grpc.WithTimeout(DefaultGRPCTimeout),
		}

		conn, err = grpc.Dial(grpcURL, grpcopts...)
		if err != nil {
			return &Client{}, fmt.Errorf("failed to connect GRPC client: %v", err)
		}
	}

	defer conn.Close()

	return &Client{conn}, nil
}

// IsNotFound returns not found status.
func IsNotFound(err error) bool {
	return status.Convert(err).Code() == codes.NotFound
}
