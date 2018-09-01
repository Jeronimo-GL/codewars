package IntegerPartitions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_examples(t *testing.T) {
	assert.Equal(t, 1, Partitions(0), "p(0)")
	assert.Equal(t, 1, Partitions(1), "p(1)")
	assert.Equal(t, 2, Partitions(2), "p(2)")
	assert.Equal(t, 3, Partitions(3), "p(3)")
	assert.Equal(t, 5, Partitions(4), "p(4)")
	assert.Equal(t, 7, Partitions(5), "p(5)")
	assert.Equal(t, 11, Partitions(6), "p(6)")
	assert.Equal(t, 15, Partitions(7), "p(7)")
	assert.Equal(t, 22, Partitions(8), "p(8)")
	assert.Equal(t, 1958, Partitions(25), "p(25)")
	assert.Equal(t, 8349, Partitions(32), "p(32)") 
//	assert.Equal(t, 173525, Partitions(49), "p(49)")
}

func Test_next(t *testing.T) {
	assert.Equal(t, -1, next(1), "next(1)")
	assert.Equal(t, 2, next(-1), "next(-1)")
	assert.Equal(t, -2, next(2), "next(2)")
	assert.Equal(t, 3, next(-2), "next(-2)")
}


func Test_sign(t *testing.T){
	assert.Equal(t, 1, sign(1), "sign(1)")
	assert.Equal(t, 1, sign(-1), "sign(-1)")
	assert.Equal(t, -1, sign(2), "sign(2)")
	assert.Equal(t, -1, sign(-2), "sign(-2)")
	assert.Equal(t, 1, sign(3), "sign(3)")
	assert.Equal(t, 1, sign(-3), "sign(-3)")
}

func Test_sub(t *testing.T){
	assert.Equal(t, 1, sub(1), "sub(1)")
	assert.Equal(t, 2, sub(-1), "sub(-1)")
	assert.Equal(t, 5, sub(2), "sub(2)")
	assert.Equal(t, 7, sub(-2), "sub(-2)")
	assert.Equal(t, 12, sub(3), "sub(3)")
	assert.Equal(t, 15, sub(-3), "sub(-3)")
}
