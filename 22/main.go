package main

import "math/big"

// 2 << 20 =

func Add(a, b int) *big.Int {
	aa := big.NewInt(int64(a))
	bb := big.NewInt(int64(b))

	return aa.Add(bb, aa)
}

func Subtract(a, b int) *big.Int {
	aa := big.NewInt(int64(a))
	bb := big.NewInt(int64(b))

	return aa.Sub(aa, bb)
}

func Multiply(a, b int) *big.Int {
	aa := big.NewInt(int64(a))
	bb := big.NewInt(int64(b))
	return aa.Mul(aa, bb)
}

func Divide(a, b int) *big.Int {
	if b == 0 {
		panic("divide by zero")
	}
	aa := big.NewInt(int64(a))
	bb := big.NewInt(int64(b))
	return aa.Div(aa, bb)
}
