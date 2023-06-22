package main

import (
	"encoding/json"
	"fmt"
	"huaweicloud.com/go-runtime/events/apig"
	"huaweicloud.com/go-runtime/go-api/context"
	"huaweicloud.com/go-runtime/pkg/runtime"
	"iot/pkg/defines"
	"net/http"
	"strings"
)

func ApigRedisTest(payload []byte, ctx context.RuntimeContext) (interface{}, error) {

	var apigEvent apig.APIGTriggerEvent
	err := json.Unmarshal(payload, &apigEvent)
	if err != nil {
		fmt.Println("Unmarshal failed")
		return "invalid data", err
	}
	rawBody := apigEvent.GetRawBody()
	rawBodyByte := []byte(rawBody)
	initInfo := &defines.InitInfo{}
	_ = json.Unmarshal(rawBodyByte, initInfo)
	mapByte, _ := json.Marshal(initInfo.MapInfo)
	graphByte, _ := json.Marshal(initInfo.GraphInfo)
	ctx.GetLogger().Logf("rawBody:%s", rawBody)
	ctx.GetLogger().Logf("mapInfo:%s", string(mapByte))
	ctx.GetLogger().Logf("rawBody:%s", string(graphByte))
	//// then store these info into Redis.
	//rd := redis.NewClient(&redis.Options{
	//	Addr:     "192.168.189.129:6379", // url
	//	Password: "",
	//	DB:       0, // No.0 DB
	//})
	//err = rd.Set("mapKey", string(mapByte), 0).Err()
	//if err != nil {
	//	ctx.GetLogger().Logf("set mapInfo error msg = %s", err)
	//}
	//err = rd.Set("graphKey", string(graphByte), 0).Err()
	//if err != nil {
	//	ctx.GetLogger().Logf("set graphInfo error msg = %s", err)
	//}
	//val, err := rd.Get("mapKey").Result()
	//if err != nil {
	//	ctx.GetLogger().Logf("get mapInfo error msg = %s", err)
	//} else {
	//	ctx.GetLogger().Logf("get mapInfo = %s", val)
	//}
	//val1, err := rd.Get("graphKey").Result()
	//if err != nil {
	//	ctx.GetLogger().Logf("get graphInfo error msg = %s", err)
	//} else {
	//	ctx.GetLogger().Logf("get graphInfo = %s", val1)
	//}
	apigResp := apig.APIGTriggerResponse{
		// Body: apigEvent.String(),
		Body: rawBody,
		Headers: map[string]string{
			"content-type": "application/json",
		},
		StatusCode: 200,
	}
	// call the demo function here.
	url := "http://05757d8afcd54fc6881df55c8b2e2908.apig.cn-east-3.huaweicloudapis.com/demo_test"
	res, err := http.Post(url, "application/json", strings.NewReader(rawBody))
	defer res.Body.Close()
	return apigResp, nil
}

func main() {
	runtime.Register(ApigRedisTest)
}
