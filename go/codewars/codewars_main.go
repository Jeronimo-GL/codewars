// You can edit this code!
// Click here and start typing.
package main;

import (
	"fmt"
	sd "codewars/seqDisguise"
)

func main() {
	for i:=0 ; i < 10000 ; i ++ {
		n:= uint(i)
		r:= sd.Fct(n)
		fmt.Printf("Fct(%d) -> %d\n", n, r)
	}
}
