package main

// Global imports
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

// Global constants
const pokeApiUrl = "https://pokeapi.co/api/v2/pokemon/"
const pokeSpriteUrl = "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/"

// Structs to receive data from the pokeapi
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

// Structs to send data to clients
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

// Main with the endpoint declarations
func main() {
	router := gin.Default()
	router.GET("/", getPokes)
	router.GET("/:name", getPokeByName)
	router.Run("localhost:8080")
}

// Home endpoint
func getPokes(c *gin.Context) {
	var listFinished bool = false
	var targetUrl string = pokeApiUrl
	var result fullPokeResult
	var countPokes int
	var pokeDraft pokeResult
	var apiResponse pokeApiList

	// Parsing query parameters
	limitStr := c.DefaultQuery("limit", "20")
	offsetStr := c.DefaultQuery("offset", "0")
	qName := c.DefaultQuery("q", "")

	// Converting limit and offset to integers
	limitInt, e1 := strconv.Atoi(limitStr)
	if e1 != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request"})
		return
	}
	offsetInt, e1 := strconv.Atoi(offsetStr)
	if e1 != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request"})
		return
	}

	// Count and limit for the response are already known
	result.Limit = limitInt
	result.Offset = offsetInt

	// Setting up target Url for fetching the list of pokes
	targetUrl = targetUrl + "?offset=" + offsetStr

	// Main loop to fetch the data
	for !listFinished {
		// Perform the http request and parse it
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

		//Loop over the list of pokes we received filtering if needed
		for _, v := range apiResponse.Results {
			vSplit := strings.Split(v.Url, "/")
			id := vSplit[len(vSplit)-2]
			pokeDraft.Id = id
			pokeDraft.Image = pokeSpriteUrl + string(id) + ".png"
			pokeDraft.Name = v.Name
			if qName != "" {
				if !strings.Contains(v.Name, qName) {
					continue
				}
			}
			result.Data = append(result.Data, pokeDraft)
			countPokes++
			if countPokes == limitInt {
				break
			}
		}

		// Check if ending condition happened (no more pokes or reached limit)
		if apiResponse.Next == "" || countPokes == limitInt {
			listFinished = true
		} else {
			targetUrl = apiResponse.Next
		}
	}

	// Sending the response
	result.Total = apiResponse.Count
	fmt.Printf("poke-server: full pokemon list fetched from server\n")
	c.IndentedJSON(http.StatusOK, result)
}

func getPokeByName(c *gin.Context) {

}
