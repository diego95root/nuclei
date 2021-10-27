package engine

import (
	"crypto/tls"
	"github.com/projectdiscovery/nuclei/v2/pkg/protocols/utils"
	"net/http"
	"net/url"
	"time"

	"github.com/projectdiscovery/nuclei/v2/pkg/protocols/common/protocolstate"
	"github.com/projectdiscovery/nuclei/v2/pkg/types"
)

// newhttpClient creates a new http client for headless communication with a timeout
func newhttpClient(options *types.Options) *http.Client {
	dialer := protocolstate.Dialer

	// Set the base TLS configuration definition
	tlsConfig := &tls.Config{
		Renegotiation:      tls.RenegotiateOnceAsClient,
		InsecureSkipVerify: true,
	}

	// Add the client certificate authentication to the request if it's configured
	tlsConfig = utils.AddConfiguredClientCertToRequest(tlsConfig, options)

	transport := &http.Transport{
		DialContext:         dialer.Dial,
		MaxIdleConns:        500,
		MaxIdleConnsPerHost: 500,
		MaxConnsPerHost:     500,
		TLSClientConfig:     tlsConfig,
	}

	if options.ProxyURL != "" {
		if proxyURL, err := url.Parse(options.ProxyURL); err == nil {
			transport.Proxy = http.ProxyURL(proxyURL)
		}
	}

	return &http.Client{Transport: transport, Timeout: time.Duration(options.Timeout*3) * time.Second}
}
