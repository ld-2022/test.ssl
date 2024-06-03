package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// 加载客户端证书和私钥
	cert, err := tls.LoadX509KeyPair("/Users/edy/Documents/GitHub/test.ssl/certs/router.crt", "/Users/edy/Documents/GitHub/test.ssl/certs/router.key")
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
		RootCAs:      caCertPool,
		ServerName:   "server",

		InsecureSkipVerify: false,
	}

	// 创建HTTP客户端
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	// 发起请求
	resp, err := client.Get("https://localhost:8443")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response from server: %s\n", body)
}
