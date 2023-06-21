package main

import (
	"encoding/json"
	"fmt"
	"huaweicloud.com/go-runtime/events/apig"
	"huaweicloud.com/go-runtime/events/cts"
	"huaweicloud.com/go-runtime/events/dds"
	"huaweicloud.com/go-runtime/events/dis"
	"huaweicloud.com/go-runtime/events/kafka"
	"huaweicloud.com/go-runtime/events/lts"
	"huaweicloud.com/go-runtime/events/smn"
	"huaweicloud.com/go-runtime/events/timer"
	"huaweicloud.com/go-runtime/go-api/context"
	"huaweicloud.com/go-runtime/pkg/runtime"
)

func ApigTest(payload []byte, ctx context.RuntimeContext) (interface{}, error) {
	var apigEvent apig.APIGTriggerEvent
	err := json.Unmarshal(payload, &apigEvent)
	if err != nil {
		fmt.Println("Unmarshal failed")
		return "invalid data", err
	}
	ctx.GetLogger().Logf("payload:%s", apigEvent.String())
	apigResp := apig.APIGTriggerResponse{
		Body: apigEvent.String(),
		Headers: map[string]string{
			"content-type": "application/json",
		},
		StatusCode: 200,
	}
	return apigResp, nil
}

func CtsTest(payload []byte, ctx context.RuntimeContext) (interface{}, error) {
	var ctsEvent cts.CTSTriggerEvent
	err := json.Unmarshal(payload, &ctsEvent)
	if err != nil {
		fmt.Println("Unmarshal failed")
		return "invalid data", err
	}
	ctx.GetLogger().Logf("payload:%s", ctsEvent.String())
	return "ok", nil
}

func DdsTest(payload []byte, ctx context.RuntimeContext) (interface{}, error) {
	var ddsEvent dds.DDSTriggerEvent
	err := json.Unmarshal(payload, &ddsEvent)
	if err != nil {
		fmt.Println("Unmarshal failed")
		return "invalid data", err
	}
	ctx.GetLogger().Logf("payload:%s", ddsEvent.String())
	return "ok", nil
}

func DisTest(payload []byte, ctx context.RuntimeContext) (interface{}, error) {
	var disEvent dis.DISTriggerEvent
	err := json.Unmarshal(payload, &disEvent)
	if err != nil {
		fmt.Println("Unmarshal failed")
		return "invalid data", err
	}
	ctx.GetLogger().Logf("payload:%s", disEvent.String())
	return "ok", nil
}

func KafkaTest(payload []byte, ctx context.RuntimeContext) (interface{}, error) {
	var kafkaEvent kafka.KAFKATriggerEvent
	err := json.Unmarshal(payload, &kafkaEvent)
	if err != nil {
		fmt.Println("Unmarshal failed")
		return "invalid data", err
	}
	ctx.GetLogger().Logf("payload:%s", kafkaEvent.String())
	return "ok", nil
}

func LtsTest(payload []byte, ctx context.RuntimeContext) (interface{}, error) {
	var ltsEvent lts.LTSTriggerEvent
	err := json.Unmarshal(payload, &ltsEvent)
	if err != nil {
		fmt.Println("Unmarshal failed")
		return "invalid data", err
	}
	ctx.GetLogger().Logf("payload:%s", ltsEvent.String())
	return "ok", nil
}

func SmnTest(payload []byte, ctx context.RuntimeContext) (interface{}, error) {
	var smnEvent smn.SMNTriggerEvent
	err := json.Unmarshal(payload, &smnEvent)
	if err != nil {
		fmt.Println("Unmarshal failed")
		return "invalid data", err
	}
	ctx.GetLogger().Logf("payload:%s", smnEvent.String())
	return "ok", nil
}

func TimerTest(payload []byte, ctx context.RuntimeContext) (interface{}, error) {
	var timerEvent timer.TimerTriggerEvent
	err := json.Unmarshal(payload, &timerEvent)
	if err != nil {
		fmt.Println("Unmarshal failed")
		return "invalid data", err
	}
	return timerEvent.String(), nil
}

func main() {
	runtime.Register(ApigTest)
}
