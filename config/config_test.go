package config_test

import (
	"testing"
	"time"

	"github.com/longbridgeapp/assert"
	"github.com/longportapp/openapi-go/config"
)

var expectedConfig = &config.Config{
	HttpURL:         "http://test",
	HTTPTimeout:     12 * time.Second,
	AppKey:          "test_app_key",
	AppSecret:       "test_app_secret",
	AccessToken:     "test_access_token",
	TradeUrl:        "http://trade_test",
	QuoteUrl:        "http://quote_test",
	EnableOvernight: true,
	AuthTimeout:     12 * time.Second,
	Timeout:         12 * time.Second,
	WriteQueueSize:  12,
	ReadQueueSize:   12,
	ReadBufferSize:  12,
	MinGzipSize:     12,
	Region:          "hk",
}

func Test_withConfigKey(t *testing.T) {
	var c, err = config.New(config.WithConfigKey("appKey", "appSecret", "accessToken"))
	assert.NoError(t, err)
	assert.Equal(t, "appKey", c.AppKey)
	assert.Equal(t, "appSecret", c.AppSecret)
	assert.Equal(t, "accessToken", c.AccessToken)
}

func Test_YamlConfig(t *testing.T) {
	c, err := config.New(config.WithFilePath("./testdata/test_config.yaml"))
	assert.NoError(t, err)
	assert.Equal(t, expectedConfig, c)
}

func Test_TomlConfig(t *testing.T) {
	c, err := config.New(config.WithFilePath("./testdata/test_config.toml"))
	assert.NoError(t, err)
	assert.Equal(t, expectedConfig, c)
}
