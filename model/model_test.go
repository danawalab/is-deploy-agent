package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestWriteJson(t *testing.T) {
	service := make([]Model, 1)

	service[0].Service = "PAS"
	service[0].PodList= make([]PodList, 2)
	service[0].PodList[0].Name = "서버1"
	service[0].PodList[0].Ip = "127.0.0.1"
	service[0].PodList[0].Port = "8080"
	service[0].PodList[0].Path = "c/apache24/conf/uriworkermap.properties"
	service[0].PodList[0].LbMap = make([]WorkerMap, 1)
	service[0].PodList[0].LbMap[0].Key = "/*"
	service[0].PodList[0].LbMap[0].Value = "load_balancer"
	service[0].PodList[0].TomcatList = make([]TomcatList, 2)
	service[0].PodList[0].TomcatList[0].Name = "WAS1"
	service[0].PodList[0].TomcatList[0].ExcludeMap = make([]WorkerMap, 2)
	service[0].PodList[0].TomcatList[0].ExcludeMap[0].Key = "/was1"
	service[0].PodList[0].TomcatList[0].ExcludeMap[0].Value = "load_balancer_1"
	service[0].PodList[0].TomcatList[0].ExcludeMap[1].Key = "/*"
	service[0].PodList[0].TomcatList[0].ExcludeMap[1].Value = "load_balancer_ex1"
	service[0].PodList[0].TomcatList[1].Name = "WAS2"
	service[0].PodList[0].TomcatList[1].ExcludeMap = make([]WorkerMap, 2)
	service[0].PodList[0].TomcatList[1].ExcludeMap[0].Key = "/was2"
	service[0].PodList[0].TomcatList[1].ExcludeMap[0].Value = "load_balancer_2"
	service[0].PodList[0].TomcatList[1].ExcludeMap[1].Key = "/*"
	service[0].PodList[0].TomcatList[1].ExcludeMap[1].Value = "load_balancer_ex2"

	service[0].PodList[1].Name = "서버2"
	service[0].PodList[1].Ip = "127.0.0.2"
	service[0].PodList[1].Port = "8080"
	service[0].PodList[1].Path = "c/apache24/conf/uriworkermap.properties"
	service[0].PodList[1].LbMap = make([]WorkerMap, 1)
	service[0].PodList[1].LbMap[0].Key = "/*"
	service[0].PodList[1].LbMap[0].Value = "load_balancer"
	service[0].PodList[1].TomcatList = make([]TomcatList, 2)
	service[0].PodList[1].TomcatList[0].Name = "WAS3"
	service[0].PodList[1].TomcatList[0].ExcludeMap = make([]WorkerMap, 2)
	service[0].PodList[1].TomcatList[0].ExcludeMap[0].Key = "/was1"
	service[0].PodList[1].TomcatList[0].ExcludeMap[0].Value = "load_balancer_3"
	service[0].PodList[1].TomcatList[0].ExcludeMap[1].Key = "/*"
	service[0].PodList[1].TomcatList[0].ExcludeMap[1].Value = "load_balancer_ex3"
	service[0].PodList[1].TomcatList[1].Name = "WAS4"
	service[0].PodList[1].TomcatList[1].ExcludeMap = make([]WorkerMap, 2)
	service[0].PodList[1].TomcatList[1].ExcludeMap[0].Key = "/was4"
	service[0].PodList[1].TomcatList[1].ExcludeMap[0].Value = "load_balancer_4"
	service[0].PodList[1].TomcatList[1].ExcludeMap[1].Key = "/*"
	service[0].PodList[1].TomcatList[1].ExcludeMap[1].Value = "load_balancer_ex4"

	doc, _ := json.Marshal(service)

	err := ioutil.WriteFile("./test.json", doc, os.FileMode(0644))
	if err != nil {
		log.Fatal(err)
		return
	}
}

func TestReadJson(t *testing.T) {
	path, err := ioutil.ReadFile("./test.json")
	if err != nil {
		log.Fatal(err)
		return
	}

	var service []Model

	json.Unmarshal(path, &service)

	fmt.Println(service[0].PodList[0].LbMap[0].Key)
	fmt.Println(service[0].PodList[0].LbMap[0].Value)
}
