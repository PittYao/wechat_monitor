package model

import (
	"encoding/json"
	"os"
)

type Data struct {
	Articles []*Wechat `json:"app_msg_list"`
}

func UnmarshalJson() Data {
	//读取json文件
	fileName := "resp.json"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	var d Data
	err = json.Unmarshal(data, &d)
	if err != nil {
		panic(err)
	}
	return d
}
