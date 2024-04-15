package types

import (
	fmt "fmt"

	"gopkg.in/yaml.v2"

	paramtypes "github.com/Finschia/finschia-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

const (
	DefaultSwapRate         uint64 = 148079656
	DefaultSwapRateDecimals int32  = 6
	DefaultNewCoinDenom     string = "PDT"
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	SwapRate uint64,
	SwapRateDecimals int32,
	NewCoinDenom string,
) Params {
	return Params{
		SwapRate:         SwapRate,
		SwapRateDecimals: SwapRateDecimals,
		NewCoinDenom:     NewCoinDenom,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultSwapRate,
		DefaultSwapRateDecimals,
		DefaultNewCoinDenom,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateSwapRate(p.SwapRate); err != nil {
		return err
	}
	if err := validateSwapRateDecimals(p.SwapRateDecimals); err != nil {
		return err
	}
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateSwapRate(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("swap rate must be positive: %d", v)
	}

	return nil
}

func validateSwapRateDecimals(i interface{}) error {
	v, ok := i.(int32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("swap rate decimals must be positive: %d", v)
	}

	return nil
}
