package types

import (
	"errors"
	fmt "fmt"
	"strings"

	"gopkg.in/yaml.v2"

	sdk "github.com/Finschia/finschia-sdk/types"
	paramtypes "github.com/Finschia/finschia-sdk/x/params/types"
)

const (
	DefaultNewCoinDenom string = "PDT"
)

var (
	ParamStoreKeyNewDenom = []byte("NewDenom")
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	newCoinDenom string,
) Params {
	return Params{
		NewCoinDenom: newCoinDenom,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultNewCoinDenom,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(ParamStoreKeyNewDenom, &p.NewCoinDenom, validateNewCoinDenom),
	}
}

func validateNewCoinDenom(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if strings.TrimSpace(v) == "" {
		return errors.New("new denom cannot be blank")
	}
	if err := sdk.ValidateDenom(v); err != nil {
		return err
	}

	return nil
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateNewCoinDenom(p.NewCoinDenom); err != nil {
		return err
	}
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
