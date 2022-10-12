package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateTran = "create_tran"

var _ sdk.Msg = &MsgCreateTran{}

func NewMsgCreateTran(creator string, vendor string, commission string) *MsgCreateTran {
	return &MsgCreateTran{
		Creator:    creator,
		Vendor:     vendor,
		Commission: commission,
	}
}

func (msg *MsgCreateTran) Route() string {
	return RouterKey
}

func (msg *MsgCreateTran) Type() string {
	return TypeMsgCreateTran
}

func (msg *MsgCreateTran) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateTran) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateTran) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
