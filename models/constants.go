package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadKey(which string) (token string) {
	Keys := Keys{}

	jsonFile, err := os.Open("configs/config.json")
	if err != nil {
		fmt.Println("Deu ruim")
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Deu ruim")
	}

	err = json.Unmarshal(byteValue, &Keys)
	if err != nil {
		fmt.Println("Deu ruim")
	}

	if which == "mongoPwd" {
		token = Keys.MongoPwd
	} else {
		fmt.Println("Deu Ruim")
	}

	return token
}

//Keys -
type Keys struct {
	MongoPwd string `json:"mongoPwd"`
}
