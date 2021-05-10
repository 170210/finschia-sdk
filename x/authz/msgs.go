package authz

import (
	"time"

	"github.com/gogo/protobuf/proto"

	cdctypes "github.com/line/lbm-sdk/codec/types"
	sdk "github.com/line/lbm-sdk/types"
	sdkerrors "github.com/line/lbm-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgGrant{}
	_ sdk.Msg = &MsgRevoke{}
	_ sdk.Msg = &MsgExec{}

	_ cdctypes.UnpackInterfacesMessage = &MsgGrant{}
	_ cdctypes.UnpackInterfacesMessage = &MsgExec{}
)

// NewMsgGrant creates a new MsgGrant
//nolint:interfacer
func NewMsgGrant(granter sdk.AccAddress, grantee sdk.AccAddress, a Authorization, expiration time.Time) (*MsgGrant, error) {
	m := &MsgGrant{
		Granter: granter.String(),
		Grantee: grantee.String(),
		Grant:   Grant{Expiration: expiration},
	}
	err := m.SetAuthorization(a)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// GetSigners implements Msg
func (msg MsgGrant) GetSigners() []sdk.AccAddress {
	err := sdk.ValidateAccAddress(msg.Granter)
	if err != nil {
		panic(err)
	}
	granter := sdk.AccAddress(msg.Granter)
	return []sdk.AccAddress{granter}
}

// ValidateBasic implements Msg
func (msg MsgGrant) ValidateBasic() error {
	err := sdk.ValidateAccAddress(msg.Granter)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid granter address")
	}
	granter := sdk.AccAddress(msg.Granter)

	err = sdk.ValidateAccAddress(msg.Grantee)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid granter address")
	}
	grantee := sdk.AccAddress(msg.Grantee)

	if granter.Equals(grantee) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "granter and grantee cannot be same")
	}
	return msg.Grant.ValidateBasic()
}

// GetAuthorization returns the cache value from the MsgGrant.Authorization if present.
func (msg *MsgGrant) GetAuthorization() Authorization {
	return msg.Grant.GetAuthorization()
}

// SetAuthorization converts Authorization to any and adds it to MsgGrant.Authorization.
func (msg *MsgGrant) SetAuthorization(a Authorization) error {
	m, ok := a.(proto.Message)
	if !ok {
		return sdkerrors.Wrapf(sdkerrors.ErrPackAny, "can't proto marshal %T", m)
	}
	any, err := cdctypes.NewAnyWithValue(m)
	if err != nil {
		return err
	}
	msg.Grant.Authorization = any
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (msg MsgExec) UnpackInterfaces(unpacker cdctypes.AnyUnpacker) error {
	for _, x := range msg.Msgs {
		var msgExecAuthorized sdk.Msg
		err := unpacker.UnpackAny(x, &msgExecAuthorized)
		if err != nil {
			return err
		}
	}

	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (msg MsgGrant) UnpackInterfaces(unpacker cdctypes.AnyUnpacker) error {
	return msg.Grant.UnpackInterfaces(unpacker)
}

// NewMsgRevoke creates a new MsgRevoke
//nolint:interfacer
func NewMsgRevoke(granter sdk.AccAddress, grantee sdk.AccAddress, msgTypeURL string) MsgRevoke {
	return MsgRevoke{
		Granter:    granter.String(),
		Grantee:    grantee.String(),
		MsgTypeUrl: msgTypeURL,
	}
}

// GetSigners implements Msg
func (msg MsgRevoke) GetSigners() []sdk.AccAddress {
	err := sdk.ValidateAccAddress(msg.Granter)
	if err != nil {
		panic(err)
	}
	granter := sdk.AccAddress(msg.Granter)
	return []sdk.AccAddress{granter}
}

// ValidateBasic implements MsgRequest.ValidateBasic
func (msg MsgRevoke) ValidateBasic() error {
	err := sdk.ValidateAccAddress(msg.Granter)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid granter address")
	}
	granter := sdk.AccAddress(msg.Granter)

	err = sdk.ValidateAccAddress(msg.Grantee)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid grantee address")
	}
	grantee := sdk.AccAddress(msg.Grantee)

	if granter.Equals(grantee) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "granter and grantee cannot be same")
	}

	if msg.MsgTypeUrl == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "missing method name")
	}

	return nil
}

// NewMsgExec creates a new MsgExecAuthorized
//nolint:interfacer
func NewMsgExec(grantee sdk.AccAddress, msgs []sdk.Msg) MsgExec {
	msgsAny := make([]*cdctypes.Any, len(msgs))
	for i, msg := range msgs {
		any, err := cdctypes.NewAnyWithValue(msg)
		if err != nil {
			panic(err)
		}

		msgsAny[i] = any
	}

	return MsgExec{
		Grantee: grantee.String(),
		Msgs:    msgsAny,
	}
}

// GetMessages returns the cache values from the MsgExecAuthorized.Msgs if present.
func (msg MsgExec) GetMessages() ([]sdk.Msg, error) {
	msgs := make([]sdk.Msg, len(msg.Msgs))
	for i, msgAny := range msg.Msgs {
		msg, ok := msgAny.GetCachedValue().(sdk.Msg)
		if !ok {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "messages contains %T which is not a sdk.MsgRequest", msgAny)
		}
		msgs[i] = msg
	}

	return msgs, nil
}

// GetSigners implements Msg
func (msg MsgExec) GetSigners() []sdk.AccAddress {
	err := sdk.ValidateAccAddress(msg.Grantee)
	if err != nil {
		panic(err)
	}
	grantee := sdk.AccAddress(msg.Grantee)
	return []sdk.AccAddress{grantee}
}

// ValidateBasic implements Msg
func (msg MsgExec) ValidateBasic() error {
	err := sdk.ValidateAccAddress(msg.Grantee)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "messages cannot be empty")
	}

	return nil
}
