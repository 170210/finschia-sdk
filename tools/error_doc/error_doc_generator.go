package main

import (
	"fmt"

	sdkerrors "github.com/Finschia/finschia-sdk/types/errors"
	_ "github.com/Finschia/finschia-sdk/x/authz"
	_ "github.com/Finschia/finschia-sdk/x/bank/types"
	_ "github.com/Finschia/finschia-sdk/x/capability/types"
	_ "github.com/Finschia/finschia-sdk/x/collection"
	_ "github.com/Finschia/finschia-sdk/x/crisis/types"
	_ "github.com/Finschia/finschia-sdk/x/distribution/types"
	_ "github.com/Finschia/finschia-sdk/x/evidence/types"
	_ "github.com/Finschia/finschia-sdk/x/feegrant"
	_ "github.com/Finschia/finschia-sdk/x/gov/types"
	_ "github.com/Finschia/finschia-sdk/x/params/types/proposal"
	_ "github.com/Finschia/finschia-sdk/x/slashing/types"
	_ "github.com/Finschia/finschia-sdk/x/staking/types"

	_ "github.com/Finschia/finschia-sdk/x/token"
	_ "github.com/Finschia/finschia-sdk/x/token/class"
)

func main() {

	errors := sdkerrors.RegisteredErrors()
	for _, err := range errors {
		if err.Codespace() != "sdk" {
			fmt.Printf("Code: %d, CodeSpace: %s, Message: %s\n", err.ABCICode(), err.Codespace(), err.Error())
		}
	}

}
