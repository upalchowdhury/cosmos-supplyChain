package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateContract = "create_contract"

var _ sdk.Msg = &MsgCreateContract{}

func NewMsgCreateContract(creator string, dealId string, consumer string, desc string, ownereta string, expiry string, fees uint64) *MsgCreateContract {
	return &MsgCreateContract{
		Creator:  creator,
		DealId:   dealId,
		Consumer: consumer,
		Desc:     desc,
		Ownereta: ownereta,
		Expiry:   expiry,
		Fees:     fees,
	}
}

func (msg *MsgCreateContract) Route() string {
	return RouterKey
}

func (msg *MsgCreateContract) Type() string {
	return TypeMsgCreateContract
}

func (msg *MsgCreateContract) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
