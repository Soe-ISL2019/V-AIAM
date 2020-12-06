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

	if len(args) != 0 {
		return shim.Error("Incorrect number of arguments. Expecting 0")
	}

	Pk := PkInfo{}
	Pk.Val_q = "0"
	Pk.Val_X = "0"
	Pk.Val_Y = "0"
	Pk.Val_Z = "0"

	CL := CL_sigInfo{}
	CL.Val_a = "0"
	CL.Val_A = "0"
	CL.Val_b = "0"
	CL.Val_B = "0"
	CL.Val_c = "0"

	Tx := TxInfo{}
	Tx.TxID = "0"
	Tx.ChID = "0"
	Tx.Timestamp = "0"

	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}

	txid := stub.GetTxID()
	Tx.TxID = txid
	jsonResp := "{\"Init_txid\":\"" + Tx.TxID + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)

	chid := stub.GetChannelID()
	Tx.ChID = chid
	jsonResp = "{\"Init_chid\":\"" + Tx.ChID + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)

	// chid := stub.GetChannelID()
	// fmt.Printf("Query Response1:%s\n", chid)
	timestamp, _ := stub.GetTxTimestamp()
	unix_time := time.Unix(timestamp.GetSeconds(), int64(timestamp.GetNanos()))
	Tx.Timestamp = unix_time.String()
	jsonResp = "{\"Init_timestamp\":\"" + unix_time.String() + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)

	PkValByte, _ := json.Marshal(Pk)
	// Write the state to the ledger
	err = stub.PutState(Pk.Val_q, PkValByte)
	if err != nil {
		return shim.Error(err.Error())
	}

	CLValByte, _ := json.Marshal(CL)
	// Write the state to the ledger
	err = stub.PutState(CL.Val_c, CLValByte)
	if err != nil {
		return shim.Error(err.Error())
	}

	TxValByte, _ := json.Marshal(Tx)
	// Write the state to the ledger
	err = stub.PutState(Tx.TxID, TxValByte)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf("TxID = %s,\nChannel_ID = %s,\nTimestamp = %s\n\n", Tx.TxID, Tx.ChID, Tx.Timestamp)
	fmt.Printf("q = %s,\nX = %s,\nY = %s,\nZ = %s\n\n", Pk.Val_q, Pk.Val_X, Pk.Val_Y, Pk.Val_Z)
	fmt.Printf("a = %s,\nA = %s,\nb = %s,\nB = %s,\nc = %s\n\n", CL.Val_a, CL.Val_A, CL.Val_b, CL.Val_B, CL.Val_c)

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
		// } else if function == "getinfo" {
		// 	// the old "Query" is now implemtned in invoke
		// 	return t.getinfo(stub, args)
		// } else if function == "verify" {
		// 	// the old "Query" is now implemtned in invoke
		// 	return t.verify(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	// var Tval, Pval, Tsval string
	var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	qvalbytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if qvalbytes == nil {
		return shim.Error("Entity not found")
	}

	Pk := PkInfo{}
	err = json.Unmarshal(qvalbytes, &Pk)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println(Pk.Val_q)
	Pk.Val_q = args[0]

	CL := CL_sigInfo{}
	Tx := TxInfo{}

	Pk, CL = cl_sig(args[1])

	txid := stub.GetTxID()
	Tx.TxID = txid
	jsonResp := "{\"Init_txid\":\"" + Tx.TxID + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)

	chid := stub.GetChannelID()
	Tx.ChID = chid
	jsonResp = "{\"Init_chid\":\"" + Tx.ChID + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)

	timestamp, _ := stub.GetTxTimestamp()
	unix_time := time.Unix(timestamp.GetSeconds(), int64(timestamp.GetNanos()))
	Tx.Timestamp = unix_time.String()
	jsonResp = "{\"Init_timestamp\":\"" + unix_time.String() + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)

	PkValByte, _ := json.Marshal(Pk)
	// Write the state to the ledger
	err = stub.PutState(Pk.Val_q, PkValByte)
	if err != nil {
		return shim.Error(err.Error())
	}

	CLValByte, _ := json.Marshal(CL)
	// Write the state to the ledger
	err = stub.PutState(CL.Val_c, CLValByte)
	if err != nil {
		return shim.Error(err.Error())
	}

	TxValByte, _ := json.Marshal(Tx)
	// Write the state to the ledger
	err = stub.PutState(Tx.TxID, TxValByte)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf("TxID = %s,\nChannel_ID = %s,\nTimestamp = %s\n\n", Tx.TxID, Tx.ChID, Tx.Timestamp)
	fmt.Printf("q = %s,\nX = %s,\nY = %s,\nZ = %s\n\n", Pk.Val_q, Pk.Val_X, Pk.Val_Y, Pk.Val_Z)
	fmt.Printf("a = %s,\nA = %s,\b = %s,\nB = %s,\nc = %s\n\n", CL.Val_a, CL.Val_A, CL.Val_b, CL.Val_B, CL.Val_c)

	return shim.Success(nil)
}

// query callback representing the query of a chaincode
func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	Pk := PkInfo{}
	Pk.Val_q = args[0]
	Tx := TxInfo{}
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	// Get the state from the ledger
	qvalbytes, err := stub.GetState(Pk.Val_q)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + Pk.Val_q + "\"}"
		return shim.Error(jsonResp)
	}

	err = json.Unmarshal(qvalbytes, &Pk)
	if err != nil {
		return shim.Error(err.Error())
	}

	if qvalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + Pk.Val_q + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"query_chid\":\"" + Tx.ChID + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)

	jsonResp2 := "{\"Pk_q\":\"" + Pk.Val_q + "\",\"contents\":\"" + string(qvalbytes)
	fmt.Printf("Query Response:%s\n", jsonResp2)
	return shim.Success(qvalbytes)
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
