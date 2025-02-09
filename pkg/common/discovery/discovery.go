package discovery

import (
	"my-im-server/pkg/common/config"
	"my-im-server/sdk/tools/discovery"
	"time"

	"my-im-server/sdk/tools/discovery/etcd"
)

func NewDiscoveryRegister(discovery *config.Discovery, runtimeEnv string) (discovery.SvcDiscoveryRegister, error) {
	switch discovery.Enable {
	case config.ETCD:
		return etcd.NewSvcDiscoveryRegister(
			discovery.Etcd.RootDirectory,
			discovery.Etcd.Address,
			etcd.WithDialTimeout(10*time.Second),
			etcd.WithMaxCallSendMsgSize(10*1024*1024),
			etcd.WithUsernameAndPassword(discovery.Etcd.Username, discovery.Etcd.Password),
		)
	}
	return nil, nil
}
