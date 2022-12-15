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

// SyncSettingJson
// Console에서 설정한 json을 받아서 해당 json 으로 변경
//
//	Agent 실행시 setting.json이 없어도 Console에서 json을 설정 후 저장하면 해당 API가 호출 되고 setting.json을 생성
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

// UpdateAgent
// 에이전트 동적 업데이트 및 다운그레이드
// 1.0.0 버전에서는 사용 금지 update.sh 미완성 및 내부 보안 정책에 의해 사용 하기 힘듬
func UpdateAgent(version string) error {
	node, err := utils.GetJson()
	if err != nil {
		log.Println(err)
		return err
	}
	port := node.Agent.Port

	cmd := exec.Command("./update.sh", port[1:], version)
	_, err = cmd.Output()

	if err != nil {
		log.Println(err)
		return err
	}

	return err
}
