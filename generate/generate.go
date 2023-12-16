package generate

import (
	"math/rand"
)

// base62 字元集
const base62Letter = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// short Id 長度
const shortIdLen = 6

func ShortId() string {
	id := make([]byte, shortIdLen)
	for i := range id {
		id[i] = base62Letter[rand.Intn(len(base62Letter))]
	}
	return string(id)
}
