package service

import (
	"encoding/json"
	"fmt"
	"is-deploy-agent/model"
	"log"
	"os"
	"testing"
)

func TestReadJsonValue(t *testing.T) {
	path, err := os.Open("../setting.json")
	if err != nil {
		log.Fatal(err)
	}

	var model []model.Model

	decoder := json.NewDecoder(path)
	decoder.Decode(&model)

	arrayLength := model[0].NodeList[0].PodList[0].ExcludeMap

	fmt.Println(len(arrayLength))

	for i := 0; i < len(arrayLength); i++ {
		key := model[0].NodeList[0].PodList[0].ExcludeMap[i].Key
		value := model[0].NodeList[0].PodList[0].ExcludeMap[i].Value

		fmt.Printf("%s = %s", key, value)
		fmt.Println()
	}
}

func TestJsonValueSave(t *testing.T) {
	path, err := os.Open("../setting.json")
	if err != nil {
		log.Fatal(err)
	}

	var model []model.Model

	decoder := json.NewDecoder(path)
	decoder.Decode(&model)

	arrayLength := model[0].NodeList[0].PodList[0].ExcludeMap

	fmt.Println(len(arrayLength))
	var newArray []ExcludeMap

	for i := 0; i < len(arrayLength); i++ {
		key := model[0].NodeList[0].PodList[0].ExcludeMap[i].Key
		value := model[0].NodeList[0].PodList[0].ExcludeMap[i].Value

		newArray = append(newArray, ExcludeMap{key, value})
	}

	fmt.Println(newArray, len(newArray))
}

type ExcludeMap struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
