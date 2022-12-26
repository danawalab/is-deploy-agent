package service

import (
	"is-deploy-agent/utils"
	"log"
	"os/exec"
)

// AgentUpdate
// 에이전트 동적 업데이트 및 다운그레이드
// agent.sh에 업데이트 기능까지 전부 넣으면 중간에 에이전트를 종료 시켜서 실행중 이던 sh도 종료됨
// agent-update.sh을 따로 만들어 agent.sh이 agent에 의해 실행 되면 인수 값을 agent-update.sh에 전달
// 해당 agent-update.sh이 기존 agent 죽여도 실행 됨
// 종종 제대로 동작 안하는 일이 있음 확인 필요 !!!!!!!!!!!!!!!!!!!!!!!!!!!!
func AgentUpdate(version string) error {
	node, err := utils.GetSettingJson()
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
