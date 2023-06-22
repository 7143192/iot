package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"iot/pkg/defines"
	"iot/pkg/roadMap"
	"log"
	"net/http"
	"strings"
)

func main() {
	url := "http://05757d8afcd54fc6881df55c8b2e2908.apig.cn-east-3.huaweicloudapis.com/go_test"
	body := defines.InitInfo{}
	body.MapInfo = roadMap.MapInit()
	body.GraphInfo = roadMap.GraphInit(body.MapInfo)
	bodyByte, _ := json.Marshal(&body)
	// send request.
	res, err := http.Post(url, "application/json", strings.NewReader(string(bodyByte)))
	defer res.Body.Close()
	if err != nil {
		log.Panicf("err = %v\n", err)
	}
	bodyInfo, err := ioutil.ReadAll(res.Body)
	fmt.Printf("response body = %v\n", string(bodyInfo))
	//rd := redis.NewClient(&redis.Options{
	//	Addr:     "127.0.0.1:6379", // url
	//	Password: "",
	//	DB:       0, // No.0 DB
	//})
	//result, err := rd.Ping().Result()
	//if err != nil {
	//	fmt.Println("ping err :", err)
	//	return
	//}
	//fmt.Println(result)
	//err = rd.Set("k1", "v1", 0).Err()
	//if err != nil {
	//	fmt.Println("set err :", err)
	//	return
	//}
	//// get
	//val, err := rd.Get("k1").Result()
	//if err != nil {
	//	fmt.Println("get err :", err)
	//	return
	//}
	//fmt.Println("k1 ==", val) // k1 == v1
	//rd.Close()
}
