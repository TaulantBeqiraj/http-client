package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const baseURL = "http://shadow.disconnect.ch:8002"

var (
	username = "taulant"
	password = "eCG58weaD6"
)

type UserInfo struct {
	Login   string `json:"Login"`
	Balance string `json:"Balance"`
	Email   string `json:"Email"`
}

type UserAssets struct {
	Assets []interface{} `json:"Assets"`
}

type MarketAsset struct {
	Name  string
	Price float64 `json:",string"` // string to float64
}

func get(c *http.Client, url string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Couldn't get the request: %v\n", err)
	}

	req.SetBasicAuth(username, password)
	req.Header.Add("Authorization", "Basic "+authorization())

	res, err := c.Do(req)
	if err != nil {
		log.Fatalf("Couldn't get the request %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Failed to close Body %v", err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusUnauthorized {
			log.Printf("unauthorized: %v\n", res.Header.Get("WWW-Authenticate"))
		}

		panic(fmt.Errorf("got HTTP error %v", res.StatusCode))
	}

	buf2, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Couldn't read response %v", buf2)
	}

	var ma []MarketAsset
	err = json.Unmarshal(buf2, &ma)
	if err != nil {
		log.Printf("Failed to Unmarshall data %v", err)
	}

	for i, v := range ma {
		fmt.Printf("Article no.%d\n Name: %v\n Price: %v\n\n", i, v.Name, v.Price)
	}
	return buf2
}

func authorization() string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func main() {

	client := &http.Client{}

	get(client, baseURL+"/rates")

}
