package tests

import (
	"gopkg.in/yaml.v2"
	"os"
	"testing"
)

type DynamicType struct {
	FileType   string `yaml:"fileType"`
	DeployType string `yaml:"deployType"`
	Model      any    `yaml:"model"`
}

type Node struct {
	Name    string    `yaml:"name"`
	Agent   Agent     `yaml:"agent"`
	Path    string    `yaml:"path"`
	LbMap   []UriMap  `yaml:"lbMap"`
	PodList []PodList `yaml:"podList"`
}

type PodList struct {
	Name    string   `yaml:"name"`
	LbMap   []UriMap `yaml:"lbMap"`
	LogPath string   `yaml:"logPath"`
	ShPath  string   `yaml:"shPath"`
}

type UriMap struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

type Agent struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func TestDynamicType(t *testing.T) {

	dt := DynamicType{
		FileType:   "war",
		DeployType: "shell",
	}

	//var model Node
	if dt.FileType == "war" {
		if dt.DeployType == "shell" {
			dt.Model = Node{
				Name: "test-ap",
				Agent: Agent{
					Host: "localhost",
					Port: ":5000",
				},
				Path: "/etc/apache2/uri.properties",
				LbMap: []UriMap{
					{
						Key:   "/*",
						Value: "lb",
					},
				},
				PodList: []PodList{
					{
						Name: "test-to1",
						LbMap: []UriMap{
							{
								Key:   "/*",
								Value: "worker2",
							},
						},
						LogPath: "/logs/te1.log",
						ShPath:  "/sr/te1.sh",
					},
					{
						Name: "test-to2",
						LbMap: []UriMap{
							{
								Key:   "/*",
								Value: "worker1",
							},
						},
						LogPath: "/logs/te2.log",
						ShPath:  "/sr/te2.sh",
					},
				},
			}
		}
	}

	//marshal, _ := yaml.Marshal(dt.Model)
	//println(string(marshal))

	var dt1 DynamicType
	file1, _ := os.Open("./config.yml")
	defer file1.Close()
	_ = yaml.NewDecoder(file1).Decode(&dt1)
	println("a = ", dt1.DeployType)

	var dt2 Node
	file2, _ := os.Open("./config.yml")
	defer file2.Close()
	_ = yaml.NewDecoder(file2).Decode(&dt2)
	println("b = ", dt2.Name)

	// file1 에서 열고 다시 file2에서 같은 파일을 여는 이유
	// file1을 DynamicType에 Deocde 후 다시 file1을 Node에 Decode시 안됨
}
