package magnetParticules

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func Test_v(t *testing.T){
	assert.InEpsilon(t, 1.0/4,  v(1,1), 1e-12, "v(1,1)")
	assert.InEpsilon(t, 1.0/9,  v(1,2), 1e-12, "v(1,2)")
	assert.InEpsilon(t, 1.0/32, v(2,1), 1e-12, "v(2,1)")
	assert.InEpsilon(t, 1.0/162, v(2,2), 1e-12, "v(2,2)")
}

func Test_examples(t *testing.T){
	assert.InEpsilon(t, 0.4236111111111111, Doubles(1, 3), 1e-12, "Doubles(1,3)")
	assert.InEpsilon(t, 0.5580321939764581, Doubles(1, 10), 1e-12, "Doubles(1,10)")
	assert.InEpsilon(t, 0.6832948559787737, Doubles(10, 100), 1e-12, "Doubles(10,100)")
}

func Test_kata(t *testing.T){
	assert.InEpsilon(t, 0.6921486500921933, Doubles(10, 1000),  1e-12,  "Doubles(10,1000)")
	assert.InEpsilon(t, 0.6930471674194457, Doubles(10, 10000), 1e-12, "Doubles(10,10000)")
}

