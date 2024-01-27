package dummies

import (
	"math/rand"
	"time"
)

func DummyDctType() string {
	seed := []string{
		"flows_category",
		"pockets_type",
		"wishlists_type",
		"flows_category",
	}

	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(seed))
	res := seed[idx]

	return res
}
