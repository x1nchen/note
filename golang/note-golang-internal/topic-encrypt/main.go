package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

func main() {
	registry := struct {
		Username      string `json:"username"`
		Password      string `json:"password"`
		Email         string `json:"email"`
		Serveraddress string `json:"serveraddress"`
	}{
		Username:      "gitlab-ci",
		Password:      "y4zYXNW99gXh",
		Email:         "",
		Serveraddress: "dockerhub.followme-internal.com",
	}

	data, _ := json.Marshal(registry)
	registryAuthToken := base64.StdEncoding.EncodeToString(data)
	fmt.Println(registryAuthToken)
}
