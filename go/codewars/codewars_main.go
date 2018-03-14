// You can edit this code!
// Click here and start typing.
package main;

import (
	"fmt"
)

func main() {
	for i:=-1; i<=1; i++ {
		for j:=-1; j<=1; j++ {
			fmt.Println(i, j)
		}
	}
}

func  printList(list [][]int ) {
	for _, n := range list {
		fmt.Printf("(%d, %d)\n", n[0], n[1])
	}
}
