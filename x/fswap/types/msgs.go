package types

import (
	sdk "github.com/Finschia/finschia-sdk/types"
	sdkerrors "github.com/Finschia/finschia-sdk/types/errors"
)

var _ sdk.Msg = &MsgSwapRequest{}

// NewMsgSwapRequest - construct a msg to swap amounts of old coin to new coin
//
//nolint:interfacer
func NewMsgSwapRequest(fromAddr, toAddr sdk.AccAddress, amount sdk.Coin) *MsgSwapRequest {
	return &MsgSwapRequest{FromAddress: fromAddr.String(), Amount: amount}
}

// ValidateBasic Implements Msg.
func (msg MsgSwapRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid address (%s)", err)
	}

	if !msg.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	if !msg.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	return nil
}

// GetSigners Implements Msg.
func (msg MsgSwapRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

var _ sdk.Msg = &MsgSwapAllRequest{}

// NewMsgSwapRequest - construct a msg to swap all old coin to new coin
//
//nolint:interfacer
func NewMsgSwapAllRequest(fromAddr, toAddr sdk.AccAddress) *MsgSwapAllRequest {
	return &MsgSwapAllRequest{FromAddress: fromAddr.String()}
}

// ValidateBasic Implements Msg.
func (msg MsgSwapAllRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid address (%s)", err)
	}

	return nil
}

// GetSigners Implements Msg.
func (msg MsgSwapAllRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}