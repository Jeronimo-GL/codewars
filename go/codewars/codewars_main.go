
package main;

import (
	"fmt"
	ja "codewars/johnAndAnn"
)

func main() {
	for i:=0 ; i < 10 ; i ++ {
		fmt.Printf("%d\t%d\t%d\n", i, ja.A(i), ja.J(i))
	}
}
