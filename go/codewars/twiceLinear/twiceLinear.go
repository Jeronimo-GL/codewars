package twiceLinear

import (
	"fmt"
)

func DblLinear(n int) int {
	twos := []int {3}
	threes := []int {4}
	lastTwo := 0
	lastThree := 0
	next:=1
	for i:=2; i<=n +1  ; i++ {
		fmt.Printf("lastTwo: %d [%d] lastThree: %d [%d]\n", lastTwo, len(twos), lastThree, len(threes))
		if twos[lastTwo] <= threes[lastThree] {
			next=twos[lastTwo]
			lastTwo++
		}else{
			next=threes[lastThree]
			lastThree++
		}
		fmt.Printf("Next: %d\n", next)
		fmt.Println(twos)
		fmt.Println(threes)
		// twos=append(twos, next*2 +1)
		// threes=append(threes, next*3 +1)
		if  (next*2 +1) > threes[len(threes)-1] {
			twos=append(twos, next*2 +1)
		}
		if (next*3 +1) > twos[len(twos)-1] {
			threes=append(threes, next*3 +1)
		}
	}
	
    return next
}
