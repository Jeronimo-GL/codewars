package coordinates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_samples(t *testing.T) {
	assert.True(t, IsValidCoordinates("-23, 25"), "Should be true for -23, 25")
	assert.True(t, IsValidCoordinates("4, -3"), "Should be true for 4, -3")
	assert.True(t, IsValidCoordinates("24.53525235, 23.45235"), "Should be true for 24.53525235, 23.45235")
	assert.True(t, IsValidCoordinates("04, -23.234235"), "Should be true for 04, -23.234235")
	assert.True(t, IsValidCoordinates("43.91343345, 143"), "Should be true for 43.91343345, 143")

	assert.False(t, IsValidCoordinates("23.234, - 23.4234"), "Should be false for 23.234, - 23.4234")
	assert.False(t, IsValidCoordinates("2342.43536, 34.324236"), "Should be false for 2342.43536, 34.324236")
	assert.False(t, IsValidCoordinates("N23.43345, E32.6457"), "Should be false for N23.43345, E32.6457")
	assert.False(t, IsValidCoordinates("99.234, 12.324"), "Should be false for 99.234, 12.324")
	assert.False(t, IsValidCoordinates("6.325624, 43.34345.345"), "Should be false for 6.325624, 43.34345.345")
	assert.False(t, IsValidCoordinates("0, 1,2"), "Should be false for 0, 1,2")
	assert.False(t, IsValidCoordinates("0.342q0832, 1.2324"), "Should be false for 0.342q0832, 1.2324")
	assert.False(t, IsValidCoordinates("23.245, 1e1"), "Should be false for 23.245, 1e1")
}

func Test_otherCases(t *testing.T) {
	assert.False(t, IsValidCoordinates(""), "Empty string")
	assert.False(t, IsValidCoordinates("   "), "Blank string")
	assert.False(t, IsValidCoordinates("1,2,3,4"), "More than two numbers in the string")
	assert.False(t, IsValidCoordinates("ñldskfj ñsdlkjf"), "Non numeric string")

	assert.False(t, IsValidCoordinates("-91.432432, 23.23432"), "Lower latitude")
	assert.False(t, IsValidCoordinates("91.432432, 23.23432"), "Higher latitude")
	assert.False(t, IsValidCoordinates("-1.432432, 223.23432"), "Higher latitude")
	assert.False(t, IsValidCoordinates("-1.432432, -223.23432"), "Lower latitude")

	assert.True(t, IsValidCoordinates("0, 0"), "Zero zero")
}
