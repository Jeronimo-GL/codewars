package seqDisguise

import "testing"

func check_uint(t *testing.T, a uint, b uint) {
	if a != b {
		t.Errorf("%d found. Expected %d", a, b)
		return
	}
}

func Test_basic_cases(t *testing.T){
	check_uint(t, Fct(17), 131072)
	check_uint(t, Fct(21), 2097152)
	check_uint(t, Fct(14), 16384)
}

func Test_initalCases(t *testing.T){
	check_uint(t, Fct(0), 1)
	check_uint(t, Fct(1), 2)
	check_uint(t, Fct(2), 4)
}
