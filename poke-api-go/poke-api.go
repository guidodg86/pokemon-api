package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const pokeApiUrl = "https://pokeapi.co/api/v2/pokemon/"
const pokeSpriteUrl = "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/"

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

type pokeResult struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type fullPokeResult struct {
	Total  int          `json:"total"`
	Limit  int          `json:"limit"`
	Offset int          `json:"offset"`
	Data   []pokeResult `json:"data"`
}

func main() {
	router := gin.Default()
	router.GET("/", getPokes)
	router.Run("localhost:8080")
}

func getPokes(c *gin.Context) {
	var listFinished bool = false
	var targetUrl string = pokeApiUrl
	var result fullPokeResult
	limitStr := c.DefaultQuery("limit", "20")
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
		var pokeDraft pokeResult
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
			vSplit := strings.Split(v.Url, "/")
			id := vSplit[len(vSplit)-2]
			pokeDraft.Id = id
			pokeDraft.Image = pokeSpriteUrl + string(id) + ".png"
			pokeDraft.Name = v.Name
			result.Data = append(result.Data, pokeDraft)
			countPokes++
			if countPokes == limitInt {
				break
			}
		}

		if apiResponse.Next == "" || countPokes == limitInt {
			listFinished = true
		} else {
			targetUrl = apiResponse.Next
		}
	}
	fmt.Printf("poke-server: full pokemon list fetched from server\n")
	c.IndentedJSON(http.StatusOK, result)

}
