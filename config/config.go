package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/viper"
)

const (
	API_KEY         = "API_KEY"
	ORGANIZATION_ID = "ORGANIZATION_ID"
	PROXY_DOMAIN    = "PROXY_DOMAIN"
)

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		color.Red("init config error: %s", err.Error())
		os.Exit(1)
	}

	cacheDir := filepath.Join(homeDir, ".aish")

	viper.SetConfigFile("./config.config")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(cacheDir)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err := os.MkdirAll(cacheDir, os.ModePerm)
			if err != nil {
				panic(err)
			}
			filePath := filepath.Join(cacheDir, "config.json")
			file, err := os.Create(filePath)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer file.Close()
			_, err = file.WriteString("{}")
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			panic(fmt.Errorf("fatal error config file: %s", err))
			// Config file was found but another error was produced
		}
	}
}

func GetAPIKey() string {
	return viper.GetString(API_KEY)
}

func SetAPIKey(value string) error {
	viper.Set(API_KEY, value)
	err := viper.WriteConfig()
	return err
}

func GetProxyDomain() string {
	return viper.GetString(PROXY_DOMAIN)
}

func SetProxyDomain(value string) error {
	viper.Set(PROXY_DOMAIN, value)
	err := viper.WriteConfig()
	return err
}

func SetOrganizationID(value string) error {
	viper.Set(ORGANIZATION_ID, value)
	err := viper.WriteConfig()
	return err
}

func GetOrganizationID() string {
	return viper.GetString(ORGANIZATION_ID)
}
