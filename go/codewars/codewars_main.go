
package main;

import (
	"fmt"
	"math/rand"
	. "codewars/ml_1"
)


func main() {
	m := NewMachine()

	for i:=0; i<100; i++{
		COMMAND:=rand.Intn(len(Actions()))
		res:=m.Command(COMMAND, i)
//		fmt.Printf("val:%d\t res:%d\texpected:%d\n", i, res, _actions[COMMAND](i))
		m.Response(res==Actions()[COMMAND](i))
		//		m.PrintState()
		if res==Actions()[COMMAND](i) {
			fmt.Print("+")
		}else{
			fmt.Print("-")
		}
	}
	fmt.Println("")
	m.PrintMatrix()
	
}
