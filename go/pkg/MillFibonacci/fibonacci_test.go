package millfibonacci

import (
	"math/big"
	"testing"
)

type testpair struct {
	fib    int64
	result *big.Int
}

func TestFib(t *testing.T) {
	var tests = []testpair{
		{0, big.NewInt(0)},
		{1, big.NewInt(1)},
		{2, big.NewInt(1)},
		{3, big.NewInt(2)},
		{4, big.NewInt(3)},
		{5, big.NewInt(5)},
		{6, big.NewInt(8)},
		{7, big.NewInt(13)},
		{8, big.NewInt(21)},
		{9, big.NewInt(34)},
		{10, big.NewInt(55)},
		{11, big.NewInt(89)},
		{12, big.NewInt(144)},
		{13, big.NewInt(233)},
		{14, big.NewInt(377)},
		{19, big.NewInt(4181)},
		{24, big.NewInt(46368)},
		{50, big.NewInt(12586269025)},
	}

	for _, pair := range tests {
		v := Fib(pair.fib)
		if v.Cmp(pair.result) != 0 {
			t.Error("For", pair.fib, "expected", pair.result, "got", v, "instead")
		}

	}
}
