package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

func getAPIToken() string {

	url := "https://bitbucket.org/site/oauth2/access_token"
	payload := strings.NewReader("grant_type=client_credentials")
	req, _ := http.NewRequest("POST", url, payload)

	bbApp := fmt.Sprintf("%v:%v", viper.GetString("BITBUCKET_APP_ID"), viper.GetString("BITBUCKET_APP_SECRET"))
	bbAppEncoded := b64.StdEncoding.EncodeToString([]byte(bbApp))
	basicAuthHeader := fmt.Sprintf("Basic %v", bbAppEncoded)

	req.Header.Add("Authorization", basicAuthHeader)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", "bitbucket.org")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", "29")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	var r apiTokenResponse
	err := json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		log.Fatalf("Error decoding API response: %v", err)
	}

	t := r.AccessToken
	// fmt.Println(t)
	return t

}
