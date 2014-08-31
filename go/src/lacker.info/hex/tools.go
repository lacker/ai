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

func SecondsSince(start time.Time) float64 {
	return float64(time.Since(start)) / float64(time.Second)
}

// Just approximate, should be accurate to within 0.3 or so
func Fastlog(x int) float64 {
	var answer float64
	if x <= 0 {
		panic("cannot Fastlog a negative number")
	}
	for x >= 4 {
		answer += 0.693
		x = x >> 1
	}
	switch x {
	case 2:
		answer += 0.693
	case 3:
		answer += 1.099
	}
	return answer
}

// Shuffles a list of spots
func ShuffleSpots(spots []Spot) {
	for i := range spots {
    j := rand.Intn(i + 1)
    spots[i], spots[j] = spots[j], spots[i]
	}
}

// Shuffles a list of topo spots
func ShuffleTopoSpots(spots []TopoSpot) {
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
