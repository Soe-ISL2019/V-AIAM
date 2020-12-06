package main

import (
	"fmt"
	"math/big"
	"os"
	"regexp"
	"strconv"
	"strings"

	snark "github.com/arnaucube/go-snark"
	"github.com/arnaucube/go-snark/circuitcompiler"
	"github.com/arnaucube/go-snark/fields"
	"github.com/arnaucube/go-snark/r1csqap"
)

//WARNING - this chaincode's ID is hard-coded in chaincode_example04 to illustrate one way of
//calling chaincode from a chaincode. If this example is modified, chaincode_example04.go has
//to be modified as well with the new ID of chaincode_example02.
//chaincode_example05 show's how chaincode ID can be passed in as a parameter instead of
//hard-coding.

// "fmt"

// "github.com/hyperledger/fabric/core/chaincode/shim"
// pb "github.com/hyperledger/fabric/protos/peer"

type ACertInfo struct {
	Token string `json:"Token"`
	// Circuit string `json:"Circuit"`
	Proof string `json:"Proof"`
	// Tsetup  string `json:"Tsetup"`
	PInputs   int    `json:"PInputs"`
	Timestamp string `json:"Timestamp"`
	TxID      string `json:"TxID"`
	ChID      string `json:"ChID"`
}

func TrimSpaceNewlineInString(s string) string {
	re := regexp.MustCompile(`LF`)
	return re.ReplaceAllString(s, "\n")
}

func zkp(xval, yval int) bool {

	// var err error
	str := "func exp3(private a):LF\tb = a * aLF\tc = a * bLF\treturn cLFLFfunc main(private s0, public s1):LF\ts3 = exp3(s0)LF\ts4 = s3 + s0LF\ts5 = s4 + 5LF\tequals(s1, s5)LF\tout = 1 * 1"
	// ACert := ACertInfo{}
	// ACert.Token = args[0]

	// Tvalbytes, err := stub.GetState(ACert.Token)
	// if err != nil {
	// 	jsonResp := "{\"Error\":\"Failed to get state for " + ACert.Token + "\"}"
	// 	return shim.Error(jsonResp)
	// }

	// err = json.Unmarshal(Tvalbytes, &ACert)

	// y, _ := strconv.Atoi(string(Tvalbytes))

	x := xval
	y := yval
	// y, err := stub.GetState(args[0])

	argCount := len(os.Args[1:])

	if argCount > 0 {
		str = os.Args[1]
	}
	if argCount > 1 {
		x, _ = strconv.Atoi(os.Args[2])
	}
	if argCount > 2 {
		y, _ = strconv.Atoi(os.Args[3])
	}

	str = TrimSpaceNewlineInString(str)

	parser := circuitcompiler.NewParser(strings.NewReader(str))
	circuit, _ := parser.Parse()

	val1 := big.NewInt(int64(x))
	privateVal := []*big.Int{val1}
	val2 := big.NewInt(int64(y))
	publicVal := []*big.Int{val2}

	// witness
	w, _ := circuit.CalculateWitness(privateVal, publicVal)

	a, b, c := circuit.GenerateR1CS()

	r, _ := new(big.Int).SetString("21888242871839275222246405745257275088548364400416034343698204186575808495617", 10)

	f := fields.NewFq(r)
	// new Polynomial Field
	pf := r1csqap.NewPolynomialField(f)
	alphas, betas, gammas, _ := pf.R1CSToQAP(a, b, c)

	_, _, _, px := pf.CombinePolynomials(w, alphas, betas, gammas)

	setup, _ := snark.GenerateTrustedSetup(len(w), *circuit, alphas, betas, gammas)

	proof, _ := snark.GenerateProofs(*circuit, setup.Pk, w, px)

	yVerif := big.NewInt(int64(y))
	publicSignalsVerif := []*big.Int{yVerif}

	rtn := snark.VerifyProof(setup.Vk, proof, publicSignalsVerif, true)
	if rtn == true {
		fmt.Printf("Valid proofs!!!")
		return true
	} else {
		return false
	}
}
