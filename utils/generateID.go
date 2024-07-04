package utils

import (
	"fmt"
	"math/rand"
)

func GenerateID() string {
	return fmt.Sprintf("P%04d", rand.Intn(10000))
}
