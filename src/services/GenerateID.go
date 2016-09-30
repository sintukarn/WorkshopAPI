package services

import (
	"time"
	"math/rand"
	"fmt"
)

func GenerateID(kind string)(string){
	output := fmt.Sprintf("%s%v%v",kind,time.Now().Unix(),rand.Intn(9999))
	return output
}
