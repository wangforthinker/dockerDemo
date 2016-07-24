package cli

import (
	"github.com/codegangsta/cli"
	"github.com/Sirupsen/logrus"
	"crypto/tls"
	"errors"
	"github.com/wangforthinker/dockerDemo/utils"
	"net/http"
	"github.com/wangforthinker/dockerDemo/handler"
	"net"
)

var (
	clientCipherSuits = []uint16 {
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
	}

	clientDefaultTlsConfig = tls.Config{
		MinVersion: tls.VersionTLS12,
		CipherSuites: clientCipherSuits,
		InsecureSkipVerify: true,
	}
)

func getTlsConfig (c *cli.Context) (*tls.Config, error){
	if !c.Bool("tls") {
		return nil,nil
	}

	keyFile := c.String("tlskey")
	certFile := c.String("tlscert")

	tlsConfig := clientDefaultTlsConfig

	if certFile != "" && keyFile != "" {
		tlsCert,err := tls.LoadX509KeyPair(certFile, keyFile)
		if err != nil {
			return nil, errors.New("Can not load x509 key pair:"+ err.Error())
		}

		tlsConfig.Certificates = []tls.Certificate{tlsCert}
	}

	return &tlsConfig,nil

}

func startCommand(c *cli.Context) {
	dockerScheme := c.String("Scheme")
	dockerEndpoint := c.String("Endpoint")

	logrus.Info(dockerEndpoint)
	logrus.Info(dockerScheme)

	tlsConfig,err :=getTlsConfig(c)

	if err != nil {
		logrus.Fatal(err)
	}

	client,err := utils.InitDockerClient(c.String("Scheme"), c.String("Endpoint"), tlsConfig)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info(client)

	containers,err := client.ListContainers(true, true, "")
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info(containers)


	server := &http.Server{
		Handler: handler.NewHandler(dockerScheme, dockerEndpoint, tlsConfig),
	}

	l,err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil{
		logrus.Fatal(err)
	}

	if err = server.Serve(l); err != nil{
		logrus.Fatal(err)
	}
}