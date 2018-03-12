// You can edit this code!
// Click here and start typing.
package main;

import (
	"fmt"
)

func main() {
	fmt.Printf("Pruebas de arrays\n")
	
	var list [][]int

	list=append(list, []int{1,1})
	list=append(list, []int{2,1})

	printList(list)
}

func  printList(list [][]int ) {
	for _, n := range list {
		fmt.Printf("(%d, %d)\n", n[0], n[1])
	}
}
