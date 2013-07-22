package engine

type Loc struct {
	x int
	y int
}

func NewLoc(x,y int) *Loc {
	loc := new(Loc)
	loc.x = x
	loc.y = y
	return loc
}
