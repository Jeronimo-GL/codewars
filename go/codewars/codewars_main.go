// You can edit this code!
// Click here and start typing.
package main;

import (
	"fmt"
	sp "codewars/stepsInPrimes"
)

func main() {
	r:= sp.Step(11, 300000, 1000000)
	fmt.Printf("(%d, %d)\n", r[0], r[1])
}

func  printList(list [][]int ) {
	for _, n := range list {
		fmt.Printf("(%d, %d)\n", n[0], n[1])
	}
}
