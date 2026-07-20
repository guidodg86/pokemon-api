package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

const pokeApiUrl = "https://pokeapi.co/api/v2/pokemon/"

type pokeApiList struct {
	Count    int
	Next     string
	Previous string
	Results  []pokeData
}

type pokeData struct {
	Name string
	Url  string
}

func main() {
	router := gin.Default()
	router.GET("/", getPokes)
	router.Run("localhost:8080")
}

func getPokes(c *gin.Context) {
	var listFinished bool = false
	var targetUrl string = pokeApiUrl
	var listPokes []pokeData
	limitStr := c.DefaultQuery("limit", "200")
	offset := c.DefaultQuery("offset", "0")

	limitInt, e1 := strconv.Atoi(limitStr)
	if e1 != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request"})
		return
	}
	targetUrl = targetUrl + "?offset=" + offset
	var countPokes int
	for !listFinished {
		var apiResponse pokeApiList
		fmt.Printf("poke-server: fetching from %s\n", targetUrl)
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
		err = json.Unmarshal([]byte(resBody), &apiResponse)
		if err != nil {
			fmt.Printf("poke-server: could not parse json data: %s\n", err)
			os.Exit(1)
		}

		for _, v := range apiResponse.Results {
			listPokes = append(listPokes, v)
			countPokes++
		}

		if apiResponse.Next == "" || countPokes > limitInt {
			listFinished = true
		} else {
			targetUrl = apiResponse.Next
		}
	}
	fmt.Printf("poke-server: full pokemon list fetched from server\n")
	c.IndentedJSON(http.StatusOK, listPokes)

}
