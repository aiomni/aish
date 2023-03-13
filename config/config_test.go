package config_test

import (
	"testing"

	"github.com/aiomni/aish/config"
)

func TestSetAPIKey(t *testing.T) {
	key := "这是一个测试Key"
	config.SetAPIKey(key)

	configKey := config.GetAPIKey()

	if configKey != key {
		t.Errorf("set api key error")
	}

}
