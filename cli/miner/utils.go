package main

import "math/big"

// BaseWeiToBaseEth is used to convert a number from it's wei representation to it's eth representation
func BaseWeiToBaseEth(x *big.Int) *big.Int {
	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	return new(big.Int).Div(x, exp)
}
