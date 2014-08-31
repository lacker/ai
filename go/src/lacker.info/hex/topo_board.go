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

// Black goes TopSide to BottomSide
const TopSide TopoSpot = 0
const BottomSide TopoSpot = 1

// White goes LeftSide to RightSide
const LeftSide TopoSpot = 2
const RightSide TopoSpot = 3

const NotASpot TopoSpot = -1

const TopLeftCorner TopoSpot = 4
const NumTopoSpots TopoSpot = TopoSpot(BoardSize * BoardSize) + TopLeftCorner
const BottomRightCorner TopoSpot = NumTopoSpots - 1


func (s TopoSpot) isOnLeftSide() bool {
	return s % BoardSize == TopLeftCorner
}

func (s TopoSpot) isOnTopSide() bool {
	return s >= TopLeftCorner && s < TopLeftCorner + BoardSize
}

func (s TopoSpot) isOnBottomSide() bool {
	return s <= BottomRightCorner && s > BottomRightCorner - BoardSize
}

func (s TopoSpot) isOnRightSide() bool {
	return s % BoardSize == TopLeftCorner - 1
}

type TopoBoard struct {
	// Contents of the board, indexed by TopoSpot
	Board [NumTopoSpots]Color

	// GroupId gives the index in GroupSpots for the group
	// that a particular spot is in.
	// The GroupId can be a TopoSpot because there are at most that many
	// groups.
	GroupId [NumTopoSpots]TopoSpot

	// Each group is a list of TopoSpots in that group.
	GroupSpots [][]TopoSpot

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
}

func (b *TopoBoard) mergeSmallGroupIntoBigGroup(
	smallGroupId TopoSpot,
	bigGroupId TopoSpot) {

	// Fix the id mapping
	for s := range b.GroupSpots[smallGroupId] {
		b.GroupId[s] = bigGroupId
	}

	// Fix the spots lists
	b.GroupSpots[bigGroupId] = append(
		b.GroupSpots[bigGroupId],
		(b.GroupSpots[smallGroupId])...)
	b.GroupSpots[smallGroupId] = nil
}

// Looks at two spots, assuming they are adjacent, and merges their
// groups if they should be merged.
func (b *TopoBoard) maybeMergeAdjacentSpots(spot1 TopoSpot, spot2 TopoSpot) {
	color1 := b.Board[spot1]
	color2 := b.Board[spot2]
	if color1 == Empty {
		return
	}
	if color1 != color2 {
		return
	}

	group1 := b.GroupId[spot1]
	group2 := b.GroupId[spot2]
	if group1 == group2 {
		return
	}

	if len(b.GroupSpots[group1]) > len(b.GroupSpots[group2]) {
		b.mergeSmallGroupIntoBigGroup(group2, group1)
	} else {
		b.mergeSmallGroupIntoBigGroup(group1, group2)
	}

	// Check win conditions
	if b.GroupId[TopSide] == b.GroupId[BottomSide] {
		b.Winner = Black
	}
	if b.GroupId[LeftSide] == b.GroupId[RightSide] {
		b.Winner = White
	}
}

func NewTopoBoard() *TopoBoard {
	b := &TopoBoard{ToMove: Black}

	// Set up the initial groups for special spots
	b.addNewGroup(TopSide, Black)
	b.addNewGroup(BottomSide, Black)
	b.addNewGroup(LeftSide, White)
	b.addNewGroup(RightSide, White)

	b.GroupSpots = [][]TopoSpot{}

	return b
}

func TopoSpotFromRowCol(row int, col int) TopoSpot {
	return TopoSpot(4 + col + BoardSize * row)
}

func (b *TopoBoard) Set(row int, col int, color Color) {
	s := TopoSpotFromRowCol(row, col)
	b.addNewGroup(s, color)

	// Update connectivity with neighbors

	// Up-left neighbor
	if s.isOnTopSide() {
		b.maybeMergeAdjacentSpots(s, TopSide)
	} else {
		b.maybeMergeAdjacentSpots(s, s - BoardSize)

		// Up-right neighbor
		if !s.isOnRightSide() {
			b.maybeMergeAdjacentSpots(s, s - BoardSize + 1)
		}
	}

	// Left neighbor
	if s.isOnLeftSide() {
		b.maybeMergeAdjacentSpots(s, LeftSide)
	} else {
		b.maybeMergeAdjacentSpots(s, s - 1)
	}

	// Right neighbor
	if s.isOnRightSide() {
		b.maybeMergeAdjacentSpots(s, RightSide)
	} else {
		b.maybeMergeAdjacentSpots(s, s + 1)
	}

	// Bottom-right neighbor
	if s.isOnBottomSide() {
		b.maybeMergeAdjacentSpots(s, BottomSide)
	} else {
		b.maybeMergeAdjacentSpots(s, s + BoardSize)

		// Bottom-left neighbor
		if !s.isOnLeftSide() {
			b.maybeMergeAdjacentSpots(s, s + BoardSize - 1)
		}
	}
}
