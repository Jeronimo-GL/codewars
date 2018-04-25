package morse_advanced

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func Test_decodeBits( t *testing.T){
	morse_1 :="1100110011001100000011000000111111001100111111001111110000000000000011001111110011111100111111000000110011001111110000001111110011001100000011"

	morse_2 :="01100110011001100000011000000111111001100111111001111110000000000000011001111110011111100111111000000110011001111110000001111110011001100000011000"

	expected := "···· · −·−− ·−−− ··− −·· ·"
	assert.Equal(t , expected , morse_1, "Original message")
	assert.Equal(t , expected , morse_2, "Message with extra spaces")
}


func Test_splitMsg(t *testing.T){
	assert.Equal(t, []string{"0", "1", "0"}, splitMessage("010"), "010")
	assert.Equal(t, []string{"00", "11", "00"}, splitMessage("001100"), "001100")
	assert.Equal(t, []string{"0"}, splitMessage("0"), "0")
	assert.Equal(t, []string{"1"}, splitMessage("1"), "1")
	assert.Equal(t, []string{"111"}, splitMessage("111"), "111")
	assert.Equal(t, []string{"00", "11"}, splitMessage("0011"), "0011")
	assert.Equal(t, []string{"11", "00"}, splitMessage("1100"), "1100")
}
