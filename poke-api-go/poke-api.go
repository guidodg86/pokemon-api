package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

const pokeApiUrl = "https://pokeapi.co/api/v2/pokemon/"

type pokeApiList struct {
	count    int
	next     string
	previous string
	results  []pokeData
}

type pokeData struct {
	name string
	url  string
}

func main() {
	var listFinished bool = false
	var offset int = 0
	var targetUrl string
	var apiResponse pokeApiList

	for !listFinished {
		targetUrl = pokeApiUrl + "?offset=" + strconv.Itoa(offset)
		res, err := http.Get(targetUrl)
		if err != nil {
			fmt.Printf("poke-server: error making http request: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("poke-server: got response!\n")
		fmt.Printf("poke-server: status code: %d\n", res.StatusCode)
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("poke-server: could not read response body: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("poke-server: response body: %s\n", resBody)
		err = json.Unmarshal([]byte(resBody), &apiResponse)
		if err != nil {
			fmt.Printf("poke-server: could not parse json data: %s\n", err)
			os.Exit(1)
		}
	}

}
