package parser

import (
	"math/big"
	"strings"
)

type Parser struct{}

func (p *Parser) ParseNumber(s string) *big.Int {
	if strings.Contains(s, "^") {
		parts := strings.Split(s, "^")
		base := new(big.Int)
		exp := new(big.Int)
		base.SetString(parts[0], 10)
		exp.SetString(parts[1], 10)
		return new(big.Int).Exp(base, exp, nil)
	}
	val := new(big.Int)
	val.SetString(s, 10)
	return val
}
