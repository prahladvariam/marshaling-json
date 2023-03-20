package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	AppName string `json:"appName"`
	AppID   int    `json:"appId"`
}

type Test struct {
	Test    string `json:"test"`
	Data    string `json:"data"`
	AppName string `json:"appName"`
}

func main() {
	jsonStr := `{"test": "abc","data": "{\"appName\": \"hello\", \"appId\": 123}","appName": ""}`

	var t Test
	err := json.Unmarshal([]byte(jsonStr), &t)
	if err != nil {
		panic(err)
	}

	var d Data
	err = json.Unmarshal([]byte(t.Data), &d)
	if err != nil {
		panic(err)
	}

	t.AppName = d.AppName

	newJsonData, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	fmt.Println("Given Json Data:", jsonStr)
	fmt.Println("New Json Data:", string(newJsonData))

}
