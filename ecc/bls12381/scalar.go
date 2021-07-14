package bls12381

import (
	"crypto/rand"
	"math/big"
)

// ScalarSize is the length in bytes of a Scalar.
const ScalarSize = 32

// Scalar represents an integer used for scalar multiplication.
type Scalar struct{ i big.Int }

func (z Scalar) String() string          { return "0x" + z.i.Text(16) }
func (z *Scalar) Set(x *Scalar)          { z.i.Set(&x.i) }
func (z *Scalar) SetString(s string)     { z.i.SetString(s, 0) }
func (z *Scalar) SetUint64(n uint64)     { z.i.SetUint64(n) }
func (z *Scalar) SetInt64(n int64)       { z.i.SetInt64(n) }
func (z *Scalar) SetZero()               { z.SetUint64(0) }
func (z *Scalar) SetOne()                { z.SetUint64(1) }
func (z *Scalar) IsZero() bool           { return z.i.Mod(&z.i, &primeOrder.i).Sign() == 0 }
func (z *Scalar) IsEqual(x *Scalar) bool { return z.i.Cmp(&x.i) == 0 }
func (z *Scalar) Neg()                   { z.i.Neg(&z.i).Mod(&z.i, &primeOrder.i) }
func (z *Scalar) Add(x, y *Scalar)       { z.i.Add(&x.i, &y.i).Mod(&z.i, &primeOrder.i) }
func (z *Scalar) Sub(x, y *Scalar)       { z.i.Sub(&x.i, &y.i).Mod(&z.i, &primeOrder.i) }
func (z *Scalar) Mul(x, y *Scalar)       { z.i.Mul(&x.i, &y.i).Mod(&z.i, &primeOrder.i) }
func (z *Scalar) Sqr(x *Scalar)          { z.i.Mul(&x.i, &x.i).Mod(&z.i, &primeOrder.i) }
func (z *Scalar) Inv(x *Scalar)          { z.i.ModInverse(&x.i, &primeOrder.i) }
func (z *Scalar) Random() {
	r, err := rand.Int(rand.Reader, &primeOrder.i)
	if err != nil {
		panic(err)
	}
	z.i.Set(r)
}

// Bytes returns a positive scalar in a slice of bytes in little-endian order.
func (z *Scalar) Bytes() []byte {
	out := make([]byte, ScalarSize)
	red := new(big.Int).Mod(&z.i, &primeOrder.i)
	b := red.Bytes()
	l := len(b)
	for i := range b {
		out[i] = b[l-1-i]
	}
	return out
}
