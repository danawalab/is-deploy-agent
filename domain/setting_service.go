package domain

import (
	"log"
	"os"
)

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
