package twiceLinear


func contains(arr []int, e int) bool {
   for _, a := range arr {
      if a == e {
         return true
      }
   }
   return false
}

func DblLinear(n int) int {
	twos := []int {3}
	threes := []int {4}
	lastTwo := 0
	lastThree := 0
	next:=1
	for i:=2; i<=n +1  ; i++ {
		if twos[lastTwo] <= threes[lastThree] {
			next=twos[lastTwo]
			lastTwo++
		}else{
			next=threes[lastThree]
			lastThree++
		}
		if !contains(threes, next*2+1){
			twos=append(twos, next*2 +1)
		}
		if !contains(twos, next*3+1){
			threes=append(threes, next*3 +1)
		}
	}
	
    return next
}
