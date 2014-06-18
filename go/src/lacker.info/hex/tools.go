package hex

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

// Generic json encoder
func ToJSON(b interface{}) string {
	j, err := json.Marshal(b)
	if err != nil {
		log.Fatal("could not encode object", err)
	}
	return string(j[:])
}

func Intmax(x int, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

// Shuffles a list of spots
func ShuffleSpots(spots []Spot) {
	for i := range spots {
    j := rand.Intn(i + 1)
    spots[i], spots[j] = spots[j], spots[i]
	}
}

// Prints a string to stderr
func Eprint(s string) {
	fmt.Fprintf(os.Stderr, s)
}

func Seed() {
	rand.Seed(time.Now().UTC().UnixNano())
}
