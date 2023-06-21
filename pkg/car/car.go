package car

import "iot/pkg/defines"
import "math"

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

func ScheduleOnePath(mapInfo *defines.RoadMap, carInfo *defines.Car, start *defines.Pos, dest *defines.Pos) []*defines.Pos {
	res := make([]*defines.Pos, 0)
	return res
}

// ScheduleOneCar start and dest pos should be center pos of one car.
func ScheduleOneCar(mapInfo *defines.RoadMap, carInfo *defines.Car, start *defines.Pos, dest *defines.Pos) []*defines.Pos {
	// the start pos should be inside the src warehouse!
	res := make([]*defines.Pos, 0)
	if !CheckInWareHouse(mapInfo.Warehouses[0], start) {
		return res
	}
	carInfo.Start = *start
	carInfo.Dest = *dest
	// schedule from init pos to one gate of src warehouse.

	// schedule from warehouse gate to the start pos.

	// schedule from warehouse back to the gate.
	// NOTE: his may be different from the previous step as there may be conflict between cars' paths.

	// schedule from src gate to the dest warehouse gate.

	// schedule from dest gate to the final dest inside the warehouse.

	// back to dest warehouse gate.

	// back to the init pos.

	return res
}
