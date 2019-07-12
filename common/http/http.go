package http

import (
	"crypto/tls"
	"net/http"
)

var Client *http.Client

func init() {
	Client = NewHttpClient()
}

func NewHttpClient() *http.Client {
	transport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	return &http.Client{Transport: transport}
}
