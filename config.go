package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/richardheath/zeal-server/data"
	"github.com/richardheath/zeal-server/storage"

	"github.com/richardheath/zeal-server/auth"
)

type zealInstance struct {
	auth    auth.Provider
	data    data.Provider
	storage storage.Provider
}

type zealConfig struct {
	Auth    providerConfig `json:"auth"`
	Data    providerConfig `json:"data"`
	Storage providerConfig `json:"storage"`
}

type providerConfig struct {
	Provider string                 `json:"provider"`
	Settings map[string]interface{} `json:"settings"`
}

func (instance *zealInstance) initialize(config zealConfig) error {
	var err error
	instance.auth, err = auth.InitializeProvider(config.Auth.Provider, config.Auth.Settings)
	if err != nil {
		return err
	}

	instance.data, err = data.InitializeProvider(config.Data.Provider, config.Data.Settings)
	if err != nil {
		return err
	}

	instance.storage, err = storage.InitializeProvider(config.Storage.Provider, config.Storage.Settings)
	if err != nil {
		return err
	}

	return nil
}

func loadConfigFile(configFile string) (zealConfig, error) {
	config := zealConfig{}
	rawConfig, err := ioutil.ReadFile(configFile)
	if err != nil {
		return zealConfig{}, err
	}

	err = json.Unmarshal(rawConfig, &config)
	if err != nil {
		return zealConfig{}, err
	}

	return config, nil
}
