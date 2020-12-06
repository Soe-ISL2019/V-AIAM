package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"go.bryk.io/miracl/core"
	"go.bryk.io/miracl/core/BN254"
	// "github.com/bryk-io/miracl/core"
	// "github.com/bryk-io/miracl/core/BN254"
)

//WARNING - this chaincode's ID is hard-coded in chaincode_example04 to illustrate one way of
//calling chaincode from a chaincode. If this example is modified, chaincode_example04.go has
//to be modified as well with the new ID of chaincode_example02.
//chaincode_example05 show's how chaincode ID can be passed in as a parameter instead of
//hard-coding.

// "fmt"

// "github.com/hyperledger/fabric/core/chaincode/shim"
// pb "github.com/hyperledger/fabric/protos/peer"

type PkInfo struct {
	Val_q string `json:"q"`
	Val_X string `json:"X"`
	Val_Y string `json:"Y"`
	Val_Z string `json:"Z"`
}

type CL_sigInfo struct {
	Val_a string `json:"a"`
	Val_A string `json:"A"`
	Val_b string `json:"b"`
	Val_B string `json:"B"`
	Val_c string `json:"c"`
}

type TxInfo struct {
	TxID      string `json:"TxID"`
	ChID      string `json:"ChID"`
	Timestamp string `json:"Timestamp"`
}

func FP12toByte(F *BN254.FP12) []byte {

	const MFS int = int(BN254.MODBYTES)
	var t [12 * MFS]byte

	F.ToBytes(t[:])
	return (t[:])
}
func randval() *core.RAND {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	rng := core.NewRAND()
	var raw [100]byte
	for i := 0; i < 100; i++ {
		raw[i] = byte(r1.Intn(255))
	}
	rng.Seed(100, raw[:])
	return rng
}

func cl_sig(msgval string) (PkInfo, CL_sigInfo) {
	mymsg := msgval

	argCount := len(os.Args[1:])

	if argCount > 0 {
		mymsg = (os.Args[1])
	}

	msg := []byte(mymsg)

	sh := core.NewHASH256()
	for i := 0; i < len(msg); i++ {
		sh.Process(msg[i])
	}

	m := BN254.FromBytes(sh.Hash())

	//     	p := BN254.NewBIGints(BN254.Modulus)
	q := BN254.NewBIGints(BN254.CURVE_Order)

	x := BN254.Randomnum(q, randval())
	y := BN254.Randomnum(q, randval())
	z := BN254.Randomnum(q, randval())
	r := BN254.Randomnum(q, randval())
	alpha := BN254.Randomnum(q, randval())

	G1 := BN254.ECP_generator() // Generator point in G1

	X := BN254.G1mul(G1, x)
	Y := BN254.G1mul(G1, y)
	Z := BN254.G1mul(G1, z)

	G2 := BN254.ECP2_generator() // Generator point in G2

	a := BN254.G2mul(G2, alpha)
	b := BN254.G2mul(a, y)

	A := BN254.G2mul(a, z)
	B := BN254.G2mul(A, y)

	Pk := PkInfo{}
	Pk.Val_q = q.ToString()
	Pk.Val_X = X.ToString()
	Pk.Val_Y = Y.ToString()
	Pk.Val_Z = Z.ToString()

	fmt.Printf("q = %s,\nq1 = %s,\nX = %s,\nY = %s,\nZ = %s\n\n", Pk.Val_q, Pk.Val_X, Pk.Val_Y, Pk.Val_Z)

	// c=a^{x+xym} A^{xyr} = a^{x+xym} a^{xyrz} = a^{x+xym+xyrz}
	e1 := BN254.Modmul(x, y, q)
	e1 = BN254.Modmul(e1, m, q)
	e1 = BN254.Modadd(e1, x, q) // (x+xym) mod q
	e2 := BN254.Modmul(x, y, q)
	e2 = BN254.Modmul(e2, r, q)
	e2 = BN254.Modmul(e2, z, q) // (xyrz) mod q

	e := BN254.Modadd(e1, e2, q)

	c := BN254.G2mul(a, e)

	CL := CL_sigInfo{}
	CL.Val_a = a.ToString()
	CL.Val_A = A.ToString()
	CL.Val_b = b.ToString()
	CL.Val_B = B.ToString()
	CL.Val_c = c.ToString()

	fmt.Printf("Message: %s\n", mymsg)

	fmt.Printf("Private key:\nx=%s,\n y=%s,\n z=%s\n\n", x.ToString(), y.ToString(), z.ToString())

	fmt.Printf("a = %s,\nA = %s,\nb = %s,\nB = %s,\nc = %s\n\n", CL.Val_a, CL.Val_A, CL.Val_b, CL.Val_B, CL.Val_c)

	LHS := BN254.Ate(a, Z)
	LHS = BN254.Fexp(LHS)
	RHS := BN254.Ate(A, G1)
	RHS = BN254.Fexp(RHS)

	fmt.Printf("Pair 1 - first 20 bytes:\t0x%x\n", FP12toByte(LHS)[:20])
	fmt.Printf("Pair 2 - first 20 bytes:\t0x%x\n", FP12toByte(RHS)[:20])

	if LHS.Equals(RHS) {
		fmt.Printf("\nPairing match: e(a,Z)=e(A,G1)\n\n")
	}

	LHS = BN254.Ate(a, Y)
	LHS = BN254.Fexp(LHS)
	RHS = BN254.Ate(b, G1)
	RHS = BN254.Fexp(RHS)

	fmt.Printf("Pair 1 - first 20 bytes:\t0x%x\n", FP12toByte(LHS)[:20])
	fmt.Printf("Pair 2 - first 20 bytes:\t0x%x\n", FP12toByte(RHS)[:20])

	if LHS.Equals(RHS) {
		fmt.Printf("\nPairing match: e(a,Y)=e(b,G1)\n\n")
	}

	LHS = BN254.Ate(A, Y)
	LHS = BN254.Fexp(LHS)
	RHS = BN254.Ate(B, G1)
	RHS = BN254.Fexp(RHS)

	fmt.Printf("Pair 1 - first 20 bytes:\t0x%x\n", FP12toByte(LHS)[:20])
	fmt.Printf("Pair 2 - first 20 bytes:\t0x%x\n", FP12toByte(RHS)[:20])

	if LHS.Equals(RHS) {
		fmt.Printf("\nPairing match: e(A,Y)=e(B,G1)\n\n")
	}

	//	 e(a,X). e(b,X)^m . e(B,X)^r=e(c,g)

	LHS = BN254.Ate(c, G1)
	LHS = BN254.Fexp(LHS)
	RHS = BN254.Ate(a, X)
	RHS = BN254.Fexp(RHS)
	RHS2 := BN254.Ate(b, X)
	RHS2 = BN254.Fexp(RHS2)
	RHS3 := BN254.Ate(B, X)
	RHS3 = BN254.Fexp(RHS3)

	RHS2 = RHS2.Pow(m)
	RHS3 = RHS3.Pow(r)
	RHS.Mul(RHS2)
	RHS.Mul(RHS3)

	fmt.Printf("Pair 1 - first 20 bytes:\t0x%x\n", FP12toByte(LHS)[:20])
	fmt.Printf("Pair 2 - first 20 bytes:\t0x%x\n", FP12toByte(RHS)[:20])

	if LHS.Equals(RHS) {
		fmt.Printf("\nPairing match: e(a,X)·e(b,X)^m·e(B,X)^r = e(c,g)\n\n")
	}

	return Pk, CL
}
