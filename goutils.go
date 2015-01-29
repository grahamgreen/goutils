package goutils

import (
	"errors"
	"log"
)

func Check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func NotEmpty(s string) {
	if len(s) == 0 {
		log.Fatal(errors.New("empty string"))
	}
}

func DefaultIfEmpty(s, dflt string) string {
	if len(s) == 0 {
		return dflt
	}
	return s
}
