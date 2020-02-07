package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	//Token
	CodeTokenExist       sdk.CodeType = 100
	CodeTokenNotExist    sdk.CodeType = 101
	CodeTokenNotMintable sdk.CodeType = 102

	//Token invalidation
	CodeTokenInvalidTokenName   sdk.CodeType = 200
	CodeTokenInvalidTokenSymbol sdk.CodeType = 201
	CodeTokenInvalidTokenID     sdk.CodeType = 202
	CodeTokenInvalidDecimals    sdk.CodeType = 203
	CodeTokenInvalidFT          sdk.CodeType = 204
	CodeTokenInvalidAmount      sdk.CodeType = 205

	//Collection
	CodeCollectionExist             sdk.CodeType = 300
	CodeCollectionNotExist          sdk.CodeType = 301
	CodeCollectionTokenTypeExist    sdk.CodeType = 302
	CodeCollectionTokenTypeNotExist sdk.CodeType = 303

	//Permission
	CodeTokenPermission sdk.CodeType = 400

	// Composability
	CodeTokenAlreadyAChild             sdk.CodeType = 500
	CodeTokenNotAChild                 sdk.CodeType = 501
	CodeTokenNotOwnedBy                sdk.CodeType = 502
	CodeTokenChildNotTransferable      sdk.CodeType = 503
	CodeTokenNotNF                     sdk.CodeType = 504
	CodeTokenNotIDNF                   sdk.CodeType = 505
	CodeTokenCannotAttachToItself      sdk.CodeType = 506
	CodeTokenCannotAttachToADescendant sdk.CodeType = 507
)

func ErrTokenExist(codespace sdk.CodespaceType, symbol string) sdk.Error {
	return sdk.NewError(codespace, CodeTokenExist, "token [%s] already exists", symbol)
}

func ErrTokenNotExist(codespace sdk.CodespaceType, symbol string) sdk.Error {
	return sdk.NewError(codespace, CodeTokenNotExist, "token [%s] does not exist", symbol)
}

func ErrTokenNotMintable(codespace sdk.CodespaceType, symbol string) sdk.Error {
	return sdk.NewError(codespace, CodeTokenNotMintable, "token [%s] is not mintable", symbol)
}

func ErrInvalidTokenName(codespace sdk.CodespaceType, name string) sdk.Error {
	return sdk.NewError(codespace, CodeTokenInvalidTokenName, "token name [%s] should not be empty", name)
}

func ErrInvalidTokenSymbol(codespace sdk.CodespaceType, msg string) sdk.Error {
	return sdk.NewError(codespace, CodeTokenInvalidTokenSymbol, msg)
}

func ErrInvalidTokenID(codespace sdk.CodespaceType, msg string) sdk.Error {
	return sdk.NewError(codespace, CodeTokenInvalidTokenID, msg)
}

func ErrInvalidTokenDecimals(codespace sdk.CodespaceType, decimals sdk.Int) sdk.Error {
	return sdk.NewError(codespace, CodeTokenInvalidDecimals, "token decimals [%s] should be within the range in 0 ~ 18", decimals.String())
}

func ErrInvalidIssueFT(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeTokenInvalidFT, "Issuing token with amount[1], decimals[0], mintable[false] prohibited. Issue nft token instead.")
}

func ErrInvalidAmount(codespace sdk.CodespaceType, amount string) sdk.Error {
	return sdk.NewError(codespace, CodeTokenInvalidAmount, "invalid token amount [%s]", amount)
}

func ErrCollectionExist(codespace sdk.CodespaceType, symbol string) sdk.Error {
	return sdk.NewError(codespace, CodeCollectionExist, "collection [%s] already exists", symbol)
}

func ErrCollectionNotExist(codespace sdk.CodespaceType, symbol string) sdk.Error {
	return sdk.NewError(codespace, CodeCollectionNotExist, "collection [%s] does not exists", symbol)
}

func ErrTokenPermission(codespace sdk.CodespaceType, account fmt.Stringer, permission PermissionI) sdk.Error {
	return sdk.NewError(codespace, CodeTokenPermission, "account [%s] does not have the permission [%s]", account.String(), permission.String())
}

func ErrCollectionTokenExist(codespace sdk.CodespaceType, symbol, tokenID string) sdk.Error {
	return sdk.NewError(codespace, CodeTokenExist, "token symbol[%s] token-id[%s] already exists", symbol, tokenID)
}

func ErrCollectionTokenNotExist(codespace sdk.CodespaceType, symbol, tokenID string) sdk.Error {
	return sdk.NewError(codespace, CodeTokenNotExist, "token symbol[%s] token-id[%s] does not exist", symbol, tokenID)
}

func ErrCollectionTokenTypeExist(codespace sdk.CodespaceType, symbol, tokenType string) sdk.Error {
	return sdk.NewError(codespace, CodeCollectionTokenTypeExist, "token type for symbol[%s] token-type[%s] already exists", symbol, tokenType)
}

func ErrCollectionTokenTypeNotExist(codespace sdk.CodespaceType, symbol, tokenType string) sdk.Error {
	return sdk.NewError(codespace, CodeCollectionTokenTypeNotExist, "token type for symbol[%s] token-type[%s] does not exist", symbol, tokenType)
}

func ErrTokenAlreadyAChild(codespace sdk.CodespaceType, denom string) sdk.Error {
	return sdk.NewError(codespace, CodeTokenAlreadyAChild, "token [%s] is already a child of some other", denom)
}

func ErrTokenNotAChild(codespace sdk.CodespaceType, denom string) sdk.Error {
	return sdk.NewError(codespace, CodeTokenNotAChild, "token [%s] is not a child of some other", denom)
}

func ErrTokenNotOwnedBy(codespace sdk.CodespaceType, denom string, owner fmt.Stringer) sdk.Error {
	return sdk.NewError(codespace, CodeTokenNotOwnedBy, "token is being not owned by [%s]", denom, owner.String())
}

func ErrTokenNotNFT(codespace sdk.CodespaceType, denom string) sdk.Error {
	return sdk.NewError(codespace, CodeTokenNotNF, "token [%s] is not a NFT", denom)
}

func ErrTokenNotCNFT(codespace sdk.CodespaceType, denom string) sdk.Error {
	return sdk.NewError(codespace, CodeTokenNotIDNF, "token [%s] is not a CNFT", denom)
}

func ErrTokenCannotTransferChildToken(codespace sdk.CodespaceType, denom string) sdk.Error {
	return sdk.NewError(codespace, CodeTokenChildNotTransferable, "cannot transfer a child token [%s]", denom)
}

func ErrCannotAttachToItself(codespace sdk.CodespaceType, denom string) sdk.Error {
	return sdk.NewError(codespace, CodeTokenCannotAttachToItself, "cannot attach token [%s] to itself", denom)
}

func ErrCannotAttachToADescendant(codespace sdk.CodespaceType, tokenDenom string, descendantDenom string) sdk.Error {
	return sdk.NewError(codespace, CodeTokenCannotAttachToADescendant, "cannot attach token [%s] to a descendant [%s]", tokenDenom, descendantDenom)
}
