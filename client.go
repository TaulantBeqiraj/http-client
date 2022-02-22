package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const baseURL = "http:///////////"

var (
	username = "taulant"
	password = "eCG58weaD6"
)

type MarketAsset struct {
	Name  string
	Price float64 `json:",string"` // string to float64
}

type article struct {
	Asset  string  `json:"asset"`
	Amount float64 `json:"amount"`
}

var (
	bw = article{Asset: "black_wool", Amount: 1}
	ot = article{Asset: "old_tires", Amount: 1}
	oo = article{Asset: "olive_oil", Amount: 1}
	tp = article{Asset: "toothpaste", Amount: 1}
	ww = article{Asset: "white_wool", Amount: 1}
)

type UserInfo struct {
	Login   string `json:"Login"`
	Balance string `json:"Balance"`
	Email   string `json:"Email"`
}

type UserAssets struct {
	Assets []interface{} `json:"Assets"`
}

func authorization() string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
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

	var jsonData UserInfo
	var jsonAssets UserAssets
	err = json.Unmarshal(buf2, &jsonData)
	err = json.Unmarshal(buf2, &jsonAssets)
	if err != nil {
		fmt.Printf("Unmarshall error: %v\n", err)
	}
	fmt.Printf("User login: %v\nCurrent Balance is: %v\nUser Email: %v\n", jsonData.Login, jsonData.Balance, jsonData.Email)
	fmt.Printf("Assets are: %v\n", jsonAssets.Assets)

	return buf2
}

func CurrentBalance(c *http.Client, url string) []byte {
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

	var jsonData UserInfo
	var jsonAssets UserAssets
	err = json.Unmarshal(buf2, &jsonData)
	err = json.Unmarshal(buf2, &jsonAssets)
	if err != nil {
		fmt.Printf("Unmarshall error: %v\n", err)
	}
	fmt.Printf("User login: %v\nCurrent Balance is: %v\nUser Email: %v\n", jsonData.Login, jsonData.Balance, jsonData.Email)
	fmt.Printf("Assets are: %v\n", jsonAssets.Assets)

	return buf2
}

func post(c *http.Client, url string, sendData interface{}) []byte {

	json_data, err := json.Marshal(sendData)
	if err != nil {
		panic(err)
	}

	reader := bytes.NewBuffer(json_data)
	req, err := http.NewRequest("POST", url, reader)
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
		panic(fmt.Errorf("got HTTP error %v", res.StatusCode))
	}

	buf2, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Couldn't read response %v", buf2)
	}

	return buf2
}

func asset(c *http.Client) {
	post(c, baseURL+"/buy", tp)
}

func sellGoods(c *http.Client) {
	post(c, baseURL+"/sell", tp)
}

func trading() {

}

func main() {

	client := &http.Client{}

	sellGoods(client)
	//asset(client)
	//post(client, baseURL+"/buy", "white_wool")
	//CurrentBalance(client, baseURL+"/account")
	//CurrentBalance(client, baseURL+"/account")

}
