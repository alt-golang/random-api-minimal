package service

import (
	"github.com/alt-golang/logger"
	"testing"
)

func TestRandomNumbe(t *testing.T) {
	randomNumber := RandomNumber{
		Logger:  logger.GetLogger("github.com/alt-golang/random-api-minimal/service/RandomNumber"),
		Maximum: 9,
	}
	result := randomNumber.Get()
	if result < 0 || result > 9 {
		t.Errorf("result < 0 || result > 9: result is: %d", result)
	}
}
