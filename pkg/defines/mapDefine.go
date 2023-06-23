package defines

const (
	UP            = 0
	LEFT          = 1
	DOWN          = 2
	RIGHT         = 3
	DEST1         = 4
	DEST2         = 5
	DEST3         = 6
	DEST4         = 7
	SRC           = 8 // the warehouse to fetch all required objects.
	VERTICAL      = 9
	HORIZONTAL    = 10
	LANE_WIDTH    = 40   // cm
	SHELVES_WIDTH = 1000 // cm
	CAR_WIDTH     = 30   // cm
	GRID_NUM      = 9    // 3 * 3, every grid contains **4** nodes.
	GRID_LINE_NUM = 3
)

type MapNode struct {
	X int
	Y int
}

type LaneInfo struct {
	// info for every lane.
	LaneType int // VERTICAL / HORIZONTAL
	// nodes in vertical lane should have the same X
	// nodes in horizontal lane should have the same Y
	CenterPos int
	Nodes     []*MapNode
	Cars      []*Car
}

type GateInfo struct {
	// info about gate for every dest and start shelves.
	// center position of one gate. one gate should have the same width with one lane.
	// center should have the same x / y with one node in this lane.
	Center Pos
	// gateType means which warehouse owns this gate.
	GateType int
}

type WarehouseInfo struct {
	Gates  []*GateInfo
	RangeX []int // [lowX, highX]
	RangeY []int // [lowY, highY]
}

type RoadMap struct {
	Nodes      []*MapNode
	Lanes      []*LaneInfo
	Warehouses []*WarehouseInfo
}

type InitInfo struct {
	MapInfo   *RoadMap
	GraphInfo *Graph
}

type InputInfo struct {
	Init InitInfo
}

type StartBody struct {
	Input InputInfo `json:"input"`
}
