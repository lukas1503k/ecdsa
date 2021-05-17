package Secp256k1

import (
	"math/big"
)



type EllipticCurve struct{
	P *big.Int
	A *big.Int
	B *big.Int
	Gx *big.Int
	Gy *big.Int
	N *big.Int
	H *big.Int
}


func GetSecp256k1() EllipticCurve{
	newP, _ := new(big.Int).SetString("0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 0)
	newN,_ := new(big.Int).SetString("0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141", 0)
	newA, _ := new(big.Int).SetString("0x0000000000000000000000000000000000000000000000000000000000000000", 0)
	newB, _ := new(big.Int).SetString("0x0000000000000000000000000000000000000000000000000000000000000007", 0)
	newGx, _ := new(big.Int).SetString("79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798", 0)
	newGy, _ := new(big.Int).SetString("0x483ADA7726A3C4655DA4FBFC0E1108A8FD17B448A68554199C47D08FFB10D4B8", 0)
	newH,_ := new(big.Int).SetString("0x01", 0)

	return EllipticCurve{
		P: newP,
		N: newN,
		A: newA,
		B: newB,
		Gx: newGx,
		Gy: newGy,
		H: newH,
	}

}

func (curve EllipticCurve) IsOnCurve(x,y *big.Int) bool{
	ySquared := new(big.Int).Mul(y,y)
	ySquared.Mod(ySquared, curve.P)

	ellipticCurveEqualtion := new(big.Int).Mul(x,x)
	ellipticCurveEqualtion.Mul(ellipticCurveEqualtion,x)

	ellipticCurveEqualtion.Add(ellipticCurveEqualtion, curve.B)
	ellipticCurveEqualtion.Mod(ellipticCurveEqualtion, curve.P)

	return ySquared.Cmp(ellipticCurveEqualtion) == 0

}

