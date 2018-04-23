
package main;

import (
	"fmt"
	ja "codewars/johnAndAnn"
)

func main() {
	k :=30
	fmt.Print("Ann  ", ja.Ann(k), "\n")
	fmt.Print("John ", ja.John(k), "\n")

	fmt.Print(ja.SumAnn(0))
}
