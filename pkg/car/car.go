package car

import (
	"container/heap"
	"fmt"
	"iot/pkg/defines"
)
import "math"

// init pos of one car.
var initPos *defines.Pos = &defines.Pos{
	X: 1040,
	Y: 5280,
}

// heap part

type NodeHeap []*defines.GraphNode

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].Weight < h[j].Weight }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*defines.GraphNode))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// GetPosInLane return laneInfo that contains this pos and the nearest node to the pos in this lane.
func GetPosInLane(mapInfo *defines.RoadMap, pos *defines.Pos) (*defines.LaneInfo, *defines.MapNode) {
	lanes := mapInfo.Lanes
	size := len(lanes)
	for i := 0; i < size; i++ {
		lane := lanes[i]
		if lane.LaneType == defines.VERTICAL && pos.X > lane.CenterPos-defines.LANE_WIDTH && pos.X < lane.CenterPos+defines.LANE_WIDTH {
			nodes := lane.Nodes
			size1 := len(nodes)
			min := 99999999
			resNode := &defines.MapNode{}
			for j := 0; j < size1; j++ {
				dis := math.Abs(float64(nodes[j].Y - pos.Y))
				if int(dis) < min {
					min = int(dis)
					resNode = nodes[j]
				}
			}
			return lane, resNode
		}
		if lane.LaneType == defines.HORIZONTAL && pos.Y > lane.CenterPos-defines.LANE_WIDTH && pos.Y < lane.CenterPos+defines.LANE_WIDTH {
			nodes := lane.Nodes
			size1 := len(nodes)
			min := 99999999
			resNode := &defines.MapNode{}
			for j := 0; j < size1; j++ {
				dis := math.Abs(float64(nodes[j].X - pos.X))
				if int(dis) < min {
					min = int(dis)
					resNode = nodes[j]
				}
			}
			return lane, resNode
		}
	}
	return nil, nil
}

func CheckInWareHouse(house *defines.WarehouseInfo, pos *defines.Pos) bool {
	if pos.X > house.RangeX[0] && pos.X < house.RangeX[1] && pos.Y > house.RangeY[0] && pos.Y < house.RangeY[1] {
		return true
	}
	return false
}

func GetGraphNodeID(pos *defines.Pos, graph *defines.Graph, mapInfo *defines.RoadMap) int {
	nodes := graph.Nodes
	found := false
	for i := 0; i < len(nodes); i++ {
		if pos.X == nodes[i].X && pos.Y == nodes[i].Y {
			found = true
			return i
		}
	}
	if found == false {
		_, resNode := GetPosInLane(mapInfo, pos)
		for i := 0; i < len(nodes); i++ {
			if resNode.X == nodes[i].X && resNode.Y == nodes[i].Y {
				return i
			}
		}
	}
	return -1
}

func Dijkstra(graph *defines.Graph, start int, end int) []*defines.Pos {
	LIMIT := 100000000
	path := make([]int, 0)
	res := make([]*defines.Pos, 0)
	dist := make([]*defines.GraphNode, 0)
	checked := make([]int, 0)
	q := NodeHeap(make([]*defines.GraphNode, 0))
	heap.Init(&q)
	for i := 0; i < len(graph.Nodes); i++ {
		checked = append(checked, 0)
		path = append(path, -1)
		dist = append(dist, &defines.GraphNode{
			ID:     i,
			Weight: LIMIT,
			X:      graph.Nodes[i].X,
			Y:      graph.Nodes[i].Y,
		})
	}
	dist[start].Weight = 0
	heap.Push(&q, dist[start])
	for q.Len() > 0 {
		node := heap.Pop(&q).(*defines.GraphNode)
		u := node.ID
		if checked[u] == 1 {
			continue
		}
		checked[u] = 1
		nodeEdges := graph.Edges[u]
		for j := 0; j < len(nodeEdges); j++ {
			tempV := nodeEdges[j].End
			tempW := nodeEdges[j].Weight
			if checked[tempV] == 0 && dist[tempV].Weight > dist[u].Weight+tempW {
				dist[tempV].Weight = dist[u].Weight + tempW
				path[tempV] = u
				heap.Push(&q, dist[tempV])
			}
		}
	}
	temp := end
	hasPath := true
	times := 0
	for temp != start {
		if times >= 50 {
			hasPath = false
			break
		}
		if temp == -1 {
			hasPath = false
			break
		}
		res = append(res, &defines.Pos{
			X: dist[temp].X,
			Y: dist[temp].Y,
		})
		temp = path[temp]
		times++
	}
	if hasPath == true {
		res = append(res, &defines.Pos{
			X: dist[start].X,
			Y: dist[start].Y,
		})
		return res
	}
	res = make([]*defines.Pos, 0)
	return res
}

