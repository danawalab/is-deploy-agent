package fetch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"is-deploy-agent/model"
	"log"
	"os"
)

// GetSettingJson
// setting.json을 읽어서 반환한다
func GetSettingJson() model.Node {
	file, err := os.ReadFile("./setting.json")
	if err != nil {
		fmt.Println(err)
	}

	var models model.Node
	err = json.NewDecoder(bytes.NewBuffer(file)).Decode(&models)
	if err != nil {
		log.Println(err)
	}

	return models
}

func SyncSettingJson(json string) error {
	file, err := os.Create("./setting.json")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	_, err = file.Write([]byte(json))
	if err != nil {
		log.Println(err)
	}
	return err
}
