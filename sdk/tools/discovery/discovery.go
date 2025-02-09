package discovery

import (
	"context"

	"google.golang.org/grpc"
)

type Conn interface {
	GetConn(ctx context.Context, serviceName string, opts ...grpc.DialOption) (*grpc.ClientConn, error)
	AddOptions(opts ...grpc.DialOption)
	CloseConn(conn *grpc.ClientConn)
}

type SvcDiscoveryRegister interface {
	Conn
}
