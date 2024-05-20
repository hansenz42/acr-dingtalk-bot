package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"time"
)

//type DingtalkMessage struct {
//	MsgType string `json:"msgtype"`
//	Text    string `json:"text"`
//}

func main() {
	TARGET_URL := "https://oapi.dingtalk.com/robot/send?access_token={YOUR DINGTALK BOT ACCESSTOKEN}"

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/OTFlNDNlOGItNzc0NS00ZTczLWFjMzYtMGEzYTI0MzExY2Vl", func(c *gin.Context) {
		// get data from request body
		data, err := c.GetRawData()
		if err != nil {
			c.String(500, err.Error())
		}

		now_datatime := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second())
		msg_str := fmt.Sprintf("%s: %s", now_datatime, "Private ACR image pushed successfully")
		json_str := fmt.Sprintf(`{"msgtype": "text","text": {"content": "%s"}}`, msg_str)

		// send out_msg to TARGET_URL use json
		post_resp, err := http.Post(TARGET_URL, "application/json", bytes.NewBuffer([]byte(json_str)))

		if err != nil {
			c.String(500, err.Error())
		}

		// print response
		defer post_resp.Body.Close()
		println(post_resp.Status)
		// print response body
		body, err := io.ReadAll(post_resp.Body)
		if err != nil {
			c.String(500, err.Error())
		}
		fmt.Println(string(body))

		// print data
		c.String(200, string(data))
	})

	//r.Run() // listen and serve on 0.0.0.0:8080
	r.Run(":8085")
}
