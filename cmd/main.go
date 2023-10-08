package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)
var (
	k8s_addr = os.Getenv("K8S_ADDRESS")
	bearer = os.Getenv("K8S_JWT_TOKEN")
)

func main() {
	resp, err := http.Get(k8s_addr+"/apis/apps/v1/deployments")
	resp.Header.Set("Authorization",bearer)
	resp.Header.Add("Accept", "application/json")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	msg, _ := io.ReadAll(resp.Body)
	fmt.Println(string(msg))
}
