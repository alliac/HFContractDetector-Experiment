package test

import (
	"fmt"
	_ "github.com/emp-toolkit/jialanli/lacia/utils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type ChainCode struct {
}

func (a *ChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Init called")
	return shim.Success(nil)
}
func (a *ChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Invoke called")
	fn, args := stub.GetFunctionAndParameters()
	if fn == "Test1" {
		return a.Test1(stub, args)
	}
	if fn == "Test2" {
		return a.Test2(stub, args)
	}
	return shim.Success(nil)
}
//len(args) >= 1
func (a *ChainCode) Test1(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("Test1 called")
	if len(args) < 2 {
	}
	if args[0] != "10" {
	}
	return shim.Success(nil)
}
func (a *ChainCode) Test2(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("Test2 called")
	if len(args) > 2 {
	}
	return shim.Success(nil)
}
