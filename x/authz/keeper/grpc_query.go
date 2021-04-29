package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/gogo/protobuf/proto"

	"github.com/line/lbm-sdk/codec"
	codectypes "github.com/line/lbm-sdk/codec/types"
	"github.com/line/lbm-sdk/store/prefix"
	sdk "github.com/line/lbm-sdk/types"
	"github.com/line/lbm-sdk/types/query"
	"github.com/line/lbm-sdk/x/authz/types"
)

var _ types.QueryServer = Keeper{}

// Authorizations implements the Query/Authorizations gRPC method.
func (k Keeper) Authorizations(c context.Context, req *types.QueryAuthorizationsRequest) (*types.QueryAuthorizationsResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	err := sdk.ValidateAccAddress(req.Granter)
	if err != nil {
		return nil, err
	}
	granter := sdk.AccAddress(req.Granter)

	err = sdk.ValidateAccAddress(req.Grantee)
	if err != nil {
		return nil, err
	}
	grantee := sdk.AccAddress(req.Grantee)

	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)
	key := types.GetAuthorizationStoreKey(grantee, granter, "")
	authStore := prefix.NewStore(store, key)
	var authorizations []*types.AuthorizationGrant
	pageRes, err := query.FilteredPaginate(authStore, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		auth, err := unmarshalAuthorization(k.cdc, value)
		if err != nil {
			return false, err
		}
		auth1 := auth.GetAuthorizationGrant()
		if accumulate {
			msg, ok := auth1.(proto.Message)
			if !ok {
				return false, status.Errorf(codes.Internal, "can't protomarshal %T", msg)
			}

			authorizationAny, err := codectypes.NewAnyWithValue(msg)
			if err != nil {
				return false, status.Errorf(codes.Internal, err.Error())
			}
			authorizations = append(authorizations, &types.AuthorizationGrant{
				Authorization: authorizationAny,
				Expiration:    auth.Expiration,
			})
		}
		return true, nil
	})
	if err != nil {
		return nil, err
	}

	return &types.QueryAuthorizationsResponse{
		Authorizations: authorizations,
		Pagination:     pageRes,
	}, nil
}

// unmarshal an authorization from a store value
func unmarshalAuthorization(cdc codec.BinaryCodec, value []byte) (v types.AuthorizationGrant, err error) {
	err = cdc.Unmarshal(value, &v)
	return v, err
}

// Authorization implements the Query/Authorization gRPC method.
func (k Keeper) Authorization(c context.Context, req *types.QueryAuthorizationRequest) (*types.QueryAuthorizationResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	if req.MethodName == "" {
		return nil, status.Errorf(codes.InvalidArgument, "empty method-name")
	}

	err := sdk.ValidateAccAddress(req.Granter)
	if err != nil {
		return nil, err
	}
	granter := sdk.AccAddress(req.Granter)

	err = sdk.ValidateAccAddress(req.Grantee)
	if err != nil {
		return nil, err
	}
	grantee := sdk.AccAddress(req.Grantee)

	ctx := sdk.UnwrapSDKContext(c)
	authorization, expiration := k.GetOrRevokeAuthorization(ctx, grantee, granter, req.MethodName)
	if authorization == nil {
		return nil, status.Errorf(codes.NotFound, "no authorization found for %s type", req.MethodName)
	}

	authorizationAny, err := codectypes.NewAnyWithValue(authorization)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &types.QueryAuthorizationResponse{
		Authorization: &types.AuthorizationGrant{
			Authorization: authorizationAny,
			Expiration:    expiration,
		},
	}, nil
}
