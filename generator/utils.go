package generator

import (
	"math/rand"
	"time"
)

func RanNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
