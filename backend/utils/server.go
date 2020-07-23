package utils

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"
)

func CreateCustomServer(cert string, key string, addr string) *http.Server {
	// Load the certificates
	certificate, err := tls.LoadX509KeyPair(cert, key)
	if err != nil {
		log.Fatal(err)
	}

	tlsConfig := &tls.Config{}
	tlsConfig.Certificates = make([]tls.Certificate, 1)
	tlsConfig.Certificates[0] = certificate

	// Create a custom server with Read and Write timeouts
	server := &http.Server{
		Addr:         addr,
		TLSConfig:    tlsConfig,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	return server
}
