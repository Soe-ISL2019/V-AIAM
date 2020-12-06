package chaincode

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"

	// "github.com/hyperledger/fabric-contract-api-go/contractapi"

	snark "github.com/arnaucube/go-snark"
	"github.com/arnaucube/go-snark/circuitcompiler"
	"github.com/arnaucube/go-snark/groth16"
	"github.com/arnaucube/go-snark/r1csqap"
	"github.com/arnaucube/go-snark/utils"
	"github.com/urfave/cli"
)

type SmartContract struct {
	contractapi.Contract
}

type ACert struct {
	Token          string `json:"Token"`
	Circuit        string `json:"Circuit"`
	Proof          string    `json:"Proof"`
	Tsetup         string `json:"Tsetup"`
	PInputs        int    `json:"PInputs"`
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	ACertinfos := []ACert{
		{Token: "1", Circuit: "1", stirng: "1", Tsetup: "1", PInputs: 1}		
	}

	for _, acertinfo := range ACertinfos {
		acJSON, err := json.Marshal(acertinfo)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(acertinfo.Token, acJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

func (s *SmartContract) CreateAcert(ctx contractapi.TransactionContextInterface, id string, circuit string, proof string, tsetup string, pinputs int) error {
	exists, err := s.AcertExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the acert %s already exists", id)
	}

	acert := ACert{
		ID:             id,
		Circuit:        circuit,
		Proof:          proof,
		Tsetup:         tsetup,
		PInputs:        pinputs,
	}
	acJSON, err := json.Marshal(acert)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, acJSON)
}

func (s *SmartContract) ReadAcert(ctx contractapi.TransactionContextInterface, id string) (*ACert, error) {
	acJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if acJSON == nil {
		return nil, fmt.Errorf("the acert %s does not exist", id)
	}

	var acert ACert
	err = json.Unmarshal(acJSON, &acert)
	if err != nil {
		return nil, err
	}

	return &acert, nil
}

func (s *SmartContract) UpdateAcert(ctx contractapi.TransactionContextInterface, id string, circuit string, proof string, tsetup string, pinputs int) error {
	exists, err := s.AcertExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the acert %s does not exist", id)
	}

	acert := ACert{
		ID:             id,
		Circuit:        circuit,
		Proof:          proof,
		Tsetup:         tsetup,
		Pinputs:        pinputs,
	}
	acJSON, err := json.Marshal(acert)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, acJSON)
}

func (s *SmartContract) DeleteAcert(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.AcertExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the acert %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

func (s *SmartContract) AcertExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	acJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return acJSON != nil, nil
}

func (s *SmartContract) TransferAcert(ctx contractapi.TransactionContextInterface, id string, newtsetup string) error {
	acert, err := s.ReadAcert(ctx, id)
	if err != nil {
		return err
	}

	acert.tsetup = newtsetup
	acJSON, err := json.Marshal(acert)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, acJSON)
}

func (s *SmartContract) GetAllAcerts(ctx contractapi.TransactionContextInterface) ([]*ACert, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var acerts []*Acert
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var acert Acert
		err = json.Unmarshal(queryResponse.Value, &acert)
		if err != nil {
			return nil, err
		}
		acerts = append(acerts, &acert)
	}

	return acerts, nil
}
/* func (s *SmartContract) Proofs(ctx contractapi.TransactionContextInterface, id string, circuit string, proof int, tsetup string, pinputs int) error
func (s *SmartContract) GetAllAssets(ctx contractapi.TransactionContextInterface, id string) ([]*Acert, error) {
	// open proofs.json
	proofsFile, err := ioutil.ReadFile(Acert[1])
	panicErr(err)
	var proof groth16.Proof
	json.Unmarshal([]byte(string(proofsFile)), &proof)
	panicErr(err)

	// open trustedsetup.json
	trustedsetupFile, err := ioutil.ReadFile("trustedsetup.json")
	panicErr(err)
	var trustedsetup groth16.Setup
	json.Unmarshal([]byte(string(trustedsetupFile)), &trustedsetup)
	panicErr(err)

	// read publicInputs file
	publicInputsFile, err := ioutil.ReadFile("publicInputs.json")
	panicErr(err)
	var publicSignals []*big.Int
	err = json.Unmarshal([]byte(string(publicInputsFile)), &publicSignals)
	panicErr(err)

	verified := groth16.VerifyProof(trustedsetup.Vk, proof, publicSignals, true)
	if !verified {
		fmt.Println("ERROR: proofs not verified")
	} else {
		fmt.Println("Proofs verified")
	}
	return nil
} */