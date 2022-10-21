package model

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestWriteJson(t *testing.T) {
	service := make([]Model, 1)

	service[0].Service = "PAS"
	service[0].NodeList = make([]NodeList, 2)
	service[0].NodeList[0].Name = "서버1"
	service[0].NodeList[0].Ip = "127.0.0.1"
	service[0].NodeList[0].Port = "8080"
	service[0].NodeList[0].Path = "c/apache24/conf/uriworkermap.properties"
	service[0].NodeList[0].LbMap = make([]WorkerMap, 1)
	service[0].NodeList[0].LbMap[0].Key = "/*"
	service[0].NodeList[0].LbMap[0].Value = "load_balancer"
	service[0].NodeList[0].PodList = make([]PodList, 2)
	service[0].NodeList[0].PodList[0].Name = "WAS1"
	service[0].NodeList[0].PodList[0].ExcludeMap = make([]WorkerMap, 2)
	service[0].NodeList[0].PodList[0].ExcludeMap[0].Key = "/was1"
	service[0].NodeList[0].PodList[0].ExcludeMap[0].Value = "load_balancer_1"
	service[0].NodeList[0].PodList[0].ExcludeMap[1].Key = "/*"
	service[0].NodeList[0].PodList[0].ExcludeMap[1].Value = "load_balancer_ex1"
	service[0].NodeList[0].PodList[1].Name = "WAS2"
	service[0].NodeList[0].PodList[1].ExcludeMap = make([]WorkerMap, 2)
	service[0].NodeList[0].PodList[1].ExcludeMap[0].Key = "/was2"
	service[0].NodeList[0].PodList[1].ExcludeMap[0].Value = "load_balancer_2"
	service[0].NodeList[0].PodList[1].ExcludeMap[1].Key = "/*"
	service[0].NodeList[0].PodList[1].ExcludeMap[1].Value = "load_balancer_ex2"

	service[0].NodeList[1].Name = "서버2"
	service[0].NodeList[1].Ip = "127.0.0.2"
	service[0].NodeList[1].Port = "8080"
	service[0].NodeList[1].Path = "c/apache24/conf/uriworkermap.properties"
	service[0].NodeList[1].LbMap = make([]WorkerMap, 1)
	service[0].NodeList[1].LbMap[0].Key = "/*"
	service[0].NodeList[1].LbMap[0].Value = "load_balancer"
	service[0].NodeList[1].PodList = make([]PodList, 2)
	service[0].NodeList[1].PodList[0].Name = "WAS3"
	service[0].NodeList[1].PodList[0].ExcludeMap = make([]WorkerMap, 2)
	service[0].NodeList[1].PodList[0].ExcludeMap[0].Key = "/was1"
	service[0].NodeList[1].PodList[0].ExcludeMap[0].Value = "load_balancer_3"
	service[0].NodeList[1].PodList[0].ExcludeMap[1].Key = "/*"
	service[0].NodeList[1].PodList[0].ExcludeMap[1].Value = "load_balancer_ex3"
	service[0].NodeList[1].PodList[1].Name = "WAS4"
	service[0].NodeList[1].PodList[1].ExcludeMap = make([]WorkerMap, 2)
	service[0].NodeList[1].PodList[1].ExcludeMap[0].Key = "/was4"
	service[0].NodeList[1].PodList[1].ExcludeMap[0].Value = "load_balancer_4"
	service[0].NodeList[1].PodList[1].ExcludeMap[1].Key = "/*"
	service[0].NodeList[1].PodList[1].ExcludeMap[1].Value = "load_balancer_ex4"

	doc, _ := json.Marshal(service)

	jsonFile, err := os.Create("./test.json")
	if err != nil {
		log.Fatal(err)
	}

	jsonFile.Write(doc)
	defer jsonFile.Close()
}

func TestReadJson(t *testing.T) {
	path, err := os.Open("./test.json")
	if err != nil {
		log.Fatal(err)
		return
	}

	var service []Model

	decoder := json.NewDecoder(path)
	decoder.Decode(&service)

	fmt.Println(service[0].NodeList[0].Ip)
}
