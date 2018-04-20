package johnAndAnn

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func Test_array(t *testing.T){	
	a := []int{1,2}
	b := []int{1,2,3}
	assert.EqualValues(t, a, b, "a2")
	assert.Equal(t, a, b, "Array")
}


func Test_j_base(t *testing.T){
	assert.Equal(t, 0, j(0), "in j(0)")
}

func Test_a_base(t *testing.T){
	assert.Equal(t, 1, a(0), "in a(0)")
}


func Test_j_and_a(t *testing.T){
	assert.Equal(t, 0, j(1), "in j(1)")
	assert.Equal(t, 1, a(1), "in a(1)")
	assert.Equal(t, 1, j(2), "in j(2)")
	assert.Equal(t, 2, a(2), "in a(2)")
}

func Test_John(t *testing.T){
	assert.Equal(t, []int{0}, John(0), "Initial case")
	
	expected := []int{0, 0, 1, 2, 2, 3, 4, 4, 5, 6, 6}
	assert.Equal(t, expected, John(11))
}

func Test_Ann(t *testing.T){
	assert.Equal(t, []int{1}, Ann(0), "Inital case")
	
	expected := []int {1, 1, 2, 2, 3, 3}
	assert.Equal(t, expected, Ann(6))
}
