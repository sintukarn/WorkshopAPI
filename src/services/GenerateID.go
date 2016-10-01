package services

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateID(kind string) string {
	output := fmt.Sprintf("%s%v%v", kind, time.Now().Unix(), rand.Intn(9999))
	return output
}
