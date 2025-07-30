package main

import (
	"os"

	awsmskiamprovider "github.com/grepplabs/kafka-proxy/pkg/libs/aws-msk-iam-provider"
	"github.com/grepplabs/kafka-proxy/plugin/token-provider/shared"
	"github.com/hashicorp/go-plugin"
	"github.com/sirupsen/logrus"
)

func main() {
	tokenProvider, err := new(awsmskiamprovider.Factory).New(os.Args[1:])

	if err != nil {
		logrus.Errorf("cannot initialize aws-msk-iam-token provider: %v", err)
		os.Exit(1)
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins: map[string]plugin.Plugin{
			"tokenProvider": &shared.TokenProviderPlugin{Impl: tokenProvider},
		},
		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
