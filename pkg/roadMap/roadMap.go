package roadMap

import (
	"iot/pkg/defines"
)

func MapInit() *defines.RoadMap {
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
		lane0.CenterPos = i*(4*defines.LANE_WIDTH+3*defines.SHELVES_WIDTH) + defines.SHELVES_WIDTH + defines.LANE_WIDTH
		lane1.CenterPos = i*(4*defines.LANE_WIDTH+3*defines.SHELVES_WIDTH) + 2*defines.SHELVES_WIDTH + 3*defines.LANE_WIDTH
		lane0.LaneType = defines.HORIZONTAL
		lane1.LaneType = defines.HORIZONTAL
		lane0.Nodes = make([]*defines.MapNode, 0)
		lane1.Nodes = make([]*defines.MapNode, 0)
		lane0.Cars = make([]*defines.Car, 0)
		lane1.Cars = make([]*defines.Car, 0)
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
			node3.Y = y + 3*defines.LANE_WIDTH + 2*defines.SHELVES_WIDTH
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
		res.HorizontalLanes = append(res.HorizontalLanes, lane0)
		res.HorizontalLanes = append(res.HorizontalLanes, lane1)
	}
	// init vertical lanes next.
	for i := 0; i < defines.GRID_LINE_NUM; i++ {
		lane0 := &defines.LaneInfo{}
		lane1 := &defines.LaneInfo{}
		lane0.CenterPos = i*(4*defines.LANE_WIDTH+3*defines.SHELVES_WIDTH) + defines.SHELVES_WIDTH + defines.LANE_WIDTH
		lane1.CenterPos = i*(4*defines.LANE_WIDTH+3*defines.SHELVES_WIDTH) + 2*defines.SHELVES_WIDTH + 3*defines.LANE_WIDTH
		lane0.LaneType = defines.VERTICAL
		lane1.LaneType = defines.VERTICAL
		lane0.Nodes = make([]*defines.MapNode, 0)
		lane1.Nodes = make([]*defines.MapNode, 0)
		lane0.Cars = make([]*defines.Car, 0)
		lane1.Cars = make([]*defines.Car, 0)
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
			node3.Y = y + 3*defines.LANE_WIDTH + 2*defines.SHELVES_WIDTH
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
		res.VerticalLanes = append(res.VerticalLanes, lane0)
		res.VerticalLanes = append(res.VerticalLanes, lane1)
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
	info0.Gates[0] = &defines.GateInfo{}
	info0.Gates[0].GateType = defines.DEST1
	info0.Gates[0].Center.X = width
	info0.Gates[0].Center.Y = width - defines.LANE_WIDTH - defines.SHELVES_WIDTH
	info1.RangeX[0] = (defines.GRID_LINE_NUM - 1) * width
	info1.RangeX[1] = defines.GRID_LINE_NUM * width
	info1.RangeY[0] = 0
	info1.RangeY[1] = width
	info1.Gates[0] = &defines.GateInfo{}
	info1.Gates[0].GateType = defines.DEST2
	info1.Gates[0].Center.X = 2*width + defines.SHELVES_WIDTH + defines.LANE_WIDTH
	info1.Gates[0].Center.Y = width
	info2.RangeX[0] = 0
	info2.RangeX[1] = width
	info2.RangeY[0] = (defines.GRID_LINE_NUM - 1) * width
	info2.RangeY[1] = defines.GRID_LINE_NUM * width
	info2.Gates[0] = &defines.GateInfo{}
	info2.Gates[0].GateType = defines.DEST3
	info2.Gates[0].Center.X = width - defines.SHELVES_WIDTH - defines.LANE_WIDTH
	info2.Gates[0].Center.Y = 2 * width
	info3.RangeX[0] = (defines.GRID_LINE_NUM - 1) * width
	info3.RangeX[1] = defines.GRID_LINE_NUM * width
	info3.RangeY[0] = (defines.GRID_LINE_NUM - 1) * width
	info3.RangeY[1] = defines.GRID_LINE_NUM * width
	info3.Gates[0] = &defines.GateInfo{}
	info3.Gates[0].GateType = defines.DEST4
	info3.Gates[0].Center.X = 2 * width
	info3.Gates[0].Center.Y = width*2 + defines.LANE_WIDTH + defines.SHELVES_WIDTH
	res.Warehouses = append(res.Warehouses, info)
	res.Warehouses = append(res.Warehouses, info0)
	res.Warehouses = append(res.Warehouses, info1)
	res.Warehouses = append(res.Warehouses, info2)
	res.Warehouses = append(res.Warehouses, info3)
	return res
}

func AddOneEdge(graph *defines.Graph, start int, end int, weight int) {
	graph.Edges[start] = append(graph.Edges[start], &defines.GraphEdge{
		End:    end,
		Weight: weight,
	})
	graph.Edges[end] = append(graph.Edges[end], &defines.GraphEdge{
		End:    start,
		Weight: weight,
	})
}

