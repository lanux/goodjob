package utils

import (
	"crypto/tls"
	"net/http"
)

var HttpClient *http.Client

func init() {
	HttpClient = newHttpClient()
}

func newHttpClient() *http.Client {
	transport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	return &http.Client{Transport: transport}
}
