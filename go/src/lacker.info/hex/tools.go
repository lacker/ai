package hex

import (
	"encoding/json"
	"log"
	"math/rand"
)

// Generic json encoder
func ToJSON(b interface{}) string {
	j, err := json.Marshal(b)
	if err != nil {
		log.Fatal("could not encode object", err)
	}
	return string(j[:])
}

// Shuffles a list of spots
func ShuffleSpots(spots []Spot) {
	for i := range spots {
    j := rand.Intn(i + 1)
    spots[i], spots[j] = spots[j], spots[i]
	}
}
