package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestConfig struct {
	Server struct {
		Port    int    `mapstructure:"port"`
		Address string `mapstructure:"address"`
	}
}

func TestLoadConfig(t *testing.T) {
	// 创建临时测试配置文件
	tmpDir, err := os.Getwd()
	assert.NoError(t, err)
	t.Log("-----")
	t.Log(tmpDir)
	configPath := filepath.Join(tmpDir, "test_config.yaml")
	configContent := []byte(`
server:
  port: 8080
  address: localhost
`)
	err = os.WriteFile(configPath, configContent, 0644)
	assert.NoError(t, err)

	// 测试正常配置文件加载
	var config TestConfig
	err = loadConfig(configPath, &config)
	assert.NoError(t, err)
	t.Log("-------")
	t.Log(config)
	assert.Equal(t, 8080, config.Server.Port)
	assert.Equal(t, "localhost", config.Server.Address)

	// 测试环境变量覆盖
	os.Setenv("SERVER_PORT", "9090")
	os.Setenv("SERVER_ADDRESS", "0.0.0.0")

	var configWithEnv TestConfig
	err = loadConfig(configPath, &configWithEnv)
	assert.NoError(t, err)
	t.Log("----")
	t.Log(configWithEnv)
	assert.Equal(t, 9090, configWithEnv.Server.Port)
	assert.Equal(t, "0.0.0.0", configWithEnv.Server.Address)

	// 测试配置文件不存在的情况
	err = loadConfig("/non/existent/path", &config)
	assert.Error(t, err)
}
