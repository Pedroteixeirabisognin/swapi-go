package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/pedroteixeirabisognin/swapi-go/dto"
)

//Retorna um planeta da api swapi
func GetPlanets() {

	var API_URL string = "https://swapi.dev/api/planets/?page=1"

	//var planetSlice []dto.Results

	var planetDTO dto.Results

	for {

		response, err := http.Get(API_URL)

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(responseData, &planetDTO)

		for x := 0; x < len(planetDTO.Results); x++ {
			fmt.Printf("O planeta é: %s\n", planetDTO.Results[x].Name)
			fmt.Printf("Seu clima é: %s\n", planetDTO.Results[x].Climate)
			fmt.Printf("Seu terreno é: %s\n", planetDTO.Results[x].Terrain)
			fmt.Printf("Apareceu em %d filme(s)", len(planetDTO.Results[x].Films))

			fmt.Print("\n")
			fmt.Print("\n")

		}

		if planetDTO.Next == API_URL {
			break
		} else {
			API_URL = planetDTO.Next
		}

	}

}
