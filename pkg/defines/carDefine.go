package defines

type Pos struct {
	X int
	Y int
}

type Car struct {
	CurX     int
	CurY     int
	Dir      int // UP LEFT DOWN RIGHT
	Start    Pos // start position
	Dest     Pos // destination shelves (in one warehouse) position
	DestType int // DEST1/2/3/4, four destination in our map.
}
