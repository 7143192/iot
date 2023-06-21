package roadMap

import (
	"iot/pkg/defines"
	"iot/pkg/roadMap"
)

func Init() *defines.RoadMap {
	res := roadMap.MapInit()
	// should store this map into cloud server's cache.
	return res
}
