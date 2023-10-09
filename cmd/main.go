package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	k8s_addr = os.Getenv("K8S_ADDRESS")
	jwt      = os.Getenv("K8S_JWT_TOKEN")
)

func Request(method, path string, body io.Reader) (*http.Request, error) {
	bearer := "Bearer " + jwt

	// Creating new request
	r, err := http.NewRequest(method, path, body)
	r.Header.Set("Authorization", bearer)
	r.Header.Add("Accept", "application/json")
	if err != nil {
		log.Println("Error creating request: ", err)
		return nil, err
	}
	return r, nil

}

func main() {
	podsPath := "/api/v1/pods"
	url := k8s_addr + podsPath

	req, err := Request("GET", url, bytes.NewBuffer(nil))
	if err != nil {
		log.Println("Error creating request: ", err)
	}

	// // Seding request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}

	// Closing request
	defer resp.Body.Close()

	// Reading th request body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	log.Println(string([]byte(body)))
}
