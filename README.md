# continent-api
## Description
This api is a basic implememtation of Rest Api in go language for managing continents with countries and city. This application is using HTTP web frame work GIN. More details about Gin can be found [here](https://gin-gonic.com/docs/). This Application is initally fetching content from JSON file when the application starts and all the changes done later are in memory.

## Running in local Environment
### Prerequisite
- docker installed. [Installation guidelines](https://docs.docker.com/get-docker/)
- docker-compose installed. [Installation guidelines](https://docs.docker.com/compose/install/)
### Instructions for setting up local Environment
- Clone the git repository into your local machine.
- Run the following command from the downloaded git repository
```bash
$ docker-compose up
```
- This should set up application in your application and start the application on local host on port 5050.

# Rest API
The REST API to the continent-api is described below.

## GET Names of All the Continent
### Request
```bash
GET /continent/names
```
```bash
curl http://localhost:5050/continents/name
```
### Response
```bash
[
    "Asia",
    "Europe",
    "Africa",
    "Oceania"
]
```
## Get All Continents and Countries with Cities
### Request
```bash
GET /continents
```
```bash
curl http://localhost:5050/continents
```
### Response
```bash
{
    "continents": [
        {
            "continentName": "Asia",
            "countries": [
                {
                    "country": "India",
                    "cities": [
                        "Mumbai",
                        "Delhi"
                    ]
                },
                {
                    "country": "Pakistan",
                    "cities": [
                        "Lahore",
                        "Karachi"
                    ]
                }
            ]
        },
        {
            "continentName": "Europe",
            "countries": [
                {
                    "country": "Finland",
                    "cities": [
                        "Tampere",
                        "Turku",
                        "Helsinki"
                    ]
                },
                {
                    "country": "Sweden",
                    "cities": [
                        "Stockholm",
                        "Malmo"
                    ]
                }
            ]
        },
        {
            "continentName": "Africa",
            "countries": [
                {
                    "country": "Algeria",
                    "cities": null
                },
                {
                    "country": "Angola",
                    "cities": null
                }
            ]
        },
        {
            "continentName": "Oceania",
            "countries": [
                {
                    "country": "Australia",
                    "cities": [
                        "Sydney",
                        "Perth"
                    ]
                }
            ]
        }
    ]
}
```
## Get all the countries
### Request
```bash
GET /countries
```
```bash
curl http://localhost:5050/countries
```
### Response
```bash
[
    "India",
    "Pakistan",
    "Finland",
    "Sweden",
    "Algeria",
    "Angola",
    "Australia"
]
```
## GET Names countries by Continent
### Request
```bash
GET /countries/continent/:name
```
```bash
curl http://localhost:5050/countries/continent/Europe
```
### Response
```bash
[
    "Finland",
    "Sweden"
]
```
## Get Continent Name by Country
### Request
```bash
GET /country/:name
```
```bash
curl http://localhost:5050/country/pakistan
```
### Response
```bash
"Asia"
```
## Get Names of Cities By Country
### Request
```bash
GET /cities/country/:name
```
```bash
curl http://localhost:5050/cities/country/finland
```
### Response
```bash
[
    "Tampere",
    "Turku",
    "Helsinki"
]
```
## Create Continent
### Request
```bash
POST /continent
```
```bash
curl -X POST http://localhost:5050/continent -H 'Content-Type:application/json' -d '{"continentName":"Antartica"}'
```
### Response
```bash
{
    "continentName": "Antartica",
    "countries": null
}
```
## Delete Continent by name
### Request
```bash
DELETE /continents/:name
```
```bash
curl -X DELETE http://localhost:5050/continents/oceania
```
### Response
```bash
{
    "continentName": "Oceania",
    "countries": [
        {
            "country": "Australia",
            "cities": [
                "Sydney",
                "Perth"
            ]
        }
    ]
}
```

