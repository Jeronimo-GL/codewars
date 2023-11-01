package twiceLinear

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_examples(t *testing.T){
	assert.Equal(t, 22, DblLinear(10), "DblLinear(10)")
}


func Test_kata(t *testing.T){
	assert.Equal(t, 175, DblLinear(50), "DblLinear(50)")
	assert.Equal(t, 447, DblLinear(100), "DblLinear(100)")
	assert.Equal(t, 3355, DblLinear(500), "DblLinear(500)")
	assert.Equal(t, 8488, DblLinear(1000), "DblLinear(1000)")
}
