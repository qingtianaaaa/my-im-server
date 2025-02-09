package etcd

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	gresolver "google.golang.org/grpc/resolver"
)

type SvcDiscoveryRegisterImpl struct {
	clinet        *clientv3.Client
	resolver      gresolver.Builder
	rootDirectory string
	connMap       map[string][]*addrConn
	dialOptions   []grpc.DialOption
	mu            sync.RWMutex
}

type addrConn struct {
	conn        *grpc.ClientConn
	addr        string
	isConnected bool
}

type Option func(*clientv3.Config)

func (s *SvcDiscoveryRegisterImpl) GetConn(ctx context.Context, serviceName string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	target := fmt.Sprintf("etcd:///%s/%s", s.rootDirectory, serviceName)
	dialOpts := append(append(s.dialOptions, opts...), grpc.WithResolvers(s.resolver))
	return grpc.DialContext(ctx, target, dialOpts...)
}

func (s *SvcDiscoveryRegisterImpl) AddOptions(opts ...grpc.DialOption) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.dialOptions = append(s.dialOptions, opts...)
}

func (s *SvcDiscoveryRegisterImpl) CloseConn(conn *grpc.ClientConn) {
	conn.Close()
}

func NewSvcDiscoveryRegister(rootDirectory string, endPoints []string, opts ...Option) (*SvcDiscoveryRegisterImpl, error) {
	cfg := clientv3.Config{
		Endpoints:           endPoints,
		DialTimeout:         5 * time.Second,
		PermitWithoutStream: true,
		Logger:              createNoOpLogger(),
		MaxCallSendMsgSize:  10 * 1024 * 1024,
	}

	for _, opt := range opts {
		opt(&cfg)
	}

	client, err := clientv3.New(cfg)
	if err != nil {
		return nil, err
	}
	r, err := resolver.NewBuilder(client)
	if err != nil {
		return nil, err
	}
	s := &SvcDiscoveryRegisterImpl{
		clinet:        client,
		resolver:      r,
		rootDirectory: rootDirectory,
		connMap:       make(map[string][]*addrConn),
	}
	return s, nil
}

func createNoOpLogger() *zap.Logger {
	noOpWriter := zapcore.AddSync(io.Discard)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		noOpWriter,
		zap.InfoLevel,
	)
	return zap.New(core)
}

func WithDialTimeout(time time.Duration) Option {
	return func(cfg *clientv3.Config) {
		cfg.DialTimeout = time
	}
}

func WithMaxCallSendMsgSize(size int) Option {
	return func(cfg *clientv3.Config) {
		cfg.MaxCallSendMsgSize = size
	}
}

func WithUsernameAndPassword(username, password string) Option {
	return func(cfg *clientv3.Config) {
		cfg.Username = username
		cfg.Password = password
	}
}
