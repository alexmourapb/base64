package generate

import (
	"log"
	"math/rand"
	"time"
)

const (
	Min = 3
	Max = 96
)


func GenerateRandom(pwdSize int)  string{

	SEED := time.Now().Unix()
	if pwdSize <= 0 {
		pwdSize= 8
		log.Println("Using default values!")
	}

	rand.Seed(SEED)
	startChar := "!"
	var newChar string
	var i int = 1
	for {
		myRand := random(Min, Max)
		newChar = newChar + string(startChar[0] + byte(myRand))
		//fmt.Print(NewChar)
		if i == pwdSize {
			break
		}
		i++
	}

	return newChar
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
