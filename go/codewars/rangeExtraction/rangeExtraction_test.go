package rangeExtraction

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_basic_cases(t *testing.T){
	assert.Equal(t, "-6,-3-1,3-5,7-11,14,15,17-20",
		Solution([]int{-6,-3,-2,-1,0,1,3,4,5,7,8,9,10,11,14,15,17,18,19,20}),
		"Sample case")
}

func Test_simple_cases(t *testing.T){
	assert.Equal(t, "3", Solution([]int{3}), "Single value list")
	assert.Equal(t, "", Solution([]int{}), "Empty list")
}

func Test_single_range(t *testing.T){
	assert.Equal(t, "1-3", Solution([]int{1,2,3}), "Single range")
}

func Test_generic_cases(t *testing.T) {
	assert.Equal(t, "-1,2-4", Solution([]int{-1,2,3,4}), "Value and range")
	assert.Equal(t, "2-4,7", Solution([]int{2,3,4,7}), "range and value")
	assert.Equal(t, "2-4,7-9", Solution([]int{2,3,4,7,8,9}), "two ranges")
}


func Test_just_is_not_a_range(t *testing.T) {
	assert.Equal(t, "1,2", Solution([]int{1,2}), "Just does not make a range")
}



