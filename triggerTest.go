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

func triggerFunc() {

}

func triggerWorkflow() {
	// url := "https://05757d8afcd54fc6881df55c8b2e2908.apig.cn-east-3.huaweicloudapis.com/carSchedule/start"
	client := &http.Client{}
	url := "https://iam.myhuaweicloud.com/v3/auth/tokens?nocatalog=true"
	scope := defines.Scope{}
	pro := defines.Project{}
	pro.Name = "cn-east-3"
	scope.Pro = pro
	// body := defines.TestEntireBody{}
	body := defines.TestAuth{}
	auth := defines.Auth{}
	domain := defines.Domain{}
	domain.Name = "zhyueyan"
	user := defines.User{}
	user.DomainInfo = domain
	user.Name = "liyuhan"
	user.Pwd = "abcLYH125558"
	password := defines.Password{}
	password.UserInfo = user
	methods := make([]string, 0)
	methods = append(methods, "password")
	identity := defines.Identity{}
	identity.Methods = methods
	identity.PwdInfo = password
	auth.IdentityInfo = identity
	auth.ScopeInfo = scope
	body.AuthInfo = auth
	// body.MapInfo = roadMap.MapInit()
	// body.GraphInfo = roadMap.GraphInit(body.MapInfo)
	bodyByte, _ := json.Marshal(&body)
	// send request.
	res, err := http.Post(url, "application/json", strings.NewReader(string(bodyByte)))
	defer res.Body.Close()
	if err != nil {
		log.Panicf("err = %v\n", err)
	}
	// bodyInfo, err := ioutil.ReadAll(res.Body)
	// fmt.Printf("response body = %v\n", string(bodyInfo))
	url1 := "https://05757d8afcd54fc6881df55c8b2e2908.apig.cn-east-3.huaweicloudapis.com/carSchedule/start"
	body1 := defines.InitInfo{}
	body1.MapInfo = roadMap.MapInit()
	body1.GraphInfo = roadMap.GraphInit(body1.MapInfo)
	body11 := defines.InputInfo{}
	body11.Init = body1
	body0 := defines.StartBody{}
	body0.Input = body11
	bodyByte1, _ := json.Marshal(&body0)
	fmt.Printf(string(bodyByte1) + "\n")
	request, _ := http.NewRequest("POST", url1, strings.NewReader(string(bodyByte1)))
	// dataByte, _ := json.Marshal(data["token"])
	// fmt.Printf("%v\n", string(dataByte))
	request.Header.Add("content-Type", "application/json")
	request.Header.Add("x-auth-token", res.Header.Get("X-Subject-Token"))
	// fmt.Printf("request = %v\n", request)
	resp, _ := client.Do(request)
	bodyInfo1, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("response body1 = %v\n", string(bodyInfo1))
}

func main() {
	triggerWorkflow()
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
