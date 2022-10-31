package utils

import (
	"encoding/json"
	"fmt"
	"is-deploy-agent/model"
	"os"
)

func GetJsonToTest() model.Model {
	path, err := os.Open("../../setting.json")
	if err != nil {
		fmt.Println(err)
		//todo log로 변경
	}

	var model model.Model

	decoder := json.NewDecoder(path)
	err = decoder.Decode(&model)
	if err != nil {
		fmt.Println(err)
		//todo log로 변경
	}

	defer path.Close()
	return model
}
