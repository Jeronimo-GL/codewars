package IntegerPartitions

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_examples(t *testing.T){
	assert.Equal(t, 1, Partitions(0), "p(0)")
	assert.Equal(t, 1, Partitions(1), "p(1)")
	assert.Equal(t, 2, Partitions(2), "p(2)")
	assert.Equal(t, 3, Partitions(3), "p(3)")
	assert.Equal(t, 5, Partitions(4), "p(4)")
	assert.Equal(t, 7, Partitions(5), "p(5)")
	assert.Equal(t, 11, Partitions(6), "p(6)")
	assert.Equal(t, 15, Partitions(7), "p(7)")
	assert.Equal(t, 22, Partitions(8), "p(8)")
}

func Test_next(t *testing.T){
	assert.Equal(t, -1, next(1), "next(1)")
	assert.Equal(t, 2, next(-1), "next(-1)")
	assert.Equal(t, -2, next(2), "next(2)")
	assert.Equal(t, 3, next(-2), "next(-2)")
}
