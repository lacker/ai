package hex

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
)

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

func (s TopoSpot) isSpecialSpot() bool {
	return s < TopLeftCorner
}

func (s TopoSpot) ToSpot() NaiveSpot {
	if s < TopLeftCorner {
		panic("special spots cannot be converted to NaiveSpot")
	}
	x := int(s - TopLeftCorner)
	col := x % BoardSize
	row := (x - col) / BoardSize
	return NaiveSpot{Row: row, Col: col}
}

func (s TopoSpot) Row() int {
	return s.ToSpot().Row
}

func (s TopoSpot) Col() int {
	return s.ToSpot().Col
}

func (s TopoSpot) String() string {
	return fmt.Sprintf("(%d, %d)", s.Row(), s.Col())
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

	// The spots that led the winner to win
	WinningPathSpots []TopoSpot

	// All the moves made in this game
	History []TopoSpot
}

// Adds a group of a single spot. Does not merge with any neighbors.
func (b *TopoBoard) addNewGroup(s TopoSpot, color Color) {
	if b.Board[s] != Empty {
		b.Eprint()
		fmt.Printf("TopoSpot to set: %d\n", s)
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
	for _, s := range b.GroupSpots[smallGroupId] {
		b.GroupId[s] = bigGroupId
	}

	// Fix the spots lists
	b.GroupSpots[bigGroupId] = append(
		b.GroupSpots[bigGroupId],
		(b.GroupSpots[smallGroupId])...)
	b.GroupSpots[smallGroupId] = nil
}

// Looks at two spots, assuming they are connected in reality but that
// may not be reflected in the groups, and merges their groups if they
// should be merged.
func (b *TopoBoard) maybeMergeSpots(spot1 TopoSpot, spot2 TopoSpot) {
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
		b.WinningPathSpots = b.GroupSpots[b.GroupId[TopSide]]
	}
	if b.GroupId[LeftSide] == b.GroupId[RightSide] {
		b.Winner = White
		b.WinningPathSpots = b.GroupSpots[b.GroupId[LeftSide]]
	}
}

