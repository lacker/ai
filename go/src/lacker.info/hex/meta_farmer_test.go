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

func TestMetaFarmerWithDemocracyOnDoomed1(t *testing.T) {
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

func TestMetaFarmerWithDeltaNetOnDoomed3(t *testing.T) {
	board := PuzzleMap["doomed3"].Board.ToTopoBoard()
	mf := &MetaFarmer{Seconds:-1, Quiet:true, QuickType:"deltanet"}
	mf.init(board)
	mf.PlayOneCycle(false)
	mf.PlayOneCycle(false)
	mf.PlayOneCycle(false)
}

func TestMetaFarmerWithDeltaNetOnDoomed6(t *testing.T) {
	board := PuzzleMap["doomed3"].Board.ToTopoBoard()
	mf := &MetaFarmer{Seconds:-1, Quiet:true, QuickType:"deltanet"}
	mf.init(board)
	mf.PlayOneCycle(false)
	mf.PlayOneCycle(false)
	mf.PlayOneCycle(false)
	mf.PlayOneCycle(false)
	mf.PlayOneCycle(false)
	mf.PlayOneCycle(false)
	mf.PlayOneCycle(false)
	mf.PlayOneCycle(false)
}

func BenchmarkDeltaNetDoomed3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		player := GetPlayer("dn5")
		puzzle := GetPuzzle("doomed3")
		_, conf := player.Play(puzzle.Board)
		if conf != 0.0 {
			log.Fatal("the player doesn't even realize doomed3 is doomed")
		}
	}
}
