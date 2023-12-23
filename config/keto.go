package config

import (
	client "github.com/ory/client-go"
)

func getKetoClientConfig() (*client.Configuration, *client.Configuration) {
	readConfiguration := client.NewConfiguration()
	readConfiguration.Servers = []client.ServerConfiguration{
		{
			URL: NymeriaConfig.URL.KetoReadURL,
		},
	}

	writeConfiguration := client.NewConfiguration()
	writeConfiguration.Servers = []client.ServerConfiguration{
		{
			URL: NymeriaConfig.URL.KetoWriteURL,
		},
	}

	return readConfiguration, writeConfiguration
}

var (
	KetoReadConfig, KetoWriteConfig = getKetoClientConfig()
	KetoReadURL                     = NymeriaConfig.URL.KetoReadURL
	KetoWriteURL                    = NymeriaConfig.URL.KetoWriteURL
)
