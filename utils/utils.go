package utils

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// GetSettingJson
// setting.json을 읽어서 반환
func GetSettingJson() (Apache, error) {
	path, err := os.Open("./setting.json")
	if err != nil {
		log.Println(err)
		return Apache{}, err
	}
	defer path.Close()

	var models Apache
	err = json.NewDecoder(path).Decode(&models)
	if err != nil {
		log.Println(err)
		return Apache{}, err
	}

	return models, err
}

func GetConfigFile() (Apache, error) {
	file, err := os.Open("./config.yml")
	if err != nil {
		log.Println(err)
		return Apache{}, err
	}
	defer file.Close()

	var model Apache
	yaml.NewDecoder(file).Decode(&model)

	return model, err
}

// IsNameEqual
// setting.json에서 지정한 node의 name 또는 podList에 name이 worker 인자와 같은 경우 true 반환
func IsNameEqual(name string, tomcat string) bool {
	if name == tomcat {
		return true
	}
	return false
}
