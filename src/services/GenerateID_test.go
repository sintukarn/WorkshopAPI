package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateID(t *testing.T) {
	assert := assert.New(t)
	dev := "DEV"
	actualDev := GenerateID(dev)
	squad := "SQUAD"
	actualSquad := GenerateID(squad)
	bu := "BU"
	actualBU := GenerateID(bu)
	assert.Equal(actualDev[0:3], "DEV", "First 3 Char of result should be DEV")
	assert.True(len(actualDev) == 17, "Test DEV GenerateID Should be length 17")
	assert.Equal(actualSquad[0:5], "SQUAD", "First 5 Char of result should be SQUAD")
	assert.True(len(actualSquad) == 19, "Test Squad GenerateID Should be length 19")
	assert.Equal(actualBU[0:2], "BU", "First 2 Char of result should be BU")
	assert.True(len(actualBU) == 16, "Test BU GenerateID Should be length 16")
}
