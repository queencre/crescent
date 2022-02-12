package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmosquad-labs/squad/x/liquidity/types"
)

func (k Keeper) ExecuteRequests(ctx sdk.Context) {
	if err := k.IterateAllPairs(ctx, func(pair types.Pair) (stop bool, err error) {
		if err := k.ExecuteMatching(ctx, pair); err != nil {
			return false, err
		}
		return false, nil
	}); err != nil {
		panic(err)
	}
	if err := k.IterateAllSwapRequests(ctx, func(req types.SwapRequest) (stop bool, err error) {
		if req.Status != types.SwapRequestStatusCompleted && !req.Status.IsCanceledOrExpired() && !ctx.BlockTime().Before(req.ExpireAt) { // ExpireAt <= BlockTime
			if err := k.FinishSwapRequest(ctx, req, types.SwapRequestStatusExpired); err != nil {
				return false, err
			}
		}
		return false, nil
	}); err != nil {
		panic(err)
	}
	if err := k.IterateAllDepositRequests(ctx, func(req types.DepositRequest) (stop bool, err error) {
		if req.Status == types.RequestStatusNotExecuted {
			if err := k.ExecuteDepositRequest(ctx, req); err != nil {
				return false, err
			}
		}
		return false, nil
	}); err != nil {
		panic(err)
	}
	if err := k.IterateAllWithdrawRequests(ctx, func(req types.WithdrawRequest) (stop bool, err error) {
		if req.Status == types.RequestStatusNotExecuted {
			if err := k.ExecuteWithdrawRequest(ctx, req); err != nil {
				return false, err
			}
		}
		return false, nil
	}); err != nil {
		panic(err)
	}
}

func (k Keeper) DeleteOutdatedRequests(ctx sdk.Context) {
	_ = k.IterateAllDepositRequests(ctx, func(req types.DepositRequest) (stop bool, err error) {
		if req.Status.ShouldBeDeleted() {
			k.DeleteDepositRequest(ctx, req.PoolId, req.Id)
		}
		return false, nil
	})
	_ = k.IterateAllWithdrawRequests(ctx, func(req types.WithdrawRequest) (stop bool, err error) {
		if req.Status.ShouldBeDeleted() {
			k.DeleteWithdrawRequest(ctx, req.PoolId, req.Id)
		}
		return false, nil
	})
	_ = k.IterateAllSwapRequests(ctx, func(req types.SwapRequest) (stop bool, err error) {
		if req.Status.ShouldBeDeleted() {
			k.DeleteSwapRequest(ctx, req.PairId, req.Id)
		}
		return false, nil
	})
}
