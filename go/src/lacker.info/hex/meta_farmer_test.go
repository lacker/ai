package hex

import (
	"log"
	"testing"
)

func TestMetaFarmerInit(t *testing.T) {
	board := NewTopoBoard()
	mf := &MetaFarmer{Seconds:-1, Quiet:true}
	mf.init(board)
}

func TestMetaFarmerIntegration(t *testing.T) {
	board := PuzzleMap["doomed1"].Board.ToTopoBoard()
	mf := &MetaFarmer{Seconds:-1, Quiet:true}
	mf.init(board)
	mf.PlayOneCycle(false)
	mf.PlayOneCycle(false)
	mf.PlayOneCycle(false)
	mf.PlayOneCycle(false)
	realMainLine := Playout(mf.whitePlayer, mf.blackPlayer, false)
	if mf.mainLine.Winner != realMainLine.Winner {
		log.Fatal("main line got corrupted")
	}
}
