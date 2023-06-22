package test

import (
	"fmt"
	"iot/pkg/car"
	"iot/pkg/roadMap"
	"testing"
)

func TestCar(t *testing.T) {
	mapInfo := roadMap.MapInit()
	graph := roadMap.GraphInit(mapInfo)
	res := car.Dijkstra(graph, 12, 36)
	for i := 0; i < len(res); i++ {
		fmt.Printf("x = %d, y = %d\n", res[i].X, res[i].Y)
	}
}
