package api

import (
	"context"
	"my-im-server/sdk/tools/discovery"

	"github.com/gin-gonic/gin"
)

func newGinRouter(ctx context.Context, client discovery.SvcDiscoveryRegister, config *Config) (*gin.Engine, error) {
	_, err := client.GetConn(ctx, config.Discovery.RpcService.Auth)
	if err != nil {
		return nil, err
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	return r, nil

}
