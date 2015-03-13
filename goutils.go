package goutils

import (
	"errors"
	"log"
	"math"
)

type Point struct {
	X, Y int64
}

func (p Point) Distance(t Point) uint64 {
	dist = math.Sqrt(math.Pow((t.X-p.X), 2) + math.Pow((t.Y-p.Y), 2))
	return dist
}

func (p Point) TravelTime(t Point, s uint64) uint64 {
	dist = p.Distance(t)
	return dist / s
}

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
