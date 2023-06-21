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
	info := &defines.WarehouseInfo{}
	info0 := &defines.WarehouseInfo{}
	info1 := &defines.WarehouseInfo{}
	info2 := &defines.WarehouseInfo{}
	info3 := &defines.WarehouseInfo{}
	width := 4*defines.LANE_WIDTH + 3*defines.SHELVES_WIDTH
	info.RangeX = make([]int, 2)
	info.RangeY = make([]int, 2)
	info0.RangeX = make([]int, 2)
	info0.RangeY = make([]int, 2)
	info1.RangeX = make([]int, 2)
	info1.RangeY = make([]int, 2)
	info2.RangeX = make([]int, 2)
	info2.RangeY = make([]int, 2)
	info3.RangeX = make([]int, 2)
	info3.RangeY = make([]int, 2)
	info.Gates = make([]*defines.GateInfo, 4)
	info0.Gates = make([]*defines.GateInfo, 1)
	info1.Gates = make([]*defines.GateInfo, 1)
	info2.Gates = make([]*defines.GateInfo, 1)
	info3.Gates = make([]*defines.GateInfo, 1)
	// the center warehouse should be the first element in the RoadMap and this warehouse is the src warehouse.
	info.RangeX[0] = (defines.GRID_LINE_NUM / 2) * width
	info.RangeX[1] = (defines.GRID_LINE_NUM/2 + 1) * width
	info.RangeY[0] = (defines.GRID_LINE_NUM / 2) * width
	info.RangeY[1] = (defines.GRID_LINE_NUM/2 + 1) * width
	for i := 0; i < 4; i++ {
		info.Gates[i] = &defines.GateInfo{}
		info.Gates[i].GateType = defines.SRC
	}
	info.Gates[0].Center.X = width
	info.Gates[0].Center.Y = 2*width - defines.LANE_WIDTH - defines.SHELVES_WIDTH
	info.Gates[1].Center.X = width + defines.SHELVES_WIDTH + defines.LANE_WIDTH
	info.Gates[1].Center.Y = width
	info.Gates[2].Center.X = width * 2
	info.Gates[2].Center.Y = width + defines.SHELVES_WIDTH + defines.LANE_WIDTH
	info.Gates[3].Center.X = 2*width - defines.LANE_WIDTH - defines.SHELVES_WIDTH
	info.Gates[3].Center.Y = 2 * width
	info0.RangeX[0] = 0
	info0.RangeX[1] = width
	info0.RangeY[0] = 0
	info0.RangeY[1] = width
	info0.Gates[0].Center.X = width
	info0.Gates[0].Center.Y = width - defines.LANE_WIDTH - defines.SHELVES_WIDTH
	info1.RangeX[0] = (defines.GRID_LINE_NUM - 1) * width
	info1.RangeX[1] = defines.GRID_LINE_NUM * width
	info1.RangeY[0] = 0
	info1.RangeY[1] = width
	info1.Gates[0].Center.X = 2*width + defines.SHELVES_WIDTH + defines.LANE_WIDTH
	info1.Gates[0].Center.Y = width
	info2.RangeX[0] = 0
	info2.RangeX[1] = width
	info2.RangeY[0] = (defines.GRID_LINE_NUM - 1) * width
	info2.RangeY[1] = defines.GRID_LINE_NUM * width
	info2.Gates[0].Center.X = width - defines.SHELVES_WIDTH - defines.LANE_WIDTH
	info2.Gates[0].Center.Y = 2 * width
	info3.RangeX[0] = (defines.GRID_LINE_NUM - 1) * width
	info3.RangeX[1] = defines.GRID_LINE_NUM * width
	info3.RangeY[0] = (defines.GRID_LINE_NUM - 1) * width
	info3.RangeY[1] = defines.GRID_LINE_NUM * width
	info3.Gates[0].Center.X = 2 * width
	info3.Gates[0].Center.Y = width*2 + defines.LANE_WIDTH + defines.SHELVES_WIDTH
	res.Warehouses = append(res.Warehouses, info)
	res.Warehouses = append(res.Warehouses, info0)
	res.Warehouses = append(res.Warehouses, info1)
	res.Warehouses = append(res.Warehouses, info2)
	res.Warehouses = append(res.Warehouses, info3)
	return res
}
