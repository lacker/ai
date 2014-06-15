package hex

import (
	"log"
)

/*
The shallow rave algorithm is that you do playouts from the given
position, and the spot with (roughly) the best win/loss record when
the player to move moves there is the best spot.
*/

type WinLossRecord struct {
	Wins int
	Losses int
}

// Predict future win-loss record with a Dirichlet distribution
func (r *WinLossRecord) Score() float64 {
	return float64(1 + r.Wins) / float64(1 + r.Wins + r.Losses)
}

type ShallowRave struct {
	NumPlayouts int
}

func (s *ShallowRave) Play(b *Board) Spot {
	records := make(map[Spot]*WinLossRecord)
	moves := b.PossibleMoves()
	for _, move := range moves {
		records[move] = new(WinLossRecord)
	}

	for i := 0; i < s.NumPlayouts; i++ {
		// To playout, first shuffle all possible moves
		moves := b.PossibleMoves()
		if len(moves) == 0 {
			log.Fatal("no possible moves")
		}
		ShuffleSpots(moves)

		// Then play moves in that order on a copy of the board.
		// Track the moves that "we" played, i.e. the player to move on b
		playout := b.Copy()
		ourMoves := make([]Spot, 0)
		for _, move := range moves {
			if playout.ToMove == b.ToMove {
				ourMoves = append(ourMoves, move)
			}
			if !playout.MakeMove(move) {
				log.Fatal("a playout somehow played an invalid move")
			}
		}

		winner := playout.Winner()
		if winner == Empty {
			playout.Eprint()
			log.Fatal("there was no winner after a full playout")
		}
		if winner == b.ToMove {
			// We won.
			for _, move := range ourMoves {
				records[move].Wins++
			}
		} else {
			// We lost.
			for _, move := range ourMoves {
				records[move].Losses++
			}
		}
	}

	// We have finished all the playouts. Now we just need to choose
	// the best-scoring move.
	bestScore := -1.0
	bestMove := Spot{-1, -1}
	for move, record := range records {
		if record.Score() > bestScore {
			bestScore = record.Score()
			bestMove = move
		}
	}
	if bestMove.Row == -1 {
		log.Fatal("there was no nonnegative score")
	}
	return bestMove
}

