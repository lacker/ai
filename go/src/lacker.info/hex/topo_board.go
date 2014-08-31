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

	// GroupId gives the index in GroupSpots and GroupSize for the group
	// that a particular spot is in.
	// The GroupId can be a TopoSpot because there are at most that many
	// groups.
	GroupId [NumTopoSpots]TopoSpot

	// Each group is a list of TopoSpots in that group.
	GroupSpots [][]TopoSpot

	// The size for the group
	GroupSize []TopoSpot

	// Whose move it is
	ToMove Color

	// Who has won the game
	Winner Color
}

// Adds a group of a single spot. Does not merge with any neighbors.
func (b *TopoBoard) addNewGroup(s TopoSpot, color Color) {
	if b.Board[s] != Empty {
		panic("TopoBoard cannot change a spot once it has something on it")
	}
	if color == Empty {
		panic("TopoBoard cannot set a spot to Empty")
	}

	newGroupId := TopoSpot(len(b.GroupSpots))
	b.Board[s] = color
	b.GroupId[s] = newGroupId
	newGroup := []TopoSpot{s}
	b.GroupSpots = append(b.GroupSpots, newGroup)
	b.GroupSize = append(b.GroupSize, 1)
}

func NewTopoBoard() *TopoBoard {
	b := &TopoBoard{ToMove: Black}

	// Set up the initial groups for special spots
	b.addNewGroup(TopSide, Black)
	b.addNewGroup(BottomSide, Black)
	b.addNewGroup(LeftSide, White)
	b.addNewGroup(RightSide, White)

	return b
}

func TopoSpotFromRowCol(row int, col int) TopoSpot {
	return TopoSpot(4 + col + BoardSize * row)
}

func (b *TopoBoard) Set(row int, col int, color Color) {
	s := TopoSpotFromRowCol(row, col)
	b.addNewGroup(s, color)

	// Update connectivity with neighbors
	// TODO.
}
