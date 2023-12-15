package utils

import (
	"math/big"
)

var zero = big.NewInt(0)
var one = big.NewInt(1)

func Fast_mod_exp(base *big.Int, power *big.Int, mod *big.Int) (result *big.Int) {
	// computes base ^ power % mod

	result = big.NewInt(1)

	for power.Cmp(zero) == 1 {

		if power.Bit(0) == 1 {
			result.Mul(result, base)
			result.Mod(result, mod)
		}

		power.Rsh(power, 1)
		base.Mul(base, base)
		base.Mod(base, mod)
	}

	return
}

func Gcd(a *big.Int, b *big.Int) (gcd *big.Int) {
	gcd = a

	for b.Cmp(zero) == 1 {
        a, b = b, a.Mod(a, b)
	}

	return
}

func Carmichaels_totient(p *big.Int, q *big.Int) (ttnt *big.Int) {
	// computes carmichaels totient of p*q. Assumes p and q are both prime

	p.Sub(p, one)
	q.Sub(q, one)

	ttnt.Mul(p, q)
	ttnt.Div(ttnt, Gcd(p, q))

	return
}

func Extended_euclidean(a *big.Int, b *big.Int) (*big.Int, *big.Int, *big.Int) {
	return zero, zero, zero
}