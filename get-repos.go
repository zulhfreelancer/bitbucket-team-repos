package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func getRepos(token string, pageNum string) repositoriesResponse {
	uuid := viper.GetString("TEAM_UUID")
	url := fmt.Sprintf("https://api.bitbucket.org/2.0/repositories/%v?page=%v", uuid, pageNum)
	req, _ := http.NewRequest("GET", url, nil)

	tokenStr := fmt.Sprintf("Bearer %v", token)
	req.Header.Add("Authorization", tokenStr)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", "api.bitbucket.org")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	var rr repositoriesResponse
	err := json.NewDecoder(res.Body).Decode(&rr)
	if err != nil {
		log.Fatalf("Error decoding API response: %v", err)
	}

	return rr
}
