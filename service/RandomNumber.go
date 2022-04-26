package service

import (
	"fmt"
	"github.com/alt-golang/logger"
	"math/rand"
)

type RandomNumber struct {
	Logger  logger.Logger
	Maximum int
}

func (randomNumber RandomNumber) Get() int {
	randomNumber.Logger.Trace("Generating random number between 0 and " + fmt.Sprint(randomNumber.Maximum))
	random := rand.Intn(randomNumber.Maximum)
	randomNumber.Logger.Trace("Result is: " + fmt.Sprint(random))
	return random
}