func GraphInit(mapInfo *defines.RoadMap) *defines.Graph {
	graph := &defines.Graph{}
	graph.Nodes = make([]*defines.GraphNode, 0)
	gateNum := 0
	for i := 0; i < len(mapInfo.Warehouses); i++ {
		gateNum += len(mapInfo.Warehouses[i].Gates)
	}
	graph.Edges = make([][]*defines.GraphEdge, len(mapInfo.Nodes)+gateNum)
	for i := 0; i < (len(mapInfo.Nodes) + gateNum); i++ {
		graph.Edges[i] = make([]*defines.GraphEdge, 0)
	}
	count := 0
	N2NDIS := defines.LANE_WIDTH*2 + defines.SHELVES_WIDTH
	N2GDIS := defines.LANE_WIDTH + defines.SHELVES_WIDTH
	// generate all nodes. (IDs)
	// and generate all Node-to-Node edges.
	for i := 0; i < defines.GRID_LINE_NUM; i++ {
		for j := 0; j < defines.GRID_LINE_NUM; j++ {
			node0 := &defines.GraphNode{}
			node1 := &defines.GraphNode{}
			node2 := &defines.GraphNode{}
			node3 := &defines.GraphNode{}
			node0.ID = count
			node0.X = mapInfo.Nodes[count].X
			node0.Y = mapInfo.Nodes[count].Y
			count++
			node1.ID = count
			node1.X = mapInfo.Nodes[count].X
			node1.Y = mapInfo.Nodes[count].Y
			count++
			node2.ID = count
			node2.X = mapInfo.Nodes[count].X
			node2.Y = mapInfo.Nodes[count].Y
			count++
			node3.ID = count
			node3.X = mapInfo.Nodes[count].X
			node3.Y = mapInfo.Nodes[count].Y
			count++
			graph.Nodes = append(graph.Nodes, node0)
			graph.Nodes = append(graph.Nodes, node1)
			graph.Nodes = append(graph.Nodes, node2)
			graph.Nodes = append(graph.Nodes, node3)
			graph.Edges[node0.ID] = append(graph.Edges[node0.ID], &defines.GraphEdge{
				End:    node1.ID,
				Weight: N2NDIS,
			})
			graph.Edges[node0.ID] = append(graph.Edges[node0.ID], &defines.GraphEdge{
				End:    node2.ID,
				Weight: N2NDIS,
			})
			graph.Edges[node1.ID] = append(graph.Edges[node1.ID], &defines.GraphEdge{
				End:    node0.ID,
				Weight: N2NDIS,
			})
			graph.Edges[node1.ID] = append(graph.Edges[node1.ID], &defines.GraphEdge{
				End:    node3.ID,
				Weight: N2NDIS,
			})
			graph.Edges[node2.ID] = append(graph.Edges[node2.ID], &defines.GraphEdge{
				End:    node0.ID,
				Weight: N2NDIS,
			})
			graph.Edges[node2.ID] = append(graph.Edges[node2.ID], &defines.GraphEdge{
				End:    node3.ID,
				Weight: N2NDIS,
			})
			graph.Edges[node3.ID] = append(graph.Edges[node3.ID], &defines.GraphEdge{
				End:    node1.ID,
				Weight: N2NDIS,
			})
			graph.Edges[node3.ID] = append(graph.Edges[node3.ID], &defines.GraphEdge{
				End:    node2.ID,
				Weight: N2NDIS,
			})
		}
	}
	for i := 0; i < gateNum; i++ {
		// src gates first.
		node := &defines.GraphNode{}
		node.ID = count
		count++
		if i < len(mapInfo.Warehouses[0].Gates) {
			node.X = mapInfo.Warehouses[0].Gates[i].Center.X
			node.Y = mapInfo.Warehouses[0].Gates[i].Center.Y
		} else {
			node.X = mapInfo.Warehouses[i-3].Gates[0].Center.X
			node.Y = mapInfo.Warehouses[i-3].Gates[0].Center.Y
		}
		graph.Nodes = append(graph.Nodes, node)
	}
	// then generate all gate-to-node edges.
	idx := defines.GRID_LINE_NUM / 2
	numOneLevel := 4 * defines.GRID_LINE_NUM
	startNum := numOneLevel*idx + 4*(idx-1)
	size := len(mapInfo.Nodes)
	AddOneEdge(graph, size, startNum+3, N2GDIS)
	AddOneEdge(graph, size, startNum+6, N2GDIS)
	AddOneEdge(graph, size+1, startNum+6-numOneLevel, N2GDIS)
	AddOneEdge(graph, size+1, startNum+4, N2GDIS)
	AddOneEdge(graph, size+2, startNum+8, N2GDIS)
	AddOneEdge(graph, size+2, startNum+5, N2GDIS)
	AddOneEdge(graph, size+3, startNum+5+numOneLevel, N2GDIS)
	AddOneEdge(graph, size+3, startNum+7, N2GDIS)
	AddOneEdge(graph, size+4, 3, N2GDIS)
	AddOneEdge(graph, size+4, 6, N2GDIS)
	AddOneEdge(graph, size+5, numOneLevel-2, N2GDIS)
	AddOneEdge(graph, size+5, numOneLevel*2-4, N2GDIS)
	AddOneEdge(graph, size+6, numOneLevel*(defines.GRID_LINE_NUM-1)+1, N2GDIS)
	AddOneEdge(graph, size+6, numOneLevel*(defines.GRID_LINE_NUM-2)+3, N2GDIS)
	AddOneEdge(graph, size+7, defines.GRID_LINE_NUM*defines.GRID_LINE_NUM*4-4, N2GDIS)
	AddOneEdge(graph, size+7, defines.GRID_LINE_NUM*defines.GRID_LINE_NUM*4-7, N2GDIS)
	return graph
}
