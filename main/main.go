package main

import (
	"encoding/json"
	"os"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

var ContinentsFromJson Continents
var IsValid = regexp.MustCompile(`^[a-zA-Z-]+$`).MatchString

func main() {
	byteValue, _ := os.ReadFile("continents.json")
	json.Unmarshal(byteValue, &ContinentsFromJson)
	router:= gin.Default()
	router.GET("/continent/names", getContinentsName)
	router.GET("/continents", getContinents)
	router.POST("/continent", addContinent)
	router.GET("/countries", listAllCountries)
	router.GET("/countries/continent/:name", listCountriesByContinent)
	router.GET("/country/:name", getContinentByCountry)
	router.GET("/cities/country/:name", listCitiesByCountry)
	router.DELETE("/continents/:name", removeContinent)
	router.Run("localhost:5050")
}

func removeContinent(c *gin.Context) {
	continent := c.Param("name")
	if !IsValid(continent) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid Continent Name"})
		return
	}
	deletedContinent := ContinentsFromJson.deleteContinent(continent)
	c.IndentedJSON(http.StatusOK, deletedContinent)
}

func listCitiesByCountry(c *gin.Context) {
	country := c.Param("name")
	if !IsValid(country) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid Country Name"})
		return
	}
	cities := ContinentsFromJson.getCitiesByCountry(country)
	if cities == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Cities not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, cities)
}

func getContinentsName(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, ContinentsFromJson.getContinents())
}

func listAllCountries(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, ContinentsFromJson.getAllCountries())
}

func getContinents (c *gin.Context) {
	c.IndentedJSON(http.StatusOK, ContinentsFromJson)
}

func addContinent(c *gin.Context) {
	var continent Continent
	err := c.BindJSON(&continent)
	if  err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid JSON"})
		return 
	}
	addedContinent := (&ContinentsFromJson).addContinent(continent)
	if addedContinent.ContinentName == "" {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid Continent JSON"})
		return 
	} 
	c.IndentedJSON(http.StatusOK, addedContinent)
}

func listCountriesByContinent(c *gin.Context) {
	continent := c.Param("name")
	if !IsValid(continent) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid Continent Name"})
		return
	}
	countries := ContinentsFromJson.getCountriesByContinent(continent)
	if countries == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Continent not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, countries)
}

func getContinentByCountry(c *gin.Context) {
	country := c.Param("name")
	if !IsValid(country) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid Country Name"})
		return
	}
	continent := ContinentsFromJson.getContinentByCountry(country)
	if continent == "" {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Country not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, continent)
}