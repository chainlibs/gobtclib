package utils

import (
	"net/url"
	"net/http"
	"crypto/tls"
	"crypto/x509"
)

type httpUtil struct {}

var Http httpUtil

/*
Description:
NewClientSimple returns a new instance of http.client with no arguments
 * Author: architect.bian
 * Date: 2018/10/05 18:08
 */
func (h httpUtil) NewClientSimple() (*http.Client, error) {
	return h.NewClient("", false, nil)
}

/*
Description:
NewClient returns a new instance of http.client with proxy, enableTLS, certificates
 * Author: architect.bian
 * Date: 2018/10/05 18:11
 */
func (h httpUtil) NewClient(proxy string, enableTLS bool, certificates []byte) (*http.Client, error) {
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