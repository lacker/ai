package hex

import (
	"encoding/json"
	"log"
)

// Generic json encoder
func ToJSON(b interface{}) string {
	j, err := json.Marshal(b)
	if err != nil {
		log.Fatal("could not encode object", err)
	}
	return string(j[:])
}

