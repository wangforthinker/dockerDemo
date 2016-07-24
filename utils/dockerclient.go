package utils

import (
	"sync"
	"crypto/tls"
	"github.com/samalba/dockerclient"
)




const (
KEY_DOCKER_SCHEME            = "DOCKER_SCHEME"
KEY_DOCKER_ENDPOINT          = "DOCKER_ENDPOINT"
KEY_DOCKER_ENDPOINT_TLS      = "DOCKER_ENDPOINT_TLS"
KEY_DOCKER_ENDPOINT_TLS_KEY  = "DOCKER_ENDPOINT_TLS_KEYFILE"
KEY_DOCKER_ENDPOINT_TLS_CERT = "DOCKER_ENDPOINT_TLS_CERTFILE"
KEY_DOCKER_DISCOVERY         = "DOCKER_DISCOVERY"
)

var client *dockerclient.DockerClient
var once *sync.Once
var initError error

func init() {
	once = &sync.Once{}
}

func InitDockerClient(scheme string, endpoint string, tlsConfig *tls.Config) (*dockerclient.DockerClient, error) {

	once.Do(func() {
		client, initError = dockerclient.NewDockerClient(
			scheme+"://"+endpoint,
			tlsConfig)
	})

	return client, initError
}