package roadMap

import (
	"iot/pkg/defines"
)

func RoadMapInit() *defines.RoadMap {
	res := &defines.RoadMap{}
	res.Nodes = make([]*defines.MapNode, 0)
	res.Lanes = make([]*defines.LaneInfo, 0)
	res.Warehouses = make([]*defines.WarehouseInfo, 0)
	x := 0
	y := 0
	// init all nodes line-by-line.
	for i := 0; i < defines.GRID_LINE_NUM; i++ {
		lane0 := &defines.LaneInfo{}
		lane1 := &defines.LaneInfo{}
		lane0.LaneType = defines.HORIZONTAL
		lane1.LaneType = defines.HORIZONTAL
		lane0.Nodes = make([]*defines.MapNode, 0)
		lane1.Nodes = make([]*defines.MapNode, 0)
		for j := 0; j < defines.GRID_LINE_NUM; j++ {
			// init four nodes in one grid.
			node0 := &defines.MapNode{}
			node1 := &defines.MapNode{}
			node2 := &defines.MapNode{}
			node3 := &defines.MapNode{}
			node0.X = x + defines.LANE_WIDTH + defines.SHELVES_WIDTH
			node1.X = x + 3*defines.LANE_WIDTH + 2*defines.SHELVES_WIDTH
			node2.X = x + defines.LANE_WIDTH + defines.SHELVES_WIDTH
			node3.X = x + 3*defines.LANE_WIDTH + 2*defines.SHELVES_WIDTH
			node0.Y = y + defines.LANE_WIDTH + defines.SHELVES_WIDTH
			node1.Y = y + defines.LANE_WIDTH + defines.SHELVES_WIDTH
			node2.Y = y + 3*defines.LANE_WIDTH + 2*defines.SHELVES_WIDTH
			node3.Y = +3*defines.LANE_WIDTH + 2*defines.SHELVES_WIDTH
			res.Nodes = append(res.Nodes, node0)
			res.Nodes = append(res.Nodes, node1)
			res.Nodes = append(res.Nodes, node2)
			res.Nodes = append(res.Nodes, node3)
			x = x + 3*defines.SHELVES_WIDTH + 4*defines.LANE_WIDTH
			lane0.Nodes = append(lane0.Nodes, node0)
			lane0.Nodes = append(lane0.Nodes, node1)
			lane1.Nodes = append(lane1.Nodes, node2)
			lane1.Nodes = append(lane1.Nodes, node3)
		}
		x = 0
		y = (i + 1) * (4*defines.LANE_WIDTH + 3*defines.SHELVES_WIDTH)
		// init horizontal lanes first.
		res.Lanes = append(res.Lanes, lane0)
		res.Lanes = append(res.Lanes, lane1)
	}
	// init vertical lanes next.
	for i := 0; i < defines.GRID_LINE_NUM; i++ {
		lane0 := &defines.LaneInfo{}
		lane1 := &defines.LaneInfo{}
		lane0.LaneType = defines.VERTICAL
		lane1.LaneType = defines.VERTICAL
		lane0.Nodes = make([]*defines.MapNode, 0)
		lane1.Nodes = make([]*defines.MapNode, 0)
		for j := 0; j < defines.GRID_LINE_NUM; j++ {
			// init four nodes in one grid.
			node0 := &defines.MapNode{}
			node1 := &defines.MapNode{}
			node2 := &defines.MapNode{}
			node3 := &defines.MapNode{}
			node0.X = x + defines.LANE_WIDTH + defines.SHELVES_WIDTH
			node1.X = x + 3*defines.LANE_WIDTH + 2*defines.SHELVES_WIDTH
			node2.X = x + defines.LANE_WIDTH + defines.SHELVES_WIDTH
			node3.X = x + 3*defines.LANE_WIDTH + 2*defines.SHELVES_WIDTH
			node0.Y = y + defines.LANE_WIDTH + defines.SHELVES_WIDTH
			node1.Y = y + defines.LANE_WIDTH + defines.SHELVES_WIDTH
			node2.Y = y + 3*defines.LANE_WIDTH + 2*defines.SHELVES_WIDTH
			node3.Y = +3*defines.LANE_WIDTH + 2*defines.SHELVES_WIDTH
			y = y + 3*defines.SHELVES_WIDTH + 4*defines.LANE_WIDTH
			lane0.Nodes = append(lane0.Nodes, node0)
			lane0.Nodes = append(lane0.Nodes, node2)
			lane1.Nodes = append(lane1.Nodes, node1)
			lane1.Nodes = append(lane1.Nodes, node3)
		}
		y = 0
		x = (i + 1) * (4*defines.LANE_WIDTH + 3*defines.SHELVES_WIDTH)
		res.Lanes = append(res.Lanes, lane0)
		res.Lanes = append(res.Lanes, lane1)
	}
	// then init warehouses info.

	return res
}
