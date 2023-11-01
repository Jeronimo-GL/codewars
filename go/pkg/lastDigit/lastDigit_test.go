package lastDigit

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func Test_examples(t *testing.T){
	assert.Equal(t, 1, LastDigit([]int{}), "Empty list")
	assert.Equal(t, 1, LastDigit([]int{0,0}), "0 ^0") // 0 ^ 0
    assert.Equal(t, 0, LastDigit( []int{0,0,0}), "0^0^0") // 0^(0 ^ 0) = 0^1 = 0
    assert.Equal(t, 1, LastDigit( []int{1,2}), "1^2")
    assert.Equal(t, 1, LastDigit( []int{3,4,5}), "3^4^5")
    assert.Equal(t, 4, LastDigit( []int{4,3,6}), "3^4^6")
    assert.Equal(t, 1, LastDigit( []int{7,6,21}), "7^6^21")
    assert.Equal(t, 6, LastDigit( []int{12,30,21}), "12^30^12")
    assert.Equal(t, 1, LastDigit( []int{2,0,1}), "2^0^1")
    assert.Equal(t, 4, LastDigit( []int{2,2,2,0}), "2^2^2^0")
    assert.Equal(t, 0, LastDigit( []int{937640,767456,981242}), "Very big 1")
    assert.Equal(t, 6, LastDigit( []int{123232,694022,140249}), "Very big 2")
    assert.Equal(t, 6, LastDigit( []int{499942,898102,846073}), "Very big 3")
}
