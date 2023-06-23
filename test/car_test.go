package test

import (
	"fmt"
	"iot/pkg/car"
	"iot/pkg/defines"
	"iot/pkg/roadMap"
	"testing"
)

func TestCar(t *testing.T) {
	mapInfo := roadMap.MapInit()
	graph := roadMap.GraphInit(mapInfo)
	//for i := 0; i < len(graph.Nodes); i++ {
	//	fmt.Printf("node_x = %d, node_y = %d\n", graph.Nodes[i].X, graph.Nodes[i].Y)
	//}
	//fmt.Println()
	//res := car.Dijkstra(graph, 12, 36)
	//for i := 0; i < len(res); i++ {
	//	fmt.Printf("x = %d, y = %d\n", res[i].X, res[i].Y)
	//}
	startPos := &defines.Pos{
		X: 4200,
		Y: 4200,
	}
	destPos := &defines.Pos{
		X: 7360,
		Y: 8440,
	}
	carInfo := &defines.Car{}
	carInfo.Dir = defines.RIGHT
	got := car.ScheduleOneCar(mapInfo, graph, carInfo, startPos, destPos)
	for i := 0; i < len(got); i++ {
		fmt.Printf("path_x = %d, path_y = %d\n", got[i].X, got[i].Y)
	}
}
