package RomanDecoder

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func Test_examples(t *testing.T){
	assert.Equal(t, 21, Decode("XXI"), "XXI")
	assert.Equal(t, 1, Decode("I"), "I")
	assert.Equal(t, 4, Decode("IV"), "IV")
	assert.Equal(t, 2008, Decode("MMVIII"), "MMVIII")
	assert.Equal(t, 1666, Decode("MDCLXVI"), "MDCLXVI")
}

func Test_complete(t *testing.T){
	assert.Equal(t, 3, Decode("III"), "III")
	assert.Equal(t, 4, Decode("IV"), "IV")
	assert.Equal(t, 5, Decode("V"), "V")
	assert.Equal(t, 6, Decode("VI"), "VI")
	assert.Equal(t, 9, Decode("IX"), "IX")
	assert.Equal(t, 10, Decode("X"), "X")
	assert.Equal(t, 20, Decode("XX"), "XX")
	assert.Equal(t, 12, Decode("XII"), "XII")
	assert.Equal(t, 14, Decode("XIV"), "XIV")
	assert.Equal(t, 16, Decode("XVI"), "XVI")
	assert.Equal(t, 19, Decode("XIX"), "XIX")
	assert.Equal(t, 40, Decode("XL"), "XL")
	assert.Equal(t, 44, Decode("XLIV"), "XLIV")
	assert.Equal(t, 49, Decode("XLIX"), "XLIX")
	assert.Equal(t, 50, Decode("L"), "L")
	assert.Equal(t, 70, Decode("LXX"), "LXX")
	assert.Equal(t, 90, Decode("XC"), "XC")
	assert.Equal(t, 100, Decode("C"), "C")
	assert.Equal(t, 400, Decode("CD"), "CD")
	assert.Equal(t, 500, Decode("D"), "D")
	assert.Equal(t, 900, Decode("CM"), "CM")
	assert.Equal(t, 901, Decode("CMI"), "CMI")
	assert.Equal(t, 1000, Decode("M"), "M")
	assert.Equal(t, 1002, Decode("MII"), "MII")
}
