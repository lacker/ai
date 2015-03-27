package meta

import (
	"log"
)

func AssertEq(expected string, actual string) {
	if expected != actual {
		log.Fatalf("expected [%s] but got [%s]", expected, actual)
	}
}
