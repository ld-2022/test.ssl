package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// 加载服务器证书和私钥
	cert, err := tls.LoadX509KeyPair("/Users/edy/Documents/GitHub/test.ssl/certs/server.crt", "/Users/edy/Documents/GitHub/test.ssl/certs/server.key")
	if err != nil {
		panic(err)
	}

	// 加载CA根证书
	caCert, err := ioutil.ReadFile("/Users/edy/Documents/GitHub/test.ssl/certs/myCA.crt")
	if err != nil {
		panic(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// 配置TLS
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    caCertPool,
		ServerName:   "router",
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}
	tlsConfig.BuildNameToCertificate()

	server := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsConfig,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s", r.TLS.PeerCertificates[0].Subject.CommonName)
		}),
	}

	fmt.Println("Starting server on https://localhost:8443")
	err = server.ListenAndServeTLS("", "")
	if err != nil {
		panic(err)
	}
}
