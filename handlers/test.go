package handlers

import (
	"encoding/json"
	"fmt"
	"huaweicloud.com/go-runtime/go-api/context"
	"iot/pkg/defines"
)

func handler(event json.RawMessage, ctx context.RuntimeContext) (interface{}, error) {
	input := defines.InputInfo{}
	_ = json.Unmarshal(event, &input)
	fmt.Printf("input = %v\n", input)
	response := map[string]interface{}{
		"result": "test",
	}
	return response, nil
}
