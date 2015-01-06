package hex

import (
	"fmt"
)

// QFeature is a spot plus a nonempty color, packed into one byte.
type QFeature uint8

// Assuming BoardSize is 11:
// 0-120: black features
// 121-241: white features
// 242: not-a-feature
const MinFeature QFeature = 0
const MaxFeature QFeature = 2 * BoardSize * BoardSize - 1
const NumFeatures QFeature = MaxFeature + 1
const NotAFeature QFeature = NumFeatures


func (qf QFeature) Color() Color {
	switch qf / (BoardSize * BoardSize) {
	case 0:
		return Black
	case 1:
		return White
	}
	panic("control should not get here")
}

func (qf QFeature) Spot() TopoSpot {
	return TopoSpot(qf % (BoardSize * BoardSize)) + TopLeftCorner
}

func (qf QFeature) String() string {
	if qf == NotAFeature {
		return "NotAFeature"
	}
	return fmt.Sprintf("%v%v", qf.Color(), qf.Spot())
}

func MakeQFeature(color Color, spot TopoSpot) QFeature {
	if spot < TopLeftCorner {
		panic("cannot make qfeature from non-spot")
	}
	if color == Empty {
		panic("cannot make qfeature from empty color")
	}
	answer := QFeature(spot - TopLeftCorner)
	if color == White {
		answer += BoardSize * BoardSize
	}
	return answer
}
