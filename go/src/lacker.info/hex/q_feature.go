package hex

import (
	"fmt"
)

// QFeature is a concise way of representing a spot plus a nonempty color.
// The first bit is color (0=black, 1=white) and the last 7 bits are spot.
type QFeature uint8

// Since 0 isn't a valid spot, 0 can be used as "not a feature".
const NotAFeature QFeature = 0

func (qf QFeature) Color() Color {
	switch qf >> 7 {
	case 0:
		return Black
	case 1:
		return White
	}
	panic("control should not get here")
}

func (qf QFeature) Spot() TopoSpot {
	return TopoSpot(qf & 127)
}

func (qf QFeature) String() string {
	if qf == 0 {
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
	answer := QFeature(spot)
	if color == White {
		answer += 128
	}
	return answer
}
