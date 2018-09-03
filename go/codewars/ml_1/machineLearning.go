package machineLearning

import (
	"math/rand"
	"fmt"
)

type Machine struct {
	actions  []func(int) int
	matrix [ACTIONS_SIZE][ACTIONS_SIZE]int
	lastCommand int
	lastActionProposed int
}

const ACTIONS_SIZE int =5

//Function called to get your machine initialised
func NewMachine() Machine {
	_actions := []func(int) int{
		func(x int) int { return x + 1 },
		func(x int) int { return 0 },
		func(x int) int { return int(x / 2) },
		func(x int) int { return x * 100 },
		func(x int) int { return x % 2 }}
	var matrix [ACTIONS_SIZE][ACTIONS_SIZE]int
	return Machine{_actions,matrix, -1, -1}
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
		i+=1
		if m.matrix[cmd][i%ACTIONS_SIZE]>=0 {
			c-=1
		}
	}
	proposedAction:=i%ACTIONS_SIZE
	m.lastCommand=cmd
	m.lastActionProposed=proposedAction
	return m.actions[proposedAction](num)
}

func (m *Machine) Response(res bool) {
	if res {
		m.matrix[m.lastCommand][m.lastActionProposed] += 1
	}else{
		m.matrix[m.lastCommand][m.lastActionProposed] -= 1
	}
}
