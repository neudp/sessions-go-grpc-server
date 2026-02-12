package grpc_client

import (
	"go-grpc-server/internal/integration/common"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	//"google.golang.org/grpc/credentials"
	//"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

func NewClientConnection(address string) (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                10 * time.Second,
			Timeout:             1 * time.Second,
			PermitWithoutStream: true,
		}),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(2*common.Megabyte),
			grpc.MaxCallSendMsgSize(2*common.Megabyte),
		),
	}

	// If endpoint uses TLS
	// opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	// If endpoint does not use TLS
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	return grpc.NewClient(
		address,
		opts...,
	)
}
