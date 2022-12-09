package fetch

import (
	"bytes"
	"encoding/json"
	"is-deploy-agent/model"
	"is-deploy-agent/utils"
	"log"
	"os"
	"os/exec"
)

// GetSettingJson
// setting.json을 읽어서 반환한다
func GetSettingJson() (model.Node, error) {
	file, err := os.ReadFile("./setting.json")
	if err != nil {
		log.Println(err)
		return model.Node{}, err
	}

	var models model.Node
	err = json.NewDecoder(bytes.NewBuffer(file)).Decode(&models)
	if err != nil {
		log.Println(err)
		return model.Node{}, err
	}

	return models, err
}

func SyncSettingJson(json string) error {
	file, err := os.Create("./setting.json")
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte(json))
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

func UpdateAgent(version string) error {
	node, err := utils.GetJson()
	if err != nil {
		log.Println(err)
		return err
	}
	port := node.Agent.Port

	cmd := exec.Command("./update.sh", port[1:], version)
	out, err := cmd.Output()
	log.Println(">>> 1 ", string(out))
	log.Println(">>> 2 ", out)
	log.Println(">>> err ", err)

	if err != nil {
		log.Println(err)
		return err
	}

	return err
}
