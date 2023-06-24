package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	url := "http://05757d8afcd54fc6881df55c8b2e2908.apig.cn-east-3.huaweicloudapis.com/redis_demo"
	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Printf("error = %v\n", err)
		return
	}
	request.Header.Add("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("an error occurs when send out new shcedule request!\n")
		return
	}
	bodyReader := response.Body
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(bodyReader)
	fmt.Printf("response body = %v\n", buf)
	return
}
