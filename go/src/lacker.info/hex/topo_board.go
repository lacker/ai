package hex

/*
TopoBoard is a Board implementation that constantly tracks which
groups of spots are connected. This makes it quick to determine when a
game is over.

TopoSpot is a representation of a single spot on the board, that
includes four special spots: the top, bottom, left, and right of the
board.

The meaning of a particular TopoSpot goes like:
0 1 2
 3 4 5
  6 7 8

except it's BoardSize by BoardSize. Indexes less than zero are the
special spots.
*/

// This is going to bite us one day for larger boards. But since
// 11 * 11 < 128 it works.
type TopoSpot int8

// Black goes TopSide to BottomSide
const TopSide TopoSpot = -1
const BottomSide TopoSpot = -2

// White goes LeftSide to RightSide
const LeftSide TopoSpot = -3
const RightSide TopoSpot = -4


type TopoBoard struct {
	// Contents of the board, indexed by TopoSpot
	Board [BoardSize * BoardSize]Color

	// Whose move it is
	ToMove Color
}


