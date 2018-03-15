// You can edit this code!
// Click here and start typing.
package main;

import (
	"fmt"
	cd "codewars/cartesianDistance"
)

func main() {
	for i:=1; i<=6; i++{
		var r=cd.CartesianNeighborsDistance(1,1,i)
		fmt.Printf("Step: %d Len: %d\n", i, len(r))
	}
}

func  printList(list [][]int ) {
	for _, n := range list {
		fmt.Printf("(%d, %d)\n", n[0], n[1])
	}
}
