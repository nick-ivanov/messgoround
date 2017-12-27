package main

import (
	"math/rand"
	"time"
)

// Returns truly random integer between 0 and upto
func trulyRandomInt(upto int) int {
	src := rand.NewSource(time.Now().UnixNano())
	return rand.New(src).Intn(upto)
}
