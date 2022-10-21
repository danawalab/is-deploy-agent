package loadbalance

import (
	"encoding/json"
	"fmt"
	"is-deploy-agent/model"
	"log"
	"os"
	"testing"
)

func TestReadJsonValue(t *testing.T) {
	model := getJson()

	arrayLength := model[0].NodeList[0].PodList[0].ExcludeMap

	fmt.Println(len(arrayLength))

	for i := 0; i < len(arrayLength); i++ {
		key := model[0].NodeList[0].PodList[0].ExcludeMap[i].Key
		value := model[0].NodeList[0].PodList[0].ExcludeMap[i].Value

		fmt.Printf("TestReadJsonValue, %s = %s", key, value)
		fmt.Println()
	}
}

func TestJsonValueSave(t *testing.T) {
	model := getJson()

	arrayLength := model[0].NodeList[0].PodList[0].ExcludeMap

	fmt.Println(len(arrayLength))
	var newArray []ExcludeMap

	for i := 0; i < len(arrayLength); i++ {
		key := model[0].NodeList[0].PodList[0].ExcludeMap[i].Key
		value := model[0].NodeList[0].PodList[0].ExcludeMap[i].Value

		newArray = append(newArray, ExcludeMap{key, value})
	}

	fmt.Println("TestJsonValueSave = ", newArray, len(newArray))
}

func TestFindByName(t *testing.T) {
	worker := "WAS1"
	model := getJson()

	length := len(model[0].NodeList[0].PodList)
	var newArray []ExcludeMap

	for i := 0; i < length; i++ {
		pod := model[0].NodeList[0].PodList[i]
		name := pod.Name

		if worker == name {
			exLength := len(pod.ExcludeMap)

			for j := 0; j < exLength; j++ {
				key := pod.ExcludeMap[j].Key
				value := pod.ExcludeMap[j].Value

				newArray = append(newArray, ExcludeMap{key, value})
			}
			break
		}
	}

	fmt.Println("TestFindByName = ", newArray, len(newArray))
}

func getJson() []model.Model {
	path, err := os.Open("../setting.json")
	if err != nil {
		log.Fatal(err)
	}

	var model []model.Model

	decoder := json.NewDecoder(path)
	decoder.Decode(&model)

	return model
}

type ExcludeMap struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
