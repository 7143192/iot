package test

import (
	"fmt"
	"iot/pkg/roadMap"
	"testing"
)

func TestMap(t *testing.T) {
	res := roadMap.MapInit()
	lanes := res.Lanes
	for i := 0; i < len(lanes); i++ {
		nodes := lanes[i].Nodes
		for j := 0; j < len(nodes); j++ {
			fmt.Printf("node info = %d %d\n", nodes[j].X, nodes[j].Y)
		}
	}
}
