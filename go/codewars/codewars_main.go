
package main;

import (
	"fmt"
	"math/rand"
	. "codewars/ml_1"
)


func main() {
	m := NewMachine()

	_actions := []func(int) int{
		func(x int) int { return x + 1 },
		func(x int) int { return 0 },
		func(x int) int { return int(x / 2) },
		func(x int) int { return x * 100 },
		func(x int) int { return x % 2 }}


	for i:=0; i<100; i++{
		COMMAND:=rand.Intn(5)
		res:=m.Command(COMMAND, i)
//		fmt.Printf("val:%d\t res:%d\texpected:%d\n", i, res, _actions[COMMAND](i))
		m.Response(res==_actions[COMMAND](i))
		//		m.PrintState()
		if res==_actions[COMMAND](i) {
			fmt.Print("+")
		}else{
			fmt.Print("-")
		}
	}
	fmt.Println("")
	m.PrintMatrix()
	
}
