package cmd

import (
	"context"
	"my-im-server/internal/api"
	"my-im-server/pkg/common/config"
	"my-im-server/sdk/tools/system/program"
	"my-im-server/version"

	"github.com/spf13/cobra"
)

type ApiCmd struct {
	rootCmd   *RootCmd
	ctx       context.Context
	configMap map[string]any
	config    *api.Config
}

func NewApiCmd() *ApiCmd {
	apiConfig := &api.Config{
		AllConfig: &config.AllConfig{},
	}
	ret := &ApiCmd{
		config: apiConfig,
	}
	ret.configMap = map[string]any{
		config.DiscoveryConfigFileName:  &apiConfig.Discovery,
		config.LogConfigName:            &apiConfig.Log,
		config.OpenIMRPCAuthCfgFileName: &apiConfig.Auth,
	}
	ret.rootCmd = NewRootCmd(program.GetProgName(), WithConfigMap(ret.configMap))
	ret.ctx = context.WithValue(context.Background(), "version", version.Version)
	ret.rootCmd.Command.RunE = func(cmd *cobra.Command, args []string) error {
		apiConfig.ConfigPath = ret.rootCmd.configPath
		return ret.runE()
	}
	return ret
}

func (a *ApiCmd) Exec() error {
	return a.rootCmd.Execute()
}

func (a *ApiCmd) runE() error {
	return api.Start(a.ctx, a.config)
}
