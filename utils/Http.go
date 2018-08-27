package utils

import (
	"net/url"
	"net/http"
	"crypto/tls"
	"crypto/x509"
)

type httpUtil struct {}

var Http httpUtil

func (this httpUtil) NewSimpleClient() (*http.Client, error) {
	return this.NewClient("", false, nil)
}

func (this httpUtil) NewClient(proxy string, enableTLS bool, certificates []byte) (*http.Client, error) {
	// Set proxy function if there is a proxy configured.
	var proxyFunc func(*http.Request) (*url.URL, error)
	if proxy != "" {
		proxyURL, err := url.Parse(proxy)
		if err != nil {
			return nil, err
		}
		proxyFunc = http.ProxyURL(proxyURL)
	}

	// Configure TLS if needed.
	var tlsConfig *tls.Config
	if enableTLS {
		if len(certificates) > 0 {
			pool := x509.NewCertPool()
			pool.AppendCertsFromPEM(certificates)
			tlsConfig = &tls.Config{
				RootCAs: pool,
			}
		}
	}

	client := http.Client{
		Transport: &http.Transport{
			Proxy:           proxyFunc,
			TLSClientConfig: tlsConfig,
		},
	}

	return &client, nil
}