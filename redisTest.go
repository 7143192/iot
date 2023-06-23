package main

import (
	"encoding/json"
	"fmt"
	"huaweicloud.com/go-runtime/go-api/context"
	"huaweicloud.com/go-runtime/pkg/runtime"
	"iot/pkg/defines"
)

func ApigRedisTest(payload []byte, ctx context.RuntimeContext) (interface{}, error) {
	//var apigEvent apig.APIGTriggerEvent
	//err := json.Unmarshal(payload, &apigEvent)
	//if err != nil {
	//	fmt.Println("Unmarshal failed")
	//	return "invalid data", err
	//}
	//
	//rawBody := apigEvent.GetRawBody()
	//rawBodyByte := []byte(rawBody)
	//initInfo := &defines.InitInfo{}
	//_ = json.Unmarshal(rawBodyByte, initInfo)
	//mapByte, _ := json.Marshal(initInfo.MapInfo)
	//graphByte, _ := json.Marshal(initInfo.GraphInfo)
	//ctx.GetLogger().Logf("userData:%s", apigEvent.UserData)
	//ctx.GetLogger().Logf("rawBody:%s", rawBody)
	//ctx.GetLogger().Logf("mapInfo:%s", string(mapByte))
	//ctx.GetLogger().Logf("graphInfo:%s", string(graphByte))
	//apigResp := apig.APIGTriggerResponse{
	//	// Body: apigEvent.String(),
	//	Body: rawBody,
	//	Headers: map[string]string{
	//		"content-type": "application/json",
	//	},
	//	StatusCode: 200,
	//}

	//// call the demo function here.
	//url := "http://05757d8afcd54fc6881df55c8b2e2908.apig.cn-east-3.huaweicloudapis.com/demo_test"
	//res, err := http.Post(url, "application/json", strings.NewReader(rawBody))
	//defer res.Body.Close()

	input := defines.InputInfo{}
	_ = json.Unmarshal(payload, &input)
	fmt.Printf("got input info = %v\n", string(payload))
	//apigResp := apig.APIGTriggerResponse{
	//	// Body: apigEvent.String(),
	//	Body: string(payload),
	//	Headers: map[string]string{
	//		"content-type": "application/json",
	//	},
	//	StatusCode: 200,
	//}
	apigResp := map[string]interface{}{
		"result": string(payload),
	}
	apigRespByte, _ := json.Marshal(&apigResp)
	return string(apigRespByte), nil
}

func main() {
	runtime.Register(ApigRedisTest)
}

//func handler(event json.RawMessage, ctx context.RuntimeContext) (interface{}, error) {
//	input := defines.InputInfo{}
//	_ = json.Unmarshal(event, &input)
//	fmt.Printf("input = %v\n", input)
//	response := map[string]interface{}{
//		"result": "test",
//	}
//	return response, nil
//}
