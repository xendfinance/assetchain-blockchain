// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sfclib100

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// SFCStateWithdrawalRequest is an auto generated low-level Go binding around an user-defined struct.
type SFCStateWithdrawalRequest struct {
	Epoch  *big.Int
	Time   *big.Int
	Amount *big.Int
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BurntRWA\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"name\":\"ChangedValidatorStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockupExtraReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockupBaseReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockedReward\",\"type\":\"uint256\"}],\"name\":\"ClaimedRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"auth\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"}],\"name\":\"CreatedValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"}],\"name\":\"DeactivatedValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Delegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LockedUpStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RefundedSlashedLegacyDelegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockupExtraReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockupBaseReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockedReward\",\"type\":\"uint256\"}],\"name\":\"RestakedRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Undelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"}],\"name\":\"UnlockedStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refundRatio\",\"type\":\"uint256\"}],\"name\":\"UpdatedSlashingRefundRatio\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"syncPubkey\",\"type\":\"bool\"}],\"name\":\"_syncValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentSealedEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getEpochSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBaseRewardWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalTxRewardWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRewardPerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"getLockedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getLockupInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lockedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getStashedLockupRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lockupExtraReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupBaseReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockedReward\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getValidator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"receivedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"auth\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getValidatorID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getValidatorPubkey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getWithdrawalRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"isLockedUp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastValidatorID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minGasPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"slashingRefundRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeTokenizerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stashedRewardsUntilEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalActiveStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSlashedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasuryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"voteBookAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getEpochValidatorIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochReceivedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochAccumulatedRewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochAccumulatedUptime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochAccumulatedOriginatedTxsFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochOfflineTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochOfflineBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"rewardsStash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"auth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"}],\"name\":\"setGenesisValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupFromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"earlyUnlockPenalty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"setGenesisDelegation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"createValidator\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getSelfStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validatorAuth\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"}],\"name\":\"recountVotes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"liquidateSRWA\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"}],\"name\":\"updateSRWAFinalizer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"isSlashed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"name\":\"deactivateValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"pendingRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"stashRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"claimRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"restakeRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnRWA\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"getUnlockedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"lockStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"relockStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unlockStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundRatio\",\"type\":\"uint256\"}],\"name\":\"updateSlashingRefundRatio\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getWrRequests\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structSFCState.WithdrawalRequest[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50617fef80620000216000396000f3fe60806040526004361061038c5760003560e01c8063854873e1116101dc578063bd14d90711610102578063cfdbb7cd116100a0578063df00c9221161006f578063df00c92214610e41578063e261641a14610e7e578063e81ee0c014610ebb578063f2fde38b14610ee45761038c565b8063cfdbb7cd14610d73578063d96ed50514610db0578063dc31e1af14610ddb578063de67f21514610e185761038c565b8063c65ee0e1116100dc578063c65ee0e114610ca5578063c7be95de14610ce2578063cc8343aa14610d0d578063cfd4766314610d365761038c565b8063bd14d90714610c14578063c3de580e14610c3d578063c5f956af14610c7a5761038c565b806396c7ee461161017a578063a86a056f11610149578063a86a056f14610b18578063b5d8962714610b55578063b810e41114610b98578063b88a37e214610bd75761038c565b806396c7ee4614610a635780639fa6dd3514610aa3578063a198d22914610abf578063a5a470ad14610afc5761038c565b80638cddb015116101b65780638cddb015146109bb5780638da5cb5b146109e45780638f32d59b14610a0f578063907754f814610a3a5761038c565b8063854873e114610928578063893675c6146109655780638b0e9f3f146109905761038c565b80634f7c4efb116102c1578063634b91e31161025f578063715018a61161022e578063715018a61461089257806376671808146108a95780637cacb1d6146108d457806381588ab6146108ff5761038c565b8063634b91e3146107b2578063670322f8146107db5780636f49866314610818578063702797e3146108555761038c565b806358f95b801161029b57806358f95b80146106d05780635fab23a81461070d5780636099ecb21461073857806361e53fcc146107755761038c565b80634f7c4efb146106415780634feb92f31461066a5780635601fe01146106935761038c565b80631d3ac42c1161032e57806320c0849d1161030857806320c0849d1461058157806328f73148146105aa57806339b80c00146105d5578063441a3e70146106185761038c565b80631d3ac42c146104dc5780631e702f83146105195780631f270152146105425761038c565b80630e559d821161036a5780630e559d821461042057806312622d0e1461044b57806318160ddd1461048857806318f628d4146104b35761038c565b80630135b1db1461039157806308c36874146103ce5780630962ef79146103f7575b600080fd5b34801561039d57600080fd5b506103b860048036036103b391908101906164b3565b610f0d565b6040516103c59190617bc7565b60405180910390f35b3480156103da57600080fd5b506103f560048036036103f0919081019061682f565b610f25565b005b34801561040357600080fd5b5061041e6004803603610419919081019061682f565b611044565b005b34801561042c57600080fd5b5061043561119e565b6040516104429190617612565b60405180910390f35b34801561045757600080fd5b50610472600480360361046d919081019061653f565b6111c4565b60405161047f9190617bc7565b60405180910390f35b34801561049457600080fd5b5061049d6112eb565b6040516104aa9190617bc7565b60405180910390f35b3480156104bf57600080fd5b506104da60048036036104d591908101906166fb565b6112f1565b005b3480156104e857600080fd5b5061050360048036036104fe91908101906168bd565b611520565b6040516105109190617bc7565b60405180910390f35b34801561052557600080fd5b50610540600480360361053b91908101906168bd565b61175e565b005b34801561054e57600080fd5b5061056960048036036105649190810190616649565b61184b565b60405161057893929190617c64565b60405180910390f35b34801561058d57600080fd5b506105a860048036036105a391908101906164dc565b61188f565b005b3480156105b657600080fd5b506105bf6119ff565b6040516105cc9190617bc7565b60405180910390f35b3480156105e157600080fd5b506105fc60048036036105f7919081019061682f565b611a05565b60405161060f9796959493929190617d4f565b60405180910390f35b34801561062457600080fd5b5061063f600480360361063a91908101906168bd565b611a47565b005b34801561064d57600080fd5b50610668600480360361066391908101906168bd565b611a57565b005b34801561067657600080fd5b50610691600480360361068c919081019061657b565b611b84565b005b34801561069f57600080fd5b506106ba60048036036106b5919081019061682f565b611c3d565b6040516106c79190617bc7565b60405180910390f35b3480156106dc57600080fd5b506106f760048036036106f291908101906168bd565b611ccd565b6040516107049190617bc7565b60405180910390f35b34801561071957600080fd5b50610722611cff565b60405161072f9190617bc7565b60405180910390f35b34801561074457600080fd5b5061075f600480360361075a919081019061653f565b611d05565b60405161076c9190617bc7565b60405180910390f35b34801561078157600080fd5b5061079c600480360361079791908101906168bd565b611d55565b6040516107a99190617bc7565b60405180910390f35b3480156107be57600080fd5b506107d960048036036107d491908101906168bd565b611d87565b005b3480156107e757600080fd5b5061080260048036036107fd919081019061653f565b612130565b60405161080f9190617bc7565b60405180910390f35b34801561082457600080fd5b5061083f600480360361083a919081019061653f565b6121a6565b60405161084c9190617bc7565b60405180910390f35b34801561086157600080fd5b5061087c60048036036108779190810190616698565b612265565b60405161088991906176c4565b60405180910390f35b34801561089e57600080fd5b506108a761238e565b005b3480156108b557600080fd5b506108be612496565b6040516108cb9190617bc7565b60405180910390f35b3480156108e057600080fd5b506108e96124a3565b6040516108f69190617bc7565b60405180910390f35b34801561090b57600080fd5b50610926600480360361092191908101906164b3565b6124a9565b005b34801561093457600080fd5b5061094f600480360361094a919081019061682f565b612534565b60405161095c9190617723565b60405180910390f35b34801561097157600080fd5b5061097a6125e4565b6040516109879190617612565b60405180910390f35b34801561099c57600080fd5b506109a561260a565b6040516109b29190617bc7565b60405180910390f35b3480156109c757600080fd5b506109e260048036036109dd919081019061653f565b612610565b005b3480156109f057600080fd5b506109f961265d565b604051610a069190617612565b60405180910390f35b348015610a1b57600080fd5b50610a24612687565b604051610a319190617708565b60405180910390f35b348015610a4657600080fd5b50610a616004803603610a5c9190810190616649565b6126df565b005b348015610a6f57600080fd5b50610a8a6004803603610a85919081019061653f565b612b39565b604051610a9a9493929190617c9b565b60405180910390f35b610abd6004803603610ab8919081019061682f565b612b76565b005b348015610acb57600080fd5b50610ae66004803603610ae191908101906168bd565b612b84565b604051610af39190617bc7565b60405180910390f35b610b166004803603610b1191908101906167ea565b612bb6565b005b348015610b2457600080fd5b50610b3f6004803603610b3a919081019061653f565b612d3d565b604051610b4c9190617bc7565b60405180910390f35b348015610b6157600080fd5b50610b7c6004803603610b77919081019061682f565b612d62565b604051610b8f9796959493929190617ce0565b60405180910390f35b348015610ba457600080fd5b50610bbf6004803603610bba919081019061653f565b612dc4565b604051610bce93929190617c64565b60405180910390f35b348015610be357600080fd5b50610bfe6004803603610bf9919081019061682f565b612dfb565b604051610c0b91906176e6565b60405180910390f35b348015610c2057600080fd5b50610c3b6004803603610c3691908101906168f9565b612e69565b005b348015610c4957600080fd5b50610c646004803603610c5f919081019061682f565b612e80565b604051610c719190617708565b60405180910390f35b348015610c8657600080fd5b50610c8f612ea6565b604051610c9c9190617612565b60405180910390f35b348015610cb157600080fd5b50610ccc6004803603610cc7919081019061682f565b612ecc565b604051610cd99190617bc7565b60405180910390f35b348015610cee57600080fd5b50610cf7612ee4565b604051610d049190617bc7565b60405180910390f35b348015610d1957600080fd5b50610d346004803603610d2f9190810190616881565b612eea565b005b348015610d4257600080fd5b50610d5d6004803603610d58919081019061653f565b6130b9565b604051610d6a9190617bc7565b60405180910390f35b348015610d7f57600080fd5b50610d9a6004803603610d95919081019061653f565b6130de565b604051610da79190617708565b60405180910390f35b348015610dbc57600080fd5b50610dc5613205565b604051610dd29190617bc7565b60405180910390f35b348015610de757600080fd5b50610e026004803603610dfd91908101906168bd565b61320b565b604051610e0f9190617bc7565b60405180910390f35b348015610e2457600080fd5b50610e3f6004803603610e3a91908101906168f9565b61323d565b005b348015610e4d57600080fd5b50610e686004803603610e6391908101906168bd565b6132e1565b604051610e759190617bc7565b60405180910390f35b348015610e8a57600080fd5b50610ea56004803603610ea091908101906168bd565b613313565b604051610eb29190617bc7565b60405180910390f35b348015610ec757600080fd5b50610ee26004803603610edd919081019061682f565b613345565b005b348015610ef057600080fd5b50610f0b6004803603610f0691908101906164b3565b613398565b005b60696020528060005260406000206000915090505481565b6000339050610f326162f8565b610f3c82846133eb565b90506000610f5b826020015183600001516135b990919063ffffffff16565b9050610f7e8385610f798560400151856135b990919063ffffffff16565b61360e565b80607460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600086815260200190815260200160002060000160008282540192505081905550838373ffffffffffffffffffffffffffffffffffffffff167f4119153d17a36f9597d40e3ab4148d03261a439dddbec4e91799ab7159608e2684600001518560200151866040015160405161103693929190617c64565b60405180910390a350505050565b60003390506110516162f8565b61105b82846133eb565b905060008273ffffffffffffffffffffffffffffffffffffffff166110a78360400151611099856020015186600001516135b990919063ffffffff16565b6135b990919063ffffffff16565b6040516110b3906175fd565b60006040518083038185875af1925050503d80600081146110f0576040519150601f19603f3d011682016040523d82523d6000602084013e6110f5565b606091505b5050905080611139576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161113090617aa7565b60405180910390fd5b838373ffffffffffffffffffffffffffffffffffffffff167fc1d8eb6e444b89fb8ff0991c19311c070df704ccb009e210d1462d5b2410bf4584600001518560200151866040015160405161119093929190617c64565b60405180910390a350505050565b607c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006111d083836130de565b61122c57607360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008381526020019081526020016000205490506112e5565b6112e2607460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600084815260200190815260200160002060000154607360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008581526020019081526020016000205461370990919063ffffffff16565b90505b92915050565b60775481565b6112fa33613753565b611339576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161133090617807565b60405180910390fd5b61134689898960006137ad565b80606f60008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a8152602001908152602001600020600201819055506113a7876139f4565b6000861461151557868611156113f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113e990617ba7565b60405180910390fd5b6000607460008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a8152602001908152602001600020905086816000018190555085816001018190555084816002018190555083816003018190555082607560008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008b815260200190815260200160002060000181905550888a73ffffffffffffffffffffffffffffffffffffffff167f138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e78868a60405161150b929190617c3b565b60405180910390a3505b505050505050505050565b6000803390506000607460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008681526020019081526020016000209050600084116115bd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115b490617787565b60405180910390fd5b6115c782866130de565b611606576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115fd90617907565b60405180910390fd5b806000015484111561164d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161164490617a87565b60405180910390fd5b6116578286613aa1565b611696576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161168d906179c7565b60405180910390fd5b6116a08286613bb8565b5060006116b38387878560000154613f30565b90506363401ec5826003015401826002015410156116d057600090505b84826000016000828254039250508190555060008114611701576116f783878360016141cd565b6117008161447c565b5b858373ffffffffffffffffffffffffffffffffffffffff167fef6c0c14fe9aa51af36acd791464dec3badbde668b63189b47bfa4e25be9b2b9878460405161174a929190617c3b565b60405180910390a380935050505092915050565b61176733613753565b6117a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161179d90617807565b60405180910390fd5b60008114156117ea576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117e1906178a7565b60405180910390fd5b6117f48282614507565b6117ff826000612eea565b60006068600084815260200190815260200160002060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050611846818260006146a5565b505050565b607160205282600052604060002060205281600052604060002060205280600052604060002060009250925050508060000154908060010154908060020154905083565b6000608360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168286866040516024016118de929190617672565b6040516020818303038152906040527f4a7702bb000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505060405161196891906175e6565b60006040518083038160008787f1925050503d80600081146119a6576040519150601f19603f3d011682016040523d82523d6000602084013e6119ab565b606091505b5050905080806119b9575082155b6119f8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119ef90617b07565b60405180910390fd5b5050505050565b606d5481565b607860205280600052604060002060009150905080600701549080600801549080600901549080600a01549080600b01549080600c01549080600d0154905087565b611a533383833361486e565b5050565b611a5f612687565b611a9e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a95906179a7565b60405180910390fd5b611aa782612e80565b611ae6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611add90617767565b60405180910390fd5b611aee614e79565b811115611b30576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b27906179e7565b60405180910390fd5b80607b600084815260200190815260200160002081905550817f047575f43f09a7a093d94ec483064acfc61b7e25c0de28017da442abf99cb91782604051611b789190617bc7565b60405180910390a25050565b611b8d33613753565b611bcc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611bc390617807565b60405180910390fd5b611c20898989898080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508888888888614e89565b606b54881115611c325787606b819055505b505050505050505050565b6000607360006068600085815260200190815260200160002060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000838152602001908152602001600020549050919050565b600060786000848152602001908152602001600020600001600083815260200190815260200160002054905092915050565b606e5481565b6000611d0f6162f8565b611d198484615132565b9050611d4c8160000151611d3e836020015184604001516135b990919063ffffffff16565b6135b990919063ffffffff16565b91505092915050565b600060786000848152602001908152602001600020600101600083815260200190815260200160002054905092915050565b6000339050611d968184613bb8565b5060008211611dda576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611dd190617787565b60405180910390fd5b611de481846111c4565b821115611e26576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e1d90617a07565b60405180910390fd5b611e308184613aa1565b611e6f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e66906179c7565b60405180910390fd5b6000607260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600085815260200190815260200160002060008154809291906001019190505590506000607160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600086815260200190815260200160002060008381526020019081526020016000206002015414611f79576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f7090617987565b60405180910390fd5b611f8682858560016141cd565b82607160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000868152602001908152602001600020600083815260200190815260200160002060020181905550611ff7612496565b607160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008681526020019081526020016000206000838152602001908152602001600020600001819055506120676151d9565b607160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008681526020019081526020016000206000838152602001908152602001600020600101819055506120da846000612eea565b80848373ffffffffffffffffffffffffffffffffffffffff167fd3bb4e423fbea695d16b982f9f682dc5f35152e5411646a8a5a79a6b02ba8d57866040516121229190617bc7565b60405180910390a450505050565b600061213c83836130de565b61214957600090506121a0565b607460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008381526020019081526020016000206000015490505b92915050565b60006121b06162f8565b606f60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000206040518060600160405290816000820154815260200160018201548152602001600282015481525050905061225c816040015161224e836000015184602001516135b990919063ffffffff16565b6135b990919063ffffffff16565b91505092915050565b606080826040519080825280602002602001820160405280156122a257816020015b61228f616319565b8152602001906001900390816122875790505b50905060008090505b8381101561238157607160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000878152602001908152602001600020600061231883886135b990919063ffffffff16565b8152602001908152602001600020604051806060016040529081600082015481526020016001820154815260200160028201548152505082828151811061235b57fe5b602002602001018190525061237a6001826135b990919063ffffffff16565b90506122ab565b5080915050949350505050565b612396612687565b6123d5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016123cc906179a7565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000603360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000600160675401905090565b60675481565b6124b1612687565b6124f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016124e7906179a7565b60405180910390fd5b80608460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b606a6020528060005260406000206000915090508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156125dc5780601f106125b1576101008083540402835291602001916125dc565b820191906000526020600020905b8154815290600101906020018083116125bf57829003601f168201915b505050505081565b608360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b606c5481565b61261a8282613bb8565b612659576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161265090617827565b60405180910390fd5b5050565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b608460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461276f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161276690617887565b60405180910390fd5b6127798383613bb8565b50600081116127bd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016127b490617787565b60405180910390fd5b607c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166378542c88338585856040518563ffffffff1660e01b815260040161281e949392919061762d565b600060405180830381600087803b15801561283857600080fd5b505af115801561284c573d6000803e3d6000fd5b50505050607360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000838152602001908152602001600020548111156128e3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128da90617ac7565b60405180910390fd5b60006128ef84846111c4565b9050808210156129c4576000607460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008581526020019081526020016000209050612966828403826000015461370990919063ffffffff16565b8160000181905550838573ffffffffffffffffffffffffffffffffffffffff167fef6c0c14fe9aa51af36acd791464dec3badbde668b63189b47bfa4e25be9b2b984860360006040516129ba929190617c12565b60405180910390a3505b6129d184848460016141cd565b6129dc836000612eea565b64ffffffffff838573ffffffffffffffffffffffffffffffffffffffff167fd3bb4e423fbea695d16b982f9f682dc5f35152e5411646a8a5a79a6b02ba8d5785604051612a299190617bc7565b60405180910390a460003373ffffffffffffffffffffffffffffffffffffffff1683604051612a57906175fd565b60006040518083038185875af1925050503d8060008114612a94576040519150601f19603f3d011682016040523d82523d6000602084013e612a99565b606091505b5050905080612add576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612ad490617aa7565b60405180910390fd5b64ffffffffff848673ffffffffffffffffffffffffffffffffffffffff167f75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a2186604051612b2a9190617bc7565b60405180910390a45050505050565b6074602052816000526040600020602052806000526040600020600091509150508060000154908060010154908060020154908060030154905084565b612b8133823461360e565b50565b600060786000848152602001908152602001600020600501600083815260200190815260200160002054905092915050565b608260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c5f530af6040518163ffffffff1660e01b815260040160206040518083038186803b158015612c1e57600080fd5b505afa158015612c32573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250612c569190810190616858565b341015612c98576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c8f90617b67565b60405180910390fd5b60008282905011612cde576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612cd590617a27565b60405180910390fd5b612d2c3383838080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050506151e1565b612d3933606b543461360e565b5050565b6070602052816000526040600020602052806000526040600020600091509150505481565b60686020528060005260406000206000915090508060000154908060010154908060020154908060030154908060040154908060050154908060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905087565b6075602052816000526040600020602052806000526040600020600091509150508060000154908060010154908060020154905083565b606060786000838152602001908152602001600020600601805480602002602001604051908101604052809291908181526020018280548015612e5d57602002820191906000526020600020905b815481526020019060010190808311612e49575b50505050509050919050565b6000339050612e7a81858585615219565b50505050565b600080608060686000858152602001908152602001600020600001541614159050919050565b608060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b607b6020528060005260406000206000915090505481565b606b5481565b612ef3826156af565b612f32576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612f2990617b27565b60405180910390fd5b6000606860008481526020019081526020016000206003015490506000606860008581526020019081526020016000206000015414612f7057600090505b606660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a4066fbe84836040518363ffffffff1660e01b8152600401612fcd929190617c3b565b600060405180830381600087803b158015612fe757600080fd5b505af1158015612ffb573d6000803e3d6000fd5b5050505081801561300d575060008114155b156130b457606660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663242a6e3f84606a60008781526020019081526020016000206040518363ffffffff1660e01b8152600401613081929190617be2565b600060405180830381600087803b15801561309b57600080fd5b505af11580156130af573d6000803e3d6000fd5b505050505b505050565b6073602052816000526040600020602052806000526040600020600091509150505481565b600080607460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000848152602001908152602001600020600201541415801561319757506000607460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000206000015414155b80156131fd5750607460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000838152602001908152602001600020600201546131fa6151d9565b11155b905092915050565b607f5481565b600060786000848152602001908152602001600020600301600083815260200190815260200160002054905092915050565b600033905060008211613285576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161327c90617787565b60405180910390fd5b61328f81856130de565b156132cf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016132c6906177c7565b60405180910390fd5b6132db81858585615219565b50505050565b600060786000848152602001908152602001600020600201600083815260200190815260200160002054905092915050565b600060786000848152602001908152602001600020600401600083815260200190815260200160002054905092915050565b61334d612687565b61338c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613383906179a7565b60405180910390fd5b6133958161447c565b50565b6133a0612687565b6133df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016133d6906179a7565b60405180910390fd5b6133e8816156d2565b50565b6133f36162f8565b6133fd8383613aa1565b61343c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613433906179c7565b60405180910390fd5b6134468383613bb8565b50606f60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008381526020019081526020016000206040518060600160405290816000820154815260200160018201548152602001600282015481525050905060006134f582600001516134e7846020015185604001516135b990919063ffffffff16565b6135b990919063ffffffff16565b9050600081141561353b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161353290617947565b60405180910390fd5b606f60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000206000808201600090556001820160009055600282016000905550506135af816139f4565b8191505092915050565b600080828401905083811015613604576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016135fb906177e7565b60405180910390fd5b8091505092915050565b613617826156af565b613656576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161364d90617b27565b60405180910390fd5b60006068600084815260200190815260200160002060000154146136af576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016136a690617867565b60405180910390fd5b6136bc83838360016137ad565b6136c582615802565b613704576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016136fb90617ae7565b60405180910390fd5b505050565b600061374b83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506158f7565b905092915050565b6000606660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16149050919050565b600082116137f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016137e790617787565b60405180910390fd5b6137fa8484613bb8565b5061385e82607360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000868152602001908152602001600020546135b990919063ffffffff16565b607360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000858152602001908152602001600020819055506000606860008581526020019081526020016000206003015490506138e083826135b990919063ffffffff16565b606860008681526020019081526020016000206003018190555061390f83606c546135b990919063ffffffff16565b606c819055506000606860008681526020019081526020016000206000015414156139505761394983606d546135b990919063ffffffff16565b606d819055505b61395d8460008314612eea565b838573ffffffffffffffffffffffffffffffffffffffff167f9a8f44850296624dadfd9c246d17e47171d35727a181bd090aa14bbbe00238bb856040516139a49190617bc7565b60405180910390a36139ed856068600087815260200190815260200160002060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846146a5565b5050505050565b606660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166366e7ea0f30836040518363ffffffff1660e01b8152600401613a5192919061769b565b600060405180830381600087803b158015613a6b57600080fd5b505af1158015613a7f573d6000803e3d6000fd5b50505050613a98816077546135b990919063ffffffff16565b60778190555050565b60008073ffffffffffffffffffffffffffffffffffffffff16607c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415613b025760019050613bb2565b607c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166321d585c384846040518363ffffffff1660e01b8152600401613b5f92919061769b565b60206040518083038186803b158015613b7757600080fd5b505afa158015613b8b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250613baf91908101906167c1565b90505b92915050565b6000613bc26162f8565b613bcc8484615952565b9050613bd783615b35565b607060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600085815260200190815260200160002081905550613cad606f60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000858152602001908152602001600020604051806060016040529081600082015481526020016001820154815260200160028201548152505082615ba7565b606f60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000858152602001908152602001600020600082015181600001556020820151816001015560408201518160020155905050613da0607560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000858152602001908152602001600020604051806060016040529081600082015481526020016001820154815260200160028201548152505082615ba7565b607560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000858152602001908152602001600020600082015181600001556020820151816001015560408201518160020155905050613e1b84846130de565b613efe57607460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000206000808201600090556001820160009055600282016000905560038201600090555050607560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000206000808201600090556001820160009055600282016000905550505b60008160200151141580613f1757506000816000015114155b80613f2757506000816040015114155b91505092915050565b600080613fab83613f9d86607560008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a815260200190815260200160002060000154615c2090919063ffffffff16565b615c9090919063ffffffff16565b905060006140278461401987607560008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008b815260200190815260200160002060010154615c2090919063ffffffff16565b615c9090919063ffffffff16565b905060006002828161403557fe5b04830190506140a083607560008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a81526020019081526020016000206000015461370990919063ffffffff16565b607560008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008981526020019081526020016000206000018190555061415d82607560008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a81526020019081526020016000206001015461370990919063ffffffff16565b607560008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000898152602001908152602001600020600101819055508581106141bf578590505b809350505050949350505050565b81607360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008581526020019081526020016000206000828254039250508190555061425482606860008681526020019081526020016000206003015461370990919063ffffffff16565b606860008581526020019081526020016000206003018190555061428382606c5461370990919063ffffffff16565b606c819055506000606860008581526020019081526020016000206000015414156142c4576142bd82606d5461370990919063ffffffff16565b606d819055505b60006142cf84611c3d565b9050600081146144285760006068600086815260200190815260200160002060000154141561442357608260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c5f530af6040518163ffffffff1660e01b815260040160206040518083038186803b15801561436057600080fd5b505afa158015614374573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506143989190810190616858565b8110156143da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016143d190617b67565b60405180910390fd5b6143e384615802565b614422576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161441990617ae7565b60405180910390fd5b5b614434565b614433846001614507565b5b614475856068600087815260200190815260200160002060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846146a5565b5050505050565b6000811461450457600073ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156144cb573d6000803e3d6000fd5b507fe733dd86f3139a17476f9309db31e45d999e6469068d5734f7d51b23110c17bc816040516144fb9190617bc7565b60405180910390a15b50565b6000606860008481526020019081526020016000206000015414801561452e575060008114155b156145655761455e6068600084815260200190815260200160002060030154606d5461370990919063ffffffff16565b606d819055505b60686000838152602001908152602001600020600001548111156146a157806068600084815260200190815260200160002060000181905550600060686000848152602001908152602001600020600201541415614668576145c5612496565b60686000848152602001908152602001600020600201819055506145e76151d9565b6068600084815260200190815260200160002060010181905550817fac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e476068600085815260200190815260200160002060020154606860008681526020019081526020016000206001015460405161465f929190617c3b565b60405180910390a25b817fcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e826040516146989190617bc7565b60405180910390a25b5050565b600073ffffffffffffffffffffffffffffffffffffffff16608360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614614869576000608360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16627a1200858560405160240161474d929190617672565b6040516020818303038152906040527f4a7702bb000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516147d791906175e6565b60006040518083038160008787f1925050503d8060008114614815576040519150601f19603f3d011682016040523d82523d6000602084013e61481a565b606091505b505090508080614828575081155b614867576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161485e90617b07565b60405180910390fd5b505b505050565b61487661633a565b607160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600085815260200190815260200160002060008481526020019081526020016000206040518060600160405290816000820154815260200160018201548152602001600282015481525050905060008160000151141561494a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161494190617a47565b60405180910390fd5b6149548585613aa1565b614993576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161498a906179c7565b60405180910390fd5b60008160200151905060008260000151905060006068600088815260200190815260200160002060010154141580156149e15750816068600088815260200190815260200160002060010154105b15614a195760686000878152602001908152602001600020600101549150606860008781526020019081526020016000206002015490505b608260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b82b84276040518163ffffffff1660e01b815260040160206040518083038186803b158015614a8157600080fd5b505afa158015614a95573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250614ab99190810190616858565b8201614ac36151d9565b1015614b04576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614afb906178c7565b60405180910390fd5b608260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663650acd666040518163ffffffff1660e01b815260040160206040518083038186803b158015614b6c57600080fd5b505afa158015614b80573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250614ba49190810190616858565b8101614bae612496565b1015614bef576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614be690617b47565b60405180910390fd5b6000607160008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600088815260200190815260200160002060008781526020019081526020016000206002015490506000614c6388612e80565b90506000614c858383607b60008d815260200190815260200160002054615cda565b9050607160008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a8152602001908152602001600020600089815260200190815260200160002060008082016000905560018201600090556002820160009055505080606e60008282540192505081905550808311614d55576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614d4c90617b87565b60405180910390fd5b60008773ffffffffffffffffffffffffffffffffffffffff16614d81838661370990919063ffffffff16565b604051614d8d906175fd565b60006040518083038185875af1925050503d8060008114614dca576040519150601f19603f3d011682016040523d82523d6000602084013e614dcf565b606091505b5050905080614e13576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614e0a90617aa7565b60405180910390fd5b614e1c8261447c565b888a8c73ffffffffffffffffffffffffffffffffffffffff167f75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a2187604051614e649190617bc7565b60405180910390a45050505050505050505050565b6000670de0b6b3a7640000905090565b6000606960008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414614f0b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614f02906178e7565b60405180910390fd5b86606960008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550846068600089815260200190815260200160002060000181905550836068600089815260200190815260200160002060040181905550826068600089815260200190815260200160002060050181905550806068600089815260200190815260200160002060010181905550816068600089815260200190815260200160002060020181905550876068600089815260200190815260200160002060060160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555085606a6000898152602001908152602001600020908051906020019061505292919061635b565b508773ffffffffffffffffffffffffffffffffffffffff16877f49bca1ed2666922f9f1690c26a569e1299c2a715fe57647d77e81adfabbf25bf868660405161509c929190617c3b565b60405180910390a3600082146150e757867fac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e4783836040516150de929190617c3b565b60405180910390a25b6000851461512857867fcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e8660405161511f9190617bc7565b60405180910390a25b5050505050505050565b61513a6162f8565b6151426162f8565b61514c8484615952565b90506151d0606f60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000858152602001908152602001600020604051806060016040529081600082015481526020016001820154815260200160028201548152505082615ba7565b91505092915050565b600042905090565b6000606b6000815460010191905081905590506152148382846000615204612496565b61520c6151d9565b600080614e89565b505050565b61522384846111c4565b811115615265576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161525c90617ac7565b60405180910390fd5b60006068600085815260200190815260200160002060000154146152be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016152b590617867565b60405180910390fd5b608260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630d7b26096040518163ffffffff1660e01b815260040160206040518083038186803b15801561532657600080fd5b505afa15801561533a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061535e9190810190616858565b821015801561540c5750608260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630d4955e36040518163ffffffff1660e01b815260040160206040518083038186803b1580156153d057600080fd5b505afa1580156153e4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506154089190810190616858565b8211155b61544b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161544290617847565b60405180910390fd5b6000615467836154596151d9565b6135b990919063ffffffff16565b905060006068600086815260200190815260200160002060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff161461556e5781607460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600087815260200190815260200160002060020154101561556d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161556490617a67565b60405180910390fd5b5b6155788686613bb8565b506000607460008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600087815260200190815260200160002090508060030154851015615614576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161560b90617967565b60405180910390fd5b61562b8482600001546135b990919063ffffffff16565b816000018190555061563b612496565b8160010181905550828160020181905550848160030181905550858773ffffffffffffffffffffffffffffffffffffffff167f138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e78878760405161569e929190617c3b565b60405180910390a350505050505050565b600080606860008481526020019081526020016000206005015414159050919050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415615742576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615739906177a7565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380603360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60006158d761580f614e79565b6158c9608260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632265f2846040518163ffffffff1660e01b815260040160206040518083038186803b15801561587a57600080fd5b505afa15801561588e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506158b29190810190616858565b6158bb86611c3d565b615c2090919063ffffffff16565b615c9090919063ffffffff16565b606860008481526020019081526020016000206003015411159050919050565b600083831115829061593f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016159369190617745565b60405180910390fd5b5060008385039050809150509392505050565b61595a6162f8565b6000607060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600084815260200190815260200160002054905060006159ba84615b35565b905060006159c88686615d63565b9050818111156159d6578190505b828110156159e2578290505b6000607460008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600087815260200190815260200160002090506000607360008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008881526020019081526020016000205490506000615aa483600001548361370990919063ffffffff16565b90506000615ab884600001548a8988615e6c565b9050615ac26162f8565b615ad0828660030154615f21565b9050615ade838b8a89615e6c565b9150615ae86162f8565b615af3836000615f21565b9050615b01858c898b615e6c565b9250615b0b6162f8565b615b16846000615f21565b9050615b2383838361618a565b9a505050505050505050505092915050565b600080606860008481526020019081526020016000206002015414615b9c5760686000838152602001908152602001600020600201546067541015615b7e576067549050615ba2565b60686000838152602001908152602001600020600201549050615ba2565b60675490505b919050565b615baf6162f8565b6040518060600160405280615bd5846000015186600001516135b990919063ffffffff16565b8152602001615bf5846020015186602001516135b990919063ffffffff16565b8152602001615c15846040015186604001516135b990919063ffffffff16565b815250905092915050565b600080831415615c335760009050615c8a565b6000828402905082848281615c4457fe5b0414615c85576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615c7c90617927565b60405180910390fd5b809150505b92915050565b6000615cd283836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f0000000000008152506161ae565b905092915050565b6000821580615cf05750615cec614e79565b8210155b15615cfe5760009050615d5c565b615d466001615d38615d0e614e79565b615d2a86615d1a614e79565b0389615c2090919063ffffffff16565b615c9090919063ffffffff16565b6135b990919063ffffffff16565b905083811115615d5857839050615d5c565b8090505b9392505050565b600080607460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600084815260200190815260200160002060010154905060006067549050615dce85858361620f565b15615ddd578092505050615e66565b615de885858461620f565b615df757600092505050615e66565b80821115615e0a57600092505050615e66565b5b80821015615e49576000600282840181615e2157fe5b049050615e2f86868361620f565b15615e3f57600181019250615e43565b8091505b50615e0b565b6000811415615e5d57600092505050615e66565b60018103925050505b92915050565b6000818310615e7e5760009050615f19565b60006078600085815260200190815260200160002060010160008681526020019081526020016000205490506000607860008581526020019081526020016000206001016000878152602001908152602001600020549050615f14615ee1614e79565b615f0689615ef8868661370990919063ffffffff16565b615c2090919063ffffffff16565b615c9090919063ffffffff16565b925050505b949350505050565b615f296162f8565b60405180606001604052806000815260200160008152602001600081525090506000608260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635e2308d26040518163ffffffff1660e01b815260040160206040518083038186803b158015615fb357600080fd5b505afa158015615fc7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250615feb9190810190616858565b90506000831461614a57600081616000614e79565b03905060006160c9608260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630d4955e36040518163ffffffff1660e01b815260040160206040518083038186803b15801561607057600080fd5b505afa158015616084573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506160a89190810190616858565b6160bb8785615c2090919063ffffffff16565b615c9090919063ffffffff16565b905060006160fb6160d8614e79565b6160ed8487018a615c2090919063ffffffff16565b615c9090919063ffffffff16565b9050616129616108614e79565b61611b868a615c2090919063ffffffff16565b615c9090919063ffffffff16565b85602001818152505084602001518103856000018181525050505050616180565b616176616155614e79565b6161688387615c2090919063ffffffff16565b615c9090919063ffffffff16565b8260400181815250505b8191505092915050565b6161926162f8565b6161a561619f8585615ba7565b83615ba7565b90509392505050565b600080831182906161f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016161ec9190617745565b60405180910390fd5b50600083858161620157fe5b049050809150509392505050565b600081607460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600085815260200190815260200160002060010154111580156162cf5750607460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000848152602001908152602001600020600201546162cc836162d8565b11155b90509392505050565b600060786000838152602001908152602001600020600701549050919050565b60405180606001604052806000815260200160008152602001600081525090565b60405180606001604052806000815260200160008152602001600081525090565b60405180606001604052806000815260200160008152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061639c57805160ff19168380011785556163ca565b828001600101855582156163ca579182015b828111156163c95782518255916020019190600101906163ae565b5b5090506163d791906163db565b5090565b6163fd91905b808211156163f95760008160009055506001016163e1565b5090565b90565b60008135905061640f81617f67565b92915050565b60008135905061642481617f7e565b92915050565b60008151905061643981617f7e565b92915050565b60008083601f84011261645157600080fd5b8235905067ffffffffffffffff81111561646a57600080fd5b60208301915083600182028301111561648257600080fd5b9250929050565b60008135905061649881617f95565b92915050565b6000815190506164ad81617f95565b92915050565b6000602082840312156164c557600080fd5b60006164d384828501616400565b91505092915050565b600080600080608085870312156164f257600080fd5b600061650087828801616400565b945050602061651187828801616400565b935050604061652287828801616415565b925050606061653387828801616489565b91505092959194509250565b6000806040838503121561655257600080fd5b600061656085828601616400565b925050602061657185828601616489565b9150509250929050565b60008060008060008060008060006101008a8c03121561659a57600080fd5b60006165a88c828d01616400565b99505060206165b98c828d01616489565b98505060408a013567ffffffffffffffff8111156165d657600080fd5b6165e28c828d0161643f565b975097505060606165f58c828d01616489565b95505060806166068c828d01616489565b94505060a06166178c828d01616489565b93505060c06166288c828d01616489565b92505060e06166398c828d01616489565b9150509295985092959850929598565b60008060006060848603121561665e57600080fd5b600061666c86828701616400565b935050602061667d86828701616489565b925050604061668e86828701616489565b9150509250925092565b600080600080608085870312156166ae57600080fd5b60006166bc87828801616400565b94505060206166cd87828801616489565b93505060406166de87828801616489565b92505060606166ef87828801616489565b91505092959194509250565b60008060008060008060008060006101208a8c03121561671a57600080fd5b60006167288c828d01616400565b99505060206167398c828d01616489565b985050604061674a8c828d01616489565b975050606061675b8c828d01616489565b965050608061676c8c828d01616489565b95505060a061677d8c828d01616489565b94505060c061678e8c828d01616489565b93505060e061679f8c828d01616489565b9250506101006167b18c828d01616489565b9150509295985092959850929598565b6000602082840312156167d357600080fd5b60006167e18482850161642a565b91505092915050565b600080602083850312156167fd57600080fd5b600083013567ffffffffffffffff81111561681757600080fd5b6168238582860161643f565b92509250509250929050565b60006020828403121561684157600080fd5b600061684f84828501616489565b91505092915050565b60006020828403121561686a57600080fd5b60006168788482850161649e565b91505092915050565b6000806040838503121561689457600080fd5b60006168a285828601616489565b92505060206168b385828601616415565b9150509250929050565b600080604083850312156168d057600080fd5b60006168de85828601616489565b92505060206168ef85828601616489565b9150509250929050565b60008060006060848603121561690e57600080fd5b600061691c86828701616489565b935050602061692d86828701616489565b925050604061693e86828701616489565b9150509250925092565b60006169548383617586565b60608301905092915050565b600061696c83836175c8565b60208301905092915050565b61698181617edb565b82525050565b61699081617e93565b82525050565b60006169a182617df3565b6169ab8185617e44565b93506169b683617dbe565b8060005b838110156169e75781516169ce8882616948565b97506169d983617e2a565b9250506001810190506169ba565b5085935050505092915050565b60006169ff82617dfe565b616a098185617e55565b9350616a1483617dce565b8060005b83811015616a45578151616a2c8882616960565b9750616a3783617e37565b925050600181019050616a18565b5085935050505092915050565b616a5b81617ea5565b82525050565b6000616a6c82617e14565b616a768185617e77565b9350616a86818560208601617f23565b80840191505092915050565b6000616a9d82617e09565b616aa78185617e66565b9350616ab7818560208601617f23565b616ac081617f56565b840191505092915050565b600081546001811660008114616ae85760018114616b0e57616b52565b607f6002830416616af98187617e66565b955060ff198316865260208601935050616b52565b60028204616b1c8187617e66565b9550616b2785617dde565b60005b82811015616b4957815481890152600182019150602081019050616b2a565b80880195505050505b505092915050565b616b6381617eed565b82525050565b6000616b7482617e1f565b616b7e8185617e82565b9350616b8e818560208601617f23565b616b9781617f56565b840191505092915050565b6000616baf601783617e82565b91507f76616c696461746f722069736e277420736c61736865640000000000000000006000830152602082019050919050565b6000616bef600b83617e82565b91507f7a65726f20616d6f756e740000000000000000000000000000000000000000006000830152602082019050919050565b6000616c2f602683617e82565b91507f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008301527f64647265737300000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000616c95601183617e82565b91507f616c7265616479206c6f636b65642075700000000000000000000000000000006000830152602082019050919050565b6000616cd5601b83617e82565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b6000616d15602983617e82565b91507f63616c6c6572206973206e6f7420746865204e6f64654472697665724175746860008301527f20636f6e747261637400000000000000000000000000000000000000000000006020830152604082019050919050565b6000616d7b601083617e82565b91507f6e6f7468696e6720746f207374617368000000000000000000000000000000006000830152602082019050919050565b6000616dbb601283617e82565b91507f696e636f7272656374206475726174696f6e00000000000000000000000000006000830152602082019050919050565b6000616dfb601683617e82565b91507f76616c696461746f722069736e277420616374697665000000000000000000006000830152602082019050919050565b6000616e3b601283617e82565b91507f6e6f7420735257412066696e616c697a657200000000000000000000000000006000830152602082019050919050565b6000616e7b600c83617e82565b91507f77726f6e672073746174757300000000000000000000000000000000000000006000830152602082019050919050565b6000616ebb601683617e82565b91507f6e6f7420656e6f7567682074696d6520706173736564000000000000000000006000830152602082019050919050565b6000616efb601883617e82565b91507f76616c696461746f7220616c72656164792065786973747300000000000000006000830152602082019050919050565b6000616f3b600d83617e82565b91507f6e6f74206c6f636b6564207570000000000000000000000000000000000000006000830152602082019050919050565b6000616f7b602183617e82565b91507f536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f60008301527f77000000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000616fe1600c83617e82565b91507f7a65726f207265776172647300000000000000000000000000000000000000006000830152602082019050919050565b6000617021601f83617e82565b91507f6c6f636b7570206475726174696f6e2063616e6e6f74206465637265617365006000830152602082019050919050565b6000617061601383617e82565b91507f7772494420616c726561647920657869737473000000000000000000000000006000830152602082019050919050565b60006170a1602083617e82565b91507f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726000830152602082019050919050565b60006170e1601883617e82565b91507f6f75747374616e64696e6720735257412062616c616e636500000000000000006000830152602082019050919050565b6000617121602183617e82565b91507f6d757374206265206c657373207468616e206f7220657175616c20746f20312e60008301527f30000000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000617187601983617e82565b91507f6e6f7420656e6f75676820756e6c6f636b6564207374616b65000000000000006000830152602082019050919050565b60006171c7600c83617e82565b91507f656d707479207075626b657900000000000000000000000000000000000000006000830152602082019050919050565b6000617207601583617e82565b91507f7265717565737420646f65736e277420657869737400000000000000000000006000830152602082019050919050565b6000617247602883617e82565b91507f76616c696461746f72206c6f636b757020706572696f642077696c6c20656e6460008301527f206561726c6965720000000000000000000000000000000000000000000000006020830152604082019050919050565b60006172ad601783617e82565b91507f6e6f7420656e6f756768206c6f636b6564207374616b650000000000000000006000830152602082019050919050565b60006172ed600083617e77565b9150600082019050919050565b6000617307601283617e82565b91507f4661696c656420746f2073656e642052574100000000000000000000000000006000830152602082019050919050565b6000617347601083617e82565b91507f6e6f7420656e6f756768207374616b65000000000000000000000000000000006000830152602082019050919050565b6000617387602983617e82565b91507f76616c696461746f7227732064656c65676174696f6e73206c696d697420697360008301527f20657863656564656400000000000000000000000000000000000000000000006020830152604082019050919050565b60006173ed601b83617e82565b91507f676f7620766f746573207265636f756e74696e67206661696c656400000000006000830152602082019050919050565b600061742d601783617e82565b91507f76616c696461746f7220646f65736e27742065786973740000000000000000006000830152602082019050919050565b600061746d601883617e82565b91507f6e6f7420656e6f7567682065706f6368732070617373656400000000000000006000830152602082019050919050565b60006174ad601783617e82565b91507f696e73756666696369656e742073656c662d7374616b650000000000000000006000830152602082019050919050565b60006174ed601683617e82565b91507f7374616b652069732066756c6c7920736c6173686564000000000000000000006000830152602082019050919050565b600061752d602c83617e82565b91507f6c6f636b6564207374616b652069732067726561746572207468616e2074686560008301527f2077686f6c65207374616b6500000000000000000000000000000000000000006020830152604082019050919050565b60608201600082015161759c60008501826175c8565b5060208201516175af60208501826175c8565b5060408201516175c260408501826175c8565b50505050565b6175d181617ed1565b82525050565b6175e081617ed1565b82525050565b60006175f28284616a61565b915081905092915050565b6000617608826172e0565b9150819050919050565b60006020820190506176276000830184616987565b92915050565b60006080820190506176426000830187616978565b61764f6020830186616987565b61765c60408301856175d7565b61766960608301846175d7565b95945050505050565b60006040820190506176876000830185616987565b6176946020830184616987565b9392505050565b60006040820190506176b06000830185616987565b6176bd60208301846175d7565b9392505050565b600060208201905081810360008301526176de8184616996565b905092915050565b6000602082019050818103600083015261770081846169f4565b905092915050565b600060208201905061771d6000830184616a52565b92915050565b6000602082019050818103600083015261773d8184616a92565b905092915050565b6000602082019050818103600083015261775f8184616b69565b905092915050565b6000602082019050818103600083015261778081616ba2565b9050919050565b600060208201905081810360008301526177a081616be2565b9050919050565b600060208201905081810360008301526177c081616c22565b9050919050565b600060208201905081810360008301526177e081616c88565b9050919050565b6000602082019050818103600083015261780081616cc8565b9050919050565b6000602082019050818103600083015261782081616d08565b9050919050565b6000602082019050818103600083015261784081616d6e565b9050919050565b6000602082019050818103600083015261786081616dae565b9050919050565b6000602082019050818103600083015261788081616dee565b9050919050565b600060208201905081810360008301526178a081616e2e565b9050919050565b600060208201905081810360008301526178c081616e6e565b9050919050565b600060208201905081810360008301526178e081616eae565b9050919050565b6000602082019050818103600083015261790081616eee565b9050919050565b6000602082019050818103600083015261792081616f2e565b9050919050565b6000602082019050818103600083015261794081616f6e565b9050919050565b6000602082019050818103600083015261796081616fd4565b9050919050565b6000602082019050818103600083015261798081617014565b9050919050565b600060208201905081810360008301526179a081617054565b9050919050565b600060208201905081810360008301526179c081617094565b9050919050565b600060208201905081810360008301526179e0816170d4565b9050919050565b60006020820190508181036000830152617a0081617114565b9050919050565b60006020820190508181036000830152617a208161717a565b9050919050565b60006020820190508181036000830152617a40816171ba565b9050919050565b60006020820190508181036000830152617a60816171fa565b9050919050565b60006020820190508181036000830152617a808161723a565b9050919050565b60006020820190508181036000830152617aa0816172a0565b9050919050565b60006020820190508181036000830152617ac0816172fa565b9050919050565b60006020820190508181036000830152617ae08161733a565b9050919050565b60006020820190508181036000830152617b008161737a565b9050919050565b60006020820190508181036000830152617b20816173e0565b9050919050565b60006020820190508181036000830152617b4081617420565b9050919050565b60006020820190508181036000830152617b6081617460565b9050919050565b60006020820190508181036000830152617b80816174a0565b9050919050565b60006020820190508181036000830152617ba0816174e0565b9050919050565b60006020820190508181036000830152617bc081617520565b9050919050565b6000602082019050617bdc60008301846175d7565b92915050565b6000604082019050617bf760008301856175d7565b8181036020830152617c098184616acb565b90509392505050565b6000604082019050617c2760008301856175d7565b617c346020830184616b5a565b9392505050565b6000604082019050617c5060008301856175d7565b617c5d60208301846175d7565b9392505050565b6000606082019050617c7960008301866175d7565b617c8660208301856175d7565b617c9360408301846175d7565b949350505050565b6000608082019050617cb060008301876175d7565b617cbd60208301866175d7565b617cca60408301856175d7565b617cd760608301846175d7565b95945050505050565b600060e082019050617cf5600083018a6175d7565b617d0260208301896175d7565b617d0f60408301886175d7565b617d1c60608301876175d7565b617d2960808301866175d7565b617d3660a08301856175d7565b617d4360c0830184616987565b98975050505050505050565b600060e082019050617d64600083018a6175d7565b617d7160208301896175d7565b617d7e60408301886175d7565b617d8b60608301876175d7565b617d9860808301866175d7565b617da560a08301856175d7565b617db260c08301846175d7565b98975050505050505050565b6000819050602082019050919050565b6000819050602082019050919050565b60008190508160005260206000209050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000617e9e82617eb1565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000617ee682617eff565b9050919050565b6000617ef882617ed1565b9050919050565b6000617f0a82617f11565b9050919050565b6000617f1c82617eb1565b9050919050565b60005b83811015617f41578082015181840152602081019050617f26565b83811115617f50576000848401525b50505050565b6000601f19601f8301169050919050565b617f7081617e93565b8114617f7b57600080fd5b50565b617f8781617ea5565b8114617f9257600080fd5b50565b617f9e81617ed1565b8114617fa957600080fd5b5056fea365627a7a72315820fd8fa742408c2b87482fc27b7e1d88ed590d49789fa62681d2f387ab02ad9ff96c6578706572696d656e74616cf564736f6c63430005100040",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Contract *ContractCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Contract *ContractSession) CurrentEpoch() (*big.Int, error) {
	return _Contract.Contract.CurrentEpoch(&_Contract.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Contract *ContractCallerSession) CurrentEpoch() (*big.Int, error) {
	return _Contract.Contract.CurrentEpoch(&_Contract.CallOpts)
}

// CurrentSealedEpoch is a free data retrieval call binding the contract method 0x7cacb1d6.
//
// Solidity: function currentSealedEpoch() view returns(uint256)
func (_Contract *ContractCaller) CurrentSealedEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "currentSealedEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentSealedEpoch is a free data retrieval call binding the contract method 0x7cacb1d6.
//
// Solidity: function currentSealedEpoch() view returns(uint256)
func (_Contract *ContractSession) CurrentSealedEpoch() (*big.Int, error) {
	return _Contract.Contract.CurrentSealedEpoch(&_Contract.CallOpts)
}

// CurrentSealedEpoch is a free data retrieval call binding the contract method 0x7cacb1d6.
//
// Solidity: function currentSealedEpoch() view returns(uint256)
func (_Contract *ContractCallerSession) CurrentSealedEpoch() (*big.Int, error) {
	return _Contract.Contract.CurrentSealedEpoch(&_Contract.CallOpts)
}

// GetEpochAccumulatedOriginatedTxsFee is a free data retrieval call binding the contract method 0xdc31e1af.
//
// Solidity: function getEpochAccumulatedOriginatedTxsFee(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCaller) GetEpochAccumulatedOriginatedTxsFee(opts *bind.CallOpts, epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEpochAccumulatedOriginatedTxsFee", epoch, validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochAccumulatedOriginatedTxsFee is a free data retrieval call binding the contract method 0xdc31e1af.
//
// Solidity: function getEpochAccumulatedOriginatedTxsFee(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractSession) GetEpochAccumulatedOriginatedTxsFee(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochAccumulatedOriginatedTxsFee(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochAccumulatedOriginatedTxsFee is a free data retrieval call binding the contract method 0xdc31e1af.
//
// Solidity: function getEpochAccumulatedOriginatedTxsFee(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCallerSession) GetEpochAccumulatedOriginatedTxsFee(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochAccumulatedOriginatedTxsFee(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochAccumulatedRewardPerToken is a free data retrieval call binding the contract method 0x61e53fcc.
//
// Solidity: function getEpochAccumulatedRewardPerToken(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCaller) GetEpochAccumulatedRewardPerToken(opts *bind.CallOpts, epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEpochAccumulatedRewardPerToken", epoch, validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochAccumulatedRewardPerToken is a free data retrieval call binding the contract method 0x61e53fcc.
//
// Solidity: function getEpochAccumulatedRewardPerToken(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractSession) GetEpochAccumulatedRewardPerToken(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochAccumulatedRewardPerToken(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochAccumulatedRewardPerToken is a free data retrieval call binding the contract method 0x61e53fcc.
//
// Solidity: function getEpochAccumulatedRewardPerToken(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCallerSession) GetEpochAccumulatedRewardPerToken(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochAccumulatedRewardPerToken(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochAccumulatedUptime is a free data retrieval call binding the contract method 0xdf00c922.
//
// Solidity: function getEpochAccumulatedUptime(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCaller) GetEpochAccumulatedUptime(opts *bind.CallOpts, epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEpochAccumulatedUptime", epoch, validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochAccumulatedUptime is a free data retrieval call binding the contract method 0xdf00c922.
//
// Solidity: function getEpochAccumulatedUptime(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractSession) GetEpochAccumulatedUptime(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochAccumulatedUptime(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochAccumulatedUptime is a free data retrieval call binding the contract method 0xdf00c922.
//
// Solidity: function getEpochAccumulatedUptime(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCallerSession) GetEpochAccumulatedUptime(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochAccumulatedUptime(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochOfflineBlocks is a free data retrieval call binding the contract method 0xa198d229.
//
// Solidity: function getEpochOfflineBlocks(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCaller) GetEpochOfflineBlocks(opts *bind.CallOpts, epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEpochOfflineBlocks", epoch, validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochOfflineBlocks is a free data retrieval call binding the contract method 0xa198d229.
//
// Solidity: function getEpochOfflineBlocks(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractSession) GetEpochOfflineBlocks(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochOfflineBlocks(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochOfflineBlocks is a free data retrieval call binding the contract method 0xa198d229.
//
// Solidity: function getEpochOfflineBlocks(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCallerSession) GetEpochOfflineBlocks(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochOfflineBlocks(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochOfflineTime is a free data retrieval call binding the contract method 0xe261641a.
//
// Solidity: function getEpochOfflineTime(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCaller) GetEpochOfflineTime(opts *bind.CallOpts, epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEpochOfflineTime", epoch, validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochOfflineTime is a free data retrieval call binding the contract method 0xe261641a.
//
// Solidity: function getEpochOfflineTime(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractSession) GetEpochOfflineTime(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochOfflineTime(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochOfflineTime is a free data retrieval call binding the contract method 0xe261641a.
//
// Solidity: function getEpochOfflineTime(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCallerSession) GetEpochOfflineTime(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochOfflineTime(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochReceivedStake is a free data retrieval call binding the contract method 0x58f95b80.
//
// Solidity: function getEpochReceivedStake(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCaller) GetEpochReceivedStake(opts *bind.CallOpts, epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEpochReceivedStake", epoch, validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochReceivedStake is a free data retrieval call binding the contract method 0x58f95b80.
//
// Solidity: function getEpochReceivedStake(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractSession) GetEpochReceivedStake(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochReceivedStake(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochReceivedStake is a free data retrieval call binding the contract method 0x58f95b80.
//
// Solidity: function getEpochReceivedStake(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCallerSession) GetEpochReceivedStake(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochReceivedStake(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochSnapshot is a free data retrieval call binding the contract method 0x39b80c00.
//
// Solidity: function getEpochSnapshot(uint256 ) view returns(uint256 endTime, uint256 epochFee, uint256 totalBaseRewardWeight, uint256 totalTxRewardWeight, uint256 baseRewardPerSecond, uint256 totalStake, uint256 totalSupply)
func (_Contract *ContractCaller) GetEpochSnapshot(opts *bind.CallOpts, arg0 *big.Int) (struct {
	EndTime               *big.Int
	EpochFee              *big.Int
	TotalBaseRewardWeight *big.Int
	TotalTxRewardWeight   *big.Int
	BaseRewardPerSecond   *big.Int
	TotalStake            *big.Int
	TotalSupply           *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEpochSnapshot", arg0)

	outstruct := new(struct {
		EndTime               *big.Int
		EpochFee              *big.Int
		TotalBaseRewardWeight *big.Int
		TotalTxRewardWeight   *big.Int
		BaseRewardPerSecond   *big.Int
		TotalStake            *big.Int
		TotalSupply           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EndTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EpochFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalBaseRewardWeight = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalTxRewardWeight = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.BaseRewardPerSecond = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TotalStake = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TotalSupply = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetEpochSnapshot is a free data retrieval call binding the contract method 0x39b80c00.
//
// Solidity: function getEpochSnapshot(uint256 ) view returns(uint256 endTime, uint256 epochFee, uint256 totalBaseRewardWeight, uint256 totalTxRewardWeight, uint256 baseRewardPerSecond, uint256 totalStake, uint256 totalSupply)
func (_Contract *ContractSession) GetEpochSnapshot(arg0 *big.Int) (struct {
	EndTime               *big.Int
	EpochFee              *big.Int
	TotalBaseRewardWeight *big.Int
	TotalTxRewardWeight   *big.Int
	BaseRewardPerSecond   *big.Int
	TotalStake            *big.Int
	TotalSupply           *big.Int
}, error) {
	return _Contract.Contract.GetEpochSnapshot(&_Contract.CallOpts, arg0)
}

// GetEpochSnapshot is a free data retrieval call binding the contract method 0x39b80c00.
//
// Solidity: function getEpochSnapshot(uint256 ) view returns(uint256 endTime, uint256 epochFee, uint256 totalBaseRewardWeight, uint256 totalTxRewardWeight, uint256 baseRewardPerSecond, uint256 totalStake, uint256 totalSupply)
func (_Contract *ContractCallerSession) GetEpochSnapshot(arg0 *big.Int) (struct {
	EndTime               *big.Int
	EpochFee              *big.Int
	TotalBaseRewardWeight *big.Int
	TotalTxRewardWeight   *big.Int
	BaseRewardPerSecond   *big.Int
	TotalStake            *big.Int
	TotalSupply           *big.Int
}, error) {
	return _Contract.Contract.GetEpochSnapshot(&_Contract.CallOpts, arg0)
}

// GetEpochValidatorIDs is a free data retrieval call binding the contract method 0xb88a37e2.
//
// Solidity: function getEpochValidatorIDs(uint256 epoch) view returns(uint256[])
func (_Contract *ContractCaller) GetEpochValidatorIDs(opts *bind.CallOpts, epoch *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEpochValidatorIDs", epoch)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetEpochValidatorIDs is a free data retrieval call binding the contract method 0xb88a37e2.
//
// Solidity: function getEpochValidatorIDs(uint256 epoch) view returns(uint256[])
func (_Contract *ContractSession) GetEpochValidatorIDs(epoch *big.Int) ([]*big.Int, error) {
	return _Contract.Contract.GetEpochValidatorIDs(&_Contract.CallOpts, epoch)
}

// GetEpochValidatorIDs is a free data retrieval call binding the contract method 0xb88a37e2.
//
// Solidity: function getEpochValidatorIDs(uint256 epoch) view returns(uint256[])
func (_Contract *ContractCallerSession) GetEpochValidatorIDs(epoch *big.Int) ([]*big.Int, error) {
	return _Contract.Contract.GetEpochValidatorIDs(&_Contract.CallOpts, epoch)
}

// GetLockedStake is a free data retrieval call binding the contract method 0x670322f8.
//
// Solidity: function getLockedStake(address delegator, uint256 toValidatorID) view returns(uint256)
func (_Contract *ContractCaller) GetLockedStake(opts *bind.CallOpts, delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getLockedStake", delegator, toValidatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLockedStake is a free data retrieval call binding the contract method 0x670322f8.
//
// Solidity: function getLockedStake(address delegator, uint256 toValidatorID) view returns(uint256)
func (_Contract *ContractSession) GetLockedStake(delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetLockedStake(&_Contract.CallOpts, delegator, toValidatorID)
}

// GetLockedStake is a free data retrieval call binding the contract method 0x670322f8.
//
// Solidity: function getLockedStake(address delegator, uint256 toValidatorID) view returns(uint256)
func (_Contract *ContractCallerSession) GetLockedStake(delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetLockedStake(&_Contract.CallOpts, delegator, toValidatorID)
}

// GetLockupInfo is a free data retrieval call binding the contract method 0x96c7ee46.
//
// Solidity: function getLockupInfo(address , uint256 ) view returns(uint256 lockedStake, uint256 fromEpoch, uint256 endTime, uint256 duration)
func (_Contract *ContractCaller) GetLockupInfo(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	LockedStake *big.Int
	FromEpoch   *big.Int
	EndTime     *big.Int
	Duration    *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getLockupInfo", arg0, arg1)

	outstruct := new(struct {
		LockedStake *big.Int
		FromEpoch   *big.Int
		EndTime     *big.Int
		Duration    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LockedStake = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FromEpoch = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetLockupInfo is a free data retrieval call binding the contract method 0x96c7ee46.
//
// Solidity: function getLockupInfo(address , uint256 ) view returns(uint256 lockedStake, uint256 fromEpoch, uint256 endTime, uint256 duration)
func (_Contract *ContractSession) GetLockupInfo(arg0 common.Address, arg1 *big.Int) (struct {
	LockedStake *big.Int
	FromEpoch   *big.Int
	EndTime     *big.Int
	Duration    *big.Int
}, error) {
	return _Contract.Contract.GetLockupInfo(&_Contract.CallOpts, arg0, arg1)
}

// GetLockupInfo is a free data retrieval call binding the contract method 0x96c7ee46.
//
// Solidity: function getLockupInfo(address , uint256 ) view returns(uint256 lockedStake, uint256 fromEpoch, uint256 endTime, uint256 duration)
func (_Contract *ContractCallerSession) GetLockupInfo(arg0 common.Address, arg1 *big.Int) (struct {
	LockedStake *big.Int
	FromEpoch   *big.Int
	EndTime     *big.Int
	Duration    *big.Int
}, error) {
	return _Contract.Contract.GetLockupInfo(&_Contract.CallOpts, arg0, arg1)
}

// GetSelfStake is a free data retrieval call binding the contract method 0x5601fe01.
//
// Solidity: function getSelfStake(uint256 validatorID) view returns(uint256)
func (_Contract *ContractCaller) GetSelfStake(opts *bind.CallOpts, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getSelfStake", validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSelfStake is a free data retrieval call binding the contract method 0x5601fe01.
//
// Solidity: function getSelfStake(uint256 validatorID) view returns(uint256)
func (_Contract *ContractSession) GetSelfStake(validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetSelfStake(&_Contract.CallOpts, validatorID)
}

// GetSelfStake is a free data retrieval call binding the contract method 0x5601fe01.
//
// Solidity: function getSelfStake(uint256 validatorID) view returns(uint256)
func (_Contract *ContractCallerSession) GetSelfStake(validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetSelfStake(&_Contract.CallOpts, validatorID)
}

// GetStake is a free data retrieval call binding the contract method 0xcfd47663.
//
// Solidity: function getStake(address , uint256 ) view returns(uint256)
func (_Contract *ContractCaller) GetStake(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getStake", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStake is a free data retrieval call binding the contract method 0xcfd47663.
//
// Solidity: function getStake(address , uint256 ) view returns(uint256)
func (_Contract *ContractSession) GetStake(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetStake(&_Contract.CallOpts, arg0, arg1)
}

// GetStake is a free data retrieval call binding the contract method 0xcfd47663.
//
// Solidity: function getStake(address , uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) GetStake(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetStake(&_Contract.CallOpts, arg0, arg1)
}

// GetStashedLockupRewards is a free data retrieval call binding the contract method 0xb810e411.
//
// Solidity: function getStashedLockupRewards(address , uint256 ) view returns(uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
func (_Contract *ContractCaller) GetStashedLockupRewards(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	LockupExtraReward *big.Int
	LockupBaseReward  *big.Int
	UnlockedReward    *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getStashedLockupRewards", arg0, arg1)

	outstruct := new(struct {
		LockupExtraReward *big.Int
		LockupBaseReward  *big.Int
		UnlockedReward    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LockupExtraReward = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LockupBaseReward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.UnlockedReward = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetStashedLockupRewards is a free data retrieval call binding the contract method 0xb810e411.
//
// Solidity: function getStashedLockupRewards(address , uint256 ) view returns(uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
func (_Contract *ContractSession) GetStashedLockupRewards(arg0 common.Address, arg1 *big.Int) (struct {
	LockupExtraReward *big.Int
	LockupBaseReward  *big.Int
	UnlockedReward    *big.Int
}, error) {
	return _Contract.Contract.GetStashedLockupRewards(&_Contract.CallOpts, arg0, arg1)
}

// GetStashedLockupRewards is a free data retrieval call binding the contract method 0xb810e411.
//
// Solidity: function getStashedLockupRewards(address , uint256 ) view returns(uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
func (_Contract *ContractCallerSession) GetStashedLockupRewards(arg0 common.Address, arg1 *big.Int) (struct {
	LockupExtraReward *big.Int
	LockupBaseReward  *big.Int
	UnlockedReward    *big.Int
}, error) {
	return _Contract.Contract.GetStashedLockupRewards(&_Contract.CallOpts, arg0, arg1)
}

// GetUnlockedStake is a free data retrieval call binding the contract method 0x12622d0e.
//
// Solidity: function getUnlockedStake(address delegator, uint256 toValidatorID) view returns(uint256)
func (_Contract *ContractCaller) GetUnlockedStake(opts *bind.CallOpts, delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getUnlockedStake", delegator, toValidatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnlockedStake is a free data retrieval call binding the contract method 0x12622d0e.
//
// Solidity: function getUnlockedStake(address delegator, uint256 toValidatorID) view returns(uint256)
func (_Contract *ContractSession) GetUnlockedStake(delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetUnlockedStake(&_Contract.CallOpts, delegator, toValidatorID)
}

// GetUnlockedStake is a free data retrieval call binding the contract method 0x12622d0e.
//
// Solidity: function getUnlockedStake(address delegator, uint256 toValidatorID) view returns(uint256)
func (_Contract *ContractCallerSession) GetUnlockedStake(delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetUnlockedStake(&_Contract.CallOpts, delegator, toValidatorID)
}

// GetValidator is a free data retrieval call binding the contract method 0xb5d89627.
//
// Solidity: function getValidator(uint256 ) view returns(uint256 status, uint256 deactivatedTime, uint256 deactivatedEpoch, uint256 receivedStake, uint256 createdEpoch, uint256 createdTime, address auth)
func (_Contract *ContractCaller) GetValidator(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Status           *big.Int
	DeactivatedTime  *big.Int
	DeactivatedEpoch *big.Int
	ReceivedStake    *big.Int
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	Auth             common.Address
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getValidator", arg0)

	outstruct := new(struct {
		Status           *big.Int
		DeactivatedTime  *big.Int
		DeactivatedEpoch *big.Int
		ReceivedStake    *big.Int
		CreatedEpoch     *big.Int
		CreatedTime      *big.Int
		Auth             common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DeactivatedTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DeactivatedEpoch = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ReceivedStake = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.CreatedEpoch = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.CreatedTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Auth = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetValidator is a free data retrieval call binding the contract method 0xb5d89627.
//
// Solidity: function getValidator(uint256 ) view returns(uint256 status, uint256 deactivatedTime, uint256 deactivatedEpoch, uint256 receivedStake, uint256 createdEpoch, uint256 createdTime, address auth)
func (_Contract *ContractSession) GetValidator(arg0 *big.Int) (struct {
	Status           *big.Int
	DeactivatedTime  *big.Int
	DeactivatedEpoch *big.Int
	ReceivedStake    *big.Int
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	Auth             common.Address
}, error) {
	return _Contract.Contract.GetValidator(&_Contract.CallOpts, arg0)
}

// GetValidator is a free data retrieval call binding the contract method 0xb5d89627.
//
// Solidity: function getValidator(uint256 ) view returns(uint256 status, uint256 deactivatedTime, uint256 deactivatedEpoch, uint256 receivedStake, uint256 createdEpoch, uint256 createdTime, address auth)
func (_Contract *ContractCallerSession) GetValidator(arg0 *big.Int) (struct {
	Status           *big.Int
	DeactivatedTime  *big.Int
	DeactivatedEpoch *big.Int
	ReceivedStake    *big.Int
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	Auth             common.Address
}, error) {
	return _Contract.Contract.GetValidator(&_Contract.CallOpts, arg0)
}

// GetValidatorID is a free data retrieval call binding the contract method 0x0135b1db.
//
// Solidity: function getValidatorID(address ) view returns(uint256)
func (_Contract *ContractCaller) GetValidatorID(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getValidatorID", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorID is a free data retrieval call binding the contract method 0x0135b1db.
//
// Solidity: function getValidatorID(address ) view returns(uint256)
func (_Contract *ContractSession) GetValidatorID(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.GetValidatorID(&_Contract.CallOpts, arg0)
}

// GetValidatorID is a free data retrieval call binding the contract method 0x0135b1db.
//
// Solidity: function getValidatorID(address ) view returns(uint256)
func (_Contract *ContractCallerSession) GetValidatorID(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.GetValidatorID(&_Contract.CallOpts, arg0)
}

// GetValidatorPubkey is a free data retrieval call binding the contract method 0x854873e1.
//
// Solidity: function getValidatorPubkey(uint256 ) view returns(bytes)
func (_Contract *ContractCaller) GetValidatorPubkey(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getValidatorPubkey", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetValidatorPubkey is a free data retrieval call binding the contract method 0x854873e1.
//
// Solidity: function getValidatorPubkey(uint256 ) view returns(bytes)
func (_Contract *ContractSession) GetValidatorPubkey(arg0 *big.Int) ([]byte, error) {
	return _Contract.Contract.GetValidatorPubkey(&_Contract.CallOpts, arg0)
}

// GetValidatorPubkey is a free data retrieval call binding the contract method 0x854873e1.
//
// Solidity: function getValidatorPubkey(uint256 ) view returns(bytes)
func (_Contract *ContractCallerSession) GetValidatorPubkey(arg0 *big.Int) ([]byte, error) {
	return _Contract.Contract.GetValidatorPubkey(&_Contract.CallOpts, arg0)
}

// GetWithdrawalRequest is a free data retrieval call binding the contract method 0x1f270152.
//
// Solidity: function getWithdrawalRequest(address , uint256 , uint256 ) view returns(uint256 epoch, uint256 time, uint256 amount)
func (_Contract *ContractCaller) GetWithdrawalRequest(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (struct {
	Epoch  *big.Int
	Time   *big.Int
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getWithdrawalRequest", arg0, arg1, arg2)

	outstruct := new(struct {
		Epoch  *big.Int
		Time   *big.Int
		Amount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Epoch = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Time = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetWithdrawalRequest is a free data retrieval call binding the contract method 0x1f270152.
//
// Solidity: function getWithdrawalRequest(address , uint256 , uint256 ) view returns(uint256 epoch, uint256 time, uint256 amount)
func (_Contract *ContractSession) GetWithdrawalRequest(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (struct {
	Epoch  *big.Int
	Time   *big.Int
	Amount *big.Int
}, error) {
	return _Contract.Contract.GetWithdrawalRequest(&_Contract.CallOpts, arg0, arg1, arg2)
}

// GetWithdrawalRequest is a free data retrieval call binding the contract method 0x1f270152.
//
// Solidity: function getWithdrawalRequest(address , uint256 , uint256 ) view returns(uint256 epoch, uint256 time, uint256 amount)
func (_Contract *ContractCallerSession) GetWithdrawalRequest(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (struct {
	Epoch  *big.Int
	Time   *big.Int
	Amount *big.Int
}, error) {
	return _Contract.Contract.GetWithdrawalRequest(&_Contract.CallOpts, arg0, arg1, arg2)
}

// GetWrRequests is a free data retrieval call binding the contract method 0x702797e3.
//
// Solidity: function getWrRequests(address delegator, uint256 validatorID, uint256 offset, uint256 limit) view returns((uint256,uint256,uint256)[])
func (_Contract *ContractCaller) GetWrRequests(opts *bind.CallOpts, delegator common.Address, validatorID *big.Int, offset *big.Int, limit *big.Int) ([]SFCStateWithdrawalRequest, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getWrRequests", delegator, validatorID, offset, limit)

	if err != nil {
		return *new([]SFCStateWithdrawalRequest), err
	}

	out0 := *abi.ConvertType(out[0], new([]SFCStateWithdrawalRequest)).(*[]SFCStateWithdrawalRequest)

	return out0, err

}

// GetWrRequests is a free data retrieval call binding the contract method 0x702797e3.
//
// Solidity: function getWrRequests(address delegator, uint256 validatorID, uint256 offset, uint256 limit) view returns((uint256,uint256,uint256)[])
func (_Contract *ContractSession) GetWrRequests(delegator common.Address, validatorID *big.Int, offset *big.Int, limit *big.Int) ([]SFCStateWithdrawalRequest, error) {
	return _Contract.Contract.GetWrRequests(&_Contract.CallOpts, delegator, validatorID, offset, limit)
}

// GetWrRequests is a free data retrieval call binding the contract method 0x702797e3.
//
// Solidity: function getWrRequests(address delegator, uint256 validatorID, uint256 offset, uint256 limit) view returns((uint256,uint256,uint256)[])
func (_Contract *ContractCallerSession) GetWrRequests(delegator common.Address, validatorID *big.Int, offset *big.Int, limit *big.Int) ([]SFCStateWithdrawalRequest, error) {
	return _Contract.Contract.GetWrRequests(&_Contract.CallOpts, delegator, validatorID, offset, limit)
}

// IsLockedUp is a free data retrieval call binding the contract method 0xcfdbb7cd.
//
// Solidity: function isLockedUp(address delegator, uint256 toValidatorID) view returns(bool)
func (_Contract *ContractCaller) IsLockedUp(opts *bind.CallOpts, delegator common.Address, toValidatorID *big.Int) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isLockedUp", delegator, toValidatorID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLockedUp is a free data retrieval call binding the contract method 0xcfdbb7cd.
//
// Solidity: function isLockedUp(address delegator, uint256 toValidatorID) view returns(bool)
func (_Contract *ContractSession) IsLockedUp(delegator common.Address, toValidatorID *big.Int) (bool, error) {
	return _Contract.Contract.IsLockedUp(&_Contract.CallOpts, delegator, toValidatorID)
}

// IsLockedUp is a free data retrieval call binding the contract method 0xcfdbb7cd.
//
// Solidity: function isLockedUp(address delegator, uint256 toValidatorID) view returns(bool)
func (_Contract *ContractCallerSession) IsLockedUp(delegator common.Address, toValidatorID *big.Int) (bool, error) {
	return _Contract.Contract.IsLockedUp(&_Contract.CallOpts, delegator, toValidatorID)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Contract *ContractCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Contract *ContractSession) IsOwner() (bool, error) {
	return _Contract.Contract.IsOwner(&_Contract.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Contract *ContractCallerSession) IsOwner() (bool, error) {
	return _Contract.Contract.IsOwner(&_Contract.CallOpts)
}

// IsSlashed is a free data retrieval call binding the contract method 0xc3de580e.
//
// Solidity: function isSlashed(uint256 validatorID) view returns(bool)
func (_Contract *ContractCaller) IsSlashed(opts *bind.CallOpts, validatorID *big.Int) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isSlashed", validatorID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSlashed is a free data retrieval call binding the contract method 0xc3de580e.
//
// Solidity: function isSlashed(uint256 validatorID) view returns(bool)
func (_Contract *ContractSession) IsSlashed(validatorID *big.Int) (bool, error) {
	return _Contract.Contract.IsSlashed(&_Contract.CallOpts, validatorID)
}

// IsSlashed is a free data retrieval call binding the contract method 0xc3de580e.
//
// Solidity: function isSlashed(uint256 validatorID) view returns(bool)
func (_Contract *ContractCallerSession) IsSlashed(validatorID *big.Int) (bool, error) {
	return _Contract.Contract.IsSlashed(&_Contract.CallOpts, validatorID)
}

// LastValidatorID is a free data retrieval call binding the contract method 0xc7be95de.
//
// Solidity: function lastValidatorID() view returns(uint256)
func (_Contract *ContractCaller) LastValidatorID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "lastValidatorID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastValidatorID is a free data retrieval call binding the contract method 0xc7be95de.
//
// Solidity: function lastValidatorID() view returns(uint256)
func (_Contract *ContractSession) LastValidatorID() (*big.Int, error) {
	return _Contract.Contract.LastValidatorID(&_Contract.CallOpts)
}

// LastValidatorID is a free data retrieval call binding the contract method 0xc7be95de.
//
// Solidity: function lastValidatorID() view returns(uint256)
func (_Contract *ContractCallerSession) LastValidatorID() (*big.Int, error) {
	return _Contract.Contract.LastValidatorID(&_Contract.CallOpts)
}

// MinGasPrice is a free data retrieval call binding the contract method 0xd96ed505.
//
// Solidity: function minGasPrice() view returns(uint256)
func (_Contract *ContractCaller) MinGasPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "minGasPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinGasPrice is a free data retrieval call binding the contract method 0xd96ed505.
//
// Solidity: function minGasPrice() view returns(uint256)
func (_Contract *ContractSession) MinGasPrice() (*big.Int, error) {
	return _Contract.Contract.MinGasPrice(&_Contract.CallOpts)
}

// MinGasPrice is a free data retrieval call binding the contract method 0xd96ed505.
//
// Solidity: function minGasPrice() view returns(uint256)
func (_Contract *ContractCallerSession) MinGasPrice() (*big.Int, error) {
	return _Contract.Contract.MinGasPrice(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// PendingRewards is a free data retrieval call binding the contract method 0x6099ecb2.
//
// Solidity: function pendingRewards(address delegator, uint256 toValidatorID) view returns(uint256)
func (_Contract *ContractCaller) PendingRewards(opts *bind.CallOpts, delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "pendingRewards", delegator, toValidatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingRewards is a free data retrieval call binding the contract method 0x6099ecb2.
//
// Solidity: function pendingRewards(address delegator, uint256 toValidatorID) view returns(uint256)
func (_Contract *ContractSession) PendingRewards(delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.PendingRewards(&_Contract.CallOpts, delegator, toValidatorID)
}

// PendingRewards is a free data retrieval call binding the contract method 0x6099ecb2.
//
// Solidity: function pendingRewards(address delegator, uint256 toValidatorID) view returns(uint256)
func (_Contract *ContractCallerSession) PendingRewards(delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.PendingRewards(&_Contract.CallOpts, delegator, toValidatorID)
}

// RewardsStash is a free data retrieval call binding the contract method 0x6f498663.
//
// Solidity: function rewardsStash(address delegator, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCaller) RewardsStash(opts *bind.CallOpts, delegator common.Address, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "rewardsStash", delegator, validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsStash is a free data retrieval call binding the contract method 0x6f498663.
//
// Solidity: function rewardsStash(address delegator, uint256 validatorID) view returns(uint256)
func (_Contract *ContractSession) RewardsStash(delegator common.Address, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.RewardsStash(&_Contract.CallOpts, delegator, validatorID)
}

// RewardsStash is a free data retrieval call binding the contract method 0x6f498663.
//
// Solidity: function rewardsStash(address delegator, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCallerSession) RewardsStash(delegator common.Address, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.RewardsStash(&_Contract.CallOpts, delegator, validatorID)
}

// SlashingRefundRatio is a free data retrieval call binding the contract method 0xc65ee0e1.
//
// Solidity: function slashingRefundRatio(uint256 ) view returns(uint256)
func (_Contract *ContractCaller) SlashingRefundRatio(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "slashingRefundRatio", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashingRefundRatio is a free data retrieval call binding the contract method 0xc65ee0e1.
//
// Solidity: function slashingRefundRatio(uint256 ) view returns(uint256)
func (_Contract *ContractSession) SlashingRefundRatio(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.SlashingRefundRatio(&_Contract.CallOpts, arg0)
}

// SlashingRefundRatio is a free data retrieval call binding the contract method 0xc65ee0e1.
//
// Solidity: function slashingRefundRatio(uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) SlashingRefundRatio(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.SlashingRefundRatio(&_Contract.CallOpts, arg0)
}

// StakeTokenizerAddress is a free data retrieval call binding the contract method 0x0e559d82.
//
// Solidity: function stakeTokenizerAddress() view returns(address)
func (_Contract *ContractCaller) StakeTokenizerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "stakeTokenizerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeTokenizerAddress is a free data retrieval call binding the contract method 0x0e559d82.
//
// Solidity: function stakeTokenizerAddress() view returns(address)
func (_Contract *ContractSession) StakeTokenizerAddress() (common.Address, error) {
	return _Contract.Contract.StakeTokenizerAddress(&_Contract.CallOpts)
}

// StakeTokenizerAddress is a free data retrieval call binding the contract method 0x0e559d82.
//
// Solidity: function stakeTokenizerAddress() view returns(address)
func (_Contract *ContractCallerSession) StakeTokenizerAddress() (common.Address, error) {
	return _Contract.Contract.StakeTokenizerAddress(&_Contract.CallOpts)
}

// StashedRewardsUntilEpoch is a free data retrieval call binding the contract method 0xa86a056f.
//
// Solidity: function stashedRewardsUntilEpoch(address , uint256 ) view returns(uint256)
func (_Contract *ContractCaller) StashedRewardsUntilEpoch(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "stashedRewardsUntilEpoch", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StashedRewardsUntilEpoch is a free data retrieval call binding the contract method 0xa86a056f.
//
// Solidity: function stashedRewardsUntilEpoch(address , uint256 ) view returns(uint256)
func (_Contract *ContractSession) StashedRewardsUntilEpoch(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Contract.Contract.StashedRewardsUntilEpoch(&_Contract.CallOpts, arg0, arg1)
}

// StashedRewardsUntilEpoch is a free data retrieval call binding the contract method 0xa86a056f.
//
// Solidity: function stashedRewardsUntilEpoch(address , uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) StashedRewardsUntilEpoch(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Contract.Contract.StashedRewardsUntilEpoch(&_Contract.CallOpts, arg0, arg1)
}

// TotalActiveStake is a free data retrieval call binding the contract method 0x28f73148.
//
// Solidity: function totalActiveStake() view returns(uint256)
func (_Contract *ContractCaller) TotalActiveStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "totalActiveStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalActiveStake is a free data retrieval call binding the contract method 0x28f73148.
//
// Solidity: function totalActiveStake() view returns(uint256)
func (_Contract *ContractSession) TotalActiveStake() (*big.Int, error) {
	return _Contract.Contract.TotalActiveStake(&_Contract.CallOpts)
}

// TotalActiveStake is a free data retrieval call binding the contract method 0x28f73148.
//
// Solidity: function totalActiveStake() view returns(uint256)
func (_Contract *ContractCallerSession) TotalActiveStake() (*big.Int, error) {
	return _Contract.Contract.TotalActiveStake(&_Contract.CallOpts)
}

// TotalSlashedStake is a free data retrieval call binding the contract method 0x5fab23a8.
//
// Solidity: function totalSlashedStake() view returns(uint256)
func (_Contract *ContractCaller) TotalSlashedStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "totalSlashedStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSlashedStake is a free data retrieval call binding the contract method 0x5fab23a8.
//
// Solidity: function totalSlashedStake() view returns(uint256)
func (_Contract *ContractSession) TotalSlashedStake() (*big.Int, error) {
	return _Contract.Contract.TotalSlashedStake(&_Contract.CallOpts)
}

// TotalSlashedStake is a free data retrieval call binding the contract method 0x5fab23a8.
//
// Solidity: function totalSlashedStake() view returns(uint256)
func (_Contract *ContractCallerSession) TotalSlashedStake() (*big.Int, error) {
	return _Contract.Contract.TotalSlashedStake(&_Contract.CallOpts)
}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_Contract *ContractCaller) TotalStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "totalStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_Contract *ContractSession) TotalStake() (*big.Int, error) {
	return _Contract.Contract.TotalStake(&_Contract.CallOpts)
}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_Contract *ContractCallerSession) TotalStake() (*big.Int, error) {
	return _Contract.Contract.TotalStake(&_Contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractSession) TotalSupply() (*big.Int, error) {
	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractCallerSession) TotalSupply() (*big.Int, error) {
	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_Contract *ContractCaller) TreasuryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "treasuryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_Contract *ContractSession) TreasuryAddress() (common.Address, error) {
	return _Contract.Contract.TreasuryAddress(&_Contract.CallOpts)
}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_Contract *ContractCallerSession) TreasuryAddress() (common.Address, error) {
	return _Contract.Contract.TreasuryAddress(&_Contract.CallOpts)
}

// VoteBookAddress is a free data retrieval call binding the contract method 0x893675c6.
//
// Solidity: function voteBookAddress() view returns(address)
func (_Contract *ContractCaller) VoteBookAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "voteBookAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoteBookAddress is a free data retrieval call binding the contract method 0x893675c6.
//
// Solidity: function voteBookAddress() view returns(address)
func (_Contract *ContractSession) VoteBookAddress() (common.Address, error) {
	return _Contract.Contract.VoteBookAddress(&_Contract.CallOpts)
}

// VoteBookAddress is a free data retrieval call binding the contract method 0x893675c6.
//
// Solidity: function voteBookAddress() view returns(address)
func (_Contract *ContractCallerSession) VoteBookAddress() (common.Address, error) {
	return _Contract.Contract.VoteBookAddress(&_Contract.CallOpts)
}

// SyncValidator is a paid mutator transaction binding the contract method 0xcc8343aa.
//
// Solidity: function _syncValidator(uint256 validatorID, bool syncPubkey) returns()
func (_Contract *ContractTransactor) SyncValidator(opts *bind.TransactOpts, validatorID *big.Int, syncPubkey bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "_syncValidator", validatorID, syncPubkey)
}

// SyncValidator is a paid mutator transaction binding the contract method 0xcc8343aa.
//
// Solidity: function _syncValidator(uint256 validatorID, bool syncPubkey) returns()
func (_Contract *ContractSession) SyncValidator(validatorID *big.Int, syncPubkey bool) (*types.Transaction, error) {
	return _Contract.Contract.SyncValidator(&_Contract.TransactOpts, validatorID, syncPubkey)
}

// SyncValidator is a paid mutator transaction binding the contract method 0xcc8343aa.
//
// Solidity: function _syncValidator(uint256 validatorID, bool syncPubkey) returns()
func (_Contract *ContractTransactorSession) SyncValidator(validatorID *big.Int, syncPubkey bool) (*types.Transaction, error) {
	return _Contract.Contract.SyncValidator(&_Contract.TransactOpts, validatorID, syncPubkey)
}

// BurnRWA is a paid mutator transaction binding the contract method 0xe81ee0c0.
//
// Solidity: function burnRWA(uint256 amount) returns()
func (_Contract *ContractTransactor) BurnRWA(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "burnRWA", amount)
}

// BurnRWA is a paid mutator transaction binding the contract method 0xe81ee0c0.
//
// Solidity: function burnRWA(uint256 amount) returns()
func (_Contract *ContractSession) BurnRWA(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.BurnRWA(&_Contract.TransactOpts, amount)
}

// BurnRWA is a paid mutator transaction binding the contract method 0xe81ee0c0.
//
// Solidity: function burnRWA(uint256 amount) returns()
func (_Contract *ContractTransactorSession) BurnRWA(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.BurnRWA(&_Contract.TransactOpts, amount)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x0962ef79.
//
// Solidity: function claimRewards(uint256 toValidatorID) returns()
func (_Contract *ContractTransactor) ClaimRewards(opts *bind.TransactOpts, toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claimRewards", toValidatorID)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x0962ef79.
//
// Solidity: function claimRewards(uint256 toValidatorID) returns()
func (_Contract *ContractSession) ClaimRewards(toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimRewards(&_Contract.TransactOpts, toValidatorID)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x0962ef79.
//
// Solidity: function claimRewards(uint256 toValidatorID) returns()
func (_Contract *ContractTransactorSession) ClaimRewards(toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimRewards(&_Contract.TransactOpts, toValidatorID)
}

// CreateValidator is a paid mutator transaction binding the contract method 0xa5a470ad.
//
// Solidity: function createValidator(bytes pubkey) payable returns()
func (_Contract *ContractTransactor) CreateValidator(opts *bind.TransactOpts, pubkey []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "createValidator", pubkey)
}

// CreateValidator is a paid mutator transaction binding the contract method 0xa5a470ad.
//
// Solidity: function createValidator(bytes pubkey) payable returns()
func (_Contract *ContractSession) CreateValidator(pubkey []byte) (*types.Transaction, error) {
	return _Contract.Contract.CreateValidator(&_Contract.TransactOpts, pubkey)
}

// CreateValidator is a paid mutator transaction binding the contract method 0xa5a470ad.
//
// Solidity: function createValidator(bytes pubkey) payable returns()
func (_Contract *ContractTransactorSession) CreateValidator(pubkey []byte) (*types.Transaction, error) {
	return _Contract.Contract.CreateValidator(&_Contract.TransactOpts, pubkey)
}

// DeactivateValidator is a paid mutator transaction binding the contract method 0x1e702f83.
//
// Solidity: function deactivateValidator(uint256 validatorID, uint256 status) returns()
func (_Contract *ContractTransactor) DeactivateValidator(opts *bind.TransactOpts, validatorID *big.Int, status *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "deactivateValidator", validatorID, status)
}

// DeactivateValidator is a paid mutator transaction binding the contract method 0x1e702f83.
//
// Solidity: function deactivateValidator(uint256 validatorID, uint256 status) returns()
func (_Contract *ContractSession) DeactivateValidator(validatorID *big.Int, status *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DeactivateValidator(&_Contract.TransactOpts, validatorID, status)
}

// DeactivateValidator is a paid mutator transaction binding the contract method 0x1e702f83.
//
// Solidity: function deactivateValidator(uint256 validatorID, uint256 status) returns()
func (_Contract *ContractTransactorSession) DeactivateValidator(validatorID *big.Int, status *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DeactivateValidator(&_Contract.TransactOpts, validatorID, status)
}

// Delegate is a paid mutator transaction binding the contract method 0x9fa6dd35.
//
// Solidity: function delegate(uint256 toValidatorID) payable returns()
func (_Contract *ContractTransactor) Delegate(opts *bind.TransactOpts, toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "delegate", toValidatorID)
}

// Delegate is a paid mutator transaction binding the contract method 0x9fa6dd35.
//
// Solidity: function delegate(uint256 toValidatorID) payable returns()
func (_Contract *ContractSession) Delegate(toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Delegate(&_Contract.TransactOpts, toValidatorID)
}

// Delegate is a paid mutator transaction binding the contract method 0x9fa6dd35.
//
// Solidity: function delegate(uint256 toValidatorID) payable returns()
func (_Contract *ContractTransactorSession) Delegate(toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Delegate(&_Contract.TransactOpts, toValidatorID)
}

// LiquidateSRWA is a paid mutator transaction binding the contract method 0x907754f8.
//
// Solidity: function liquidateSRWA(address delegator, uint256 toValidatorID, uint256 amount) returns()
func (_Contract *ContractTransactor) LiquidateSRWA(opts *bind.TransactOpts, delegator common.Address, toValidatorID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "liquidateSRWA", delegator, toValidatorID, amount)
}

// LiquidateSRWA is a paid mutator transaction binding the contract method 0x907754f8.
//
// Solidity: function liquidateSRWA(address delegator, uint256 toValidatorID, uint256 amount) returns()
func (_Contract *ContractSession) LiquidateSRWA(delegator common.Address, toValidatorID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LiquidateSRWA(&_Contract.TransactOpts, delegator, toValidatorID, amount)
}

// LiquidateSRWA is a paid mutator transaction binding the contract method 0x907754f8.
//
// Solidity: function liquidateSRWA(address delegator, uint256 toValidatorID, uint256 amount) returns()
func (_Contract *ContractTransactorSession) LiquidateSRWA(delegator common.Address, toValidatorID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LiquidateSRWA(&_Contract.TransactOpts, delegator, toValidatorID, amount)
}

// LockStake is a paid mutator transaction binding the contract method 0xde67f215.
//
// Solidity: function lockStake(uint256 toValidatorID, uint256 lockupDuration, uint256 amount) returns()
func (_Contract *ContractTransactor) LockStake(opts *bind.TransactOpts, toValidatorID *big.Int, lockupDuration *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "lockStake", toValidatorID, lockupDuration, amount)
}

// LockStake is a paid mutator transaction binding the contract method 0xde67f215.
//
// Solidity: function lockStake(uint256 toValidatorID, uint256 lockupDuration, uint256 amount) returns()
func (_Contract *ContractSession) LockStake(toValidatorID *big.Int, lockupDuration *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LockStake(&_Contract.TransactOpts, toValidatorID, lockupDuration, amount)
}

// LockStake is a paid mutator transaction binding the contract method 0xde67f215.
//
// Solidity: function lockStake(uint256 toValidatorID, uint256 lockupDuration, uint256 amount) returns()
func (_Contract *ContractTransactorSession) LockStake(toValidatorID *big.Int, lockupDuration *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LockStake(&_Contract.TransactOpts, toValidatorID, lockupDuration, amount)
}

// RecountVotes is a paid mutator transaction binding the contract method 0x20c0849d.
//
// Solidity: function recountVotes(address delegator, address validatorAuth, bool strict, uint256 gas) returns()
func (_Contract *ContractTransactor) RecountVotes(opts *bind.TransactOpts, delegator common.Address, validatorAuth common.Address, strict bool, gas *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "recountVotes", delegator, validatorAuth, strict, gas)
}

// RecountVotes is a paid mutator transaction binding the contract method 0x20c0849d.
//
// Solidity: function recountVotes(address delegator, address validatorAuth, bool strict, uint256 gas) returns()
func (_Contract *ContractSession) RecountVotes(delegator common.Address, validatorAuth common.Address, strict bool, gas *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RecountVotes(&_Contract.TransactOpts, delegator, validatorAuth, strict, gas)
}

// RecountVotes is a paid mutator transaction binding the contract method 0x20c0849d.
//
// Solidity: function recountVotes(address delegator, address validatorAuth, bool strict, uint256 gas) returns()
func (_Contract *ContractTransactorSession) RecountVotes(delegator common.Address, validatorAuth common.Address, strict bool, gas *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RecountVotes(&_Contract.TransactOpts, delegator, validatorAuth, strict, gas)
}

// RelockStake is a paid mutator transaction binding the contract method 0xbd14d907.
//
// Solidity: function relockStake(uint256 toValidatorID, uint256 lockupDuration, uint256 amount) returns()
func (_Contract *ContractTransactor) RelockStake(opts *bind.TransactOpts, toValidatorID *big.Int, lockupDuration *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "relockStake", toValidatorID, lockupDuration, amount)
}

// RelockStake is a paid mutator transaction binding the contract method 0xbd14d907.
//
// Solidity: function relockStake(uint256 toValidatorID, uint256 lockupDuration, uint256 amount) returns()
func (_Contract *ContractSession) RelockStake(toValidatorID *big.Int, lockupDuration *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RelockStake(&_Contract.TransactOpts, toValidatorID, lockupDuration, amount)
}

// RelockStake is a paid mutator transaction binding the contract method 0xbd14d907.
//
// Solidity: function relockStake(uint256 toValidatorID, uint256 lockupDuration, uint256 amount) returns()
func (_Contract *ContractTransactorSession) RelockStake(toValidatorID *big.Int, lockupDuration *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RelockStake(&_Contract.TransactOpts, toValidatorID, lockupDuration, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RestakeRewards is a paid mutator transaction binding the contract method 0x08c36874.
//
// Solidity: function restakeRewards(uint256 toValidatorID) returns()
func (_Contract *ContractTransactor) RestakeRewards(opts *bind.TransactOpts, toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "restakeRewards", toValidatorID)
}

// RestakeRewards is a paid mutator transaction binding the contract method 0x08c36874.
//
// Solidity: function restakeRewards(uint256 toValidatorID) returns()
func (_Contract *ContractSession) RestakeRewards(toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RestakeRewards(&_Contract.TransactOpts, toValidatorID)
}

// RestakeRewards is a paid mutator transaction binding the contract method 0x08c36874.
//
// Solidity: function restakeRewards(uint256 toValidatorID) returns()
func (_Contract *ContractTransactorSession) RestakeRewards(toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RestakeRewards(&_Contract.TransactOpts, toValidatorID)
}

// SetGenesisDelegation is a paid mutator transaction binding the contract method 0x18f628d4.
//
// Solidity: function setGenesisDelegation(address delegator, uint256 toValidatorID, uint256 stake, uint256 lockedStake, uint256 lockupFromEpoch, uint256 lockupEndTime, uint256 lockupDuration, uint256 earlyUnlockPenalty, uint256 rewards) returns()
func (_Contract *ContractTransactor) SetGenesisDelegation(opts *bind.TransactOpts, delegator common.Address, toValidatorID *big.Int, stake *big.Int, lockedStake *big.Int, lockupFromEpoch *big.Int, lockupEndTime *big.Int, lockupDuration *big.Int, earlyUnlockPenalty *big.Int, rewards *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setGenesisDelegation", delegator, toValidatorID, stake, lockedStake, lockupFromEpoch, lockupEndTime, lockupDuration, earlyUnlockPenalty, rewards)
}

// SetGenesisDelegation is a paid mutator transaction binding the contract method 0x18f628d4.
//
// Solidity: function setGenesisDelegation(address delegator, uint256 toValidatorID, uint256 stake, uint256 lockedStake, uint256 lockupFromEpoch, uint256 lockupEndTime, uint256 lockupDuration, uint256 earlyUnlockPenalty, uint256 rewards) returns()
func (_Contract *ContractSession) SetGenesisDelegation(delegator common.Address, toValidatorID *big.Int, stake *big.Int, lockedStake *big.Int, lockupFromEpoch *big.Int, lockupEndTime *big.Int, lockupDuration *big.Int, earlyUnlockPenalty *big.Int, rewards *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetGenesisDelegation(&_Contract.TransactOpts, delegator, toValidatorID, stake, lockedStake, lockupFromEpoch, lockupEndTime, lockupDuration, earlyUnlockPenalty, rewards)
}

// SetGenesisDelegation is a paid mutator transaction binding the contract method 0x18f628d4.
//
// Solidity: function setGenesisDelegation(address delegator, uint256 toValidatorID, uint256 stake, uint256 lockedStake, uint256 lockupFromEpoch, uint256 lockupEndTime, uint256 lockupDuration, uint256 earlyUnlockPenalty, uint256 rewards) returns()
func (_Contract *ContractTransactorSession) SetGenesisDelegation(delegator common.Address, toValidatorID *big.Int, stake *big.Int, lockedStake *big.Int, lockupFromEpoch *big.Int, lockupEndTime *big.Int, lockupDuration *big.Int, earlyUnlockPenalty *big.Int, rewards *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetGenesisDelegation(&_Contract.TransactOpts, delegator, toValidatorID, stake, lockedStake, lockupFromEpoch, lockupEndTime, lockupDuration, earlyUnlockPenalty, rewards)
}

// SetGenesisValidator is a paid mutator transaction binding the contract method 0x4feb92f3.
//
// Solidity: function setGenesisValidator(address auth, uint256 validatorID, bytes pubkey, uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime) returns()
func (_Contract *ContractTransactor) SetGenesisValidator(opts *bind.TransactOpts, auth common.Address, validatorID *big.Int, pubkey []byte, status *big.Int, createdEpoch *big.Int, createdTime *big.Int, deactivatedEpoch *big.Int, deactivatedTime *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setGenesisValidator", auth, validatorID, pubkey, status, createdEpoch, createdTime, deactivatedEpoch, deactivatedTime)
}

// SetGenesisValidator is a paid mutator transaction binding the contract method 0x4feb92f3.
//
// Solidity: function setGenesisValidator(address auth, uint256 validatorID, bytes pubkey, uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime) returns()
func (_Contract *ContractSession) SetGenesisValidator(auth common.Address, validatorID *big.Int, pubkey []byte, status *big.Int, createdEpoch *big.Int, createdTime *big.Int, deactivatedEpoch *big.Int, deactivatedTime *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetGenesisValidator(&_Contract.TransactOpts, auth, validatorID, pubkey, status, createdEpoch, createdTime, deactivatedEpoch, deactivatedTime)
}

// SetGenesisValidator is a paid mutator transaction binding the contract method 0x4feb92f3.
//
// Solidity: function setGenesisValidator(address auth, uint256 validatorID, bytes pubkey, uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime) returns()
func (_Contract *ContractTransactorSession) SetGenesisValidator(auth common.Address, validatorID *big.Int, pubkey []byte, status *big.Int, createdEpoch *big.Int, createdTime *big.Int, deactivatedEpoch *big.Int, deactivatedTime *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetGenesisValidator(&_Contract.TransactOpts, auth, validatorID, pubkey, status, createdEpoch, createdTime, deactivatedEpoch, deactivatedTime)
}

// StashRewards is a paid mutator transaction binding the contract method 0x8cddb015.
//
// Solidity: function stashRewards(address delegator, uint256 toValidatorID) returns()
func (_Contract *ContractTransactor) StashRewards(opts *bind.TransactOpts, delegator common.Address, toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "stashRewards", delegator, toValidatorID)
}

// StashRewards is a paid mutator transaction binding the contract method 0x8cddb015.
//
// Solidity: function stashRewards(address delegator, uint256 toValidatorID) returns()
func (_Contract *ContractSession) StashRewards(delegator common.Address, toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.StashRewards(&_Contract.TransactOpts, delegator, toValidatorID)
}

// StashRewards is a paid mutator transaction binding the contract method 0x8cddb015.
//
// Solidity: function stashRewards(address delegator, uint256 toValidatorID) returns()
func (_Contract *ContractTransactorSession) StashRewards(delegator common.Address, toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.StashRewards(&_Contract.TransactOpts, delegator, toValidatorID)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// Undelegate is a paid mutator transaction binding the contract method 0x634b91e3.
//
// Solidity: function undelegate(uint256 toValidatorID, uint256 amount) returns()
func (_Contract *ContractTransactor) Undelegate(opts *bind.TransactOpts, toValidatorID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "undelegate", toValidatorID, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x634b91e3.
//
// Solidity: function undelegate(uint256 toValidatorID, uint256 amount) returns()
func (_Contract *ContractSession) Undelegate(toValidatorID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Undelegate(&_Contract.TransactOpts, toValidatorID, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x634b91e3.
//
// Solidity: function undelegate(uint256 toValidatorID, uint256 amount) returns()
func (_Contract *ContractTransactorSession) Undelegate(toValidatorID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Undelegate(&_Contract.TransactOpts, toValidatorID, amount)
}

// UnlockStake is a paid mutator transaction binding the contract method 0x1d3ac42c.
//
// Solidity: function unlockStake(uint256 toValidatorID, uint256 amount) returns(uint256)
func (_Contract *ContractTransactor) UnlockStake(opts *bind.TransactOpts, toValidatorID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "unlockStake", toValidatorID, amount)
}

// UnlockStake is a paid mutator transaction binding the contract method 0x1d3ac42c.
//
// Solidity: function unlockStake(uint256 toValidatorID, uint256 amount) returns(uint256)
func (_Contract *ContractSession) UnlockStake(toValidatorID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UnlockStake(&_Contract.TransactOpts, toValidatorID, amount)
}

// UnlockStake is a paid mutator transaction binding the contract method 0x1d3ac42c.
//
// Solidity: function unlockStake(uint256 toValidatorID, uint256 amount) returns(uint256)
func (_Contract *ContractTransactorSession) UnlockStake(toValidatorID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UnlockStake(&_Contract.TransactOpts, toValidatorID, amount)
}

// UpdateSRWAFinalizer is a paid mutator transaction binding the contract method 0x81588ab6.
//
// Solidity: function updateSRWAFinalizer(address v) returns()
func (_Contract *ContractTransactor) UpdateSRWAFinalizer(opts *bind.TransactOpts, v common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateSRWAFinalizer", v)
}

// UpdateSRWAFinalizer is a paid mutator transaction binding the contract method 0x81588ab6.
//
// Solidity: function updateSRWAFinalizer(address v) returns()
func (_Contract *ContractSession) UpdateSRWAFinalizer(v common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateSRWAFinalizer(&_Contract.TransactOpts, v)
}

// UpdateSRWAFinalizer is a paid mutator transaction binding the contract method 0x81588ab6.
//
// Solidity: function updateSRWAFinalizer(address v) returns()
func (_Contract *ContractTransactorSession) UpdateSRWAFinalizer(v common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateSRWAFinalizer(&_Contract.TransactOpts, v)
}

// UpdateSlashingRefundRatio is a paid mutator transaction binding the contract method 0x4f7c4efb.
//
// Solidity: function updateSlashingRefundRatio(uint256 validatorID, uint256 refundRatio) returns()
func (_Contract *ContractTransactor) UpdateSlashingRefundRatio(opts *bind.TransactOpts, validatorID *big.Int, refundRatio *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateSlashingRefundRatio", validatorID, refundRatio)
}

// UpdateSlashingRefundRatio is a paid mutator transaction binding the contract method 0x4f7c4efb.
//
// Solidity: function updateSlashingRefundRatio(uint256 validatorID, uint256 refundRatio) returns()
func (_Contract *ContractSession) UpdateSlashingRefundRatio(validatorID *big.Int, refundRatio *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateSlashingRefundRatio(&_Contract.TransactOpts, validatorID, refundRatio)
}

// UpdateSlashingRefundRatio is a paid mutator transaction binding the contract method 0x4f7c4efb.
//
// Solidity: function updateSlashingRefundRatio(uint256 validatorID, uint256 refundRatio) returns()
func (_Contract *ContractTransactorSession) UpdateSlashingRefundRatio(validatorID *big.Int, refundRatio *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateSlashingRefundRatio(&_Contract.TransactOpts, validatorID, refundRatio)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 toValidatorID, uint256 wrID) returns()
func (_Contract *ContractTransactor) Withdraw(opts *bind.TransactOpts, toValidatorID *big.Int, wrID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdraw", toValidatorID, wrID)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 toValidatorID, uint256 wrID) returns()
func (_Contract *ContractSession) Withdraw(toValidatorID *big.Int, wrID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, toValidatorID, wrID)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 toValidatorID, uint256 wrID) returns()
func (_Contract *ContractTransactorSession) Withdraw(toValidatorID *big.Int, wrID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, toValidatorID, wrID)
}

// ContractBurntRWAIterator is returned from FilterBurntRWA and is used to iterate over the raw logs and unpacked data for BurntRWA events raised by the Contract contract.
type ContractBurntRWAIterator struct {
	Event *ContractBurntRWA // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractBurntRWAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBurntRWA)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractBurntRWA)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractBurntRWAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBurntRWAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBurntRWA represents a BurntRWA event raised by the Contract contract.
type ContractBurntRWA struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurntRWA is a free log retrieval operation binding the contract event 0xe733dd86f3139a17476f9309db31e45d999e6469068d5734f7d51b23110c17bc.
//
// Solidity: event BurntRWA(uint256 amount)
func (_Contract *ContractFilterer) FilterBurntRWA(opts *bind.FilterOpts) (*ContractBurntRWAIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "BurntRWA")
	if err != nil {
		return nil, err
	}
	return &ContractBurntRWAIterator{contract: _Contract.contract, event: "BurntRWA", logs: logs, sub: sub}, nil
}

// WatchBurntRWA is a free log subscription operation binding the contract event 0xe733dd86f3139a17476f9309db31e45d999e6469068d5734f7d51b23110c17bc.
//
// Solidity: event BurntRWA(uint256 amount)
func (_Contract *ContractFilterer) WatchBurntRWA(opts *bind.WatchOpts, sink chan<- *ContractBurntRWA) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "BurntRWA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBurntRWA)
				if err := _Contract.contract.UnpackLog(event, "BurntRWA", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBurntRWA is a log parse operation binding the contract event 0xe733dd86f3139a17476f9309db31e45d999e6469068d5734f7d51b23110c17bc.
//
// Solidity: event BurntRWA(uint256 amount)
func (_Contract *ContractFilterer) ParseBurntRWA(log types.Log) (*ContractBurntRWA, error) {
	event := new(ContractBurntRWA)
	if err := _Contract.contract.UnpackLog(event, "BurntRWA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractChangedValidatorStatusIterator is returned from FilterChangedValidatorStatus and is used to iterate over the raw logs and unpacked data for ChangedValidatorStatus events raised by the Contract contract.
type ContractChangedValidatorStatusIterator struct {
	Event *ContractChangedValidatorStatus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractChangedValidatorStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractChangedValidatorStatus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractChangedValidatorStatus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractChangedValidatorStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractChangedValidatorStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractChangedValidatorStatus represents a ChangedValidatorStatus event raised by the Contract contract.
type ContractChangedValidatorStatus struct {
	ValidatorID *big.Int
	Status      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterChangedValidatorStatus is a free log retrieval operation binding the contract event 0xcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e.
//
// Solidity: event ChangedValidatorStatus(uint256 indexed validatorID, uint256 status)
func (_Contract *ContractFilterer) FilterChangedValidatorStatus(opts *bind.FilterOpts, validatorID []*big.Int) (*ContractChangedValidatorStatusIterator, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ChangedValidatorStatus", validatorIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractChangedValidatorStatusIterator{contract: _Contract.contract, event: "ChangedValidatorStatus", logs: logs, sub: sub}, nil
}

// WatchChangedValidatorStatus is a free log subscription operation binding the contract event 0xcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e.
//
// Solidity: event ChangedValidatorStatus(uint256 indexed validatorID, uint256 status)
func (_Contract *ContractFilterer) WatchChangedValidatorStatus(opts *bind.WatchOpts, sink chan<- *ContractChangedValidatorStatus, validatorID []*big.Int) (event.Subscription, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ChangedValidatorStatus", validatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractChangedValidatorStatus)
				if err := _Contract.contract.UnpackLog(event, "ChangedValidatorStatus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChangedValidatorStatus is a log parse operation binding the contract event 0xcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e.
//
// Solidity: event ChangedValidatorStatus(uint256 indexed validatorID, uint256 status)
func (_Contract *ContractFilterer) ParseChangedValidatorStatus(log types.Log) (*ContractChangedValidatorStatus, error) {
	event := new(ContractChangedValidatorStatus)
	if err := _Contract.contract.UnpackLog(event, "ChangedValidatorStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractClaimedRewardsIterator is returned from FilterClaimedRewards and is used to iterate over the raw logs and unpacked data for ClaimedRewards events raised by the Contract contract.
type ContractClaimedRewardsIterator struct {
	Event *ContractClaimedRewards // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractClaimedRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractClaimedRewards)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractClaimedRewards)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractClaimedRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractClaimedRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractClaimedRewards represents a ClaimedRewards event raised by the Contract contract.
type ContractClaimedRewards struct {
	Delegator         common.Address
	ToValidatorID     *big.Int
	LockupExtraReward *big.Int
	LockupBaseReward  *big.Int
	UnlockedReward    *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterClaimedRewards is a free log retrieval operation binding the contract event 0xc1d8eb6e444b89fb8ff0991c19311c070df704ccb009e210d1462d5b2410bf45.
//
// Solidity: event ClaimedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
func (_Contract *ContractFilterer) FilterClaimedRewards(opts *bind.FilterOpts, delegator []common.Address, toValidatorID []*big.Int) (*ContractClaimedRewardsIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ClaimedRewards", delegatorRule, toValidatorIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractClaimedRewardsIterator{contract: _Contract.contract, event: "ClaimedRewards", logs: logs, sub: sub}, nil
}

// WatchClaimedRewards is a free log subscription operation binding the contract event 0xc1d8eb6e444b89fb8ff0991c19311c070df704ccb009e210d1462d5b2410bf45.
//
// Solidity: event ClaimedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
func (_Contract *ContractFilterer) WatchClaimedRewards(opts *bind.WatchOpts, sink chan<- *ContractClaimedRewards, delegator []common.Address, toValidatorID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ClaimedRewards", delegatorRule, toValidatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractClaimedRewards)
				if err := _Contract.contract.UnpackLog(event, "ClaimedRewards", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimedRewards is a log parse operation binding the contract event 0xc1d8eb6e444b89fb8ff0991c19311c070df704ccb009e210d1462d5b2410bf45.
//
// Solidity: event ClaimedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
func (_Contract *ContractFilterer) ParseClaimedRewards(log types.Log) (*ContractClaimedRewards, error) {
	event := new(ContractClaimedRewards)
	if err := _Contract.contract.UnpackLog(event, "ClaimedRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractCreatedValidatorIterator is returned from FilterCreatedValidator and is used to iterate over the raw logs and unpacked data for CreatedValidator events raised by the Contract contract.
type ContractCreatedValidatorIterator struct {
	Event *ContractCreatedValidator // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractCreatedValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCreatedValidator)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractCreatedValidator)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractCreatedValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCreatedValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCreatedValidator represents a CreatedValidator event raised by the Contract contract.
type ContractCreatedValidator struct {
	ValidatorID  *big.Int
	Auth         common.Address
	CreatedEpoch *big.Int
	CreatedTime  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCreatedValidator is a free log retrieval operation binding the contract event 0x49bca1ed2666922f9f1690c26a569e1299c2a715fe57647d77e81adfabbf25bf.
//
// Solidity: event CreatedValidator(uint256 indexed validatorID, address indexed auth, uint256 createdEpoch, uint256 createdTime)
func (_Contract *ContractFilterer) FilterCreatedValidator(opts *bind.FilterOpts, validatorID []*big.Int, auth []common.Address) (*ContractCreatedValidatorIterator, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}
	var authRule []interface{}
	for _, authItem := range auth {
		authRule = append(authRule, authItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "CreatedValidator", validatorIDRule, authRule)
	if err != nil {
		return nil, err
	}
	return &ContractCreatedValidatorIterator{contract: _Contract.contract, event: "CreatedValidator", logs: logs, sub: sub}, nil
}

// WatchCreatedValidator is a free log subscription operation binding the contract event 0x49bca1ed2666922f9f1690c26a569e1299c2a715fe57647d77e81adfabbf25bf.
//
// Solidity: event CreatedValidator(uint256 indexed validatorID, address indexed auth, uint256 createdEpoch, uint256 createdTime)
func (_Contract *ContractFilterer) WatchCreatedValidator(opts *bind.WatchOpts, sink chan<- *ContractCreatedValidator, validatorID []*big.Int, auth []common.Address) (event.Subscription, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}
	var authRule []interface{}
	for _, authItem := range auth {
		authRule = append(authRule, authItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "CreatedValidator", validatorIDRule, authRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCreatedValidator)
				if err := _Contract.contract.UnpackLog(event, "CreatedValidator", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreatedValidator is a log parse operation binding the contract event 0x49bca1ed2666922f9f1690c26a569e1299c2a715fe57647d77e81adfabbf25bf.
//
// Solidity: event CreatedValidator(uint256 indexed validatorID, address indexed auth, uint256 createdEpoch, uint256 createdTime)
func (_Contract *ContractFilterer) ParseCreatedValidator(log types.Log) (*ContractCreatedValidator, error) {
	event := new(ContractCreatedValidator)
	if err := _Contract.contract.UnpackLog(event, "CreatedValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDeactivatedValidatorIterator is returned from FilterDeactivatedValidator and is used to iterate over the raw logs and unpacked data for DeactivatedValidator events raised by the Contract contract.
type ContractDeactivatedValidatorIterator struct {
	Event *ContractDeactivatedValidator // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDeactivatedValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDeactivatedValidator)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDeactivatedValidator)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDeactivatedValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDeactivatedValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDeactivatedValidator represents a DeactivatedValidator event raised by the Contract contract.
type ContractDeactivatedValidator struct {
	ValidatorID      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDeactivatedValidator is a free log retrieval operation binding the contract event 0xac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e47.
//
// Solidity: event DeactivatedValidator(uint256 indexed validatorID, uint256 deactivatedEpoch, uint256 deactivatedTime)
func (_Contract *ContractFilterer) FilterDeactivatedValidator(opts *bind.FilterOpts, validatorID []*big.Int) (*ContractDeactivatedValidatorIterator, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DeactivatedValidator", validatorIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractDeactivatedValidatorIterator{contract: _Contract.contract, event: "DeactivatedValidator", logs: logs, sub: sub}, nil
}

// WatchDeactivatedValidator is a free log subscription operation binding the contract event 0xac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e47.
//
// Solidity: event DeactivatedValidator(uint256 indexed validatorID, uint256 deactivatedEpoch, uint256 deactivatedTime)
func (_Contract *ContractFilterer) WatchDeactivatedValidator(opts *bind.WatchOpts, sink chan<- *ContractDeactivatedValidator, validatorID []*big.Int) (event.Subscription, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DeactivatedValidator", validatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDeactivatedValidator)
				if err := _Contract.contract.UnpackLog(event, "DeactivatedValidator", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeactivatedValidator is a log parse operation binding the contract event 0xac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e47.
//
// Solidity: event DeactivatedValidator(uint256 indexed validatorID, uint256 deactivatedEpoch, uint256 deactivatedTime)
func (_Contract *ContractFilterer) ParseDeactivatedValidator(log types.Log) (*ContractDeactivatedValidator, error) {
	event := new(ContractDeactivatedValidator)
	if err := _Contract.contract.UnpackLog(event, "DeactivatedValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegatedIterator is returned from FilterDelegated and is used to iterate over the raw logs and unpacked data for Delegated events raised by the Contract contract.
type ContractDelegatedIterator struct {
	Event *ContractDelegated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDelegated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegated represents a Delegated event raised by the Contract contract.
type ContractDelegated struct {
	Delegator     common.Address
	ToValidatorID *big.Int
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDelegated is a free log retrieval operation binding the contract event 0x9a8f44850296624dadfd9c246d17e47171d35727a181bd090aa14bbbe00238bb.
//
// Solidity: event Delegated(address indexed delegator, uint256 indexed toValidatorID, uint256 amount)
func (_Contract *ContractFilterer) FilterDelegated(opts *bind.FilterOpts, delegator []common.Address, toValidatorID []*big.Int) (*ContractDelegatedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Delegated", delegatorRule, toValidatorIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegatedIterator{contract: _Contract.contract, event: "Delegated", logs: logs, sub: sub}, nil
}

// WatchDelegated is a free log subscription operation binding the contract event 0x9a8f44850296624dadfd9c246d17e47171d35727a181bd090aa14bbbe00238bb.
//
// Solidity: event Delegated(address indexed delegator, uint256 indexed toValidatorID, uint256 amount)
func (_Contract *ContractFilterer) WatchDelegated(opts *bind.WatchOpts, sink chan<- *ContractDelegated, delegator []common.Address, toValidatorID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Delegated", delegatorRule, toValidatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegated)
				if err := _Contract.contract.UnpackLog(event, "Delegated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegated is a log parse operation binding the contract event 0x9a8f44850296624dadfd9c246d17e47171d35727a181bd090aa14bbbe00238bb.
//
// Solidity: event Delegated(address indexed delegator, uint256 indexed toValidatorID, uint256 amount)
func (_Contract *ContractFilterer) ParseDelegated(log types.Log) (*ContractDelegated, error) {
	event := new(ContractDelegated)
	if err := _Contract.contract.UnpackLog(event, "Delegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLockedUpStakeIterator is returned from FilterLockedUpStake and is used to iterate over the raw logs and unpacked data for LockedUpStake events raised by the Contract contract.
type ContractLockedUpStakeIterator struct {
	Event *ContractLockedUpStake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractLockedUpStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLockedUpStake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractLockedUpStake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractLockedUpStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLockedUpStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLockedUpStake represents a LockedUpStake event raised by the Contract contract.
type ContractLockedUpStake struct {
	Delegator   common.Address
	ValidatorID *big.Int
	Duration    *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLockedUpStake is a free log retrieval operation binding the contract event 0x138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e78.
//
// Solidity: event LockedUpStake(address indexed delegator, uint256 indexed validatorID, uint256 duration, uint256 amount)
func (_Contract *ContractFilterer) FilterLockedUpStake(opts *bind.FilterOpts, delegator []common.Address, validatorID []*big.Int) (*ContractLockedUpStakeIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LockedUpStake", delegatorRule, validatorIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractLockedUpStakeIterator{contract: _Contract.contract, event: "LockedUpStake", logs: logs, sub: sub}, nil
}

// WatchLockedUpStake is a free log subscription operation binding the contract event 0x138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e78.
//
// Solidity: event LockedUpStake(address indexed delegator, uint256 indexed validatorID, uint256 duration, uint256 amount)
func (_Contract *ContractFilterer) WatchLockedUpStake(opts *bind.WatchOpts, sink chan<- *ContractLockedUpStake, delegator []common.Address, validatorID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LockedUpStake", delegatorRule, validatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLockedUpStake)
				if err := _Contract.contract.UnpackLog(event, "LockedUpStake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLockedUpStake is a log parse operation binding the contract event 0x138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e78.
//
// Solidity: event LockedUpStake(address indexed delegator, uint256 indexed validatorID, uint256 duration, uint256 amount)
func (_Contract *ContractFilterer) ParseLockedUpStake(log types.Log) (*ContractLockedUpStake, error) {
	event := new(ContractLockedUpStake)
	if err := _Contract.contract.UnpackLog(event, "LockedUpStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRefundedSlashedLegacyDelegationIterator is returned from FilterRefundedSlashedLegacyDelegation and is used to iterate over the raw logs and unpacked data for RefundedSlashedLegacyDelegation events raised by the Contract contract.
type ContractRefundedSlashedLegacyDelegationIterator struct {
	Event *ContractRefundedSlashedLegacyDelegation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractRefundedSlashedLegacyDelegationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRefundedSlashedLegacyDelegation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractRefundedSlashedLegacyDelegation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractRefundedSlashedLegacyDelegationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRefundedSlashedLegacyDelegationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRefundedSlashedLegacyDelegation represents a RefundedSlashedLegacyDelegation event raised by the Contract contract.
type ContractRefundedSlashedLegacyDelegation struct {
	Delegator   common.Address
	ValidatorID *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRefundedSlashedLegacyDelegation is a free log retrieval operation binding the contract event 0x172fdfaf5222519d28d2794b7617be033f46d954f9b6c3896e7d2611ff444252.
//
// Solidity: event RefundedSlashedLegacyDelegation(address indexed delegator, uint256 indexed validatorID, uint256 amount)
func (_Contract *ContractFilterer) FilterRefundedSlashedLegacyDelegation(opts *bind.FilterOpts, delegator []common.Address, validatorID []*big.Int) (*ContractRefundedSlashedLegacyDelegationIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RefundedSlashedLegacyDelegation", delegatorRule, validatorIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractRefundedSlashedLegacyDelegationIterator{contract: _Contract.contract, event: "RefundedSlashedLegacyDelegation", logs: logs, sub: sub}, nil
}

// WatchRefundedSlashedLegacyDelegation is a free log subscription operation binding the contract event 0x172fdfaf5222519d28d2794b7617be033f46d954f9b6c3896e7d2611ff444252.
//
// Solidity: event RefundedSlashedLegacyDelegation(address indexed delegator, uint256 indexed validatorID, uint256 amount)
func (_Contract *ContractFilterer) WatchRefundedSlashedLegacyDelegation(opts *bind.WatchOpts, sink chan<- *ContractRefundedSlashedLegacyDelegation, delegator []common.Address, validatorID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RefundedSlashedLegacyDelegation", delegatorRule, validatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRefundedSlashedLegacyDelegation)
				if err := _Contract.contract.UnpackLog(event, "RefundedSlashedLegacyDelegation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRefundedSlashedLegacyDelegation is a log parse operation binding the contract event 0x172fdfaf5222519d28d2794b7617be033f46d954f9b6c3896e7d2611ff444252.
//
// Solidity: event RefundedSlashedLegacyDelegation(address indexed delegator, uint256 indexed validatorID, uint256 amount)
func (_Contract *ContractFilterer) ParseRefundedSlashedLegacyDelegation(log types.Log) (*ContractRefundedSlashedLegacyDelegation, error) {
	event := new(ContractRefundedSlashedLegacyDelegation)
	if err := _Contract.contract.UnpackLog(event, "RefundedSlashedLegacyDelegation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRestakedRewardsIterator is returned from FilterRestakedRewards and is used to iterate over the raw logs and unpacked data for RestakedRewards events raised by the Contract contract.
type ContractRestakedRewardsIterator struct {
	Event *ContractRestakedRewards // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractRestakedRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRestakedRewards)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractRestakedRewards)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractRestakedRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRestakedRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRestakedRewards represents a RestakedRewards event raised by the Contract contract.
type ContractRestakedRewards struct {
	Delegator         common.Address
	ToValidatorID     *big.Int
	LockupExtraReward *big.Int
	LockupBaseReward  *big.Int
	UnlockedReward    *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRestakedRewards is a free log retrieval operation binding the contract event 0x4119153d17a36f9597d40e3ab4148d03261a439dddbec4e91799ab7159608e26.
//
// Solidity: event RestakedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
func (_Contract *ContractFilterer) FilterRestakedRewards(opts *bind.FilterOpts, delegator []common.Address, toValidatorID []*big.Int) (*ContractRestakedRewardsIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RestakedRewards", delegatorRule, toValidatorIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractRestakedRewardsIterator{contract: _Contract.contract, event: "RestakedRewards", logs: logs, sub: sub}, nil
}

// WatchRestakedRewards is a free log subscription operation binding the contract event 0x4119153d17a36f9597d40e3ab4148d03261a439dddbec4e91799ab7159608e26.
//
// Solidity: event RestakedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
func (_Contract *ContractFilterer) WatchRestakedRewards(opts *bind.WatchOpts, sink chan<- *ContractRestakedRewards, delegator []common.Address, toValidatorID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RestakedRewards", delegatorRule, toValidatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRestakedRewards)
				if err := _Contract.contract.UnpackLog(event, "RestakedRewards", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRestakedRewards is a log parse operation binding the contract event 0x4119153d17a36f9597d40e3ab4148d03261a439dddbec4e91799ab7159608e26.
//
// Solidity: event RestakedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
func (_Contract *ContractFilterer) ParseRestakedRewards(log types.Log) (*ContractRestakedRewards, error) {
	event := new(ContractRestakedRewards)
	if err := _Contract.contract.UnpackLog(event, "RestakedRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUndelegatedIterator is returned from FilterUndelegated and is used to iterate over the raw logs and unpacked data for Undelegated events raised by the Contract contract.
type ContractUndelegatedIterator struct {
	Event *ContractUndelegated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUndelegated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractUndelegated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUndelegated represents a Undelegated event raised by the Contract contract.
type ContractUndelegated struct {
	Delegator     common.Address
	ToValidatorID *big.Int
	WrID          *big.Int
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUndelegated is a free log retrieval operation binding the contract event 0xd3bb4e423fbea695d16b982f9f682dc5f35152e5411646a8a5a79a6b02ba8d57.
//
// Solidity: event Undelegated(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func (_Contract *ContractFilterer) FilterUndelegated(opts *bind.FilterOpts, delegator []common.Address, toValidatorID []*big.Int, wrID []*big.Int) (*ContractUndelegatedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}
	var wrIDRule []interface{}
	for _, wrIDItem := range wrID {
		wrIDRule = append(wrIDRule, wrIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Undelegated", delegatorRule, toValidatorIDRule, wrIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractUndelegatedIterator{contract: _Contract.contract, event: "Undelegated", logs: logs, sub: sub}, nil
}

// WatchUndelegated is a free log subscription operation binding the contract event 0xd3bb4e423fbea695d16b982f9f682dc5f35152e5411646a8a5a79a6b02ba8d57.
//
// Solidity: event Undelegated(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func (_Contract *ContractFilterer) WatchUndelegated(opts *bind.WatchOpts, sink chan<- *ContractUndelegated, delegator []common.Address, toValidatorID []*big.Int, wrID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}
	var wrIDRule []interface{}
	for _, wrIDItem := range wrID {
		wrIDRule = append(wrIDRule, wrIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Undelegated", delegatorRule, toValidatorIDRule, wrIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUndelegated)
				if err := _Contract.contract.UnpackLog(event, "Undelegated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUndelegated is a log parse operation binding the contract event 0xd3bb4e423fbea695d16b982f9f682dc5f35152e5411646a8a5a79a6b02ba8d57.
//
// Solidity: event Undelegated(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func (_Contract *ContractFilterer) ParseUndelegated(log types.Log) (*ContractUndelegated, error) {
	event := new(ContractUndelegated)
	if err := _Contract.contract.UnpackLog(event, "Undelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUnlockedStakeIterator is returned from FilterUnlockedStake and is used to iterate over the raw logs and unpacked data for UnlockedStake events raised by the Contract contract.
type ContractUnlockedStakeIterator struct {
	Event *ContractUnlockedStake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractUnlockedStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUnlockedStake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractUnlockedStake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractUnlockedStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUnlockedStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUnlockedStake represents a UnlockedStake event raised by the Contract contract.
type ContractUnlockedStake struct {
	Delegator   common.Address
	ValidatorID *big.Int
	Amount      *big.Int
	Penalty     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnlockedStake is a free log retrieval operation binding the contract event 0xef6c0c14fe9aa51af36acd791464dec3badbde668b63189b47bfa4e25be9b2b9.
//
// Solidity: event UnlockedStake(address indexed delegator, uint256 indexed validatorID, uint256 amount, uint256 penalty)
func (_Contract *ContractFilterer) FilterUnlockedStake(opts *bind.FilterOpts, delegator []common.Address, validatorID []*big.Int) (*ContractUnlockedStakeIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UnlockedStake", delegatorRule, validatorIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractUnlockedStakeIterator{contract: _Contract.contract, event: "UnlockedStake", logs: logs, sub: sub}, nil
}

// WatchUnlockedStake is a free log subscription operation binding the contract event 0xef6c0c14fe9aa51af36acd791464dec3badbde668b63189b47bfa4e25be9b2b9.
//
// Solidity: event UnlockedStake(address indexed delegator, uint256 indexed validatorID, uint256 amount, uint256 penalty)
func (_Contract *ContractFilterer) WatchUnlockedStake(opts *bind.WatchOpts, sink chan<- *ContractUnlockedStake, delegator []common.Address, validatorID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UnlockedStake", delegatorRule, validatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUnlockedStake)
				if err := _Contract.contract.UnpackLog(event, "UnlockedStake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnlockedStake is a log parse operation binding the contract event 0xef6c0c14fe9aa51af36acd791464dec3badbde668b63189b47bfa4e25be9b2b9.
//
// Solidity: event UnlockedStake(address indexed delegator, uint256 indexed validatorID, uint256 amount, uint256 penalty)
func (_Contract *ContractFilterer) ParseUnlockedStake(log types.Log) (*ContractUnlockedStake, error) {
	event := new(ContractUnlockedStake)
	if err := _Contract.contract.UnpackLog(event, "UnlockedStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUpdatedSlashingRefundRatioIterator is returned from FilterUpdatedSlashingRefundRatio and is used to iterate over the raw logs and unpacked data for UpdatedSlashingRefundRatio events raised by the Contract contract.
type ContractUpdatedSlashingRefundRatioIterator struct {
	Event *ContractUpdatedSlashingRefundRatio // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractUpdatedSlashingRefundRatioIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpdatedSlashingRefundRatio)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractUpdatedSlashingRefundRatio)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractUpdatedSlashingRefundRatioIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpdatedSlashingRefundRatioIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpdatedSlashingRefundRatio represents a UpdatedSlashingRefundRatio event raised by the Contract contract.
type ContractUpdatedSlashingRefundRatio struct {
	ValidatorID *big.Int
	RefundRatio *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedSlashingRefundRatio is a free log retrieval operation binding the contract event 0x047575f43f09a7a093d94ec483064acfc61b7e25c0de28017da442abf99cb917.
//
// Solidity: event UpdatedSlashingRefundRatio(uint256 indexed validatorID, uint256 refundRatio)
func (_Contract *ContractFilterer) FilterUpdatedSlashingRefundRatio(opts *bind.FilterOpts, validatorID []*big.Int) (*ContractUpdatedSlashingRefundRatioIterator, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdatedSlashingRefundRatio", validatorIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractUpdatedSlashingRefundRatioIterator{contract: _Contract.contract, event: "UpdatedSlashingRefundRatio", logs: logs, sub: sub}, nil
}

// WatchUpdatedSlashingRefundRatio is a free log subscription operation binding the contract event 0x047575f43f09a7a093d94ec483064acfc61b7e25c0de28017da442abf99cb917.
//
// Solidity: event UpdatedSlashingRefundRatio(uint256 indexed validatorID, uint256 refundRatio)
func (_Contract *ContractFilterer) WatchUpdatedSlashingRefundRatio(opts *bind.WatchOpts, sink chan<- *ContractUpdatedSlashingRefundRatio, validatorID []*big.Int) (event.Subscription, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdatedSlashingRefundRatio", validatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpdatedSlashingRefundRatio)
				if err := _Contract.contract.UnpackLog(event, "UpdatedSlashingRefundRatio", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedSlashingRefundRatio is a log parse operation binding the contract event 0x047575f43f09a7a093d94ec483064acfc61b7e25c0de28017da442abf99cb917.
//
// Solidity: event UpdatedSlashingRefundRatio(uint256 indexed validatorID, uint256 refundRatio)
func (_Contract *ContractFilterer) ParseUpdatedSlashingRefundRatio(log types.Log) (*ContractUpdatedSlashingRefundRatio, error) {
	event := new(ContractUpdatedSlashingRefundRatio)
	if err := _Contract.contract.UnpackLog(event, "UpdatedSlashingRefundRatio", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Contract contract.
type ContractWithdrawnIterator struct {
	Event *ContractWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractWithdrawn represents a Withdrawn event raised by the Contract contract.
type ContractWithdrawn struct {
	Delegator     common.Address
	ToValidatorID *big.Int
	WrID          *big.Int
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a21.
//
// Solidity: event Withdrawn(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func (_Contract *ContractFilterer) FilterWithdrawn(opts *bind.FilterOpts, delegator []common.Address, toValidatorID []*big.Int, wrID []*big.Int) (*ContractWithdrawnIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}
	var wrIDRule []interface{}
	for _, wrIDItem := range wrID {
		wrIDRule = append(wrIDRule, wrIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Withdrawn", delegatorRule, toValidatorIDRule, wrIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractWithdrawnIterator{contract: _Contract.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a21.
//
// Solidity: event Withdrawn(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func (_Contract *ContractFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractWithdrawn, delegator []common.Address, toValidatorID []*big.Int, wrID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}
	var wrIDRule []interface{}
	for _, wrIDItem := range wrID {
		wrIDRule = append(wrIDRule, wrIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Withdrawn", delegatorRule, toValidatorIDRule, wrIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractWithdrawn)
				if err := _Contract.contract.UnpackLog(event, "Withdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawn is a log parse operation binding the contract event 0x75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a21.
//
// Solidity: event Withdrawn(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func (_Contract *ContractFilterer) ParseWithdrawn(log types.Log) (*ContractWithdrawn, error) {
	event := new(ContractWithdrawn)
	if err := _Contract.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var ContractBinRuntime = "0x60806040526004361061038c5760003560e01c8063854873e1116101dc578063bd14d90711610102578063cfdbb7cd116100a0578063df00c9221161006f578063df00c92214610e41578063e261641a14610e7e578063e81ee0c014610ebb578063f2fde38b14610ee45761038c565b8063cfdbb7cd14610d73578063d96ed50514610db0578063dc31e1af14610ddb578063de67f21514610e185761038c565b8063c65ee0e1116100dc578063c65ee0e114610ca5578063c7be95de14610ce2578063cc8343aa14610d0d578063cfd4766314610d365761038c565b8063bd14d90714610c14578063c3de580e14610c3d578063c5f956af14610c7a5761038c565b806396c7ee461161017a578063a86a056f11610149578063a86a056f14610b18578063b5d8962714610b55578063b810e41114610b98578063b88a37e214610bd75761038c565b806396c7ee4614610a635780639fa6dd3514610aa3578063a198d22914610abf578063a5a470ad14610afc5761038c565b80638cddb015116101b65780638cddb015146109bb5780638da5cb5b146109e45780638f32d59b14610a0f578063907754f814610a3a5761038c565b8063854873e114610928578063893675c6146109655780638b0e9f3f146109905761038c565b80634f7c4efb116102c1578063634b91e31161025f578063715018a61161022e578063715018a61461089257806376671808146108a95780637cacb1d6146108d457806381588ab6146108ff5761038c565b8063634b91e3146107b2578063670322f8146107db5780636f49866314610818578063702797e3146108555761038c565b806358f95b801161029b57806358f95b80146106d05780635fab23a81461070d5780636099ecb21461073857806361e53fcc146107755761038c565b80634f7c4efb146106415780634feb92f31461066a5780635601fe01146106935761038c565b80631d3ac42c1161032e57806320c0849d1161030857806320c0849d1461058157806328f73148146105aa57806339b80c00146105d5578063441a3e70146106185761038c565b80631d3ac42c146104dc5780631e702f83146105195780631f270152146105425761038c565b80630e559d821161036a5780630e559d821461042057806312622d0e1461044b57806318160ddd1461048857806318f628d4146104b35761038c565b80630135b1db1461039157806308c36874146103ce5780630962ef79146103f7575b600080fd5b34801561039d57600080fd5b506103b860048036036103b391908101906164b3565b610f0d565b6040516103c59190617bc7565b60405180910390f35b3480156103da57600080fd5b506103f560048036036103f0919081019061682f565b610f25565b005b34801561040357600080fd5b5061041e6004803603610419919081019061682f565b611044565b005b34801561042c57600080fd5b5061043561119e565b6040516104429190617612565b60405180910390f35b34801561045757600080fd5b50610472600480360361046d919081019061653f565b6111c4565b60405161047f9190617bc7565b60405180910390f35b34801561049457600080fd5b5061049d6112eb565b6040516104aa9190617bc7565b60405180910390f35b3480156104bf57600080fd5b506104da60048036036104d591908101906166fb565b6112f1565b005b3480156104e857600080fd5b5061050360048036036104fe91908101906168bd565b611520565b6040516105109190617bc7565b60405180910390f35b34801561052557600080fd5b50610540600480360361053b91908101906168bd565b61175e565b005b34801561054e57600080fd5b5061056960048036036105649190810190616649565b61184b565b60405161057893929190617c64565b60405180910390f35b34801561058d57600080fd5b506105a860048036036105a391908101906164dc565b61188f565b005b3480156105b657600080fd5b506105bf6119ff565b6040516105cc9190617bc7565b60405180910390f35b3480156105e157600080fd5b506105fc60048036036105f7919081019061682f565b611a05565b60405161060f9796959493929190617d4f565b60405180910390f35b34801561062457600080fd5b5061063f600480360361063a91908101906168bd565b611a47565b005b34801561064d57600080fd5b50610668600480360361066391908101906168bd565b611a57565b005b34801561067657600080fd5b50610691600480360361068c919081019061657b565b611b84565b005b34801561069f57600080fd5b506106ba60048036036106b5919081019061682f565b611c3d565b6040516106c79190617bc7565b60405180910390f35b3480156106dc57600080fd5b506106f760048036036106f291908101906168bd565b611ccd565b6040516107049190617bc7565b60405180910390f35b34801561071957600080fd5b50610722611cff565b60405161072f9190617bc7565b60405180910390f35b34801561074457600080fd5b5061075f600480360361075a919081019061653f565b611d05565b60405161076c9190617bc7565b60405180910390f35b34801561078157600080fd5b5061079c600480360361079791908101906168bd565b611d55565b6040516107a99190617bc7565b60405180910390f35b3480156107be57600080fd5b506107d960048036036107d491908101906168bd565b611d87565b005b3480156107e757600080fd5b5061080260048036036107fd919081019061653f565b612130565b60405161080f9190617bc7565b60405180910390f35b34801561082457600080fd5b5061083f600480360361083a919081019061653f565b6121a6565b60405161084c9190617bc7565b60405180910390f35b34801561086157600080fd5b5061087c60048036036108779190810190616698565b612265565b60405161088991906176c4565b60405180910390f35b34801561089e57600080fd5b506108a761238e565b005b3480156108b557600080fd5b506108be612496565b6040516108cb9190617bc7565b60405180910390f35b3480156108e057600080fd5b506108e96124a3565b6040516108f69190617bc7565b60405180910390f35b34801561090b57600080fd5b50610926600480360361092191908101906164b3565b6124a9565b005b34801561093457600080fd5b5061094f600480360361094a919081019061682f565b612534565b60405161095c9190617723565b60405180910390f35b34801561097157600080fd5b5061097a6125e4565b6040516109879190617612565b60405180910390f35b34801561099c57600080fd5b506109a561260a565b6040516109b29190617bc7565b60405180910390f35b3480156109c757600080fd5b506109e260048036036109dd919081019061653f565b612610565b005b3480156109f057600080fd5b506109f961265d565b604051610a069190617612565b60405180910390f35b348015610a1b57600080fd5b50610a24612687565b604051610a319190617708565b60405180910390f35b348015610a4657600080fd5b50610a616004803603610a5c9190810190616649565b6126df565b005b348015610a6f57600080fd5b50610a8a6004803603610a85919081019061653f565b612b39565b604051610a9a9493929190617c9b565b60405180910390f35b610abd6004803603610ab8919081019061682f565b612b76565b005b348015610acb57600080fd5b50610ae66004803603610ae191908101906168bd565b612b84565b604051610af39190617bc7565b60405180910390f35b610b166004803603610b1191908101906167ea565b612bb6565b005b348015610b2457600080fd5b50610b3f6004803603610b3a919081019061653f565b612d3d565b604051610b4c9190617bc7565b60405180910390f35b348015610b6157600080fd5b50610b7c6004803603610b77919081019061682f565b612d62565b604051610b8f9796959493929190617ce0565b60405180910390f35b348015610ba457600080fd5b50610bbf6004803603610bba919081019061653f565b612dc4565b604051610bce93929190617c64565b60405180910390f35b348015610be357600080fd5b50610bfe6004803603610bf9919081019061682f565b612dfb565b604051610c0b91906176e6565b60405180910390f35b348015610c2057600080fd5b50610c3b6004803603610c3691908101906168f9565b612e69565b005b348015610c4957600080fd5b50610c646004803603610c5f919081019061682f565b612e80565b604051610c719190617708565b60405180910390f35b348015610c8657600080fd5b50610c8f612ea6565b604051610c9c9190617612565b60405180910390f35b348015610cb157600080fd5b50610ccc6004803603610cc7919081019061682f565b612ecc565b604051610cd99190617bc7565b60405180910390f35b348015610cee57600080fd5b50610cf7612ee4565b604051610d049190617bc7565b60405180910390f35b348015610d1957600080fd5b50610d346004803603610d2f9190810190616881565b612eea565b005b348015610d4257600080fd5b50610d5d6004803603610d58919081019061653f565b6130b9565b604051610d6a9190617bc7565b60405180910390f35b348015610d7f57600080fd5b50610d9a6004803603610d95919081019061653f565b6130de565b604051610da79190617708565b60405180910390f35b348015610dbc57600080fd5b50610dc5613205565b604051610dd29190617bc7565b60405180910390f35b348015610de757600080fd5b50610e026004803603610dfd91908101906168bd565b61320b565b604051610e0f9190617bc7565b60405180910390f35b348015610e2457600080fd5b50610e3f6004803603610e3a91908101906168f9565b61323d565b005b348015610e4d57600080fd5b50610e686004803603610e6391908101906168bd565b6132e1565b604051610e759190617bc7565b60405180910390f35b348015610e8a57600080fd5b50610ea56004803603610ea091908101906168bd565b613313565b604051610eb29190617bc7565b60405180910390f35b348015610ec757600080fd5b50610ee26004803603610edd919081019061682f565b613345565b005b348015610ef057600080fd5b50610f0b6004803603610f0691908101906164b3565b613398565b005b60696020528060005260406000206000915090505481565b6000339050610f326162f8565b610f3c82846133eb565b90506000610f5b826020015183600001516135b990919063ffffffff16565b9050610f7e8385610f798560400151856135b990919063ffffffff16565b61360e565b80607460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600086815260200190815260200160002060000160008282540192505081905550838373ffffffffffffffffffffffffffffffffffffffff167f4119153d17a36f9597d40e3ab4148d03261a439dddbec4e91799ab7159608e2684600001518560200151866040015160405161103693929190617c64565b60405180910390a350505050565b60003390506110516162f8565b61105b82846133eb565b905060008273ffffffffffffffffffffffffffffffffffffffff166110a78360400151611099856020015186600001516135b990919063ffffffff16565b6135b990919063ffffffff16565b6040516110b3906175fd565b60006040518083038185875af1925050503d80600081146110f0576040519150601f19603f3d011682016040523d82523d6000602084013e6110f5565b606091505b5050905080611139576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161113090617aa7565b60405180910390fd5b838373ffffffffffffffffffffffffffffffffffffffff167fc1d8eb6e444b89fb8ff0991c19311c070df704ccb009e210d1462d5b2410bf4584600001518560200151866040015160405161119093929190617c64565b60405180910390a350505050565b607c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006111d083836130de565b61122c57607360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008381526020019081526020016000205490506112e5565b6112e2607460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600084815260200190815260200160002060000154607360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008581526020019081526020016000205461370990919063ffffffff16565b90505b92915050565b60775481565b6112fa33613753565b611339576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161133090617807565b60405180910390fd5b61134689898960006137ad565b80606f60008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a8152602001908152602001600020600201819055506113a7876139f4565b6000861461151557868611156113f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113e990617ba7565b60405180910390fd5b6000607460008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a8152602001908152602001600020905086816000018190555085816001018190555084816002018190555083816003018190555082607560008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008b815260200190815260200160002060000181905550888a73ffffffffffffffffffffffffffffffffffffffff167f138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e78868a60405161150b929190617c3b565b60405180910390a3505b505050505050505050565b6000803390506000607460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008681526020019081526020016000209050600084116115bd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115b490617787565b60405180910390fd5b6115c782866130de565b611606576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115fd90617907565b60405180910390fd5b806000015484111561164d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161164490617a87565b60405180910390fd5b6116578286613aa1565b611696576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161168d906179c7565b60405180910390fd5b6116a08286613bb8565b5060006116b38387878560000154613f30565b90506363401ec5826003015401826002015410156116d057600090505b84826000016000828254039250508190555060008114611701576116f783878360016141cd565b6117008161447c565b5b858373ffffffffffffffffffffffffffffffffffffffff167fef6c0c14fe9aa51af36acd791464dec3badbde668b63189b47bfa4e25be9b2b9878460405161174a929190617c3b565b60405180910390a380935050505092915050565b61176733613753565b6117a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161179d90617807565b60405180910390fd5b60008114156117ea576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117e1906178a7565b60405180910390fd5b6117f48282614507565b6117ff826000612eea565b60006068600084815260200190815260200160002060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050611846818260006146a5565b505050565b607160205282600052604060002060205281600052604060002060205280600052604060002060009250925050508060000154908060010154908060020154905083565b6000608360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168286866040516024016118de929190617672565b6040516020818303038152906040527f4a7702bb000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505060405161196891906175e6565b60006040518083038160008787f1925050503d80600081146119a6576040519150601f19603f3d011682016040523d82523d6000602084013e6119ab565b606091505b5050905080806119b9575082155b6119f8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119ef90617b07565b60405180910390fd5b5050505050565b606d5481565b607860205280600052604060002060009150905080600701549080600801549080600901549080600a01549080600b01549080600c01549080600d0154905087565b611a533383833361486e565b5050565b611a5f612687565b611a9e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a95906179a7565b60405180910390fd5b611aa782612e80565b611ae6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611add90617767565b60405180910390fd5b611aee614e79565b811115611b30576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b27906179e7565b60405180910390fd5b80607b600084815260200190815260200160002081905550817f047575f43f09a7a093d94ec483064acfc61b7e25c0de28017da442abf99cb91782604051611b789190617bc7565b60405180910390a25050565b611b8d33613753565b611bcc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611bc390617807565b60405180910390fd5b611c20898989898080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508888888888614e89565b606b54881115611c325787606b819055505b505050505050505050565b6000607360006068600085815260200190815260200160002060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000838152602001908152602001600020549050919050565b600060786000848152602001908152602001600020600001600083815260200190815260200160002054905092915050565b606e5481565b6000611d0f6162f8565b611d198484615132565b9050611d4c8160000151611d3e836020015184604001516135b990919063ffffffff16565b6135b990919063ffffffff16565b91505092915050565b600060786000848152602001908152602001600020600101600083815260200190815260200160002054905092915050565b6000339050611d968184613bb8565b5060008211611dda576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611dd190617787565b60405180910390fd5b611de481846111c4565b821115611e26576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e1d90617a07565b60405180910390fd5b611e308184613aa1565b611e6f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e66906179c7565b60405180910390fd5b6000607260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600085815260200190815260200160002060008154809291906001019190505590506000607160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600086815260200190815260200160002060008381526020019081526020016000206002015414611f79576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f7090617987565b60405180910390fd5b611f8682858560016141cd565b82607160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000868152602001908152602001600020600083815260200190815260200160002060020181905550611ff7612496565b607160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008681526020019081526020016000206000838152602001908152602001600020600001819055506120676151d9565b607160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008681526020019081526020016000206000838152602001908152602001600020600101819055506120da846000612eea565b80848373ffffffffffffffffffffffffffffffffffffffff167fd3bb4e423fbea695d16b982f9f682dc5f35152e5411646a8a5a79a6b02ba8d57866040516121229190617bc7565b60405180910390a450505050565b600061213c83836130de565b61214957600090506121a0565b607460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008381526020019081526020016000206000015490505b92915050565b60006121b06162f8565b606f60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000206040518060600160405290816000820154815260200160018201548152602001600282015481525050905061225c816040015161224e836000015184602001516135b990919063ffffffff16565b6135b990919063ffffffff16565b91505092915050565b606080826040519080825280602002602001820160405280156122a257816020015b61228f616319565b8152602001906001900390816122875790505b50905060008090505b8381101561238157607160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000878152602001908152602001600020600061231883886135b990919063ffffffff16565b8152602001908152602001600020604051806060016040529081600082015481526020016001820154815260200160028201548152505082828151811061235b57fe5b602002602001018190525061237a6001826135b990919063ffffffff16565b90506122ab565b5080915050949350505050565b612396612687565b6123d5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016123cc906179a7565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000603360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000600160675401905090565b60675481565b6124b1612687565b6124f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016124e7906179a7565b60405180910390fd5b80608460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b606a6020528060005260406000206000915090508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156125dc5780601f106125b1576101008083540402835291602001916125dc565b820191906000526020600020905b8154815290600101906020018083116125bf57829003601f168201915b505050505081565b608360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b606c5481565b61261a8282613bb8565b612659576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161265090617827565b60405180910390fd5b5050565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b608460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461276f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161276690617887565b60405180910390fd5b6127798383613bb8565b50600081116127bd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016127b490617787565b60405180910390fd5b607c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166378542c88338585856040518563ffffffff1660e01b815260040161281e949392919061762d565b600060405180830381600087803b15801561283857600080fd5b505af115801561284c573d6000803e3d6000fd5b50505050607360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000838152602001908152602001600020548111156128e3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128da90617ac7565b60405180910390fd5b60006128ef84846111c4565b9050808210156129c4576000607460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008581526020019081526020016000209050612966828403826000015461370990919063ffffffff16565b8160000181905550838573ffffffffffffffffffffffffffffffffffffffff167fef6c0c14fe9aa51af36acd791464dec3badbde668b63189b47bfa4e25be9b2b984860360006040516129ba929190617c12565b60405180910390a3505b6129d184848460016141cd565b6129dc836000612eea565b64ffffffffff838573ffffffffffffffffffffffffffffffffffffffff167fd3bb4e423fbea695d16b982f9f682dc5f35152e5411646a8a5a79a6b02ba8d5785604051612a299190617bc7565b60405180910390a460003373ffffffffffffffffffffffffffffffffffffffff1683604051612a57906175fd565b60006040518083038185875af1925050503d8060008114612a94576040519150601f19603f3d011682016040523d82523d6000602084013e612a99565b606091505b5050905080612add576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612ad490617aa7565b60405180910390fd5b64ffffffffff848673ffffffffffffffffffffffffffffffffffffffff167f75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a2186604051612b2a9190617bc7565b60405180910390a45050505050565b6074602052816000526040600020602052806000526040600020600091509150508060000154908060010154908060020154908060030154905084565b612b8133823461360e565b50565b600060786000848152602001908152602001600020600501600083815260200190815260200160002054905092915050565b608260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c5f530af6040518163ffffffff1660e01b815260040160206040518083038186803b158015612c1e57600080fd5b505afa158015612c32573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250612c569190810190616858565b341015612c98576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c8f90617b67565b60405180910390fd5b60008282905011612cde576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612cd590617a27565b60405180910390fd5b612d2c3383838080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050506151e1565b612d3933606b543461360e565b5050565b6070602052816000526040600020602052806000526040600020600091509150505481565b60686020528060005260406000206000915090508060000154908060010154908060020154908060030154908060040154908060050154908060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905087565b6075602052816000526040600020602052806000526040600020600091509150508060000154908060010154908060020154905083565b606060786000838152602001908152602001600020600601805480602002602001604051908101604052809291908181526020018280548015612e5d57602002820191906000526020600020905b815481526020019060010190808311612e49575b50505050509050919050565b6000339050612e7a81858585615219565b50505050565b600080608060686000858152602001908152602001600020600001541614159050919050565b608060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b607b6020528060005260406000206000915090505481565b606b5481565b612ef3826156af565b612f32576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612f2990617b27565b60405180910390fd5b6000606860008481526020019081526020016000206003015490506000606860008581526020019081526020016000206000015414612f7057600090505b606660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a4066fbe84836040518363ffffffff1660e01b8152600401612fcd929190617c3b565b600060405180830381600087803b158015612fe757600080fd5b505af1158015612ffb573d6000803e3d6000fd5b5050505081801561300d575060008114155b156130b457606660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663242a6e3f84606a60008781526020019081526020016000206040518363ffffffff1660e01b8152600401613081929190617be2565b600060405180830381600087803b15801561309b57600080fd5b505af11580156130af573d6000803e3d6000fd5b505050505b505050565b6073602052816000526040600020602052806000526040600020600091509150505481565b600080607460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000848152602001908152602001600020600201541415801561319757506000607460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000206000015414155b80156131fd5750607460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000838152602001908152602001600020600201546131fa6151d9565b11155b905092915050565b607f5481565b600060786000848152602001908152602001600020600301600083815260200190815260200160002054905092915050565b600033905060008211613285576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161327c90617787565b60405180910390fd5b61328f81856130de565b156132cf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016132c6906177c7565b60405180910390fd5b6132db81858585615219565b50505050565b600060786000848152602001908152602001600020600201600083815260200190815260200160002054905092915050565b600060786000848152602001908152602001600020600401600083815260200190815260200160002054905092915050565b61334d612687565b61338c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613383906179a7565b60405180910390fd5b6133958161447c565b50565b6133a0612687565b6133df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016133d6906179a7565b60405180910390fd5b6133e8816156d2565b50565b6133f36162f8565b6133fd8383613aa1565b61343c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613433906179c7565b60405180910390fd5b6134468383613bb8565b50606f60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008381526020019081526020016000206040518060600160405290816000820154815260200160018201548152602001600282015481525050905060006134f582600001516134e7846020015185604001516135b990919063ffffffff16565b6135b990919063ffffffff16565b9050600081141561353b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161353290617947565b60405180910390fd5b606f60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000206000808201600090556001820160009055600282016000905550506135af816139f4565b8191505092915050565b600080828401905083811015613604576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016135fb906177e7565b60405180910390fd5b8091505092915050565b613617826156af565b613656576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161364d90617b27565b60405180910390fd5b60006068600084815260200190815260200160002060000154146136af576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016136a690617867565b60405180910390fd5b6136bc83838360016137ad565b6136c582615802565b613704576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016136fb90617ae7565b60405180910390fd5b505050565b600061374b83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506158f7565b905092915050565b6000606660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16149050919050565b600082116137f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016137e790617787565b60405180910390fd5b6137fa8484613bb8565b5061385e82607360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000868152602001908152602001600020546135b990919063ffffffff16565b607360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000858152602001908152602001600020819055506000606860008581526020019081526020016000206003015490506138e083826135b990919063ffffffff16565b606860008681526020019081526020016000206003018190555061390f83606c546135b990919063ffffffff16565b606c819055506000606860008681526020019081526020016000206000015414156139505761394983606d546135b990919063ffffffff16565b606d819055505b61395d8460008314612eea565b838573ffffffffffffffffffffffffffffffffffffffff167f9a8f44850296624dadfd9c246d17e47171d35727a181bd090aa14bbbe00238bb856040516139a49190617bc7565b60405180910390a36139ed856068600087815260200190815260200160002060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846146a5565b5050505050565b606660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166366e7ea0f30836040518363ffffffff1660e01b8152600401613a5192919061769b565b600060405180830381600087803b158015613a6b57600080fd5b505af1158015613a7f573d6000803e3d6000fd5b50505050613a98816077546135b990919063ffffffff16565b60778190555050565b60008073ffffffffffffffffffffffffffffffffffffffff16607c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415613b025760019050613bb2565b607c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166321d585c384846040518363ffffffff1660e01b8152600401613b5f92919061769b565b60206040518083038186803b158015613b7757600080fd5b505afa158015613b8b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250613baf91908101906167c1565b90505b92915050565b6000613bc26162f8565b613bcc8484615952565b9050613bd783615b35565b607060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600085815260200190815260200160002081905550613cad606f60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000858152602001908152602001600020604051806060016040529081600082015481526020016001820154815260200160028201548152505082615ba7565b606f60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000858152602001908152602001600020600082015181600001556020820151816001015560408201518160020155905050613da0607560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000858152602001908152602001600020604051806060016040529081600082015481526020016001820154815260200160028201548152505082615ba7565b607560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000858152602001908152602001600020600082015181600001556020820151816001015560408201518160020155905050613e1b84846130de565b613efe57607460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000206000808201600090556001820160009055600282016000905560038201600090555050607560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000206000808201600090556001820160009055600282016000905550505b60008160200151141580613f1757506000816000015114155b80613f2757506000816040015114155b91505092915050565b600080613fab83613f9d86607560008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a815260200190815260200160002060000154615c2090919063ffffffff16565b615c9090919063ffffffff16565b905060006140278461401987607560008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008b815260200190815260200160002060010154615c2090919063ffffffff16565b615c9090919063ffffffff16565b905060006002828161403557fe5b04830190506140a083607560008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a81526020019081526020016000206000015461370990919063ffffffff16565b607560008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008981526020019081526020016000206000018190555061415d82607560008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a81526020019081526020016000206001015461370990919063ffffffff16565b607560008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000898152602001908152602001600020600101819055508581106141bf578590505b809350505050949350505050565b81607360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008581526020019081526020016000206000828254039250508190555061425482606860008681526020019081526020016000206003015461370990919063ffffffff16565b606860008581526020019081526020016000206003018190555061428382606c5461370990919063ffffffff16565b606c819055506000606860008581526020019081526020016000206000015414156142c4576142bd82606d5461370990919063ffffffff16565b606d819055505b60006142cf84611c3d565b9050600081146144285760006068600086815260200190815260200160002060000154141561442357608260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c5f530af6040518163ffffffff1660e01b815260040160206040518083038186803b15801561436057600080fd5b505afa158015614374573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506143989190810190616858565b8110156143da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016143d190617b67565b60405180910390fd5b6143e384615802565b614422576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161441990617ae7565b60405180910390fd5b5b614434565b614433846001614507565b5b614475856068600087815260200190815260200160002060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846146a5565b5050505050565b6000811461450457600073ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156144cb573d6000803e3d6000fd5b507fe733dd86f3139a17476f9309db31e45d999e6469068d5734f7d51b23110c17bc816040516144fb9190617bc7565b60405180910390a15b50565b6000606860008481526020019081526020016000206000015414801561452e575060008114155b156145655761455e6068600084815260200190815260200160002060030154606d5461370990919063ffffffff16565b606d819055505b60686000838152602001908152602001600020600001548111156146a157806068600084815260200190815260200160002060000181905550600060686000848152602001908152602001600020600201541415614668576145c5612496565b60686000848152602001908152602001600020600201819055506145e76151d9565b6068600084815260200190815260200160002060010181905550817fac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e476068600085815260200190815260200160002060020154606860008681526020019081526020016000206001015460405161465f929190617c3b565b60405180910390a25b817fcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e826040516146989190617bc7565b60405180910390a25b5050565b600073ffffffffffffffffffffffffffffffffffffffff16608360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614614869576000608360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16627a1200858560405160240161474d929190617672565b6040516020818303038152906040527f4a7702bb000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516147d791906175e6565b60006040518083038160008787f1925050503d8060008114614815576040519150601f19603f3d011682016040523d82523d6000602084013e61481a565b606091505b505090508080614828575081155b614867576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161485e90617b07565b60405180910390fd5b505b505050565b61487661633a565b607160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600085815260200190815260200160002060008481526020019081526020016000206040518060600160405290816000820154815260200160018201548152602001600282015481525050905060008160000151141561494a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161494190617a47565b60405180910390fd5b6149548585613aa1565b614993576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161498a906179c7565b60405180910390fd5b60008160200151905060008260000151905060006068600088815260200190815260200160002060010154141580156149e15750816068600088815260200190815260200160002060010154105b15614a195760686000878152602001908152602001600020600101549150606860008781526020019081526020016000206002015490505b608260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b82b84276040518163ffffffff1660e01b815260040160206040518083038186803b158015614a8157600080fd5b505afa158015614a95573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250614ab99190810190616858565b8201614ac36151d9565b1015614b04576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614afb906178c7565b60405180910390fd5b608260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663650acd666040518163ffffffff1660e01b815260040160206040518083038186803b158015614b6c57600080fd5b505afa158015614b80573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250614ba49190810190616858565b8101614bae612496565b1015614bef576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614be690617b47565b60405180910390fd5b6000607160008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600088815260200190815260200160002060008781526020019081526020016000206002015490506000614c6388612e80565b90506000614c858383607b60008d815260200190815260200160002054615cda565b9050607160008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a8152602001908152602001600020600089815260200190815260200160002060008082016000905560018201600090556002820160009055505080606e60008282540192505081905550808311614d55576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614d4c90617b87565b60405180910390fd5b60008773ffffffffffffffffffffffffffffffffffffffff16614d81838661370990919063ffffffff16565b604051614d8d906175fd565b60006040518083038185875af1925050503d8060008114614dca576040519150601f19603f3d011682016040523d82523d6000602084013e614dcf565b606091505b5050905080614e13576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614e0a90617aa7565b60405180910390fd5b614e1c8261447c565b888a8c73ffffffffffffffffffffffffffffffffffffffff167f75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a2187604051614e649190617bc7565b60405180910390a45050505050505050505050565b6000670de0b6b3a7640000905090565b6000606960008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414614f0b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614f02906178e7565b60405180910390fd5b86606960008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550846068600089815260200190815260200160002060000181905550836068600089815260200190815260200160002060040181905550826068600089815260200190815260200160002060050181905550806068600089815260200190815260200160002060010181905550816068600089815260200190815260200160002060020181905550876068600089815260200190815260200160002060060160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555085606a6000898152602001908152602001600020908051906020019061505292919061635b565b508773ffffffffffffffffffffffffffffffffffffffff16877f49bca1ed2666922f9f1690c26a569e1299c2a715fe57647d77e81adfabbf25bf868660405161509c929190617c3b565b60405180910390a3600082146150e757867fac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e4783836040516150de929190617c3b565b60405180910390a25b6000851461512857867fcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e8660405161511f9190617bc7565b60405180910390a25b5050505050505050565b61513a6162f8565b6151426162f8565b61514c8484615952565b90506151d0606f60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000858152602001908152602001600020604051806060016040529081600082015481526020016001820154815260200160028201548152505082615ba7565b91505092915050565b600042905090565b6000606b6000815460010191905081905590506152148382846000615204612496565b61520c6151d9565b600080614e89565b505050565b61522384846111c4565b811115615265576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161525c90617ac7565b60405180910390fd5b60006068600085815260200190815260200160002060000154146152be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016152b590617867565b60405180910390fd5b608260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630d7b26096040518163ffffffff1660e01b815260040160206040518083038186803b15801561532657600080fd5b505afa15801561533a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061535e9190810190616858565b821015801561540c5750608260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630d4955e36040518163ffffffff1660e01b815260040160206040518083038186803b1580156153d057600080fd5b505afa1580156153e4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506154089190810190616858565b8211155b61544b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161544290617847565b60405180910390fd5b6000615467836154596151d9565b6135b990919063ffffffff16565b905060006068600086815260200190815260200160002060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff161461556e5781607460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600087815260200190815260200160002060020154101561556d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161556490617a67565b60405180910390fd5b5b6155788686613bb8565b506000607460008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600087815260200190815260200160002090508060030154851015615614576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161560b90617967565b60405180910390fd5b61562b8482600001546135b990919063ffffffff16565b816000018190555061563b612496565b8160010181905550828160020181905550848160030181905550858773ffffffffffffffffffffffffffffffffffffffff167f138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e78878760405161569e929190617c3b565b60405180910390a350505050505050565b600080606860008481526020019081526020016000206005015414159050919050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415615742576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615739906177a7565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380603360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60006158d761580f614e79565b6158c9608260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632265f2846040518163ffffffff1660e01b815260040160206040518083038186803b15801561587a57600080fd5b505afa15801561588e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506158b29190810190616858565b6158bb86611c3d565b615c2090919063ffffffff16565b615c9090919063ffffffff16565b606860008481526020019081526020016000206003015411159050919050565b600083831115829061593f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016159369190617745565b60405180910390fd5b5060008385039050809150509392505050565b61595a6162f8565b6000607060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600084815260200190815260200160002054905060006159ba84615b35565b905060006159c88686615d63565b9050818111156159d6578190505b828110156159e2578290505b6000607460008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600087815260200190815260200160002090506000607360008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008881526020019081526020016000205490506000615aa483600001548361370990919063ffffffff16565b90506000615ab884600001548a8988615e6c565b9050615ac26162f8565b615ad0828660030154615f21565b9050615ade838b8a89615e6c565b9150615ae86162f8565b615af3836000615f21565b9050615b01858c898b615e6c565b9250615b0b6162f8565b615b16846000615f21565b9050615b2383838361618a565b9a505050505050505050505092915050565b600080606860008481526020019081526020016000206002015414615b9c5760686000838152602001908152602001600020600201546067541015615b7e576067549050615ba2565b60686000838152602001908152602001600020600201549050615ba2565b60675490505b919050565b615baf6162f8565b6040518060600160405280615bd5846000015186600001516135b990919063ffffffff16565b8152602001615bf5846020015186602001516135b990919063ffffffff16565b8152602001615c15846040015186604001516135b990919063ffffffff16565b815250905092915050565b600080831415615c335760009050615c8a565b6000828402905082848281615c4457fe5b0414615c85576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615c7c90617927565b60405180910390fd5b809150505b92915050565b6000615cd283836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f0000000000008152506161ae565b905092915050565b6000821580615cf05750615cec614e79565b8210155b15615cfe5760009050615d5c565b615d466001615d38615d0e614e79565b615d2a86615d1a614e79565b0389615c2090919063ffffffff16565b615c9090919063ffffffff16565b6135b990919063ffffffff16565b905083811115615d5857839050615d5c565b8090505b9392505050565b600080607460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600084815260200190815260200160002060010154905060006067549050615dce85858361620f565b15615ddd578092505050615e66565b615de885858461620f565b615df757600092505050615e66565b80821115615e0a57600092505050615e66565b5b80821015615e49576000600282840181615e2157fe5b049050615e2f86868361620f565b15615e3f57600181019250615e43565b8091505b50615e0b565b6000811415615e5d57600092505050615e66565b60018103925050505b92915050565b6000818310615e7e5760009050615f19565b60006078600085815260200190815260200160002060010160008681526020019081526020016000205490506000607860008581526020019081526020016000206001016000878152602001908152602001600020549050615f14615ee1614e79565b615f0689615ef8868661370990919063ffffffff16565b615c2090919063ffffffff16565b615c9090919063ffffffff16565b925050505b949350505050565b615f296162f8565b60405180606001604052806000815260200160008152602001600081525090506000608260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635e2308d26040518163ffffffff1660e01b815260040160206040518083038186803b158015615fb357600080fd5b505afa158015615fc7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250615feb9190810190616858565b90506000831461614a57600081616000614e79565b03905060006160c9608260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630d4955e36040518163ffffffff1660e01b815260040160206040518083038186803b15801561607057600080fd5b505afa158015616084573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506160a89190810190616858565b6160bb8785615c2090919063ffffffff16565b615c9090919063ffffffff16565b905060006160fb6160d8614e79565b6160ed8487018a615c2090919063ffffffff16565b615c9090919063ffffffff16565b9050616129616108614e79565b61611b868a615c2090919063ffffffff16565b615c9090919063ffffffff16565b85602001818152505084602001518103856000018181525050505050616180565b616176616155614e79565b6161688387615c2090919063ffffffff16565b615c9090919063ffffffff16565b8260400181815250505b8191505092915050565b6161926162f8565b6161a561619f8585615ba7565b83615ba7565b90509392505050565b600080831182906161f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016161ec9190617745565b60405180910390fd5b50600083858161620157fe5b049050809150509392505050565b600081607460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600085815260200190815260200160002060010154111580156162cf5750607460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000848152602001908152602001600020600201546162cc836162d8565b11155b90509392505050565b600060786000838152602001908152602001600020600701549050919050565b60405180606001604052806000815260200160008152602001600081525090565b60405180606001604052806000815260200160008152602001600081525090565b60405180606001604052806000815260200160008152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061639c57805160ff19168380011785556163ca565b828001600101855582156163ca579182015b828111156163c95782518255916020019190600101906163ae565b5b5090506163d791906163db565b5090565b6163fd91905b808211156163f95760008160009055506001016163e1565b5090565b90565b60008135905061640f81617f67565b92915050565b60008135905061642481617f7e565b92915050565b60008151905061643981617f7e565b92915050565b60008083601f84011261645157600080fd5b8235905067ffffffffffffffff81111561646a57600080fd5b60208301915083600182028301111561648257600080fd5b9250929050565b60008135905061649881617f95565b92915050565b6000815190506164ad81617f95565b92915050565b6000602082840312156164c557600080fd5b60006164d384828501616400565b91505092915050565b600080600080608085870312156164f257600080fd5b600061650087828801616400565b945050602061651187828801616400565b935050604061652287828801616415565b925050606061653387828801616489565b91505092959194509250565b6000806040838503121561655257600080fd5b600061656085828601616400565b925050602061657185828601616489565b9150509250929050565b60008060008060008060008060006101008a8c03121561659a57600080fd5b60006165a88c828d01616400565b99505060206165b98c828d01616489565b98505060408a013567ffffffffffffffff8111156165d657600080fd5b6165e28c828d0161643f565b975097505060606165f58c828d01616489565b95505060806166068c828d01616489565b94505060a06166178c828d01616489565b93505060c06166288c828d01616489565b92505060e06166398c828d01616489565b9150509295985092959850929598565b60008060006060848603121561665e57600080fd5b600061666c86828701616400565b935050602061667d86828701616489565b925050604061668e86828701616489565b9150509250925092565b600080600080608085870312156166ae57600080fd5b60006166bc87828801616400565b94505060206166cd87828801616489565b93505060406166de87828801616489565b92505060606166ef87828801616489565b91505092959194509250565b60008060008060008060008060006101208a8c03121561671a57600080fd5b60006167288c828d01616400565b99505060206167398c828d01616489565b985050604061674a8c828d01616489565b975050606061675b8c828d01616489565b965050608061676c8c828d01616489565b95505060a061677d8c828d01616489565b94505060c061678e8c828d01616489565b93505060e061679f8c828d01616489565b9250506101006167b18c828d01616489565b9150509295985092959850929598565b6000602082840312156167d357600080fd5b60006167e18482850161642a565b91505092915050565b600080602083850312156167fd57600080fd5b600083013567ffffffffffffffff81111561681757600080fd5b6168238582860161643f565b92509250509250929050565b60006020828403121561684157600080fd5b600061684f84828501616489565b91505092915050565b60006020828403121561686a57600080fd5b60006168788482850161649e565b91505092915050565b6000806040838503121561689457600080fd5b60006168a285828601616489565b92505060206168b385828601616415565b9150509250929050565b600080604083850312156168d057600080fd5b60006168de85828601616489565b92505060206168ef85828601616489565b9150509250929050565b60008060006060848603121561690e57600080fd5b600061691c86828701616489565b935050602061692d86828701616489565b925050604061693e86828701616489565b9150509250925092565b60006169548383617586565b60608301905092915050565b600061696c83836175c8565b60208301905092915050565b61698181617edb565b82525050565b61699081617e93565b82525050565b60006169a182617df3565b6169ab8185617e44565b93506169b683617dbe565b8060005b838110156169e75781516169ce8882616948565b97506169d983617e2a565b9250506001810190506169ba565b5085935050505092915050565b60006169ff82617dfe565b616a098185617e55565b9350616a1483617dce565b8060005b83811015616a45578151616a2c8882616960565b9750616a3783617e37565b925050600181019050616a18565b5085935050505092915050565b616a5b81617ea5565b82525050565b6000616a6c82617e14565b616a768185617e77565b9350616a86818560208601617f23565b80840191505092915050565b6000616a9d82617e09565b616aa78185617e66565b9350616ab7818560208601617f23565b616ac081617f56565b840191505092915050565b600081546001811660008114616ae85760018114616b0e57616b52565b607f6002830416616af98187617e66565b955060ff198316865260208601935050616b52565b60028204616b1c8187617e66565b9550616b2785617dde565b60005b82811015616b4957815481890152600182019150602081019050616b2a565b80880195505050505b505092915050565b616b6381617eed565b82525050565b6000616b7482617e1f565b616b7e8185617e82565b9350616b8e818560208601617f23565b616b9781617f56565b840191505092915050565b6000616baf601783617e82565b91507f76616c696461746f722069736e277420736c61736865640000000000000000006000830152602082019050919050565b6000616bef600b83617e82565b91507f7a65726f20616d6f756e740000000000000000000000000000000000000000006000830152602082019050919050565b6000616c2f602683617e82565b91507f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008301527f64647265737300000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000616c95601183617e82565b91507f616c7265616479206c6f636b65642075700000000000000000000000000000006000830152602082019050919050565b6000616cd5601b83617e82565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b6000616d15602983617e82565b91507f63616c6c6572206973206e6f7420746865204e6f64654472697665724175746860008301527f20636f6e747261637400000000000000000000000000000000000000000000006020830152604082019050919050565b6000616d7b601083617e82565b91507f6e6f7468696e6720746f207374617368000000000000000000000000000000006000830152602082019050919050565b6000616dbb601283617e82565b91507f696e636f7272656374206475726174696f6e00000000000000000000000000006000830152602082019050919050565b6000616dfb601683617e82565b91507f76616c696461746f722069736e277420616374697665000000000000000000006000830152602082019050919050565b6000616e3b601283617e82565b91507f6e6f7420735257412066696e616c697a657200000000000000000000000000006000830152602082019050919050565b6000616e7b600c83617e82565b91507f77726f6e672073746174757300000000000000000000000000000000000000006000830152602082019050919050565b6000616ebb601683617e82565b91507f6e6f7420656e6f7567682074696d6520706173736564000000000000000000006000830152602082019050919050565b6000616efb601883617e82565b91507f76616c696461746f7220616c72656164792065786973747300000000000000006000830152602082019050919050565b6000616f3b600d83617e82565b91507f6e6f74206c6f636b6564207570000000000000000000000000000000000000006000830152602082019050919050565b6000616f7b602183617e82565b91507f536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f60008301527f77000000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000616fe1600c83617e82565b91507f7a65726f207265776172647300000000000000000000000000000000000000006000830152602082019050919050565b6000617021601f83617e82565b91507f6c6f636b7570206475726174696f6e2063616e6e6f74206465637265617365006000830152602082019050919050565b6000617061601383617e82565b91507f7772494420616c726561647920657869737473000000000000000000000000006000830152602082019050919050565b60006170a1602083617e82565b91507f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726000830152602082019050919050565b60006170e1601883617e82565b91507f6f75747374616e64696e6720735257412062616c616e636500000000000000006000830152602082019050919050565b6000617121602183617e82565b91507f6d757374206265206c657373207468616e206f7220657175616c20746f20312e60008301527f30000000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000617187601983617e82565b91507f6e6f7420656e6f75676820756e6c6f636b6564207374616b65000000000000006000830152602082019050919050565b60006171c7600c83617e82565b91507f656d707479207075626b657900000000000000000000000000000000000000006000830152602082019050919050565b6000617207601583617e82565b91507f7265717565737420646f65736e277420657869737400000000000000000000006000830152602082019050919050565b6000617247602883617e82565b91507f76616c696461746f72206c6f636b757020706572696f642077696c6c20656e6460008301527f206561726c6965720000000000000000000000000000000000000000000000006020830152604082019050919050565b60006172ad601783617e82565b91507f6e6f7420656e6f756768206c6f636b6564207374616b650000000000000000006000830152602082019050919050565b60006172ed600083617e77565b9150600082019050919050565b6000617307601283617e82565b91507f4661696c656420746f2073656e642052574100000000000000000000000000006000830152602082019050919050565b6000617347601083617e82565b91507f6e6f7420656e6f756768207374616b65000000000000000000000000000000006000830152602082019050919050565b6000617387602983617e82565b91507f76616c696461746f7227732064656c65676174696f6e73206c696d697420697360008301527f20657863656564656400000000000000000000000000000000000000000000006020830152604082019050919050565b60006173ed601b83617e82565b91507f676f7620766f746573207265636f756e74696e67206661696c656400000000006000830152602082019050919050565b600061742d601783617e82565b91507f76616c696461746f7220646f65736e27742065786973740000000000000000006000830152602082019050919050565b600061746d601883617e82565b91507f6e6f7420656e6f7567682065706f6368732070617373656400000000000000006000830152602082019050919050565b60006174ad601783617e82565b91507f696e73756666696369656e742073656c662d7374616b650000000000000000006000830152602082019050919050565b60006174ed601683617e82565b91507f7374616b652069732066756c6c7920736c6173686564000000000000000000006000830152602082019050919050565b600061752d602c83617e82565b91507f6c6f636b6564207374616b652069732067726561746572207468616e2074686560008301527f2077686f6c65207374616b6500000000000000000000000000000000000000006020830152604082019050919050565b60608201600082015161759c60008501826175c8565b5060208201516175af60208501826175c8565b5060408201516175c260408501826175c8565b50505050565b6175d181617ed1565b82525050565b6175e081617ed1565b82525050565b60006175f28284616a61565b915081905092915050565b6000617608826172e0565b9150819050919050565b60006020820190506176276000830184616987565b92915050565b60006080820190506176426000830187616978565b61764f6020830186616987565b61765c60408301856175d7565b61766960608301846175d7565b95945050505050565b60006040820190506176876000830185616987565b6176946020830184616987565b9392505050565b60006040820190506176b06000830185616987565b6176bd60208301846175d7565b9392505050565b600060208201905081810360008301526176de8184616996565b905092915050565b6000602082019050818103600083015261770081846169f4565b905092915050565b600060208201905061771d6000830184616a52565b92915050565b6000602082019050818103600083015261773d8184616a92565b905092915050565b6000602082019050818103600083015261775f8184616b69565b905092915050565b6000602082019050818103600083015261778081616ba2565b9050919050565b600060208201905081810360008301526177a081616be2565b9050919050565b600060208201905081810360008301526177c081616c22565b9050919050565b600060208201905081810360008301526177e081616c88565b9050919050565b6000602082019050818103600083015261780081616cc8565b9050919050565b6000602082019050818103600083015261782081616d08565b9050919050565b6000602082019050818103600083015261784081616d6e565b9050919050565b6000602082019050818103600083015261786081616dae565b9050919050565b6000602082019050818103600083015261788081616dee565b9050919050565b600060208201905081810360008301526178a081616e2e565b9050919050565b600060208201905081810360008301526178c081616e6e565b9050919050565b600060208201905081810360008301526178e081616eae565b9050919050565b6000602082019050818103600083015261790081616eee565b9050919050565b6000602082019050818103600083015261792081616f2e565b9050919050565b6000602082019050818103600083015261794081616f6e565b9050919050565b6000602082019050818103600083015261796081616fd4565b9050919050565b6000602082019050818103600083015261798081617014565b9050919050565b600060208201905081810360008301526179a081617054565b9050919050565b600060208201905081810360008301526179c081617094565b9050919050565b600060208201905081810360008301526179e0816170d4565b9050919050565b60006020820190508181036000830152617a0081617114565b9050919050565b60006020820190508181036000830152617a208161717a565b9050919050565b60006020820190508181036000830152617a40816171ba565b9050919050565b60006020820190508181036000830152617a60816171fa565b9050919050565b60006020820190508181036000830152617a808161723a565b9050919050565b60006020820190508181036000830152617aa0816172a0565b9050919050565b60006020820190508181036000830152617ac0816172fa565b9050919050565b60006020820190508181036000830152617ae08161733a565b9050919050565b60006020820190508181036000830152617b008161737a565b9050919050565b60006020820190508181036000830152617b20816173e0565b9050919050565b60006020820190508181036000830152617b4081617420565b9050919050565b60006020820190508181036000830152617b6081617460565b9050919050565b60006020820190508181036000830152617b80816174a0565b9050919050565b60006020820190508181036000830152617ba0816174e0565b9050919050565b60006020820190508181036000830152617bc081617520565b9050919050565b6000602082019050617bdc60008301846175d7565b92915050565b6000604082019050617bf760008301856175d7565b8181036020830152617c098184616acb565b90509392505050565b6000604082019050617c2760008301856175d7565b617c346020830184616b5a565b9392505050565b6000604082019050617c5060008301856175d7565b617c5d60208301846175d7565b9392505050565b6000606082019050617c7960008301866175d7565b617c8660208301856175d7565b617c9360408301846175d7565b949350505050565b6000608082019050617cb060008301876175d7565b617cbd60208301866175d7565b617cca60408301856175d7565b617cd760608301846175d7565b95945050505050565b600060e082019050617cf5600083018a6175d7565b617d0260208301896175d7565b617d0f60408301886175d7565b617d1c60608301876175d7565b617d2960808301866175d7565b617d3660a08301856175d7565b617d4360c0830184616987565b98975050505050505050565b600060e082019050617d64600083018a6175d7565b617d7160208301896175d7565b617d7e60408301886175d7565b617d8b60608301876175d7565b617d9860808301866175d7565b617da560a08301856175d7565b617db260c08301846175d7565b98975050505050505050565b6000819050602082019050919050565b6000819050602082019050919050565b60008190508160005260206000209050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000617e9e82617eb1565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000617ee682617eff565b9050919050565b6000617ef882617ed1565b9050919050565b6000617f0a82617f11565b9050919050565b6000617f1c82617eb1565b9050919050565b60005b83811015617f41578082015181840152602081019050617f26565b83811115617f50576000848401525b50505050565b6000601f19601f8301169050919050565b617f7081617e93565b8114617f7b57600080fd5b50565b617f8781617ea5565b8114617f9257600080fd5b50565b617f9e81617ed1565b8114617fa957600080fd5b5056fea365627a7a72315820fd8fa742408c2b87482fc27b7e1d88ed590d49789fa62681d2f387ab02ad9ff96c6578706572696d656e74616cf564736f6c63430005100040"
