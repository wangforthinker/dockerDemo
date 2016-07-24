package handler

import (
	"net/http"
	"crypto/tls"
	"github.com/gorilla/mux"
	"github.com/Sirupsen/logrus"
)

type Context struct {
	scheme string
	endpoint string
	tls *tls.Config
}

func NewHandler(scheme string, endpoint string, tlsConfig *tls.Config) http.Handler {
	context := &Context{
		scheme: scheme,
		endpoint: endpoint,
		tls : tlsConfig,
	}

	logrus.Info(context)

	r := mux.NewRouter()


	return r
}