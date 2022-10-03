package config

import (
	client "github.com/ory/kratos-client-go"
)

func getKratosClientConfig() *client.Configuration {
	configuration := client.NewConfiguration()
	configuration.Servers = []client.ServerConfiguration{
		{
			URL: NymeriaConfig.URL.KratosURL,
		},
	}

	return configuration
}

var (
	KratosClientConfig = getKratosClientConfig()
)