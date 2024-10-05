package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/yueijdguedh/ReflexNet/x/modelregistry/types"
)

type queryServer struct {
	Keeper
}

// NewQueryServerImpl returns an implementation of the QueryServer interface
func NewQueryServerImpl(keeper Keeper) types.QueryServer {
	return &queryServer{Keeper: keeper}
}

var _ types.QueryServer = queryServer{}

// Params queries the parameters of the module
func (k queryServer) Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}

// Model queries a model by ID
func (k queryServer) Model(goCtx context.Context, req *types.QueryModelRequest) (*types.QueryModelResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	model, found := k.GetModel(ctx, req.ModelId)
	if !found {
		return nil, status.Error(codes.NotFound, "model not found")
	}

	return &types.QueryModelResponse{Model: model}, nil
}

// Models queries all models with pagination
func (k queryServer) Models(goCtx context.Context, req *types.QueryModelsRequest) (*types.QueryModelsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(k.storeKey)
	modelStore := prefix.NewStore(store, types.ModelKey)

	var models []types.Model
	pageRes, err := query.Paginate(modelStore, req.Pagination, func(key []byte, value []byte) error {
		var model types.Model
		if err := k.cdc.Unmarshal(value, &model); err != nil {
			return err
		}
		models = append(models, model)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryModelsResponse{
		Models:     models,
		Pagination: pageRes,
	}, nil
}

// ModelsByOwner queries all models owned by a specific address
func (k queryServer) ModelsByOwner(goCtx context.Context, req *types.QueryModelsByOwnerRequest) (*types.QueryModelsByOwnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	models := k.GetModelsByOwner(ctx, req.Owner)

	return &types.QueryModelsByOwnerResponse{
		Models: models,
	}, nil
}

