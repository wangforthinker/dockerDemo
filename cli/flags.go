package cli

import (
	"github.com/codegangsta/cli"
	"github.com/wangforthinker/dockerDemo/utils"
)

var (
	flDockerEndpoint = cli.StringFlag{
		Name:   "Endpoint",
		Value:  "100.69.196.212:2376",
		EnvVar: utils.KEY_DOCKER_ENDPOINT,
		Usage:  "endpoint",
	}

	flDockerScheme = cli.StringFlag{
		Name:   "Scheme",
		Value:  "http",
		EnvVar: utils.KEY_DOCKER_SCHEME,
		Usage:  "scheme",
	}

	flDockerTls = cli.BoolFlag{
		Name:   "tls",
		EnvVar: utils.KEY_DOCKER_ENDPOINT_TLS,
		Usage:  "use TLS to connect to swarm/docker",
	}

	flDockerTlsKeyFile = cli.StringFlag{
		Name:   "tlskey",
		EnvVar: utils.KEY_DOCKER_ENDPOINT_TLS_KEY,
		Usage:  "path to TLS key file",
	}

	flDockerTlsCertFile = cli.StringFlag{
		Name:   "tlscert",
		EnvVar: utils.KEY_DOCKER_ENDPOINT_TLS_CERT,
		Usage:  "path to TLS cert file",
	}

)
