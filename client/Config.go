package client

const (
	// requestsChanBufferSize is the number of elements the HTTP POST send
	// channel can queue before blocking.
	requestsChanBufferSize = 100
)

/*
Description:
ConnConfig describes the connection configuration parameters for the client.
 * Author: architect.bian
 * Date: 2018/08/23 15:32
 */
type Config struct {
	// Host is the IP address and port of the RPC server you want to connect to.
	Host string

	// User is the username to use to authenticate to the RPC server.
	User string

	// Pass is the passphrase to use to authenticate to the RPC server.
	Pass string

	// EnableTLS specifies whether transport layer security should be disabled.
	// It is recommended to always use TLS if the RPC server
	// supports it as otherwise your username and password is sent across the wire in cleartext.
	EnableTLS bool

	// Certificates are the bytes for a PEM-encoded certificate chain used
	// for the TLS connection.  It only has effect if the EnableTLS parameter is true.
	Certificates []byte

	// Proxy specifies to connect through a SOCKS 5 proxy server.  It may
	// be an empty string if a proxy is not required.
	Proxy string

	// ProxyUser is an optional username to use for the proxy server if it
	// requires authentication.  It has no effect if the Proxy parameter is not set.
	ProxyUser string

	// ProxyPass is an optional password to use for the proxy server if it
	// requires authentication.  It has no effect if the Proxy parameter is not set.
	ProxyPass string

	// DisableAutoReconnect specifies the client should not automatically
	// try to reconnect to the server when it has been disconnected.
	DisableAutoReconnect bool

	// HTTPPostMode instructs the client to run using multiple independent
	// connections issuing HTTP POST requests instead of using the default
	// of websockets.  Websockets are generally preferred as some of the
	// features of the client such notifications only work with websockets,
	// however, not all servers support the websocket extensions, so this
	// flag can be set to true to use basic HTTP POST requests instead.
	//HTTPPostMode bool
}