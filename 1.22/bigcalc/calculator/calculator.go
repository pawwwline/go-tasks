package calculator

import "math/big"

type Calculator struct{}

func (c *Calculator) Add(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func (c *Calculator) Sub(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func (c *Calculator) Mul(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func (c *Calculator) Div(a, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b)
}

func (c *Calculator) Pow(a, b *big.Int) *big.Int {
	return new(big.Int).Exp(a, b, nil)
}
