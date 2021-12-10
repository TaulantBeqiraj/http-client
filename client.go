package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const baseURL = "http://shadow.disconnect.ch:8002"

type MarketAsset struct {
	Name  string
	Price float64 `json:",string"` // string to float64
}

func get(c *http.Client, url string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	res, err := c.Do(req)
	if err != nil {
		log.Fatalf("Couldn't get the request %v", err)
	}
	defer res.Body.Close()

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

func main() {

	client := &http.Client{}

	get(client, baseURL+"/rates")

}
