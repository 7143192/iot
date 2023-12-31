package main

import (
	"encoding/json"
	"fmt"
	"huaweicloud.com/go-runtime/events/apig"
	"huaweicloud.com/go-runtime/go-api/context"
	"huaweicloud.com/go-runtime/pkg/runtime"
	"iot/pkg/car"
	"iot/pkg/defines"
)

func ApigDemoTest(payload []byte, ctx context.RuntimeContext) (interface{}, error) {
	//var apigEvent apig.APIGTriggerEvent
	//err := json.Unmarshal(payload, &apigEvent)
	//if err != nil {
	//	fmt.Println("Unmarshal failed")
	//	return "invalid data", err
	//}
	//rawBody := apigEvent.GetRawBody()
	//rawBodyByte := []byte(rawBody)
	//initInfo := &defines.InitInfo{}
	//_ = json.Unmarshal(rawBodyByte, initInfo)
	//mapByte, _ := json.Marshal(initInfo.MapInfo)
	//graphByte, _ := json.Marshal(initInfo.GraphInfo)
	//ctx.GetLogger().Logf("demo rawBody:%s", rawBody)
	//ctx.GetLogger().Logf("demo mapInfo:%s", string(mapByte))
	//ctx.GetLogger().Logf("demo rawBody:%s", string(graphByte))
	//// ctx.GetLogger().Logf("payload:%s", apigEvent.String())
	//apigResp := apig.APIGTriggerResponse{
	//	Body: apigEvent.String(),
	//	Headers: map[string]string{
	//		"content-type": "application/json",
	//	},
	//	StatusCode: 200,
	//}
	//return apigResp, nil

	fmt.Printf("demo got payload = %v\n", string(payload))
	gotMap := make(map[string]string)
	input := defines.InputInfo{}
	_ = json.Unmarshal(payload, &gotMap)
	inputStr := gotMap["input"]
	_ = json.Unmarshal([]byte(inputStr), &input)
	mapInfo := input.Init.MapInfo
	mapByte, _ := json.Marshal(mapInfo)
	graphInfo := input.Init.GraphInfo
	graphByte, _ := json.Marshal(graphInfo)
	carInfo := input.Init.CarInfo
	start := input.Init.Start
	dest := input.Init.Dest
	startByte, _ := json.Marshal(start)
	destByte, _ := json.Marshal(dest)
	fmt.Printf("demo got mapInfo = %v\n", string(mapByte))
	fmt.Printf("demo got graphInfo = %v\n", string(graphByte))
	fmt.Printf("demo got startInfo = %v\n", string(startByte))
	fmt.Printf("demo got destInfo = %v\n", string(destByte))
	res := car.ScheduleOneCar(mapInfo, graphInfo, carInfo, start, dest)
	for _, val := range res {
		fmt.Printf("x = %v, y = %v\n", val.X, val.Y)
	}
	apigResp := apig.APIGTriggerResponse{
		Body: string(payload),
		Headers: map[string]string{
			"content-type": "application/json",
		},
		StatusCode: 200,
	}

	return apigResp, nil
}

func main() {
	runtime.Register(ApigDemoTest)
}
