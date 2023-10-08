package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)
var (
	k8s_addr = os.Getenv("K8S_ADDRESS")
)
// Authenticate K8S
func httpsClient() *http.Client {
	// load tls certificates
	clientTLSCert, err := tls.LoadX509KeyPair(os.Getenv("K8S_CERT"), os.Getenv("K8S_KEY_CERT"))
	if err != nil {
		log.Fatalf("Error loading certificate and key file: %v", err)
		return nil
	}
	// Configure the client to trust TLS server certs issued by a CA.
	certPool, err := x509.SystemCertPool()
	if err != nil {
		panic(err)
	}
	if caCertPEM, err := os.ReadFile(os.Getenv("K8S_CA_CERT")); err != nil {
		panic(err)
	} else if ok := certPool.AppendCertsFromPEM(caCertPEM); !ok {
		panic("invalid cert in CA PEM")
	}
	tlsConfig := &tls.Config{
		RootCAs:      certPool,
		Certificates: []tls.Certificate{clientTLSCert},
	}
	tr := &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	client := &http.Client{Transport: tr}

	return client
}

func main() {
	client := httpsClient()
	resp, err := client.Get(k8s_addr+"/version")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	msg, _ := io.ReadAll(resp.Body)
	fmt.Println(string(msg))
}
