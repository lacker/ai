package main

import (
	"fmt"
)

const BoardSize = 11;

// The Color pseudo-enum
const Black = -1;
const White = 1;
const Empty = 0;

/*
The board is a grid. Each value is either BLACK, WHITE, or EMPTY.
In the external world, spots are typically referred to with a
letter plus a number.
Display would look like a diamond, e.g. for a 4x4 board:

ABCD
-----
\    \       1
 \    \      2
  \    \     3
   \    \    4
    -----

To represent a spot, we do row and column like it's a matrix.
Black goes top to bottom; White goes left to right.
So Black could win with a single column; White could win with a single row.
*/

type Spot struct {
	Row, Col int8
}

func main() {
	fmt.Printf("sup\n");
}
