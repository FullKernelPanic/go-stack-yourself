package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go-stack-yourself/src/roll/services"
)

func TestRollDice(t *testing.T) {
	a := assert.New(t)

	for i := 0; i < 10; i++ {
		result := services.RollDice("foo")
		a.True(result <= 6 && result >= 1, "RollDice was incorrect")
	}
}

func TestIfTheHouseWinsAlways(t *testing.T) {
	a := assert.New(t)
	for i := 0; i < 10; i++ {
		a.Equal(6, services.RollDice("FullKernelPanic"), "RollDice was incorrect")
	}
}
