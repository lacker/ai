package hex

/*
TopoBoard is a Board implementation that constantly tracks which
groups of spots are connected. This makes it quick to determine when a
game is over.

TopoSpot is a representation of a single spot on the board, that
includes four special spots: the top, bottom, left, and right of the
board.

The meaning of a particular TopoSpot goes like:
4  5  6
 7  8  9
  10 11 12

It starts at 4 because 0, 1, 2, 3 are taken up by the special spots.
*/

// This is going to bite us one day for larger boards. But since
// 11 * 11 < 128 it works.
type TopoSpot int8

const NumTopoSpots int = BoardSize * BoardSize + 4

// Black goes TopSide to BottomSide
const TopSide TopoSpot = 0
const BottomSide TopoSpot = 1

// White goes LeftSide to RightSide
const LeftSide TopoSpot = 2
const RightSide TopoSpot = 3

const NotASpot TopoSpot = -1


type TopoBoard struct {
	// Contents of the board, indexed by TopoSpot
	Board [NumTopoSpots]Color

	// Whose move it is
	ToMove Color
}


