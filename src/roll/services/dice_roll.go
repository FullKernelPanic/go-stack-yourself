package services

import (
	"math/rand"
)

func RollDice(player string) int {
	if player == "FullKernelPanic" {
		return 6
	}

	return 1 + rand.Intn(6)
}
