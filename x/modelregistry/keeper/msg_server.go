package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/yueijdguedh/ReflexNet/x/modelregistry/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// RegisterModel handles the registration of a new AI model
func (k msgServer) RegisterModel(goCtx context.Context, msg *types.MsgRegisterModel) (*types.MsgRegisterModelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate message
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Get module parameters
	params := k.GetParams(ctx)

	// Validate shard count
	if msg.ShardCount > params.MaxShardCount {
		return nil, types.ErrShardCountExceedsMax
	}

	// Get next model ID
	modelID := k.GetNextModelID(ctx)
	k.SetNextModelID(ctx, modelID+1)

	// Create model
	model := types.Model{
		Id:             modelID,
		Owner:          msg.Owner,
		Name:           msg.Name,
		MetadataCid:    msg.MetadataCid,
		ShardCount:     msg.ShardCount,
		Version:        msg.Version,
		CreatedAt:      ctx.BlockHeight(),
		UpdatedAt:      ctx.BlockHeight(),
		Status:         types.ModelStatus_MODEL_STATUS_ACTIVE,
		InferenceCount: 0,
	}

	// Store model
	k.SetModel(ctx, model)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"model_registered",
			sdk.NewAttribute("model_id", fmt.Sprintf("%d", modelID)),
			sdk.NewAttribute("owner", msg.Owner),
			sdk.NewAttribute("name", msg.Name),
		),
	)

	return &types.MsgRegisterModelResponse{
		ModelId: modelID,
	}, nil
}

// UpdateModelVersion handles updating a model's version
func (k msgServer) UpdateModelVersion(goCtx context.Context, msg *types.MsgUpdateModelVersion) (*types.MsgUpdateModelVersionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate message
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Get existing model
	model, found := k.GetModel(ctx, msg.ModelId)
	if !found {
		return nil, types.ErrModelNotFound
	}

	// Check ownership
	if model.Owner != msg.Owner {
		return nil, types.ErrUnauthorized
	}

	// Update model
	model.Version = msg.NewVersion
	model.MetadataCid = msg.NewMetadataCid
	model.UpdatedAt = ctx.BlockHeight()

	// Store updated model
	k.SetModel(ctx, model)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"model_version_updated",
			sdk.NewAttribute("model_id", fmt.Sprintf("%d", msg.ModelId)),
			sdk.NewAttribute("new_version", msg.NewVersion),
		),
	)

	return &types.MsgUpdateModelVersionResponse{}, nil
}

// UpdateModelStatus handles updating a model's status
func (k msgServer) UpdateModelStatus(goCtx context.Context, msg *types.MsgUpdateModelStatus) (*types.MsgUpdateModelStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate message
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Get existing model
	model, found := k.GetModel(ctx, msg.ModelId)
	if !found {
		return nil, types.ErrModelNotFound
	}

	// Check ownership
	if model.Owner != msg.Owner {
		return nil, types.ErrUnauthorized
	}

	// Update model status
	model.Status = msg.NewStatus
	model.UpdatedAt = ctx.BlockHeight()

	// Store updated model
	k.SetModel(ctx, model)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"model_status_updated",
			sdk.NewAttribute("model_id", fmt.Sprintf("%d", msg.ModelId)),
			sdk.NewAttribute("new_status", msg.NewStatus.String()),
		),
	)

	return &types.MsgUpdateModelStatusResponse{}, nil
}

