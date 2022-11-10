package fetch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"is-deploy-agent/model"
	"os"
)

// GetSettingJson
// setting.json을 읽어서 반환한다
func GetSettingJson() model.Model {
	j, err := os.ReadFile("./setting.json")
	if err != nil {
		fmt.Println(err)
	}

	var models model.Model
	json.NewDecoder(bytes.NewBuffer(j)).Decode(&models)

	return models
}

func SyncSettingJson() {

}
