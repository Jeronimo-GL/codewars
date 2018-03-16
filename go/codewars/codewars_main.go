// You can edit this code!
// Click here and start typing.
package main;

import (
	"fmt"
	sp "codewars/stepsInPrimes"
)

func main() {
	r:= sp.Step(4, 100, 110)
	fmt.Printf("(%d, %d)", r[0], r[1])
}

func  printList(list [][]int ) {
	for _, n := range list {
		fmt.Printf("(%d, %d)\n", n[0], n[1])
	}
}
