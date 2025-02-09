package config

import (
	"my-im-server/sdk/tools/errs"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func LoadConfig(configDir, configName string, configStruct any, runtimeEnv string) error {
	return loadConfig(filepath.Join(configDir, configName), configStruct)
}

func loadConfig(configPath string, configStruct any) error {
	v := viper.New()
	v.SetConfigFile(configPath)
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := v.ReadInConfig(); err != nil {
		return errs.WrapMsg(err, "failed to read config file", "path", configPath)
	}

	if err := v.Unmarshal(configStruct); err != nil {
		return errs.WrapMsg(err, "failed to unmarshal config", "path", configPath)
	}
	return nil
}
