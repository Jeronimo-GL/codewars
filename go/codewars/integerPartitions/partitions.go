package IntegerPartitions

func Partitions(n int) int {
	return 1
}

func next(n int) int{
	if n>=0 {
		return -n
	} else {
		return -n+1
	}
}
