package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// 打包代码：  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
func main() {
	text := fmt.Sprintf("告警消息 ：time %v ", time.Now())

	// 在钉钉群中创建机器人并复制其token替换xx
	dingUrl := "https://oapi.dingtalk.com/robot/send?access_token=xx"
	data := make(map[string]interface{})

	data["msgtype"] = "text"
	data["text"] = map[string]string{"content": text}
	// isAtAll 是否@所有人
	data["at"] = map[string]interface{}{"atMobiles": [0]string{}, "isAtAll": false}

	bytePayload, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, postErr := http.Post(dingUrl, "application/json", bytes.NewBuffer(bytePayload))
	if postErr != nil {
		fmt.Println(err)
		return
	}
}
