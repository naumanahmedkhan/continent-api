package main

import "strings"

type Continents struct {
	Continents []Continent `json:"continents"`
}

type Continent struct {
	ContinentName string `json:"continentName"`
	Countries []Country   `json:"countries"`
}

type Country struct {
	CountryName string   `json:"country"`
	Cities []string      `json:"cities"`
}

func (c *Continents) getContinents() []string {
	var continents []string
	for _,continent := range (*c).Continents {
		continents = append(continents, continent.ContinentName)
	}
	return continents
}

func (c *Continents) getCountriesByContinent(continentName string) []string {
	var countries []string
	for _, continent := range (*c).Continents {
		if strings.ToLower(continentName) == strings.ToLower(continent.ContinentName) {
			for _, country := range continent.Countries {
				countries = append(countries, country.CountryName)
			}
		}
	}
	return countries
}

func (c *Continents) getAllCountries() []string {
	var countries []string
	for _, continent := range (*c).Continents {
		for _, country := range continent.Countries {
		countries = append(countries, country.CountryName)	
		}
	}
	return countries
}

func (c *Continents) getContinentByCountry(countryName string) string {
	var continentName string
	for _, continent := range (*c).Continents {
		found := false
		for _, country := range continent.Countries {
			if strings.ToLower(country.CountryName) == strings.ToLower(countryName) {
				found = true
			}
		}
		if found {
			continentName = continent.ContinentName
		}
	}
	return continentName
}

func (c *Continents) getCitiesByCountry(countryName string) []string {
	var cities []string
	for _, continent := range (*c).Continents {
		for _, country := range continent.Countries {
			if strings.ToLower(country.CountryName) == strings.ToLower(countryName) {
				cities = country.Cities
			}
		}
	}
	return cities
}

func (c *Continents) deleteContinent(continent string) Continent {
	var deletedContinent Continent
	for index, existingContinent := range (*c).Continents {
		if (strings.ToLower(existingContinent.ContinentName) == strings.ToLower(continent)) {
			deletedContinent = existingContinent
			(*c).Continents = append((*c).Continents[:index], (*c).Continents[index+1:]...)
		}
	}
	return deletedContinent
}

func (c *Continents) addContinent(continent Continent) Continent {
	if continent.ContinentName == "" {
		return Continent{}
	}
	if !contains((*c).Continents, continent) {
		(*c).Continents = append ((*c).Continents, continent)
	}
	return continent
}

func contains(continents []Continent, continent Continent) bool {
	for _, existingContinent := range continents {
		if strings.ToLower(existingContinent.ContinentName) == strings.ToLower(continent.ContinentName) {
			return true
		}
	}
	return false
} 