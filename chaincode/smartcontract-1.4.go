/*
Copyright IBM Corp. 2016 All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
		 http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

//WARNING - this chaincode's ID is hard-coded in chaincode_example04 to illustrate one way of
//calling chaincode from a chaincode. If this example is modified, chaincode_example04.go has
//to be modified as well with the new ID of chaincode_example02.
//chaincode_example05 show's how chaincode ID can be passed in as a parameter instead of
//hard-coding.

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("chaincode Init")
	_, args := stub.GetFunctionAndParameters()
	var err error

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	ACert := ACertInfo{}
	ACert.Token = args[0]
	ACert.Proof = args[1]
	ACert.PInputs, err = strconv.Atoi(args[2])

	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	// now := time.Now().UTC()
	// fmt.Printf("Time method:%s\n", now)

	txid := stub.GetTxID()
	ACert.TxID = txid
	jsonResp := "{\"Init_txid\":\"" + ACert.TxID + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)

	chid := stub.GetChannelID()
	ACert.ChID = chid
	jsonResp = "{\"Init_chid\":\"" + ACert.ChID + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)

	timestamp, _ := stub.GetTxTimestamp()
	unix_time := time.Unix(timestamp.GetSeconds(), int64(timestamp.GetNanos()))
	ACert.Timestamp = unix_time.String()
	jsonResp = "{\"Init_timestamp\":\"" + unix_time.String() + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)

	ACertValByte, _ := json.Marshal(ACert)
	// Write the state to the ledger
	err = stub.PutState(ACert.Token, ACertValByte)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf("TxID = %s,\nChannel_ID = %s,\nToken = %s,\nProof = %s,\nPInputs = %s,\nTimestamp = %s\n\n", ACert.TxID, ACert.ChID, ACert.Token, ACert.Proof, ACert.PInputs, ACert.Timestamp)

	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "invoke" {
		// Make payment of X units from A to B
		return t.invoke(stub, args)
	} else if function == "query" {
		// the old "Query" is now implemtned in invoke
		return t.query(stub, args)
	} else if function == "verify" {
		// the old "Query" is now implemtned in invoke
		return t.verify(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	// var Tval, Pval, Tsval string
	var err error

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	ACert := ACertInfo{}
	ACert.Token = args[0]
	ACert.Proof = args[1]
	ACert.PInputs, err = strconv.Atoi(args[2])

	Tvalbytes, err := stub.GetState(ACert.Token)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Tvalbytes == nil {
		return shim.Error("Entity not found")
	}

	Pvalbytes, err := stub.GetState(ACert.Proof)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Pvalbytes == nil {
		return shim.Error("Entity not found")
	}

	PIsvalbytes, err := stub.GetState(string(ACert.PInputs))
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if PIsvalbytes == nil {
		return shim.Error("Entity not found")
	}

	TSvalbytes, err := stub.GetState(ACert.Timestamp)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if TSvalbytes == nil {
		return shim.Error("Entity not found")
	}

	Chvalbytes, err := stub.GetState(ACert.ChID)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Chvalbytes == nil {
		return shim.Error("Entity not found")
	}

	return shim.Success(nil)
}

// query callback representing the query of a chaincode
func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	ACert := ACertInfo{}
	ACert.Token = args[0]
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	// Get the state from the ledger
	Tvalbytes, err := stub.GetState(ACert.Token)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + ACert.Token + "\"}"
		return shim.Error(jsonResp)
	}

	err = json.Unmarshal(Tvalbytes, &ACert)
	if err != nil {
		return shim.Error(err.Error())
	}

	if Tvalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + ACert.Token + "\"}"
		return shim.Error(jsonResp)
	}
	
	jsonResp := "{\"query_chid\":\"" + ACert.ChID + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)

	jsonResp2 := "{\"Token\":\"" + ACert.Token + "\",\"contents\":\"" + string(Tvalbytes)
	fmt.Printf("Query Response:%s\n", jsonResp2)
	return shim.Success(Tvalbytes)
}

func (t *SimpleChaincode) verify(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error
	
	ACert := ACertInfo{}
	ACert.Token = args[0]

	Tvalbytes, err := stub.GetState(ACert.Token)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + ACert.Token + "\"}"
		return shim.Error(jsonResp)
	}

	res := zkp(3, 35)

	if res == true {
		return shim.Success(Tvalbytes)
	} else {
		return shim.Error(err.Error())
	}

}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
