package main

import (
	"fmt"
	"reflect"

	sdkerrors "github.com/Finschia/finschia-sdk/types/errors"
	authz "github.com/Finschia/finschia-sdk/x/authz"
	bank "github.com/Finschia/finschia-sdk/x/bank/types"
)

func main() {
	_ = authz.ErrInvalidExpirationTime
	_ = bank.ErrNoInputs
	test := reflect.TypeOf(struct{}{}).PkgPath(bank)
	errors := sdkerrors.RegisteredErrors()
	for _, err := range errors {
		fmt.Printf("Code: %d, CodeSpace: %s, Message: %s\n", err.ABCICode(), err.Codespace(), err.Error())
	}

}
