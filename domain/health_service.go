package domain

import (
	"log"
	"os/exec"
)

func HealthCheck(tomcat string) (string, error) {

	cmd := exec.Command("bash", "-c", "ps -ef | grep "+tomcat)

	output, err := cmd.Output()
	if err != nil {
		log.Println(err)
		return string(output), err
	}

	return string(output), err
}
