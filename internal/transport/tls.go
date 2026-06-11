package transport

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
)

// TLSOptions configures optional TLS behaviour for the HTTP transport.
type TLSOptions struct {
	InsecureSkipVerify bool
	CACertPEM          string
	CACertFile         string
}

func buildTLSConfig(opts TLSOptions) (*tls.Config, error) {
	if !opts.InsecureSkipVerify && opts.CACertPEM == "" && opts.CACertFile == "" {
		return nil, nil
	}

	tlsCfg := &tls.Config{MinVersion: tls.VersionTLS12}

	if opts.InsecureSkipVerify {
		tlsCfg.InsecureSkipVerify = true
	}

	caPEM := opts.CACertPEM
	if caPEM == "" && opts.CACertFile != "" {
		raw, err := os.ReadFile(opts.CACertFile)
		if err != nil {
			return nil, fmt.Errorf("transport: read CACertFile: %w", err)
		}
		caPEM = string(raw)
	}

	if caPEM != "" {
		pool := x509.NewCertPool()
		if !pool.AppendCertsFromPEM([]byte(caPEM)) {
			return nil, fmt.Errorf("transport: CACertPEM/CACertFile contains no valid CA certificates")
		}
		tlsCfg.RootCAs = pool
		if opts.InsecureSkipVerify {
			tlsCfg.InsecureSkipVerify = false
		}
	}

	return tlsCfg, nil
}