// GetCurCarDir when the car is not located at any of the nodes, we should make car run to the start node first.
func GetCurCarDir(mapInfo *defines.RoadMap, carInfo *defines.Car, pos *defines.Pos, start *defines.Pos) {
	carInfo.CurX = start.X
	carInfo.CurY = start.Y
	lane, _ := GetPosInLane(mapInfo, pos)
	if lane.LaneType == defines.VERTICAL && pos.Y < start.Y {
		// TODO: logic to make the car move.
		carInfo.Dir = defines.UP
		StraightMoveCar(mapInfo, carInfo, pos, start)
	}
	if lane.LaneType == defines.VERTICAL && pos.Y > start.Y {
		// TODO: logic to make the car move.
		carInfo.Dir = defines.DOWN
		StraightMoveCar(mapInfo, carInfo, pos, start)
	}
	if lane.LaneType == defines.HORIZONTAL && pos.X < start.X {
		// TODO: logic to make the car move.
		carInfo.Dir = defines.RIGHT
		StraightMoveCar(mapInfo, carInfo, pos, start)
	}
	if lane.LaneType == defines.HORIZONTAL && pos.X < start.X {
		// TODO: logic to make the car move.
		carInfo.Dir = defines.LEFT
		StraightMoveCar(mapInfo, carInfo, pos, start)
	}
	return
}

// StraightMoveCar make the car go straightly. UP / DOWN / LEFT / RIGHT
func StraightMoveCar(mapInfo *defines.RoadMap, carInfo *defines.Car, start *defines.Pos, dest *defines.Pos) {

}

// TurnDirMoveCar the case that the car turn around.
func TurnDirMoveCar(mapInfo *defines.RoadMap, carInfo *defines.Car, start *defines.Pos, dest *defines.Pos) {

}

// SortDestWarehouse this function should select a dest warehouse for this scheduled car.
// And maybe can choose a suitable destination pos according to current storage situation?
func SortDestWarehouse() int {
	return 1
}

func GetCarRunningLane(mapInfo *defines.RoadMap, carInfo *defines.Car, pos *defines.Pos) {
	lane, _ := GetPosInLane(mapInfo, pos)
	carInfo.LaneInfo = lane
	if lane.LaneType == defines.VERTICAL {
		X := lane.Nodes[0].X
		if pos.X >= X-defines.LANE_WIDTH && pos.X <= X {
			carInfo.RunningLane = 0
		} else {
			carInfo.RunningLane = 1
		}
	} else {
		Y := lane.Nodes[0].Y
		if pos.Y >= Y-defines.LANE_WIDTH && pos.Y <= Y {
			carInfo.RunningLane = 0
		} else {
			carInfo.RunningLane = 1
		}
	}
}

func ScheduleOnePath(mapInfo *defines.RoadMap, graph *defines.Graph,
	carInfo *defines.Car, start *defines.Pos, dest *defines.Pos) []*defines.Pos {
	res := make([]*defines.Pos, 0)
	// start and dest are the same position.
	if start.X == dest.X && start.Y == dest.Y {
		return res
	}
	sID := GetGraphNodeID(start, graph, mapInfo)
	eID := GetGraphNodeID(dest, graph, mapInfo)
	got := Dijkstra(graph, sID, eID)
	if len(got) > 0 {
		startPos := got[len(got)-1]
		res = append(res, got...)
		if !(start.X == startPos.X && start.Y == startPos.Y) {
			GetCurCarDir(mapInfo, carInfo, start, startPos)
			res = append(res, start)
		}
		return res
	}
	fmt.Printf("startID = %d, destID = %d\n", sID, eID)
	fmt.Printf("there is no path between node %v and %v!\n", *start, *dest)
	return res
}

func GetDestWarehouseID(dest *defines.Pos, mapInfo *defines.RoadMap) int {
	for i := 0; i < 4; i++ {
		if dest.X > mapInfo.Warehouses[i+1].RangeX[0] && dest.X < mapInfo.Warehouses[i+1].RangeX[1] &&
			dest.Y > mapInfo.Warehouses[i+1].RangeY[0] && dest.Y < mapInfo.Warehouses[i+1].RangeY[1] {
			return i + 1
		}
	}
	return -1
}

// ScheduleOneCar start and dest pos should be center pos of one car.
func ScheduleOneCar(mapInfo *defines.RoadMap, graph *defines.Graph,
	carInfo *defines.Car, start *defines.Pos, dest *defines.Pos) []*defines.Pos {
	// the start pos should be inside the src warehouse!
	res := make([]*defines.Pos, 0)
	if !CheckInWareHouse(mapInfo.Warehouses[0], start) {
		return res
	}
	carInfo.Start = *start
	carInfo.Dest = *dest
	//// schedule from init pos to one gate of src warehouse.
	//srcGatePos := &defines.Pos{
	//	X: 3160,
	//	Y: 5280,
	//}
	// schedule from init pos to start pos inside the src warehouse.
	got0 := ScheduleOnePath(mapInfo, graph, carInfo, initPos, start)
	if len(got0) > 0 {
		for i := len(got0) - 1; i >= 0; i-- {
			res = append(res, got0[i])
		}
	}
	// schedule from start pos to dest pos inside the dest warehouse.
	got1 := ScheduleOnePath(mapInfo, graph, carInfo, start, dest)
	if len(got1) > 0 {
		for i := len(got1) - 1; i >= 0; i-- {
			res = append(res, got1[i])
		}
	}
	// schedule from dest pos to init pos of the cars.
	got2 := ScheduleOnePath(mapInfo, graph, carInfo, dest, initPos)
	if len(got2) > 0 {
		for i := len(got2) - 1; i >= 0; i-- {
			res = append(res, got2[i])
		}
	}
	return res
}
