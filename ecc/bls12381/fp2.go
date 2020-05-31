package bls12381

import "fmt"

type fp2 [2]fp

func (z fp2) String() string       { return fmt.Sprintf("0: %v\n   1: %v", z[0], z[1]) }
func (z *fp2) Set(x *fp2)          { z[0].Set(&x[0]); z[1].Set(&x[1]) }
func (z *fp2) SetZero()            { z[0].SetZero(); z[1].SetZero() }
func (z *fp2) SetOne()             { z[0].SetOne(); z[1].SetZero() }
func (z *fp2) IsZero() bool        { return z[0].IsZero() && z[1].IsZero() }
func (z *fp2) IsEqual(x *fp2) bool { return z[0].IsEqual(&x[0]) && z[1].IsEqual(&x[1]) }
func (z *fp2) MulBeta()            { var t fp; t.Set(&z[0]); z[0].Sub(&z[0], &z[1]); z[1].Add(&t, &z[1]) }
func (z *fp2) Frob()               { z.Cjg() }
func (z *fp2) Cjg()                { z[1].Neg() }
func (z *fp2) Neg()                { z[0].Neg(); z[1].Neg() }
func (z *fp2) Add(x, y *fp2)       { z[0].Add(&x[0], &y[0]); z[1].Add(&x[1], &y[1]) }
func (z *fp2) Sub(x, y *fp2)       { z[0].Sub(&x[0], &y[0]); z[1].Sub(&x[1], &y[1]) }
func (z *fp2) Mul(x, y *fp2) {
	var x0y0, x1y1, sx, sy, k fp
	x0y0.Mul(&x[0], &y[0])
	x1y1.Mul(&x[1], &y[1])
	sx.Add(&x[0], &x[1])
	sy.Add(&y[0], &y[1])
	k.Mul(&sx, &sy)
	z[0].Sub(&x0y0, &x1y1)
	z[1].Sub(&k, &x0y0)
	z[1].Sub(&z[1], &x1y1)
}
func (z *fp2) Sqr(x *fp2) {
	var x02, x12, k fp
	x02.Sqr(&x[0])
	x12.Sqr(&x[1])
	k.Mul(&x[0], &x[1])
	z[0].Sub(&x02, &x12)
	z[1].Add(&k, &k)
}
func (z *fp2) Inv(x *fp2) {
	var x02, x12, den fp
	x02.Sqr(&x[0])
	x12.Sqr(&x[1])
	den.Add(&x02, &x12)
	den.Inv(&den)
	z[0].Mul(&x[0], &den)
	z[1].Mul(&x[1], &den)
	z[1].Neg()
}