package johnAndAnn


func j(n int) int{
	if n==0 {
		return 0
	}else {
		return n - a(j(n-1))
	}
}

func a(n int) int {
	if n==0 {
		return 1
	}else{
		return n - j(a(n-1))
	}
}

func aux(n int)([]int, []int) {
	if n == 0  {
		return []int{1}, []int{0}
	} else {
		a_prev := []int{1, 1}
		j_prev := []int{0, 0}
		for i:=2; i<n; i++{
			a_prev = append(a_prev, i - j_prev[a_prev[i-1]])
			j_prev = append(j_prev, i - a_prev[j_prev[i-1]])
		}
		return a_prev, j_prev
	}
}

func Ann(n int) []int {
	resp, _ := aux(n)
	return resp
}

func John(n int) []int {
	_, resp := aux(n)
	return resp 
}

func SumJohn(n int) int {
	list :=John(n)
	sum := 0
	for _, e := range list {
		sum += e
	}
	
	return sum
}

func SumAnn(n int) int {
	list :=Ann(n)
	sum := 0
	for _, e := range list {
		sum += e
	}
	
	return sum
}
