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

It starts at 4 because 0-3 are taken up by the special spots.
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

	// NextSpot maps each spot to the next spot in its same group.
	// This should create a singly linked loop for each group.
	// A group is a set of spots that are all connected.
	// For empty spots, it doesn't matter what these point to.
	NextSpot [NumTopoSpots]TopoSpot

	// GroupSize maps each spot to the size of its group.
	// An empty spot is defined to be a group of size zero.
	GroupSize [NumTopoSpots]uint8

	// Whose move it is
	ToMove Color

	// Who has won the game
	Winner Color
}

func NewTopoBoard() *TopoBoard {
	b := &TopoBoard{ToMove: Black}

	// Board starts off with zeros which are Empty so only the special
	// spots need to be set
	b.Board[TopSide] = Black
	b.Board[BottomSide] = Black
	b.Board[LeftSide] = White
	b.Board[RightSide] = White

	// For empty spots, it doesn't matter where NextSpot points. But for
	// the special spots, they should be groups of size 1 and NextSpot
	// should point to themselves.
	b.NextSpot[TopSide] = TopSide
	b.NextSpot[BottomSide] = BottomSide
	b.NextSpot[LeftSide] = LeftSide
	b.NextSpot[RightSide] = RightSide

	// Set GroupSize for special spots
	b.GroupSize[TopSide] = 1
	b.GroupSize[BottomSide] = 1
	b.GroupSize[LeftSide] = 1
	b.GroupSize[RightSide] = 1

	return b
}

func TopoSpotFromRowCol(row int, col int) TopoSpot {
	return TopoSpot(4 + col + BoardSize * row)
}

func (b *TopoBoard) Set(row int, col int, color Color) {

}
