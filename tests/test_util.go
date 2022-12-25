package tests

import (
	"encoding/json"
	"is-deploy-agent/utils"
	"log"
	"os"
)

func GetJsonToTest() utils.Apache {
	path, err := os.Open("../../setting.json")
	if err != nil {
		log.Println(err)
	}

	var models utils.Apache

	decoder := json.NewDecoder(path)
	err = decoder.Decode(&models)
	if err != nil {
		log.Println(err)
	}

	defer path.Close()
	return models
}
