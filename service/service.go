package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"is-deploy-agent/model"
)
func readFile(path string) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(file))
}

func writeFile(path string, lb string) {
	ioutil.WriteFile(path, []byte(lb), os.FileMode(644))
}

func readJson() string {
	path, err := ioutil.ReadFile("./setting.json")
	if err != nil {
		log.Fatal(err)
	}

	var model []model.Model
	json.Unmarshal(path, &model)

	key := model[0].PodList[0].LbMap[0].Key
	value := model[0].PodList[0].LbMap[0].Value

	result := key + "=" + value
	return result
}

func Restore()  {
	path := "../ex.properties"
	lb := readJson()
	readFile(path)
	writeFile(path, lb)
	readFile(path)
}