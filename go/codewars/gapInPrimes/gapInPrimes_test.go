package gapInPrimes

import (
	"testing"
	"fmt"
)

func check_pairs(t *testing.T, a []int, b[]int) {
	if a == nil && b == nil {
		return
	}

	if b == nil && a!= nil {
		t.Errorf("Expected nil but (%d, %d) found", a[0], a[1])
		return
	}

	if a == nil && b!= nil {
		t.Errorf("nil found, expected (%d, %d) found", b[0], b[1])
		return
	}
	
	if a[0]!=b[0] || a[1]!=b[1] {
		t.Errorf("(%d, %d) found. Expected (%d, %d)", a[0], a[1], b[0], b[1])
		return
	}
}

func Test_basic_cases(t *testing.T){
	fmt.Printf("Testing ...")
	check_pairs(t, Gap(2, 100, 110), []int{101, 103})
	check_pairs(t, Gap(4, 100, 110), []int{103, 107})
	check_pairs(t, Gap(8, 300, 400), []int{359, 367})
	check_pairs(t, Gap(10, 300, 400), []int{337, 347})

}


func Test_nil_response(t *testing.T){
	check_pairs(t, Gap(6,100,110), nil)
	check_pairs(t, Gap(6,100000,100000), nil)
}