func NewTopoBoard() *TopoBoard {
	b := &TopoBoard{ToMove: Black}

	b.GroupSpots = [][]TopoSpot{}
	b.History = make([]TopoSpot, 0)

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

func TopoSpotFromSpot(s NaiveSpot) TopoSpot {
	return TopoSpotFromRowCol(s.Row, s.Col)
}

func (b *TopoBoard) Get(s NaiveSpot) Color {
	return b.GetByRowCol(s.Row, s.Col)
}

func (b *TopoBoard) GetTopoSpot(s TopoSpot) Color {
	return b.Board[s]
}

func (b *TopoBoard) GetByRowCol(row int, col int) Color {
	s := TopoSpotFromRowCol(row, col)
	return b.Board[s]
}

func (b *TopoBoard) ToNaiveBoard() *NaiveBoard {
	c := NewNaiveBoard()
	c.ToMove = b.ToMove
	for _, spot := range AllSpots() {
		c.Set(spot, b.Get(spot))
	}
	return c
}

func (b *TopoBoard) ToTopoBoard() *TopoBoard {
	return b.ToNaiveBoard().ToTopoBoard()
}

func (b *TopoBoard) Copy() Board {
	return b.ToTopoBoard()
}

func (b *TopoBoard) Eprint() {
	b.ToNaiveBoard().Eprint()
}

// The number of non-empty groups.
func (b *TopoBoard) NumGroups() int {
	answer := 0
	for _, group := range b.GroupSpots {
		if group != nil {
			answer++
		}
	}
	return answer
}

func (b *TopoBoard) Set(row int, col int, color Color) {
	s := TopoSpotFromRowCol(row, col)
	b.SetTopoSpot(s, color)
}

// Cannot set things to empty or change the color of stones
func (b *TopoBoard) SetTopoSpot(s TopoSpot, color Color) {
	b.addNewGroup(s, color)

	// Update connectivity with neighbors

	// Up-left neighbor
	if s.isOnTopSide() {
		b.maybeMergeSpots(s, TopSide)
	} else {
		b.maybeMergeSpots(s, s - BoardSize)

		// Up-right neighbor
		if !s.isOnRightSide() {
			b.maybeMergeSpots(s, s - BoardSize + 1)
		}
	}

	// Left neighbor
	if s.isOnLeftSide() {
		b.maybeMergeSpots(s, LeftSide)
	} else {
		b.maybeMergeSpots(s, s - 1)
	}

	// Right neighbor
	if s.isOnRightSide() {
		b.maybeMergeSpots(s, RightSide)
	} else {
		b.maybeMergeSpots(s, s + 1)
	}

	// Bottom-right neighbor
	if s.isOnBottomSide() {
		b.maybeMergeSpots(s, BottomSide)
	} else {
		b.maybeMergeSpots(s, s + BoardSize)

		// Bottom-left neighbor
		if !s.isOnLeftSide() {
			b.maybeMergeSpots(s, s + BoardSize - 1)
		}
	}
}

// Returns a zobrist hash of the board state.
var blackZobrist [NumTopoSpots]int64
var whiteZobrist [NumTopoSpots]int64
var zobristInitialized bool = false
func (b TopoBoard) Zobrist() int64 {
	var spot TopoSpot
	if !zobristInitialized {
		for spot = TopLeftCorner; spot <= BottomRightCorner; spot++ {
			blackZobrist[spot] = rand.Int63()
			whiteZobrist[spot] = rand.Int63()
		}
	}
	var answer int64 = 0
	for spot = TopLeftCorner; spot <= BottomRightCorner; spot++ {
		switch b.Board[spot] {
		case Black:
			answer ^= blackZobrist[spot]
		case White:
			answer ^= whiteZobrist[spot]
		}
	}
	return answer
}

func (b *TopoBoard) PossibleTopoSpotMoves() []TopoSpot {
	answer := make([]TopoSpot, 0)
	var spot TopoSpot
	for spot = 0; spot < NumTopoSpots; spot++ {
		color := b.Board[spot]
		if color == Empty {
			answer = append(answer, spot)
		}
	}
	return answer
}

func (b *TopoBoard) PossibleMoves() []NaiveSpot {
	topo := b.PossibleTopoSpotMoves()
	answer := make([]NaiveSpot, len(topo))
	for i, s := range topo {
		answer[i] = s.ToSpot()
	}
	return answer
}

func (b *TopoBoard) MakeMoveWithNaiveSpot(s NaiveSpot) {
	b.MakeMove(TopoSpotFromSpot(s))
}

func (b *TopoBoard) MakeMove(s TopoSpot) {
	if b.ToMove == Empty {
		log.Fatal("this isn't a valid topo board, there is nobody to move")
	}
	b.SetTopoSpot(s, b.ToMove)
	b.ToMove = -b.ToMove
	b.History = append(b.History, s)
}

// Makes moves repeatedly. When this stops the game is over.
// Returns the winner.
// This mutates the board.
func (b *TopoBoard) Playout() Color {
	moves := b.PossibleMoves()
	ShuffleSpots(moves)

	for _, move := range moves {
		b.MakeMoveWithNaiveSpot(move)
		if b.Winner != Empty {
			return b.Winner
		}
	}

	panic("played all moves and still no winner")
}

func (b *TopoBoard) GetToMove() Color {
	return b.ToMove
}

func (b *TopoBoard) GetWinningPathSpots() []NaiveSpot {
	if b.Winner == Empty {
		panic("cannot GetWinningPathSpots with no winner")
	}
	answer := make([]NaiveSpot, 0)
	for _, spot := range b.WinningPathSpots {
		if spot.isSpecialSpot() {
			continue
		}
		answer = append(answer, spot.ToSpot())
	}
	return answer
}

// Log the state of the board
func (b *TopoBoard) Log() {
	log.Printf("%s to move\n", b.GetToMove().Name())
	for r := 0; r < BoardSize; r++ {
		line := strings.Repeat(" ", r)
		for c := 0; c < BoardSize; c++ {
			switch b.GetByRowCol(r, c) {
			case Black:
				line += "B"
			case White:
				line += "W"
			case Empty:
				line += "."
			}
			if c == BoardSize - 1 {
				line += "\n"
			} else {
				line += " "
			}
		}
		log.Printf(line)
	}
}
