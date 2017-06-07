package random

import (
	"crypto/sha256"
	"math/rand"
	"time"
	"fmt"
)

func RandomString(length int) string {
	return fmt.Sprintf("%x", generate())[:length]
}

func generate() [32]byte {
	input := time.Now().UnixNano()
	salt := rand.Int63()
	sum := sha256.Sum256([]byte(fmt.Sprintf("%x%x", input, salt)))
	return sum
}
