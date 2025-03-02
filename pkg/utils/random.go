package utils

import (
	"time"

	"golang.org/x/exp/rand"
)

func GenerateSixDigitsOTP() int {
	rng := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
	otp := 100000 + rng.Intn(900000)
	return otp
}
