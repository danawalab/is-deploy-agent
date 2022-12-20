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
// agent.sh에 업데이트 기능까지 전부 넣으면 중간에 에이전트를 종료 시켜서 실행중 이던 sh도 종료됨
// agent-update.sh을 따로 만들어 agent.sh이 agent에 의해 실행 되면 인수 값을 agent-update.sh에 전달
// 해당 agent-update.sh이 기존 agent 죽여도 실행 됨
// 종종 제대로 동작 안하는 일이 있음 확인 필요 !!!!!!!!!!!!!!!!!!!!!!!!!!!!
func UpdateAgent(version string) error {
	node, err := utils.GetJson()
	if err != nil {
		log.Println(err)
		return err
	}
	port := node.Agent.Port

	cmd := exec.Command("./agent.sh", port[1:], version)
	_, err = cmd.Output()

	if err != nil {
		log.Println(err)
		return err
	}

	return err
}
