package service

import (
	"log"
	"os/exec"
)

func TomcatHealthCheck(worker string) (string, error) {

	//cmd := exec.Command("bash", "-c", "ps -ef | grep "+worker)
	cmd := exec.Command("bash", "-c", "ps -ef | grep "+worker+" | grep -v grep")

	output, err := cmd.Output()
	if err != nil {
		log.Println(err)
		return string(output), err
	}

	return string(output), err
}
