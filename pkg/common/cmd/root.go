package cmd

import (
	"fmt"
	"my-im-server/pkg/common/config"
	"my-im-server/sdk/tools/utils/runtime"

	"github.com/spf13/cobra"
)

type RootCmd struct {
	Command     *cobra.Command
	log         config.Log
	processName string
	configPath  string
	port        int
}

type CmdOpts struct {
	loggerPrefix string
	ConfigMap    map[string]any
}

func NewRootCmd(programName string, opts ...func(*CmdOpts)) *RootCmd {
	rootCmd := &RootCmd{processName: programName}
	cmd := &cobra.Command{
		Use:  "Start IM Application",
		Long: fmt.Sprintf(`Start %s `, programName),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return rootCmd.presistentPreRunE(cmd, opts...)
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	cmd.Flags().StringP(config.FlagConf, "c", "", "config path")
	rootCmd.Command = cmd
	return rootCmd
}

func WithConfigMap(configMap map[string]any) func(*CmdOpts) {
	return func(c *CmdOpts) {
		c.ConfigMap = configMap
	}
}

func (r *RootCmd) initConfiguration(cmd *cobra.Command, cmdOpts *CmdOpts) error {
	configDir, err := r.getFlag(cmd)
	if err != nil {
		return err
	}
	runtimeEnv := runtime.PrintRuntimeEnv()

	for configName, configStruct := range cmdOpts.ConfigMap {
		config.LoadConfig(configDir, configName, configStruct, runtimeEnv)
	}
	return config.LoadConfig(configDir, config.LogConfigName, &r.log, runtimeEnv)
}

func (r *RootCmd) getFlag(cmd *cobra.Command) (string, error) {
	configDir, err := cmd.Flags().GetString(config.FlagConf)
	if err != nil {
		return "", err
	}
	r.configPath = configDir
	return configDir, nil
}

func (r *RootCmd) presistentPreRunE(cmd *cobra.Command, opts ...func(*CmdOpts)) error {
	cmdOpts := r.applyOpts(opts...)
	if err := r.initConfiguration(cmd, cmdOpts); err != nil {
		return err
	}
	return nil
}

func (r *RootCmd) applyOpts(opts ...func(*CmdOpts)) *CmdOpts {
	cmdOpts := &CmdOpts{}
	for _, opt := range opts {
		opt(cmdOpts)
	}
	return cmdOpts
}

func (r *RootCmd) Execute() error {
	return r.Command.Execute()
}
