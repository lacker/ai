package hex

import (
	"fmt"
)

type NaiveSpot struct {
	row, col int
}

func (s NaiveSpot) Row() int {
	return s.row
}

func (s NaiveSpot) Col() int {
	return s.col
}

func (s NaiveSpot) IsNotASpot() bool {
	return s.Row() < 0 || s.Row() >= BoardSize ||
		s.Col() < 0 || s.Col() >= BoardSize
}

func (s NaiveSpot) NaiveSpot() NaiveSpot {
	return s
}

func (s NaiveSpot) TopoSpot() TopoSpot {
	return MakeTopoSpot(s.Row(), s.Col())
}

func MakeNaiveSpot(row int, col int) NaiveSpot {
	return NaiveSpot{row: row, col: col}
}

func AllSpots() [NumSpots]NaiveSpot {
	var answer [NumSpots]NaiveSpot
	for r := 0; r < BoardSize; r++ {
		for c := 0; c < BoardSize; c++ {
			spot := MakeNaiveSpot(r, c)
			answer[spot.Index()] = spot
		}
	}
	return answer
}

func (s NaiveSpot) Index() int {
	return s.Col() + BoardSize * s.Row()
}

func (s NaiveSpot) String() string {
	return fmt.Sprintf("(%d, %d)", s.Row(), s.Col())
}

func (s NaiveSpot) Transpose() NaiveSpot {
	return MakeNaiveSpot(s.Col(), s.Row())
}

func (s NaiveSpot) ApplyToNeighbors(f func(NaiveSpot)) {
	if s.Row() > 0 {
		f(MakeNaiveSpot(s.Row() - 1, s.Col()))
	}
	if s.Row() + 1 < BoardSize {
		f(MakeNaiveSpot(s.Row() + 1, s.Col()))
		if s.Col() > 0 {
			f(MakeNaiveSpot(s.Row() + 1, s.Col() - 1))
		}
	}
	if s.Col() > 0 {
		f(MakeNaiveSpot(s.Row(), s.Col() - 1))
	}
	if s.Col() + 1 < BoardSize {
		f(MakeNaiveSpot(s.Row(), s.Col() + 1))
		if s.Row() > 0 {
			f(MakeNaiveSpot(s.Row() - 1, s.Col() + 1))
		}
	}
}

func (s NaiveSpot) Neighbors() []NaiveSpot {
	answer := make([]NaiveSpot, 0)
	possible := []NaiveSpot{
		NaiveSpot{s.Row() - 1, s.Col()},
		NaiveSpot{s.Row() + 1, s.Col()},
		NaiveSpot{s.Row(), s.Col() - 1},
		NaiveSpot{s.Row(), s.Col() + 1},
		NaiveSpot{s.Row() + 1, s.Col() - 1},
		NaiveSpot{s.Row() - 1, s.Col() + 1},
	}
	for _, spot := range possible {
		if spot.IsNotASpot() {
			continue
		}
		answer = append(answer, spot)
	}
	return answer
}


