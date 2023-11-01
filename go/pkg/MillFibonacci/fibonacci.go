package millfibonacci

import (
	"fmt"
	"math/big"
)

func debugPrint(a, b, c, d *big.Int) {
	fmt.Printf("(%d, %d)\n(%d, %d)\n", a, b, c, d)
}

func Fib(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(0)
	} else if n < 0 {
		sign := big.NewInt(-1)
		sign.Exp(sign, big.NewInt(-n+1), big.NewInt(1))
		f := new(big.Int)
		return f.Mul(Fib(-n), sign)
	} else {
		i := n - 1
		a := big.NewInt(1)
		b := big.NewInt(0)
		c := big.NewInt(0)
		d := big.NewInt(1)
		for i > 0 {
			if i%2 == 1 {
				bd := new(big.Int)
				bd.Mul(b, d)
				ca := new(big.Int)
				ca.Mul(c, a)
				ba := new(big.Int)
				ba.Add(b, a)
				bda := new(big.Int)
				bda.Mul(d, ba)
				cb := new(big.Int)
				cb.Mul(c, b)
				a.Add(bd, ca)
				b.Add(bda, cb)
			}
			c2 := new(big.Int)
			c2.Mul(c, c)
			d2 := new(big.Int)
			d2.Mul(d, d)
			cc := new(big.Int)
			cc.Add(c, c)
			ccd := new(big.Int)
			ccd.Add(cc, d)
			c.Add(c2, d2)
			d.Mul(d, ccd)
			i = i / 2
		}
		return a.Add(a, b)
	}
}

func fast_fib(n int64) *big.Int {
	a := big.NewInt(0)
	b := big.NewInt(1)
	for i := int64(0); i < n; i++ {
		a.Add(a, b)
		a, b = b, a
	}
	return a
}

func fib_clasic(n int64) *big.Int {
	// Code here :)
	if n == 0 {
		return big.NewInt(0)
	} else if n == 1 {
		return big.NewInt(1)
	} else {
		return fib_clasic(n-1).Add(fib_clasic(n-2), fib_clasic(n-1))
	}
}
