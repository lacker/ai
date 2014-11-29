package hex

import (
	"log"
	"testing"
)

func TestMetaFarmerInit(t *testing.T) {
	board := NewTopoBoard()
	mf := &MetaFarmer{Seconds:-1, Quiet:true, QuickType:"democracy"}
	mf.init(board)
}

func TestMetaFarmerIntegration(t *testing.T) {
	board := PuzzleMap["doomed1"].Board.ToTopoBoard()
	mf := &MetaFarmer{Seconds:-1, Quiet:true, QuickType:"democracy"}
	mf.init(board)
	mf.PlayOneCycle(false)
	mf.PlayOneCycle(false)
	mf.PlayOneCycle(false)

	realMainLine := Playout(mf.whitePlayer, mf.blackPlayer, false)
	AssertHistoriesEqual(mf.mainLine.History, realMainLine.History)

	mf.PlayOneCycle(false)
	realMainLine = Playout(mf.whitePlayer, mf.blackPlayer, false)
	if mf.mainLine.Winner != realMainLine.Winner {
		log.Printf("metafarmer main line:\n")
		mf.mainLine.Debug()
		log.Printf("main line from actual players:\n")
		realMainLine.Debug()
		log.Fatal("main line got corrupted")
	}
}
