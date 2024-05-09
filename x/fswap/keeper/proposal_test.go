package keeper_test

import (
	sdk "github.com/Finschia/finschia-sdk/types"
	sdkerrors "github.com/Finschia/finschia-sdk/types/errors"
	bank "github.com/Finschia/finschia-sdk/x/bank/types"
	"github.com/Finschia/finschia-sdk/x/fswap/testutil"
	"github.com/Finschia/finschia-sdk/x/fswap/types"
	"github.com/golang/mock/gomock"
)

func (s *KeeperTestSuite) TestMakeSwapProposal() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()

	bankKeeper := testutil.NewMockBankKeeper(ctrl)
	bankKeeper.EXPECT().HasSupply(gomock.Any(), "fromdenom").Return(true).AnyTimes()
	bankKeeper.EXPECT().HasSupply(gomock.Any(), gomock.Any()).Return(false).AnyTimes()
	bankKeeper.EXPECT().GetDenomMetaData(gomock.Any(), "todenom").Return(bank.Metadata{}, false).AnyTimes()
	bankKeeper.EXPECT().GetDenomMetaData(gomock.Any(), gomock.Any()).Return(s.toDenomMetadata, true).AnyTimes()
	bankKeeper.EXPECT().SetDenomMetaData(gomock.Any(), s.toDenomMetadata).Times(1)
	s.keeper.BankKeeper = bankKeeper

	testCases := map[string]struct {
		swap          types.Swap
		toDenomMeta   bank.Metadata
		expectedError error
	}{
		"valid": {
			types.Swap{
				FromDenom:           "fromdenom",
				ToDenom:             "todenom",
				AmountCapForToDenom: sdk.OneInt(),
				SwapRate:            sdk.OneDec(),
			},
			s.toDenomMetadata,
			nil,
		},
		"from-denom does not exist": {
			types.Swap{
				FromDenom:           "fakedenom",
				ToDenom:             "todenom",
				AmountCapForToDenom: sdk.OneInt(),
				SwapRate:            sdk.OneDec(),
			},
			s.toDenomMetadata,
			sdkerrors.ErrInvalidRequest,
		},
		"to-denom does not equal with metadata": {
			types.Swap{
				FromDenom:           "fromdenom",
				ToDenom:             "fakedenom",
				AmountCapForToDenom: sdk.OneInt(),
				SwapRate:            sdk.OneDec(),
			},
			s.toDenomMetadata,
			sdkerrors.ErrInvalidRequest,
		},
		"to-denom metadata change not allowed": {
			types.Swap{
				FromDenom:           "fromdenom",
				ToDenom:             "change",
				AmountCapForToDenom: sdk.OneInt(),
				SwapRate:            sdk.OneDec(),
			},
			bank.Metadata{
				Description: s.toDenomMetadata.Description,
				DenomUnits:  s.toDenomMetadata.DenomUnits,
				Base:        "change",
				Display:     s.toDenomMetadata.Display,
				Name:        s.toDenomMetadata.Name,
				Symbol:      s.toDenomMetadata.Symbol,
			},
			sdkerrors.ErrInvalidRequest,
		},
	}
	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()
			err := s.keeper.MakeSwap(ctx, tc.swap, tc.toDenomMeta)
			s.Require().ErrorIs(err, tc.expectedError)
		})
	}
}

func (s *KeeperTestSuite) TestSwapValidateBasic() {
	testCases := map[string]struct {
		swap             types.Swap
		shouldThrowError bool
		expectedError    error
	}{
		"valid": {
			types.Swap{
				FromDenom:           "fromD",
				ToDenom:             "toD",
				AmountCapForToDenom: sdk.OneInt(),
				SwapRate:            sdk.OneDec(),
			},
			false,
			nil,
		},
		"invalid empty from-denom": {
			types.Swap{
				FromDenom:           "",
				ToDenom:             "toD",
				AmountCapForToDenom: sdk.OneInt(),
				SwapRate:            sdk.OneDec(),
			},
			true,
			sdkerrors.ErrInvalidRequest,
		},
		"invalid empty to-denom": {
			types.Swap{
				FromDenom:           "fromD",
				ToDenom:             "",
				AmountCapForToDenom: sdk.OneInt(),
				SwapRate:            sdk.OneDec(),
			},
			true,
			sdkerrors.ErrInvalidRequest,
		},
		"invalid zero amount cap for to-denom": {
			types.Swap{
				FromDenom:           "fromD",
				ToDenom:             "toD",
				AmountCapForToDenom: sdk.ZeroInt(),
				SwapRate:            sdk.OneDec(),
			},
			true,
			sdkerrors.ErrInvalidRequest,
		},
		"invalid zero swap-rate": {
			types.Swap{
				FromDenom:           "fromD",
				ToDenom:             "toD",
				AmountCapForToDenom: sdk.OneInt(),
				SwapRate:            sdk.ZeroDec(),
			},
			true,
			sdkerrors.ErrInvalidRequest,
		},
		"invalid the same from-denom and to-denom": {
			types.Swap{
				FromDenom:           "same",
				ToDenom:             "same",
				AmountCapForToDenom: sdk.OneInt(),
				SwapRate:            sdk.OneDec(),
			},
			true,
			sdkerrors.ErrInvalidRequest,
		},
	}
	for name, tc := range testCases {
		s.Run(name, func() {
			err := tc.swap.ValidateBasic()
			if tc.shouldThrowError {
				s.Require().ErrorIs(err, tc.expectedError)
				return
			}
			s.Require().NoError(err)
		})
	}
}
