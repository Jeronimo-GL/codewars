package machineLearning

import (
	"math/rand"
	"fmt"
)

type Machine struct {
//	actions  []func(int) int
	matrix [][]int
	lastCommand int
	lastActionProposed int
}


func Actions()[]func(int) int{
	_actions := []func(int) int{
		func(x int) int { return x + 1 },
		func(x int) int { return 0 },
		func(x int) int { return (x / 2) },
		func(x int) int { return x * 100 },
		func(x int) int { return x % 2 }}

	return _actions
}

var ACTIONS_SIZE int =len(Actions())

//Function called to get your machine initialised
func NewMachine() Machine {

	matrix := make([][]int, ACTIONS_SIZE)
	for i := range matrix {
		matrix[i] = make([]int, ACTIONS_SIZE)
	}
	return Machine{matrix, -1, -1}
} 

func (m *Machine) PrintState(){
	fmt.Printf("LastCommand:%d\tlastProposedAction:%d\n", m.lastCommand, m.lastActionProposed)
	for _,r:= range m.matrix[m.lastCommand]{
		fmt.Print(r, " ")
	}
	fmt.Println("")
}


func (m *Machine ) PrintMatrix(){
	fmt.Println("Matrix:")
	for _, act:=range m.matrix{
		for _,com:=range act{
			fmt.Print(com, " ")
		}
		fmt.Println("")
	}
}

func (m *Machine) Command(cmd int, num int) int {
	c:=rand.Intn(5)
	i:=0
	for c>=0{
		i ++
		if m.matrix[cmd][i%ACTIONS_SIZE]>=0 {
			c--
		}
	}
	proposedAction:=i%ACTIONS_SIZE
	m.lastCommand=cmd
	m.lastActionProposed=proposedAction
	return Actions()[proposedAction](num)
}

func (m *Machine) Response(res bool) {
	if res {
		m.matrix[m.lastCommand][m.lastActionProposed] ++
	}else{
		m.matrix[m.lastCommand][m.lastActionProposed] --
	}
}
