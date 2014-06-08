package hex

/*
Player is an interface for a hex player.
*/

type Player interface {
	Play(b *Board) Spot
}
