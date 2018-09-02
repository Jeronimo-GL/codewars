package machineLearning


import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_command_0(t *testing.T){
	m:=NewMachine()
	for i:=0; i<5; i++{
		res:=m.Command(0, i)
		m.Response(res==i+1)
	}
	assert.Equal(t, 4, m.Command(0, 3), "command 0")
}

func Test_command_1(t *testing.T){
	m:=NewMachine()
	for i:=0; i<5; i++{
		res:=m.Command(1, i)
		m.Response(res==0)
	}
	assert.Equal(t, 0, m.Command(0, 3), "command 1")
}
