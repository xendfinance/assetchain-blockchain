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
)

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BurntVITRA\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"name\":\"ChangedValidatorStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockupExtraReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockupBaseReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockedReward\",\"type\":\"uint256\"}],\"name\":\"ClaimedRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"auth\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"}],\"name\":\"CreatedValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"}],\"name\":\"DeactivatedValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Delegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LockedUpStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RefundedSlashedLegacyDelegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockupExtraReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockupBaseReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockedReward\",\"type\":\"uint256\"}],\"name\":\"RestakedRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Undelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"}],\"name\":\"UnlockedStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refundRatio\",\"type\":\"uint256\"}],\"name\":\"UpdatedSlashingRefundRatio\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"syncPubkey\",\"type\":\"bool\"}],\"name\":\"_syncValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentSealedEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getEpochSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBaseRewardWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalTxRewardWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRewardPerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"getLockedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getLockupInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lockedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getStashedLockupRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lockupExtraReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupBaseReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockedReward\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getValidator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"receivedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"auth\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getValidatorID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getValidatorPubkey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getWithdrawalRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"isLockedUp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastValidatorID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minGasPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"slashingRefundRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeTokenizerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stashedRewardsUntilEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalActiveStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSlashedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasuryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"voteBookAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getEpochValidatorIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochReceivedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochAccumulatedRewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochAccumulatedUptime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochAccumulatedOriginatedTxsFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochOfflineTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochOfflineBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"rewardsStash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"auth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"}],\"name\":\"setGenesisValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupFromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"earlyUnlockPenalty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"setGenesisDelegation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"createValidator\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getSelfStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validatorAuth\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"}],\"name\":\"recountVotes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"liquidateSVITRA\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"}],\"name\":\"updateSVITRAFinalizer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"isSlashed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"name\":\"deactivateValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"pendingRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"stashRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"claimRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"restakeRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnVITRA\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"getUnlockedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"lockStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"relockStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unlockStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundRatio\",\"type\":\"uint256\"}],\"name\":\"updateSlashingRefundRatio\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getWrRequests\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structSFCState.WithdrawalRequest[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061559280620000216000396000f3fe60806040526004361061038c5760003560e01c8063854873e1116101dc578063bd14d90711610102578063cfdbb7cd116100a0578063df00c9221161006f578063df00c92214610a5b578063e261641a14610a7b578063f2fde38b14610a9b578063f6a3250314610abb5761038c565b8063cfdbb7cd146109e6578063d96ed50514610a06578063dc31e1af14610a1b578063de67f21514610a3b5761038c565b8063c65ee0e1116100dc578063c65ee0e114610971578063c7be95de14610991578063cc8343aa146109a6578063cfd47663146109c65761038c565b8063bd14d9071461091c578063c3de580e1461093c578063c5f956af1461095c5761038c565b80639fa6dd351161017a578063b334566b11610149578063b334566b1461087c578063b5d896271461089c578063b810e411146108cf578063b88a37e2146108ef5761038c565b80639fa6dd3514610816578063a198d22914610829578063a5a470ad14610849578063a86a056f1461085c5761038c565b80638cddb015116101b65780638cddb0151461078f5780638da5cb5b146107af5780638f32d59b146107c457806396c7ee46146107e65761038c565b8063854873e114610738578063893675c6146107655780638b0e9f3f1461077a5761038c565b80634f7c4efb116102c1578063634b91e31161025f578063715018a61161022e578063715018a6146106d957806376671808146106ee5780637aa353bd146107035780637cacb1d6146107235761038c565b8063634b91e31461064c578063670322f81461066c5780636f4986631461068c578063702797e3146106ac5761038c565b806358f95b801161029b57806358f95b80146105d75780635fab23a8146105f75780636099ecb21461060c57806361e53fcc1461062c5761038c565b80634f7c4efb146105775780634feb92f3146105975780635601fe01146105b75761038c565b80631d3ac42c1161032e57806320c0849d1161030857806320c0849d146104ef57806328f731481461050f57806339b80c0014610524578063441a3e70146105575761038c565b80631d3ac42c146104805780631e702f83146104a05780631f270152146104c05761038c565b80630e559d821161036a5780630e559d821461040957806312622d0e1461042b57806318160ddd1461044b57806318f628d4146104605761038c565b80630135b1db1461039157806308c36874146103c75780630962ef79146103e9575b600080fd5b34801561039d57600080fd5b506103b16103ac366004614115565b610adb565b6040516103be919061531b565b60405180910390f35b3480156103d357600080fd5b506103e76103e236600461443f565b610aed565b005b3480156103f557600080fd5b506103e761040436600461443f565b610bb6565b34801561041557600080fd5b5061041e610cc5565b6040516103be9190615028565b34801561043757600080fd5b506103b1610446366004614194565b610cd4565b34801561045757600080fd5b506103b1610d5d565b34801561046c57600080fd5b506103e761047b36600461432b565b610d63565b34801561048c57600080fd5b506103b161049b3660046144ab565b610e89565b3480156104ac57600080fd5b506103e76104bb3660046144ab565b610fd9565b3480156104cc57600080fd5b506104e06104db36600461429a565b61105e565b6040516103be93929190615372565b3480156104fb57600080fd5b506103e761050a366004614133565b611090565b34801561051b57600080fd5b506103b16111aa565b34801561053057600080fd5b5061054461053f36600461443f565b6111b0565b6040516103be979695949392919061541d565b34801561056357600080fd5b506103e76105723660046144ab565b6111f2565b34801561058357600080fd5b506103e76105923660046144ab565b611202565b3480156105a357600080fd5b506103e76105b23660046141ce565b6112c1565b3480156105c357600080fd5b506103b16105d236600461443f565b611348565b3480156105e357600080fd5b506103b16105f23660046144ab565b61137e565b34801561060357600080fd5b506103b161139b565b34801561061857600080fd5b506103b1610627366004614194565b6113a1565b34801561063857600080fd5b506103b16106473660046144ab565b6113df565b34801561065857600080fd5b506103e76106673660046144ab565b611400565b34801561067857600080fd5b506103b1610687366004614194565b6115e2565b34801561069857600080fd5b506103b16106a7366004614194565b611623565b3480156106b857600080fd5b506106cc6106c73660046142e7565b61168d565b6040516103be91906150aa565b3480156106e557600080fd5b506103e761177c565b3480156106fa57600080fd5b506103b1611802565b34801561070f57600080fd5b506103e761071e36600461443f565b61180c565b34801561072f57600080fd5b506103b161183c565b34801561074457600080fd5b5061075861075336600461443f565b611842565b6040516103be91906150da565b34801561077157600080fd5b5061041e6118fb565b34801561078657600080fd5b506103b161190a565b34801561079b57600080fd5b506103e76107aa366004614194565b611910565b3480156107bb57600080fd5b5061041e611936565b3480156107d057600080fd5b506107d9611945565b6040516103be91906150cc565b3480156107f257600080fd5b50610806610801366004614194565b611956565b6040516103be949392919061539a565b6103e761082436600461443f565b611988565b34801561083557600080fd5b506103b16108443660046144ab565b611993565b6103e76108573660046143fd565b6119b4565b34801561086857600080fd5b506103b1610877366004614194565b611ac3565b34801561088857600080fd5b506103e7610897366004614115565b611ae0565b3480156108a857600080fd5b506108bc6108b736600461443f565b611b3e565b6040516103be97969594939291906153b5565b3480156108db57600080fd5b506104e06108ea366004614194565b611b84565b3480156108fb57600080fd5b5061090f61090a36600461443f565b611bb0565b6040516103be91906150bb565b34801561092857600080fd5b506103e76109373660046144ca565b611c15565b34801561094857600080fd5b506107d961095736600461443f565b611c28565b34801561096857600080fd5b5061041e611c3f565b34801561097d57600080fd5b506103b161098c36600461443f565b611c4e565b34801561099d57600080fd5b506103b1611c60565b3480156109b257600080fd5b506103e76109c136600461447b565b611c66565b3480156109d257600080fd5b506103b16109e1366004614194565b611dc7565b3480156109f257600080fd5b506107d9610a01366004614194565b611de4565b348015610a1257600080fd5b506103b1611e7a565b348015610a2757600080fd5b506103b1610a363660046144ab565b611e80565b348015610a4757600080fd5b506103e7610a563660046144ca565b611ea1565b348015610a6757600080fd5b506103b1610a763660046144ab565b611ef2565b348015610a8757600080fd5b506103b1610a963660046144ab565b611f13565b348015610aa757600080fd5b506103e7610ab6366004614115565b611f34565b348015610ac757600080fd5b506103e7610ad636600461429a565b611f61565b60696020526000908152604090205481565b33610af6613fdc565b610b008284612240565b60208101518151919250600091610b1c9163ffffffff61233616565b9050610b3f8385610b3a85604001518561233690919063ffffffff16565b61235b565b6001600160a01b03831660008181526074602090815260408083208884528252918290208054850190558451908501518583015192518894937f4119153d17a36f9597d40e3ab4148d03261a439dddbec4e91799ab7159608e2693610ba8939092909190615372565b60405180910390a350505050565b33610bbf613fdc565b610bc98284612240565b90506000826001600160a01b0316610c068360400151610bfa8560200151866000015161233690919063ffffffff16565b9063ffffffff61233616565b604051610c129061501d565b60006040518083038185875af1925050503d8060008114610c4f576040519150601f19603f3d011682016040523d82523d6000602084013e610c54565b606091505b5050905080610c7e5760405162461bcd60e51b8152600401610c75906152fb565b60405180910390fd5b81516020830151604080850151905187936001600160a01b038816937fc1d8eb6e444b89fb8ff0991c19311c070df704ccb009e210d1462d5b2410bf4593610ba893615372565b607c546001600160a01b031681565b6000610ce08383611de4565b610d0e57506001600160a01b0382166000908152607360209081526040808320848452909152902054610d57565b6001600160a01b038316600081815260746020908152604080832086845282528083205493835260738252808320868452909152902054610d549163ffffffff6123de16565b90505b92915050565b60775481565b610d6c33612420565b610d885760405162461bcd60e51b8152600401610c759061513b565b610d958989896000612434565b6001600160a01b0389166000908152606f602090815260408083208b84529091529020600201819055610dc787612599565b8515610e7e5786861115610ded5760405162461bcd60e51b8152600401610c759061530b565b6001600160a01b03891660008181526074602090815260408083208c845282528083208a8155600181018a90556002810189905560038101889055848452607583528184208d855290925291829020859055905190918a917f138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e7890610e749088908c90615364565b60405180910390a3505b505050505050505050565b336000818152607460209081526040808320868452909152812090919083610ec35760405162461bcd60e51b8152600401610c75906150fb565b610ecd8286611de4565b610ee95760405162461bcd60e51b8152600401610c75906151ab565b8054841115610f0a5760405162461bcd60e51b8152600401610c759061526b565b610f148286612630565b610f305760405162461bcd60e51b8152600401610c75906151bb565b610f3a82866126e6565b506000610f4d83878785600001546128b1565b905081600301546363401ec50182600201541015610f69575060005b815485900382558015610f8c57610f8383878360016129e2565b610f8c81612ba0565b85836001600160a01b03167fef6c0c14fe9aa51af36acd791464dec3badbde668b63189b47bfa4e25be9b2b98784604051610fc8929190615364565b60405180910390a395945050505050565b610fe233612420565b610ffe5760405162461bcd60e51b8152600401610c759061513b565b8061101b5760405162461bcd60e51b8152600401610c759061517b565b6110258282612c0e565b611030826000611c66565b6000828152606860205260408120600601546001600160a01b0316906110599082908190612d2d565b505050565b607160209081526000938452604080852082529284528284209052825290208054600182015460029092015490919083565b6083546040516000916001600160a01b03169083906110b59088908890602401615074565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f4a7702bb00000000000000000000000000000000000000000000000000000000179052516111369190615011565b60006040518083038160008787f1925050503d8060008114611174576040519150601f19603f3d011682016040523d82523d6000602084013e611179565b606091505b505090508080611187575082155b6111a35760405162461bcd60e51b8152600401610c759061529b565b5050505050565b606d5481565b607860205280600052604060002060009150905080600701549080600801549080600901549080600a01549080600b01549080600c01549080600d0154905087565b6111fe33838333612e54565b5050565b61120a611945565b6112265760405162461bcd60e51b8152600401610c759061520b565b61122f82611c28565b61124b5760405162461bcd60e51b8152600401610c75906150eb565b61125361324b565b8111156112725760405162461bcd60e51b8152600401610c759061521b565b6000828152607b6020526040908190208290555182907f047575f43f09a7a093d94ec483064acfc61b7e25c0de28017da442abf99cb917906112b590849061531b565b60405180910390a25050565b6112ca33612420565b6112e65760405162461bcd60e51b8152600401610c759061513b565b61132e898989898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508b92508a91508990508888613257565b606b54881115610e7e57606b889055505050505050505050565b6000818152606860209081526040808320600601546001600160a01b03168352607382528083208484529091529020545b919050565b600091825260786020908152604080842092845291905290205490565b606e5481565b60006113ab613fdc565b6113b584846133e5565b8051602082015160408301519293506113d792610bfa9163ffffffff61233616565b949350505050565b60009182526078602090815260408084209284526001909201905290205490565b3361140b81846126e6565b506000821161142c5760405162461bcd60e51b8152600401610c75906150fb565b6114368184610cd4565b8211156114555760405162461bcd60e51b8152600401610c759061522b565b61145f8184612630565b61147b5760405162461bcd60e51b8152600401610c75906151bb565b6001600160a01b03811660008181526072602090815260408083208784528252808320805460018101909155938352607182528083208784528252808320848452909152902060020154156114e25760405162461bcd60e51b8152600401610c75906151fb565b6114ef82858560016129e2565b6001600160a01b038216600090815260716020908152604080832087845282528083208484529091529020600201839055611528611802565b6001600160a01b0383166000908152607160209081526040808320888452825280832085845290915290205561155c613453565b6001600160a01b03831660009081526071602090815260408083208884528252808320858452909152812060010191909155611599908590611c66565b8084836001600160a01b03167fd3bb4e423fbea695d16b982f9f682dc5f35152e5411646a8a5a79a6b02ba8d57866040516115d4919061531b565b60405180910390a450505050565b60006115ee8383611de4565b6115fa57506000610d57565b506001600160a01b03919091166000908152607460209081526040808320938352929052205490565b600061162d613fdc565b506001600160a01b0383166000908152606f60209081526040808320858452825291829020825160608101845281548082526001830154938201849052600290920154938101849052926113d7929091610bfa919063ffffffff61233616565b606080826040519080825280602002602001820160405280156116ca57816020015b6116b7613fdc565b8152602001906001900390816116af5790505b50905060005b83811015611772576001600160a01b038716600090815260716020908152604080832089845290915281209061170c878463ffffffff61233616565b8152602001908152602001600020604051806060016040529081600082015481526020016001820154815260200160028201548152505082828151811061174f57fe5b602090810291909101015261176b81600163ffffffff61233616565b90506116d0565b5095945050505050565b611784611945565b6117a05760405162461bcd60e51b8152600401610c759061520b565b6033546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3603380547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055565b6067546001015b90565b611814611945565b6118305760405162461bcd60e51b8152600401610c759061520b565b61183981612ba0565b50565b60675481565b606a6020908152600091825260409182902080548351601f60027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff610100600186161502019093169290920491820184900484028101840190945280845290918301828280156118f35780601f106118c8576101008083540402835291602001916118f3565b820191906000526020600020905b8154815290600101906020018083116118d657829003601f168201915b505050505081565b6083546001600160a01b031681565b606c5481565b61191a82826126e6565b6111fe5760405162461bcd60e51b8152600401610c759061514b565b6033546001600160a01b031690565b6033546001600160a01b0316331490565b607460209081526000928352604080842090915290825290208054600182015460028301546003909301549192909184565b61183933823461235b565b60009182526078602090815260408084209284526005909201905290205490565b608260009054906101000a90046001600160a01b03166001600160a01b031663c5f530af6040518163ffffffff1660e01b815260040160206040518083038186803b158015611a0257600080fd5b505afa158015611a16573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250611a3a919081019061445d565b341015611a595760405162461bcd60e51b8152600401610c75906152db565b80611a765760405162461bcd60e51b8152600401610c759061523b565b611ab63383838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061345792505050565b6111fe33606b543461235b565b607060209081526000928352604080842090915290825290205481565b611ae8611945565b611b045760405162461bcd60e51b8152600401610c759061520b565b608480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b606860205260009081526040902080546001820154600283015460038401546004850154600586015460069096015494959394929391929091906001600160a01b031687565b607560209081526000928352604080842090915290825290208054600182015460029092015490919083565b600081815260786020908152604091829020600601805483518184028101840190945280845260609392830182828015611c0957602002820191906000526020600020905b815481526020019060010190808311611bf5575b50505050509050919050565b33611c2281858585613482565b50505050565b600090815260686020526040902054608016151590565b6080546001600160a01b031681565b607b6020526000908152604090205481565b606b5481565b611c6f8261375b565b611c8b5760405162461bcd60e51b8152600401610c75906152ab565b60008281526068602052604090206003810154905415611ca9575060005b6066546040517fa4066fbe0000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063a4066fbe90611cf49086908590600401615364565b600060405180830381600087803b158015611d0e57600080fd5b505af1158015611d22573d6000803e3d6000fd5b50505050818015611d3257508015155b15611059576066546000848152606a60205260409081902090517f242a6e3f0000000000000000000000000000000000000000000000000000000081526001600160a01b039092169163242a6e3f91611d9091879190600401615329565b600060405180830381600087803b158015611daa57600080fd5b505af1158015611dbe573d6000803e3d6000fd5b50505050505050565b607360209081526000928352604080842090915290825290205481565b6001600160a01b038216600090815260746020908152604080832084845290915281206002015415801590611e3b57506001600160a01b038316600090815260746020908152604080832085845290915290205415155b8015610d5457506001600160a01b0383166000908152607460209081526040808320858452909152902060020154611e71613453565b11159392505050565b607f5481565b60009182526078602090815260408084209284526003909201905290205490565b3381611ebf5760405162461bcd60e51b8152600401610c75906150fb565b611ec98185611de4565b15611ee65760405162461bcd60e51b8152600401610c759061511b565b611c2281858585613482565b60009182526078602090815260408084209284526002909201905290205490565b60009182526078602090815260408084209284526004909201905290205490565b611f3c611945565b611f585760405162461bcd60e51b8152600401610c759061520b565b61183981613772565b6084546001600160a01b03163314611f8b5760405162461bcd60e51b8152600401610c75906152cb565b611f9583836126e6565b5060008111611fb65760405162461bcd60e51b8152600401610c75906150fb565b607c546040517f34f5b34f0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116906334f5b34f90612005903390879087908790600401615036565b600060405180830381600087803b15801561201f57600080fd5b505af1158015612033573d6000803e3d6000fd5b505050506001600160a01b038316600090815260736020908152604080832085845290915290205481111561207a5760405162461bcd60e51b8152600401610c759061527b565b60006120868484610cd4565b905080821015612115576001600160a01b0384166000908152607460209081526040808320868452909152902080546120c79083850363ffffffff6123de16565b815560405184906001600160a01b038716907fef6c0c14fe9aa51af36acd791464dec3badbde668b63189b47bfa4e25be9b2b99061210b9086880390600090615349565b60405180910390a3505b61212284848460016129e2565b61212d836000611c66565b64ffffffffff83856001600160a01b03167fd3bb4e423fbea695d16b982f9f682dc5f35152e5411646a8a5a79a6b02ba8d578560405161216d919061531b565b60405180910390a46000336001600160a01b03168360405161218e9061501d565b60006040518083038185875af1925050503d80600081146121cb576040519150601f19603f3d011682016040523d82523d6000602084013e6121d0565b606091505b50509050806121f15760405162461bcd60e51b8152600401610c75906152fb565b64ffffffffff84866001600160a01b03167f75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a2186604051612231919061531b565b60405180910390a45050505050565b612248613fdc565b6122528383612630565b61226e5760405162461bcd60e51b8152600401610c75906151bb565b61227883836126e6565b50506001600160a01b0382166000908152606f60209081526040808320848452825280832081516060810183528154808252600183015494820185905260029092015492810183905293926122d692610bfa9163ffffffff61233616565b9050806122f55760405162461bcd60e51b8152600401610c75906151db565b6001600160a01b0384166000908152606f602090815260408083208684529091528120818155600181018290556002015561232f81612599565b5092915050565b600082820183811015610d545760405162461bcd60e51b8152600401610c759061512b565b6123648261375b565b6123805760405162461bcd60e51b8152600401610c75906152ab565b600082815260686020526040902054156123ac5760405162461bcd60e51b8152600401610c759061516b565b6123b98383836001612434565b6123c28261380c565b6110595760405162461bcd60e51b8152600401610c759061528b565b6000610d5483836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506138e0565b6066546001600160a01b0390811691161490565b600082116124545760405162461bcd60e51b8152600401610c75906150fb565b61245e84846126e6565b506001600160a01b0384166000908152607360209081526040808320868452909152902054612493908363ffffffff61233616565b6001600160a01b03851660009081526073602090815260408083208784528252808320939093556068905220600301546124d3818463ffffffff61233616565b600085815260686020526040902060030155606c546124f8908463ffffffff61233616565b606c5560008481526068602052604090205461252557606d54612521908463ffffffff61233616565b606d555b612530848215611c66565b83856001600160a01b03167f9a8f44850296624dadfd9c246d17e47171d35727a181bd090aa14bbbe00238bb8560405161256a919061531b565b60405180910390a36000848152606860205260409020600601546111a39086906001600160a01b031684612d2d565b6066546040517f66e7ea0f0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116906366e7ea0f906125e4903090859060040161508f565b600060405180830381600087803b1580156125fe57600080fd5b505af1158015612612573d6000803e3d6000fd5b505060775461262a925090508263ffffffff61233616565b60775550565b607c546000906001600160a01b031661264b57506001610d57565b607c546040517f21d585c30000000000000000000000000000000000000000000000000000000081526001600160a01b03909116906321d585c390612696908690869060040161508f565b60206040518083038186803b1580156126ae57600080fd5b505afa1580156126c2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610d5491908101906143df565b60006126f0613fdc565b6126fa8484613911565b905061270583613a49565b6001600160a01b0385166000818152607060209081526040808320888452825280832094909455918152606f82528281208682528252829020825160608101845281548152600182015492810192909252600201549181019190915261276b9082613aa4565b6001600160a01b0385166000818152606f60209081526040808320888452825280832085518155858301516001808301919091559582015160029182015593835260758252808320888452825291829020825160608101845281548152948101549185019190915290910154908201526127e59082613aa4565b6001600160a01b0385166000908152607560209081526040808320878452825291829020835181559083015160018201559101516002909101556128298484611de4565b61288c576001600160a01b0384166000818152607460209081526040808320878452825280832083815560018082018590556002808301869055600390920185905594845260758352818420888552909252822082815592830182905591909101555b602081015115158061289e5750805115155b806113d757506040015115159392505050565b6001600160a01b038416600090815260756020908152604080832086845290915281205481906128f99084906128ed908763ffffffff613b1616565b9063ffffffff613b5016565b6001600160a01b03871660009081526075602090815260408083208984529091528120600101549192509061293a9085906128ed908863ffffffff613b1616565b6001600160a01b03881660009081526075602090815260408083208a84529091529020549091506002820483019061297290846123de565b6001600160a01b03891660009081526075602090815260408083208b84529091529020908155600101546129a690836123de565b6001600160a01b03891660009081526075602090815260408083208b84529091529020600101558581106129d75750845b979650505050505050565b6001600160a01b03841660009081526073602090815260408083208684528252808320805486900390556068909152902060030154612a27908363ffffffff6123de16565b600084815260686020526040902060030155606c54612a4c908363ffffffff6123de16565b606c55600083815260686020526040902054612a7957606d54612a75908363ffffffff6123de16565b606d555b6000612a8484611348565b90508015612b6e57600084815260686020526040902054612b6957608260009054906101000a90046001600160a01b03166001600160a01b031663c5f530af6040518163ffffffff1660e01b815260040160206040518083038186803b158015612aed57600080fd5b505afa158015612b01573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250612b25919081019061445d565b811015612b445760405162461bcd60e51b8152600401610c75906152db565b612b4d8461380c565b612b695760405162461bcd60e51b8152600401610c759061528b565b612b79565b612b79846001612c0e565b6000848152606860205260409020600601546111a39086906001600160a01b031684612d2d565b80156118395760405160009082156108fc0290839083818181858288f19350505050158015612bd3573d6000803e3d6000fd5b507fbd4ad73ad077fa0484c5c286b4c58aeb7302cceae2cb60113ce46a6510f47add81604051612c03919061531b565b60405180910390a150565b600082815260686020526040902054158015612c2957508015155b15612c5657600082815260686020526040902060030154606d54612c529163ffffffff6123de16565b606d555b6000828152606860205260409020548111156111fe57600082815260686020526040902081815560020154612cfd57612c8d611802565b600083815260686020526040902060020155612ca7613453565b600083815260686020526040908190206001810183905560020154905184927fac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e4792612cf492909190615364565b60405180910390a25b817fcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e826040516112b5919061531b565b6083546001600160a01b031615611059576083546040516000916001600160a01b031690627a120090612d669087908790602401615074565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f4a7702bb0000000000000000000000000000000000000000000000000000000017905251612de79190615011565b60006040518083038160008787f1925050503d8060008114612e25576040519150601f19603f3d011682016040523d82523d6000602084013e612e2a565b606091505b505090508080612e38575081155b611c225760405162461bcd60e51b8152600401610c759061529b565b612e5c613fdc565b506001600160a01b038416600090815260716020908152604080832086845282528083208584528252918290208251606081018452815480825260018301549382019390935260029091015492810192909252612ecb5760405162461bcd60e51b8152600401610c759061524b565b612ed58585612630565b612ef15760405162461bcd60e51b8152600401610c75906151bb565b60208082015182516000878152606890935260409092206001015490919015801590612f2d575060008681526068602052604090206001015482115b15612f4e575050600084815260686020526040902060018101546002909101545b608260009054906101000a90046001600160a01b03166001600160a01b031663b82b84276040518163ffffffff1660e01b815260040160206040518083038186803b158015612f9c57600080fd5b505afa158015612fb0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250612fd4919081019061445d565b8201612fde613453565b1015612ffc5760405162461bcd60e51b8152600401610c759061518b565b608260009054906101000a90046001600160a01b03166001600160a01b031663650acd666040518163ffffffff1660e01b815260040160206040518083038186803b15801561304a57600080fd5b505afa15801561305e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250613082919081019061445d565b810161308c611802565b10156130aa5760405162461bcd60e51b8152600401610c75906152bb565b6001600160a01b03871660009081526071602090815260408083208984528252808320888452909152812060020154906130e388611c28565b905060006131058383607b60008d815260200190815260200160002054613b92565b6001600160a01b038b1660009081526071602090815260408083208d845282528083208c845290915281208181556001810182905560020155606e80548201905590508083116131675760405162461bcd60e51b8152600401610c75906152eb565b60006001600160a01b038816613183858463ffffffff6123de16565b60405161318f9061501d565b60006040518083038185875af1925050503d80600081146131cc576040519150601f19603f3d011682016040523d82523d6000602084013e6131d1565b606091505b50509050806131f25760405162461bcd60e51b8152600401610c75906152fb565b6131fb82612ba0565b888a8c6001600160a01b03167f75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a2187604051613236919061531b565b60405180910390a45050505050505050505050565b670de0b6b3a764000090565b6001600160a01b0388166000908152606960205260409020541561328d5760405162461bcd60e51b8152600401610c759061519b565b6001600160a01b03881660008181526069602090815260408083208b90558a8352606882528083208981556004810189905560058101889055600181018690556002810187905560060180547fffffffffffffffffffffffff000000000000000000000000000000000000000016909417909355606a8152919020875161331692890190613ffd565b50876001600160a01b0316877f49bca1ed2666922f9f1690c26a569e1299c2a715fe57647d77e81adfabbf25bf8686604051613353929190615364565b60405180910390a3811561339c57867fac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e478383604051613393929190615364565b60405180910390a25b84156133db57867fcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e866040516133d2919061531b565b60405180910390a25b5050505050505050565b6133ed613fdc565b6133f5613fdc565b6133ff8484613911565b6001600160a01b0385166000908152606f6020908152604080832087845282529182902082516060810184528154815260018201549281019290925260020154918101919091529091506113d79082613aa4565b4290565b606b8054600101908190556110598382846000613472611802565b61347a613453565b600080613257565b61348c8484610cd4565b8111156134ab5760405162461bcd60e51b8152600401610c759061527b565b600083815260686020526040902054156134d75760405162461bcd60e51b8152600401610c759061516b565b608260009054906101000a90046001600160a01b03166001600160a01b0316630d7b26096040518163ffffffff1660e01b815260040160206040518083038186803b15801561352557600080fd5b505afa158015613539573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061355d919081019061445d565b82101580156135f15750608260009054906101000a90046001600160a01b03166001600160a01b0316630d4955e36040518163ffffffff1660e01b815260040160206040518083038186803b1580156135b557600080fd5b505afa1580156135c9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506135ed919081019061445d565b8211155b61360d5760405162461bcd60e51b8152600401610c759061515b565b600061361b83610bfa613453565b6000858152606860205260409020600601549091506001600160a01b03908116908616811461368a576001600160a01b038116600090815260746020908152604080832088845290915290206002015482111561368a5760405162461bcd60e51b8152600401610c759061525b565b61369486866126e6565b506001600160a01b0386166000908152607460209081526040808320888452909152902060038101548510156136dc5760405162461bcd60e51b8152600401610c75906151eb565b80546136ee908563ffffffff61233616565b81556136f8611802565b6001820155600281018390556003810185905560405186906001600160a01b038916907f138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e789061374a9089908990615364565b60405180910390a350505050505050565b600090815260686020526040902060050154151590565b6001600160a01b0381166137985760405162461bcd60e51b8152600401610c759061510b565b6033546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3603380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b60006138c561381961324b565b608254604080517f2265f28400000000000000000000000000000000000000000000000000000000815290516128ed926001600160a01b031691632265f284916004808301926020929190829003018186803b15801561387857600080fd5b505afa15801561388c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506138b0919081019061445d565b6138b986611348565b9063ffffffff613b1616565b60008381526068602052604090206003015411159050919050565b600081848411156139045760405162461bcd60e51b8152600401610c7591906150da565b50508183035b9392505050565b613919613fdc565b6001600160a01b03831660009081526070602090815260408083208584529091528120549061394784613a49565b905060006139558686613bf1565b9050818111156139625750805b8281101561396d5750815b6001600160a01b038616600081815260746020908152604080832089845282528083209383526073825280832089845290915281205482549091906139b990839063ffffffff6123de16565b905060006139cd84600001548a8988613cce565b90506139d7613fdc565b6139e5828660030154613d31565b90506139f3838b8a89613cce565b91506139fd613fdc565b613a08836000613d31565b9050613a16858c898b613cce565b9250613a20613fdc565b613a2b846000613d31565b9050613a38838383613f0a565b9d9c50505050505050505050505050565b60008181526068602052604081206002015415613a9c576000828152606860205260409020600201546067541015613a845750606754611379565b50600081815260686020526040902060020154611379565b505060675490565b613aac613fdc565b6040805160608101909152825184518291613acd919063ffffffff61233616565b8152602001613aed8460200151866020015161233690919063ffffffff16565b8152602001613b0d8460400151866040015161233690919063ffffffff16565b90529392505050565b600082613b2557506000610d57565b82820282848281613b3257fe5b0414610d545760405162461bcd60e51b8152600401610c75906151cb565b6000610d5483836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f000000000000815250613f25565b6000821580613ba85750613ba461324b565b8210155b15613bb55750600061390a565b613be06001610bfa613bc561324b565b6128ed86613bd161324b565b8a91900363ffffffff613b1616565b90508381111561390a57508261390a565b6001600160a01b0382166000908152607460209081526040808320848452909152812060010154606754613c26858583613f5c565b15613c34579150610d579050565b613c3f858584613f5c565b613c4e57600092505050610d57565b80821115613c6157600092505050610d57565b80821015613c9457600281830104613c7a868683613f5c565b15613c8a57806001019250613c8e565b8091505b50613c61565b80613ca457600092505050610d57565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01949350505050565b6000818310613cdf575060006113d7565b6000838152607860208181526040808420888552600190810183528185205487865293835281852089865201909152909120546129d7613d1d61324b565b6128ed896138b9858763ffffffff6123de16565b613d39613fdc565b60405180606001604052806000815260200160008152602001600081525090506000608260009054906101000a90046001600160a01b03166001600160a01b0316635e2308d26040518163ffffffff1660e01b815260040160206040518083038186803b158015613da957600080fd5b505afa158015613dbd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250613de1919081019061445d565b90508215613ee357600081613df461324b565b0390506000613e92608260009054906101000a90046001600160a01b03166001600160a01b0316630d4955e36040518163ffffffff1660e01b815260040160206040518083038186803b158015613e4a57600080fd5b505afa158015613e5e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250613e82919081019061445d565b6128ed848863ffffffff613b1616565b90506000613eb3613ea161324b565b6128ed8987860163ffffffff613b1616565b9050613ed0613ec061324b565b6128ed898763ffffffff613b1616565b60208601819052900384525061232f9050565b613efe613eee61324b565b6128ed868463ffffffff613b1616565b60408301525092915050565b613f12613fdc565b6113d7613f1f8585613aa4565b83613aa4565b60008183613f465760405162461bcd60e51b8152600401610c7591906150da565b506000838581613f5257fe5b0495945050505050565b6001600160a01b038316600090815260746020908152604080832085845290915281206001015482108015906113d757506001600160a01b0384166000908152607460209081526040808320868452909152902060020154613fbd83613fc7565b1115949350505050565b60009081526078602052604090206007015490565b60405180606001604052806000815260200160008152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061403e57805160ff191683800117855561406b565b8280016001018555821561406b579182015b8281111561406b578251825591602001919060010190614050565b5061407792915061407b565b5090565b61180991905b808211156140775760008155600101614081565b8035610d5781615529565b8035610d578161553d565b8051610d578161553d565b60008083601f8401126140c857600080fd5b50813567ffffffffffffffff8111156140e057600080fd5b6020830191508360018202830111156140f857600080fd5b9250929050565b8035610d5781615546565b8051610d5781615546565b60006020828403121561412757600080fd5b60006113d78484614095565b6000806000806080858703121561414957600080fd5b60006141558787614095565b945050602061416687828801614095565b9350506040614177878288016140a0565b9250506060614188878288016140ff565b91505092959194509250565b600080604083850312156141a757600080fd5b60006141b38585614095565b92505060206141c4858286016140ff565b9150509250929050565b60008060008060008060008060006101008a8c0312156141ed57600080fd5b60006141f98c8c614095565b995050602061420a8c828d016140ff565b98505060408a013567ffffffffffffffff81111561422757600080fd5b6142338c828d016140b6565b975097505060606142468c828d016140ff565b95505060806142578c828d016140ff565b94505060a06142688c828d016140ff565b93505060c06142798c828d016140ff565b92505060e061428a8c828d016140ff565b9150509295985092959850929598565b6000806000606084860312156142af57600080fd5b60006142bb8686614095565b93505060206142cc868287016140ff565b92505060406142dd868287016140ff565b9150509250925092565b600080600080608085870312156142fd57600080fd5b60006143098787614095565b945050602061431a878288016140ff565b9350506040614177878288016140ff565b60008060008060008060008060006101208a8c03121561434a57600080fd5b60006143568c8c614095565b99505060206143678c828d016140ff565b98505060406143788c828d016140ff565b97505060606143898c828d016140ff565b965050608061439a8c828d016140ff565b95505060a06143ab8c828d016140ff565b94505060c06143bc8c828d016140ff565b93505060e06143cd8c828d016140ff565b92505061010061428a8c828d016140ff565b6000602082840312156143f157600080fd5b60006113d784846140ab565b6000806020838503121561441057600080fd5b823567ffffffffffffffff81111561442757600080fd5b614433858286016140b6565b92509250509250929050565b60006020828403121561445157600080fd5b60006113d784846140ff565b60006020828403121561446f57600080fd5b60006113d7848461410a565b6000806040838503121561448e57600080fd5b600061449a85856140ff565b92505060206141c4858286016140a0565b600080604083850312156144be57600080fd5b60006141b385856140ff565b6000806000606084860312156144df57600080fd5b60006142bb86866140ff565b60006144f78383614fd5565b505060600190565b600061450b8383615008565b505060200190565b61451c816154b4565b82525050565b61451c81615498565b60006145368261548b565b614540818561548f565b935061454b83615479565b8060005b8381101561457957815161456388826144eb565b975061456e83615479565b92505060010161454f565b509495945050505050565b600061458f8261548b565b614599818561548f565b93506145a483615479565b8060005b838110156145795781516145bc88826144ff565b97506145c783615479565b9250506001016145a8565b61451c816154a3565b60006145e68261548b565b6145f08185611379565b93506146008185602086016154d5565b9290920192915050565b60006146158261548b565b61461f818561548f565b935061462f8185602086016154d5565b61463881615501565b9093019392505050565b60008154600181166000811461465f57600181146146a3576146e2565b607f6002830416614670818761548f565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00841681529550506020850192506146e2565b600282046146b1818761548f565b95506146bc8561547f565b60005b828110156146db578154888201526001909101906020016146bf565b8701945050505b505092915050565b61451c816154bf565b600061470060178361548f565b7f76616c696461746f722069736e277420736c6173686564000000000000000000815260200192915050565b6000614739600b8361548f565b7f7a65726f20616d6f756e74000000000000000000000000000000000000000000815260200192915050565b600061477260268361548f565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206181527f6464726573730000000000000000000000000000000000000000000000000000602082015260400192915050565b60006147d160118361548f565b7f616c7265616479206c6f636b6564207570000000000000000000000000000000815260200192915050565b600061480a601b8361548f565b7f536166654d6174683a206164646974696f6e206f766572666c6f770000000000815260200192915050565b600061484360298361548f565b7f63616c6c6572206973206e6f7420746865204e6f64654472697665724175746881527f20636f6e74726163740000000000000000000000000000000000000000000000602082015260400192915050565b60006148a260108361548f565b7f6e6f7468696e6720746f20737461736800000000000000000000000000000000815260200192915050565b60006148db60128361548f565b7f696e636f7272656374206475726174696f6e0000000000000000000000000000815260200192915050565b600061491460168361548f565b7f76616c696461746f722069736e27742061637469766500000000000000000000815260200192915050565b600061494d600c8361548f565b7f77726f6e67207374617475730000000000000000000000000000000000000000815260200192915050565b600061498660168361548f565b7f6e6f7420656e6f7567682074696d652070617373656400000000000000000000815260200192915050565b60006149bf60188361548f565b7f76616c696461746f7220616c7265616479206578697374730000000000000000815260200192915050565b60006149f8600d8361548f565b7f6e6f74206c6f636b656420757000000000000000000000000000000000000000815260200192915050565b6000614a31601a8361548f565b7f6f75747374616e64696e67207356495452412062616c616e6365000000000000815260200192915050565b6000614a6a60218361548f565b7f536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f81527f7700000000000000000000000000000000000000000000000000000000000000602082015260400192915050565b6000614ac9600c8361548f565b7f7a65726f20726577617264730000000000000000000000000000000000000000815260200192915050565b6000614b02601f8361548f565b7f6c6f636b7570206475726174696f6e2063616e6e6f7420646563726561736500815260200192915050565b6000614b3b60138361548f565b7f7772494420616c72656164792065786973747300000000000000000000000000815260200192915050565b6000614b7460208361548f565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572815260200192915050565b6000614bad60218361548f565b7f6d757374206265206c657373207468616e206f7220657175616c20746f20312e81527f3000000000000000000000000000000000000000000000000000000000000000602082015260400192915050565b6000614c0c60198361548f565b7f6e6f7420656e6f75676820756e6c6f636b6564207374616b6500000000000000815260200192915050565b6000614c45600c8361548f565b7f656d707479207075626b65790000000000000000000000000000000000000000815260200192915050565b6000614c7e60158361548f565b7f7265717565737420646f65736e27742065786973740000000000000000000000815260200192915050565b6000614cb760288361548f565b7f76616c696461746f72206c6f636b757020706572696f642077696c6c20656e6481527f206561726c696572000000000000000000000000000000000000000000000000602082015260400192915050565b6000614d1660178361548f565b7f6e6f7420656e6f756768206c6f636b6564207374616b65000000000000000000815260200192915050565b6000610d57600083611379565b6000614d5c60108361548f565b7f6e6f7420656e6f756768207374616b6500000000000000000000000000000000815260200192915050565b6000614d9560298361548f565b7f76616c696461746f7227732064656c65676174696f6e73206c696d697420697381527f2065786365656465640000000000000000000000000000000000000000000000602082015260400192915050565b6000614df4601b8361548f565b7f676f7620766f746573207265636f756e74696e67206661696c65640000000000815260200192915050565b6000614e2d60178361548f565b7f76616c696461746f7220646f65736e2774206578697374000000000000000000815260200192915050565b6000614e6660188361548f565b7f6e6f7420656e6f7567682065706f636873207061737365640000000000000000815260200192915050565b6000614e9f60148361548f565b7f6e6f74207356495452412066696e616c697a6572000000000000000000000000815260200192915050565b6000614ed860178361548f565b7f696e73756666696369656e742073656c662d7374616b65000000000000000000815260200192915050565b6000614f1160168361548f565b7f7374616b652069732066756c6c7920736c617368656400000000000000000000815260200192915050565b6000614f4a60148361548f565b7f4661696c656420746f2073656e64205649545241000000000000000000000000815260200192915050565b6000614f83602c8361548f565b7f6c6f636b6564207374616b652069732067726561746572207468616e2074686581527f2077686f6c65207374616b650000000000000000000000000000000000000000602082015260400192915050565b80516060830190614fe68482615008565b506020820151614ff96020850182615008565b506040820151611c2260408501825b61451c81611809565b600061390a82846145db565b6000610d5782614d42565b60208101610d578284614522565b608081016150448287614513565b6150516020830186614522565b61505e6040830185615008565b61506b6060830184615008565b95945050505050565b604081016150828285614522565b61390a6020830184614522565b6040810161509d8285614522565b61390a6020830184615008565b60208082528101610d54818461452b565b60208082528101610d548184614584565b60208101610d5782846145d2565b60208082528101610d54818461460a565b60208082528101610d57816146f3565b60208082528101610d578161472c565b60208082528101610d5781614765565b60208082528101610d57816147c4565b60208082528101610d57816147fd565b60208082528101610d5781614836565b60208082528101610d5781614895565b60208082528101610d57816148ce565b60208082528101610d5781614907565b60208082528101610d5781614940565b60208082528101610d5781614979565b60208082528101610d57816149b2565b60208082528101610d57816149eb565b60208082528101610d5781614a24565b60208082528101610d5781614a5d565b60208082528101610d5781614abc565b60208082528101610d5781614af5565b60208082528101610d5781614b2e565b60208082528101610d5781614b67565b60208082528101610d5781614ba0565b60208082528101610d5781614bff565b60208082528101610d5781614c38565b60208082528101610d5781614c71565b60208082528101610d5781614caa565b60208082528101610d5781614d09565b60208082528101610d5781614d4f565b60208082528101610d5781614d88565b60208082528101610d5781614de7565b60208082528101610d5781614e20565b60208082528101610d5781614e59565b60208082528101610d5781614e92565b60208082528101610d5781614ecb565b60208082528101610d5781614f04565b60208082528101610d5781614f3d565b60208082528101610d5781614f76565b60208101610d578284615008565b604081016153378285615008565b81810360208301526113d78184614642565b604081016153578285615008565b61390a60208301846146ea565b6040810161509d8285615008565b606081016153808286615008565b61538d6020830185615008565b6113d76040830184615008565b608081016153a88287615008565b6150516020830186615008565b60e081016153c3828a615008565b6153d06020830189615008565b6153dd6040830188615008565b6153ea6060830187615008565b6153f76080830186615008565b61540460a0830185615008565b61541160c0830184614522565b98975050505050505050565b60e0810161542b828a615008565b6154386020830189615008565b6154456040830188615008565b6154526060830187615008565b61545f6080830186615008565b61546c60a0830185615008565b61541160c0830184615008565b60200190565b60009081526020902090565b5190565b90815260200190565b6000610d57826154a8565b151590565b6001600160a01b031690565b6000610d57826154ca565b6000610d5782611809565b6000610d5782615498565b60005b838110156154f05781810151838201526020016154d8565b83811115611c225750506000910152565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01690565b61553281615498565b811461183957600080fd5b615532816154a3565b6155328161180956fea365627a7a723158205b69aa0e5aa05624737f0852357b99c9bf5a285c499fc3be09a88b2945f31aeb6c6578706572696d656e74616cf564736f6c63430005110040",
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
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// BurnFTM is a paid mutator transaction binding the contract method 0x90a6c475.
//
// Solidity: function burnFTM(uint256 amount) returns()
func (_Contract *ContractTransactor) BurnFTM(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "burnFTM", amount)
}

// BurnFTM is a paid mutator transaction binding the contract method 0x90a6c475.
//
// Solidity: function burnFTM(uint256 amount) returns()
func (_Contract *ContractSession) BurnFTM(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.BurnFTM(&_Contract.TransactOpts, amount)
}

// BurnFTM is a paid mutator transaction binding the contract method 0x90a6c475.
//
// Solidity: function burnFTM(uint256 amount) returns()
func (_Contract *ContractTransactorSession) BurnFTM(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.BurnFTM(&_Contract.TransactOpts, amount)
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

// MintFTM is a paid mutator transaction binding the contract method 0xe2f8c336.
//
// Solidity: function mintFTM(address receiver, uint256 amount, string justification) returns()
func (_Contract *ContractTransactor) MintFTM(opts *bind.TransactOpts, receiver common.Address, amount *big.Int, justification string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "mintFTM", receiver, amount, justification)
}

// MintFTM is a paid mutator transaction binding the contract method 0xe2f8c336.
//
// Solidity: function mintFTM(address receiver, uint256 amount, string justification) returns()
func (_Contract *ContractSession) MintFTM(receiver common.Address, amount *big.Int, justification string) (*types.Transaction, error) {
	return _Contract.Contract.MintFTM(&_Contract.TransactOpts, receiver, amount, justification)
}

// MintFTM is a paid mutator transaction binding the contract method 0xe2f8c336.
//
// Solidity: function mintFTM(address receiver, uint256 amount, string justification) returns()
func (_Contract *ContractTransactorSession) MintFTM(receiver common.Address, amount *big.Int, justification string) (*types.Transaction, error) {
	return _Contract.Contract.MintFTM(&_Contract.TransactOpts, receiver, amount, justification)
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

// Undelegate is a paid mutator transaction binding the contract method 0x4f864df4.
//
// Solidity: function undelegate(uint256 toValidatorID, uint256 wrID, uint256 amount) returns()
func (_Contract *ContractTransactor) Undelegate(opts *bind.TransactOpts, toValidatorID *big.Int, wrID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "undelegate", toValidatorID, wrID, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4f864df4.
//
// Solidity: function undelegate(uint256 toValidatorID, uint256 wrID, uint256 amount) returns()
func (_Contract *ContractSession) Undelegate(toValidatorID *big.Int, wrID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Undelegate(&_Contract.TransactOpts, toValidatorID, wrID, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4f864df4.
//
// Solidity: function undelegate(uint256 toValidatorID, uint256 wrID, uint256 amount) returns()
func (_Contract *ContractTransactorSession) Undelegate(toValidatorID *big.Int, wrID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Undelegate(&_Contract.TransactOpts, toValidatorID, wrID, amount)
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

// ContractBurntFTMIterator is returned from FilterBurntFTM and is used to iterate over the raw logs and unpacked data for BurntFTM events raised by the Contract contract.
type ContractBurntFTMIterator struct {
	Event *ContractBurntFTM // Event containing the contract specifics and raw log

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
func (it *ContractBurntFTMIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBurntFTM)
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
		it.Event = new(ContractBurntFTM)
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
func (it *ContractBurntFTMIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBurntFTMIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBurntFTM represents a BurntFTM event raised by the Contract contract.
type ContractBurntFTM struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurntFTM is a free log retrieval operation binding the contract event 0x8918bd6046d08b314e457977f29562c5d76a7030d79b1edba66e8a5da0b77ae8.
//
// Solidity: event BurntFTM(uint256 amount)
func (_Contract *ContractFilterer) FilterBurntFTM(opts *bind.FilterOpts) (*ContractBurntFTMIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "BurntFTM")
	if err != nil {
		return nil, err
	}
	return &ContractBurntFTMIterator{contract: _Contract.contract, event: "BurntFTM", logs: logs, sub: sub}, nil
}

// WatchBurntFTM is a free log subscription operation binding the contract event 0x8918bd6046d08b314e457977f29562c5d76a7030d79b1edba66e8a5da0b77ae8.
//
// Solidity: event BurntFTM(uint256 amount)
func (_Contract *ContractFilterer) WatchBurntFTM(opts *bind.WatchOpts, sink chan<- *ContractBurntFTM) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "BurntFTM")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBurntFTM)
				if err := _Contract.contract.UnpackLog(event, "BurntFTM", log); err != nil {
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

// ParseBurntFTM is a log parse operation binding the contract event 0x8918bd6046d08b314e457977f29562c5d76a7030d79b1edba66e8a5da0b77ae8.
//
// Solidity: event BurntFTM(uint256 amount)
func (_Contract *ContractFilterer) ParseBurntFTM(log types.Log) (*ContractBurntFTM, error) {
	event := new(ContractBurntFTM)
	if err := _Contract.contract.UnpackLog(event, "BurntFTM", log); err != nil {
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

// ContractInflatedFTMIterator is returned from FilterInflatedFTM and is used to iterate over the raw logs and unpacked data for InflatedFTM events raised by the Contract contract.
type ContractInflatedFTMIterator struct {
	Event *ContractInflatedFTM // Event containing the contract specifics and raw log

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
func (it *ContractInflatedFTMIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractInflatedFTM)
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
		it.Event = new(ContractInflatedFTM)
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
func (it *ContractInflatedFTMIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractInflatedFTMIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractInflatedFTM represents a InflatedFTM event raised by the Contract contract.
type ContractInflatedFTM struct {
	Receiver      common.Address
	Amount        *big.Int
	Justification string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInflatedFTM is a free log retrieval operation binding the contract event 0x9eec469b348bcf64bbfb60e46ce7b160e2e09bf5421496a2cdbc43714c28b8ad.
//
// Solidity: event InflatedFTM(address indexed receiver, uint256 amount, string justification)
func (_Contract *ContractFilterer) FilterInflatedFTM(opts *bind.FilterOpts, receiver []common.Address) (*ContractInflatedFTMIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "InflatedFTM", receiverRule)
	if err != nil {
		return nil, err
	}
	return &ContractInflatedFTMIterator{contract: _Contract.contract, event: "InflatedFTM", logs: logs, sub: sub}, nil
}

// WatchInflatedFTM is a free log subscription operation binding the contract event 0x9eec469b348bcf64bbfb60e46ce7b160e2e09bf5421496a2cdbc43714c28b8ad.
//
// Solidity: event InflatedFTM(address indexed receiver, uint256 amount, string justification)
func (_Contract *ContractFilterer) WatchInflatedFTM(opts *bind.WatchOpts, sink chan<- *ContractInflatedFTM, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "InflatedFTM", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractInflatedFTM)
				if err := _Contract.contract.UnpackLog(event, "InflatedFTM", log); err != nil {
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

// ParseInflatedFTM is a log parse operation binding the contract event 0x9eec469b348bcf64bbfb60e46ce7b160e2e09bf5421496a2cdbc43714c28b8ad.
//
// Solidity: event InflatedFTM(address indexed receiver, uint256 amount, string justification)
func (_Contract *ContractFilterer) ParseInflatedFTM(log types.Log) (*ContractInflatedFTM, error) {
	event := new(ContractInflatedFTM)
	if err := _Contract.contract.UnpackLog(event, "InflatedFTM", log); err != nil {
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

var ContractBinRuntime = "0x60806040526004361061038c5760003560e01c8063854873e1116101dc578063bd14d90711610102578063cfdbb7cd116100a0578063df00c9221161006f578063df00c92214610a5b578063e261641a14610a7b578063f2fde38b14610a9b578063f6a3250314610abb5761038c565b8063cfdbb7cd146109e6578063d96ed50514610a06578063dc31e1af14610a1b578063de67f21514610a3b5761038c565b8063c65ee0e1116100dc578063c65ee0e114610971578063c7be95de14610991578063cc8343aa146109a6578063cfd47663146109c65761038c565b8063bd14d9071461091c578063c3de580e1461093c578063c5f956af1461095c5761038c565b80639fa6dd351161017a578063b334566b11610149578063b334566b1461087c578063b5d896271461089c578063b810e411146108cf578063b88a37e2146108ef5761038c565b80639fa6dd3514610816578063a198d22914610829578063a5a470ad14610849578063a86a056f1461085c5761038c565b80638cddb015116101b65780638cddb0151461078f5780638da5cb5b146107af5780638f32d59b146107c457806396c7ee46146107e65761038c565b8063854873e114610738578063893675c6146107655780638b0e9f3f1461077a5761038c565b80634f7c4efb116102c1578063634b91e31161025f578063715018a61161022e578063715018a6146106d957806376671808146106ee5780637aa353bd146107035780637cacb1d6146107235761038c565b8063634b91e31461064c578063670322f81461066c5780636f4986631461068c578063702797e3146106ac5761038c565b806358f95b801161029b57806358f95b80146105d75780635fab23a8146105f75780636099ecb21461060c57806361e53fcc1461062c5761038c565b80634f7c4efb146105775780634feb92f3146105975780635601fe01146105b75761038c565b80631d3ac42c1161032e57806320c0849d1161030857806320c0849d146104ef57806328f731481461050f57806339b80c0014610524578063441a3e70146105575761038c565b80631d3ac42c146104805780631e702f83146104a05780631f270152146104c05761038c565b80630e559d821161036a5780630e559d821461040957806312622d0e1461042b57806318160ddd1461044b57806318f628d4146104605761038c565b80630135b1db1461039157806308c36874146103c75780630962ef79146103e9575b600080fd5b34801561039d57600080fd5b506103b16103ac366004614115565b610adb565b6040516103be919061531b565b60405180910390f35b3480156103d357600080fd5b506103e76103e236600461443f565b610aed565b005b3480156103f557600080fd5b506103e761040436600461443f565b610bb6565b34801561041557600080fd5b5061041e610cc5565b6040516103be9190615028565b34801561043757600080fd5b506103b1610446366004614194565b610cd4565b34801561045757600080fd5b506103b1610d5d565b34801561046c57600080fd5b506103e761047b36600461432b565b610d63565b34801561048c57600080fd5b506103b161049b3660046144ab565b610e89565b3480156104ac57600080fd5b506103e76104bb3660046144ab565b610fd9565b3480156104cc57600080fd5b506104e06104db36600461429a565b61105e565b6040516103be93929190615372565b3480156104fb57600080fd5b506103e761050a366004614133565b611090565b34801561051b57600080fd5b506103b16111aa565b34801561053057600080fd5b5061054461053f36600461443f565b6111b0565b6040516103be979695949392919061541d565b34801561056357600080fd5b506103e76105723660046144ab565b6111f2565b34801561058357600080fd5b506103e76105923660046144ab565b611202565b3480156105a357600080fd5b506103e76105b23660046141ce565b6112c1565b3480156105c357600080fd5b506103b16105d236600461443f565b611348565b3480156105e357600080fd5b506103b16105f23660046144ab565b61137e565b34801561060357600080fd5b506103b161139b565b34801561061857600080fd5b506103b1610627366004614194565b6113a1565b34801561063857600080fd5b506103b16106473660046144ab565b6113df565b34801561065857600080fd5b506103e76106673660046144ab565b611400565b34801561067857600080fd5b506103b1610687366004614194565b6115e2565b34801561069857600080fd5b506103b16106a7366004614194565b611623565b3480156106b857600080fd5b506106cc6106c73660046142e7565b61168d565b6040516103be91906150aa565b3480156106e557600080fd5b506103e761177c565b3480156106fa57600080fd5b506103b1611802565b34801561070f57600080fd5b506103e761071e36600461443f565b61180c565b34801561072f57600080fd5b506103b161183c565b34801561074457600080fd5b5061075861075336600461443f565b611842565b6040516103be91906150da565b34801561077157600080fd5b5061041e6118fb565b34801561078657600080fd5b506103b161190a565b34801561079b57600080fd5b506103e76107aa366004614194565b611910565b3480156107bb57600080fd5b5061041e611936565b3480156107d057600080fd5b506107d9611945565b6040516103be91906150cc565b3480156107f257600080fd5b50610806610801366004614194565b611956565b6040516103be949392919061539a565b6103e761082436600461443f565b611988565b34801561083557600080fd5b506103b16108443660046144ab565b611993565b6103e76108573660046143fd565b6119b4565b34801561086857600080fd5b506103b1610877366004614194565b611ac3565b34801561088857600080fd5b506103e7610897366004614115565b611ae0565b3480156108a857600080fd5b506108bc6108b736600461443f565b611b3e565b6040516103be97969594939291906153b5565b3480156108db57600080fd5b506104e06108ea366004614194565b611b84565b3480156108fb57600080fd5b5061090f61090a36600461443f565b611bb0565b6040516103be91906150bb565b34801561092857600080fd5b506103e76109373660046144ca565b611c15565b34801561094857600080fd5b506107d961095736600461443f565b611c28565b34801561096857600080fd5b5061041e611c3f565b34801561097d57600080fd5b506103b161098c36600461443f565b611c4e565b34801561099d57600080fd5b506103b1611c60565b3480156109b257600080fd5b506103e76109c136600461447b565b611c66565b3480156109d257600080fd5b506103b16109e1366004614194565b611dc7565b3480156109f257600080fd5b506107d9610a01366004614194565b611de4565b348015610a1257600080fd5b506103b1611e7a565b348015610a2757600080fd5b506103b1610a363660046144ab565b611e80565b348015610a4757600080fd5b506103e7610a563660046144ca565b611ea1565b348015610a6757600080fd5b506103b1610a763660046144ab565b611ef2565b348015610a8757600080fd5b506103b1610a963660046144ab565b611f13565b348015610aa757600080fd5b506103e7610ab6366004614115565b611f34565b348015610ac757600080fd5b506103e7610ad636600461429a565b611f61565b60696020526000908152604090205481565b33610af6613fdc565b610b008284612240565b60208101518151919250600091610b1c9163ffffffff61233616565b9050610b3f8385610b3a85604001518561233690919063ffffffff16565b61235b565b6001600160a01b03831660008181526074602090815260408083208884528252918290208054850190558451908501518583015192518894937f4119153d17a36f9597d40e3ab4148d03261a439dddbec4e91799ab7159608e2693610ba8939092909190615372565b60405180910390a350505050565b33610bbf613fdc565b610bc98284612240565b90506000826001600160a01b0316610c068360400151610bfa8560200151866000015161233690919063ffffffff16565b9063ffffffff61233616565b604051610c129061501d565b60006040518083038185875af1925050503d8060008114610c4f576040519150601f19603f3d011682016040523d82523d6000602084013e610c54565b606091505b5050905080610c7e5760405162461bcd60e51b8152600401610c75906152fb565b60405180910390fd5b81516020830151604080850151905187936001600160a01b038816937fc1d8eb6e444b89fb8ff0991c19311c070df704ccb009e210d1462d5b2410bf4593610ba893615372565b607c546001600160a01b031681565b6000610ce08383611de4565b610d0e57506001600160a01b0382166000908152607360209081526040808320848452909152902054610d57565b6001600160a01b038316600081815260746020908152604080832086845282528083205493835260738252808320868452909152902054610d549163ffffffff6123de16565b90505b92915050565b60775481565b610d6c33612420565b610d885760405162461bcd60e51b8152600401610c759061513b565b610d958989896000612434565b6001600160a01b0389166000908152606f602090815260408083208b84529091529020600201819055610dc787612599565b8515610e7e5786861115610ded5760405162461bcd60e51b8152600401610c759061530b565b6001600160a01b03891660008181526074602090815260408083208c845282528083208a8155600181018a90556002810189905560038101889055848452607583528184208d855290925291829020859055905190918a917f138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e7890610e749088908c90615364565b60405180910390a3505b505050505050505050565b336000818152607460209081526040808320868452909152812090919083610ec35760405162461bcd60e51b8152600401610c75906150fb565b610ecd8286611de4565b610ee95760405162461bcd60e51b8152600401610c75906151ab565b8054841115610f0a5760405162461bcd60e51b8152600401610c759061526b565b610f148286612630565b610f305760405162461bcd60e51b8152600401610c75906151bb565b610f3a82866126e6565b506000610f4d83878785600001546128b1565b905081600301546363401ec50182600201541015610f69575060005b815485900382558015610f8c57610f8383878360016129e2565b610f8c81612ba0565b85836001600160a01b03167fef6c0c14fe9aa51af36acd791464dec3badbde668b63189b47bfa4e25be9b2b98784604051610fc8929190615364565b60405180910390a395945050505050565b610fe233612420565b610ffe5760405162461bcd60e51b8152600401610c759061513b565b8061101b5760405162461bcd60e51b8152600401610c759061517b565b6110258282612c0e565b611030826000611c66565b6000828152606860205260408120600601546001600160a01b0316906110599082908190612d2d565b505050565b607160209081526000938452604080852082529284528284209052825290208054600182015460029092015490919083565b6083546040516000916001600160a01b03169083906110b59088908890602401615074565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f4a7702bb00000000000000000000000000000000000000000000000000000000179052516111369190615011565b60006040518083038160008787f1925050503d8060008114611174576040519150601f19603f3d011682016040523d82523d6000602084013e611179565b606091505b505090508080611187575082155b6111a35760405162461bcd60e51b8152600401610c759061529b565b5050505050565b606d5481565b607860205280600052604060002060009150905080600701549080600801549080600901549080600a01549080600b01549080600c01549080600d0154905087565b6111fe33838333612e54565b5050565b61120a611945565b6112265760405162461bcd60e51b8152600401610c759061520b565b61122f82611c28565b61124b5760405162461bcd60e51b8152600401610c75906150eb565b61125361324b565b8111156112725760405162461bcd60e51b8152600401610c759061521b565b6000828152607b6020526040908190208290555182907f047575f43f09a7a093d94ec483064acfc61b7e25c0de28017da442abf99cb917906112b590849061531b565b60405180910390a25050565b6112ca33612420565b6112e65760405162461bcd60e51b8152600401610c759061513b565b61132e898989898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508b92508a91508990508888613257565b606b54881115610e7e57606b889055505050505050505050565b6000818152606860209081526040808320600601546001600160a01b03168352607382528083208484529091529020545b919050565b600091825260786020908152604080842092845291905290205490565b606e5481565b60006113ab613fdc565b6113b584846133e5565b8051602082015160408301519293506113d792610bfa9163ffffffff61233616565b949350505050565b60009182526078602090815260408084209284526001909201905290205490565b3361140b81846126e6565b506000821161142c5760405162461bcd60e51b8152600401610c75906150fb565b6114368184610cd4565b8211156114555760405162461bcd60e51b8152600401610c759061522b565b61145f8184612630565b61147b5760405162461bcd60e51b8152600401610c75906151bb565b6001600160a01b03811660008181526072602090815260408083208784528252808320805460018101909155938352607182528083208784528252808320848452909152902060020154156114e25760405162461bcd60e51b8152600401610c75906151fb565b6114ef82858560016129e2565b6001600160a01b038216600090815260716020908152604080832087845282528083208484529091529020600201839055611528611802565b6001600160a01b0383166000908152607160209081526040808320888452825280832085845290915290205561155c613453565b6001600160a01b03831660009081526071602090815260408083208884528252808320858452909152812060010191909155611599908590611c66565b8084836001600160a01b03167fd3bb4e423fbea695d16b982f9f682dc5f35152e5411646a8a5a79a6b02ba8d57866040516115d4919061531b565b60405180910390a450505050565b60006115ee8383611de4565b6115fa57506000610d57565b506001600160a01b03919091166000908152607460209081526040808320938352929052205490565b600061162d613fdc565b506001600160a01b0383166000908152606f60209081526040808320858452825291829020825160608101845281548082526001830154938201849052600290920154938101849052926113d7929091610bfa919063ffffffff61233616565b606080826040519080825280602002602001820160405280156116ca57816020015b6116b7613fdc565b8152602001906001900390816116af5790505b50905060005b83811015611772576001600160a01b038716600090815260716020908152604080832089845290915281209061170c878463ffffffff61233616565b8152602001908152602001600020604051806060016040529081600082015481526020016001820154815260200160028201548152505082828151811061174f57fe5b602090810291909101015261176b81600163ffffffff61233616565b90506116d0565b5095945050505050565b611784611945565b6117a05760405162461bcd60e51b8152600401610c759061520b565b6033546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3603380547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055565b6067546001015b90565b611814611945565b6118305760405162461bcd60e51b8152600401610c759061520b565b61183981612ba0565b50565b60675481565b606a6020908152600091825260409182902080548351601f60027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff610100600186161502019093169290920491820184900484028101840190945280845290918301828280156118f35780601f106118c8576101008083540402835291602001916118f3565b820191906000526020600020905b8154815290600101906020018083116118d657829003601f168201915b505050505081565b6083546001600160a01b031681565b606c5481565b61191a82826126e6565b6111fe5760405162461bcd60e51b8152600401610c759061514b565b6033546001600160a01b031690565b6033546001600160a01b0316331490565b607460209081526000928352604080842090915290825290208054600182015460028301546003909301549192909184565b61183933823461235b565b60009182526078602090815260408084209284526005909201905290205490565b608260009054906101000a90046001600160a01b03166001600160a01b031663c5f530af6040518163ffffffff1660e01b815260040160206040518083038186803b158015611a0257600080fd5b505afa158015611a16573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250611a3a919081019061445d565b341015611a595760405162461bcd60e51b8152600401610c75906152db565b80611a765760405162461bcd60e51b8152600401610c759061523b565b611ab63383838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061345792505050565b6111fe33606b543461235b565b607060209081526000928352604080842090915290825290205481565b611ae8611945565b611b045760405162461bcd60e51b8152600401610c759061520b565b608480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b606860205260009081526040902080546001820154600283015460038401546004850154600586015460069096015494959394929391929091906001600160a01b031687565b607560209081526000928352604080842090915290825290208054600182015460029092015490919083565b600081815260786020908152604091829020600601805483518184028101840190945280845260609392830182828015611c0957602002820191906000526020600020905b815481526020019060010190808311611bf5575b50505050509050919050565b33611c2281858585613482565b50505050565b600090815260686020526040902054608016151590565b6080546001600160a01b031681565b607b6020526000908152604090205481565b606b5481565b611c6f8261375b565b611c8b5760405162461bcd60e51b8152600401610c75906152ab565b60008281526068602052604090206003810154905415611ca9575060005b6066546040517fa4066fbe0000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063a4066fbe90611cf49086908590600401615364565b600060405180830381600087803b158015611d0e57600080fd5b505af1158015611d22573d6000803e3d6000fd5b50505050818015611d3257508015155b15611059576066546000848152606a60205260409081902090517f242a6e3f0000000000000000000000000000000000000000000000000000000081526001600160a01b039092169163242a6e3f91611d9091879190600401615329565b600060405180830381600087803b158015611daa57600080fd5b505af1158015611dbe573d6000803e3d6000fd5b50505050505050565b607360209081526000928352604080842090915290825290205481565b6001600160a01b038216600090815260746020908152604080832084845290915281206002015415801590611e3b57506001600160a01b038316600090815260746020908152604080832085845290915290205415155b8015610d5457506001600160a01b0383166000908152607460209081526040808320858452909152902060020154611e71613453565b11159392505050565b607f5481565b60009182526078602090815260408084209284526003909201905290205490565b3381611ebf5760405162461bcd60e51b8152600401610c75906150fb565b611ec98185611de4565b15611ee65760405162461bcd60e51b8152600401610c759061511b565b611c2281858585613482565b60009182526078602090815260408084209284526002909201905290205490565b60009182526078602090815260408084209284526004909201905290205490565b611f3c611945565b611f585760405162461bcd60e51b8152600401610c759061520b565b61183981613772565b6084546001600160a01b03163314611f8b5760405162461bcd60e51b8152600401610c75906152cb565b611f9583836126e6565b5060008111611fb65760405162461bcd60e51b8152600401610c75906150fb565b607c546040517f34f5b34f0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116906334f5b34f90612005903390879087908790600401615036565b600060405180830381600087803b15801561201f57600080fd5b505af1158015612033573d6000803e3d6000fd5b505050506001600160a01b038316600090815260736020908152604080832085845290915290205481111561207a5760405162461bcd60e51b8152600401610c759061527b565b60006120868484610cd4565b905080821015612115576001600160a01b0384166000908152607460209081526040808320868452909152902080546120c79083850363ffffffff6123de16565b815560405184906001600160a01b038716907fef6c0c14fe9aa51af36acd791464dec3badbde668b63189b47bfa4e25be9b2b99061210b9086880390600090615349565b60405180910390a3505b61212284848460016129e2565b61212d836000611c66565b64ffffffffff83856001600160a01b03167fd3bb4e423fbea695d16b982f9f682dc5f35152e5411646a8a5a79a6b02ba8d578560405161216d919061531b565b60405180910390a46000336001600160a01b03168360405161218e9061501d565b60006040518083038185875af1925050503d80600081146121cb576040519150601f19603f3d011682016040523d82523d6000602084013e6121d0565b606091505b50509050806121f15760405162461bcd60e51b8152600401610c75906152fb565b64ffffffffff84866001600160a01b03167f75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a2186604051612231919061531b565b60405180910390a45050505050565b612248613fdc565b6122528383612630565b61226e5760405162461bcd60e51b8152600401610c75906151bb565b61227883836126e6565b50506001600160a01b0382166000908152606f60209081526040808320848452825280832081516060810183528154808252600183015494820185905260029092015492810183905293926122d692610bfa9163ffffffff61233616565b9050806122f55760405162461bcd60e51b8152600401610c75906151db565b6001600160a01b0384166000908152606f602090815260408083208684529091528120818155600181018290556002015561232f81612599565b5092915050565b600082820183811015610d545760405162461bcd60e51b8152600401610c759061512b565b6123648261375b565b6123805760405162461bcd60e51b8152600401610c75906152ab565b600082815260686020526040902054156123ac5760405162461bcd60e51b8152600401610c759061516b565b6123b98383836001612434565b6123c28261380c565b6110595760405162461bcd60e51b8152600401610c759061528b565b6000610d5483836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506138e0565b6066546001600160a01b0390811691161490565b600082116124545760405162461bcd60e51b8152600401610c75906150fb565b61245e84846126e6565b506001600160a01b0384166000908152607360209081526040808320868452909152902054612493908363ffffffff61233616565b6001600160a01b03851660009081526073602090815260408083208784528252808320939093556068905220600301546124d3818463ffffffff61233616565b600085815260686020526040902060030155606c546124f8908463ffffffff61233616565b606c5560008481526068602052604090205461252557606d54612521908463ffffffff61233616565b606d555b612530848215611c66565b83856001600160a01b03167f9a8f44850296624dadfd9c246d17e47171d35727a181bd090aa14bbbe00238bb8560405161256a919061531b565b60405180910390a36000848152606860205260409020600601546111a39086906001600160a01b031684612d2d565b6066546040517f66e7ea0f0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116906366e7ea0f906125e4903090859060040161508f565b600060405180830381600087803b1580156125fe57600080fd5b505af1158015612612573d6000803e3d6000fd5b505060775461262a925090508263ffffffff61233616565b60775550565b607c546000906001600160a01b031661264b57506001610d57565b607c546040517f21d585c30000000000000000000000000000000000000000000000000000000081526001600160a01b03909116906321d585c390612696908690869060040161508f565b60206040518083038186803b1580156126ae57600080fd5b505afa1580156126c2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610d5491908101906143df565b60006126f0613fdc565b6126fa8484613911565b905061270583613a49565b6001600160a01b0385166000818152607060209081526040808320888452825280832094909455918152606f82528281208682528252829020825160608101845281548152600182015492810192909252600201549181019190915261276b9082613aa4565b6001600160a01b0385166000818152606f60209081526040808320888452825280832085518155858301516001808301919091559582015160029182015593835260758252808320888452825291829020825160608101845281548152948101549185019190915290910154908201526127e59082613aa4565b6001600160a01b0385166000908152607560209081526040808320878452825291829020835181559083015160018201559101516002909101556128298484611de4565b61288c576001600160a01b0384166000818152607460209081526040808320878452825280832083815560018082018590556002808301869055600390920185905594845260758352818420888552909252822082815592830182905591909101555b602081015115158061289e5750805115155b806113d757506040015115159392505050565b6001600160a01b038416600090815260756020908152604080832086845290915281205481906128f99084906128ed908763ffffffff613b1616565b9063ffffffff613b5016565b6001600160a01b03871660009081526075602090815260408083208984529091528120600101549192509061293a9085906128ed908863ffffffff613b1616565b6001600160a01b03881660009081526075602090815260408083208a84529091529020549091506002820483019061297290846123de565b6001600160a01b03891660009081526075602090815260408083208b84529091529020908155600101546129a690836123de565b6001600160a01b03891660009081526075602090815260408083208b84529091529020600101558581106129d75750845b979650505050505050565b6001600160a01b03841660009081526073602090815260408083208684528252808320805486900390556068909152902060030154612a27908363ffffffff6123de16565b600084815260686020526040902060030155606c54612a4c908363ffffffff6123de16565b606c55600083815260686020526040902054612a7957606d54612a75908363ffffffff6123de16565b606d555b6000612a8484611348565b90508015612b6e57600084815260686020526040902054612b6957608260009054906101000a90046001600160a01b03166001600160a01b031663c5f530af6040518163ffffffff1660e01b815260040160206040518083038186803b158015612aed57600080fd5b505afa158015612b01573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250612b25919081019061445d565b811015612b445760405162461bcd60e51b8152600401610c75906152db565b612b4d8461380c565b612b695760405162461bcd60e51b8152600401610c759061528b565b612b79565b612b79846001612c0e565b6000848152606860205260409020600601546111a39086906001600160a01b031684612d2d565b80156118395760405160009082156108fc0290839083818181858288f19350505050158015612bd3573d6000803e3d6000fd5b507fbd4ad73ad077fa0484c5c286b4c58aeb7302cceae2cb60113ce46a6510f47add81604051612c03919061531b565b60405180910390a150565b600082815260686020526040902054158015612c2957508015155b15612c5657600082815260686020526040902060030154606d54612c529163ffffffff6123de16565b606d555b6000828152606860205260409020548111156111fe57600082815260686020526040902081815560020154612cfd57612c8d611802565b600083815260686020526040902060020155612ca7613453565b600083815260686020526040908190206001810183905560020154905184927fac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e4792612cf492909190615364565b60405180910390a25b817fcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e826040516112b5919061531b565b6083546001600160a01b031615611059576083546040516000916001600160a01b031690627a120090612d669087908790602401615074565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f4a7702bb0000000000000000000000000000000000000000000000000000000017905251612de79190615011565b60006040518083038160008787f1925050503d8060008114612e25576040519150601f19603f3d011682016040523d82523d6000602084013e612e2a565b606091505b505090508080612e38575081155b611c225760405162461bcd60e51b8152600401610c759061529b565b612e5c613fdc565b506001600160a01b038416600090815260716020908152604080832086845282528083208584528252918290208251606081018452815480825260018301549382019390935260029091015492810192909252612ecb5760405162461bcd60e51b8152600401610c759061524b565b612ed58585612630565b612ef15760405162461bcd60e51b8152600401610c75906151bb565b60208082015182516000878152606890935260409092206001015490919015801590612f2d575060008681526068602052604090206001015482115b15612f4e575050600084815260686020526040902060018101546002909101545b608260009054906101000a90046001600160a01b03166001600160a01b031663b82b84276040518163ffffffff1660e01b815260040160206040518083038186803b158015612f9c57600080fd5b505afa158015612fb0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250612fd4919081019061445d565b8201612fde613453565b1015612ffc5760405162461bcd60e51b8152600401610c759061518b565b608260009054906101000a90046001600160a01b03166001600160a01b031663650acd666040518163ffffffff1660e01b815260040160206040518083038186803b15801561304a57600080fd5b505afa15801561305e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250613082919081019061445d565b810161308c611802565b10156130aa5760405162461bcd60e51b8152600401610c75906152bb565b6001600160a01b03871660009081526071602090815260408083208984528252808320888452909152812060020154906130e388611c28565b905060006131058383607b60008d815260200190815260200160002054613b92565b6001600160a01b038b1660009081526071602090815260408083208d845282528083208c845290915281208181556001810182905560020155606e80548201905590508083116131675760405162461bcd60e51b8152600401610c75906152eb565b60006001600160a01b038816613183858463ffffffff6123de16565b60405161318f9061501d565b60006040518083038185875af1925050503d80600081146131cc576040519150601f19603f3d011682016040523d82523d6000602084013e6131d1565b606091505b50509050806131f25760405162461bcd60e51b8152600401610c75906152fb565b6131fb82612ba0565b888a8c6001600160a01b03167f75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a2187604051613236919061531b565b60405180910390a45050505050505050505050565b670de0b6b3a764000090565b6001600160a01b0388166000908152606960205260409020541561328d5760405162461bcd60e51b8152600401610c759061519b565b6001600160a01b03881660008181526069602090815260408083208b90558a8352606882528083208981556004810189905560058101889055600181018690556002810187905560060180547fffffffffffffffffffffffff000000000000000000000000000000000000000016909417909355606a8152919020875161331692890190613ffd565b50876001600160a01b0316877f49bca1ed2666922f9f1690c26a569e1299c2a715fe57647d77e81adfabbf25bf8686604051613353929190615364565b60405180910390a3811561339c57867fac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e478383604051613393929190615364565b60405180910390a25b84156133db57867fcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e866040516133d2919061531b565b60405180910390a25b5050505050505050565b6133ed613fdc565b6133f5613fdc565b6133ff8484613911565b6001600160a01b0385166000908152606f6020908152604080832087845282529182902082516060810184528154815260018201549281019290925260020154918101919091529091506113d79082613aa4565b4290565b606b8054600101908190556110598382846000613472611802565b61347a613453565b600080613257565b61348c8484610cd4565b8111156134ab5760405162461bcd60e51b8152600401610c759061527b565b600083815260686020526040902054156134d75760405162461bcd60e51b8152600401610c759061516b565b608260009054906101000a90046001600160a01b03166001600160a01b0316630d7b26096040518163ffffffff1660e01b815260040160206040518083038186803b15801561352557600080fd5b505afa158015613539573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061355d919081019061445d565b82101580156135f15750608260009054906101000a90046001600160a01b03166001600160a01b0316630d4955e36040518163ffffffff1660e01b815260040160206040518083038186803b1580156135b557600080fd5b505afa1580156135c9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506135ed919081019061445d565b8211155b61360d5760405162461bcd60e51b8152600401610c759061515b565b600061361b83610bfa613453565b6000858152606860205260409020600601549091506001600160a01b03908116908616811461368a576001600160a01b038116600090815260746020908152604080832088845290915290206002015482111561368a5760405162461bcd60e51b8152600401610c759061525b565b61369486866126e6565b506001600160a01b0386166000908152607460209081526040808320888452909152902060038101548510156136dc5760405162461bcd60e51b8152600401610c75906151eb565b80546136ee908563ffffffff61233616565b81556136f8611802565b6001820155600281018390556003810185905560405186906001600160a01b038916907f138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e789061374a9089908990615364565b60405180910390a350505050505050565b600090815260686020526040902060050154151590565b6001600160a01b0381166137985760405162461bcd60e51b8152600401610c759061510b565b6033546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3603380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b60006138c561381961324b565b608254604080517f2265f28400000000000000000000000000000000000000000000000000000000815290516128ed926001600160a01b031691632265f284916004808301926020929190829003018186803b15801561387857600080fd5b505afa15801561388c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506138b0919081019061445d565b6138b986611348565b9063ffffffff613b1616565b60008381526068602052604090206003015411159050919050565b600081848411156139045760405162461bcd60e51b8152600401610c7591906150da565b50508183035b9392505050565b613919613fdc565b6001600160a01b03831660009081526070602090815260408083208584529091528120549061394784613a49565b905060006139558686613bf1565b9050818111156139625750805b8281101561396d5750815b6001600160a01b038616600081815260746020908152604080832089845282528083209383526073825280832089845290915281205482549091906139b990839063ffffffff6123de16565b905060006139cd84600001548a8988613cce565b90506139d7613fdc565b6139e5828660030154613d31565b90506139f3838b8a89613cce565b91506139fd613fdc565b613a08836000613d31565b9050613a16858c898b613cce565b9250613a20613fdc565b613a2b846000613d31565b9050613a38838383613f0a565b9d9c50505050505050505050505050565b60008181526068602052604081206002015415613a9c576000828152606860205260409020600201546067541015613a845750606754611379565b50600081815260686020526040902060020154611379565b505060675490565b613aac613fdc565b6040805160608101909152825184518291613acd919063ffffffff61233616565b8152602001613aed8460200151866020015161233690919063ffffffff16565b8152602001613b0d8460400151866040015161233690919063ffffffff16565b90529392505050565b600082613b2557506000610d57565b82820282848281613b3257fe5b0414610d545760405162461bcd60e51b8152600401610c75906151cb565b6000610d5483836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f000000000000815250613f25565b6000821580613ba85750613ba461324b565b8210155b15613bb55750600061390a565b613be06001610bfa613bc561324b565b6128ed86613bd161324b565b8a91900363ffffffff613b1616565b90508381111561390a57508261390a565b6001600160a01b0382166000908152607460209081526040808320848452909152812060010154606754613c26858583613f5c565b15613c34579150610d579050565b613c3f858584613f5c565b613c4e57600092505050610d57565b80821115613c6157600092505050610d57565b80821015613c9457600281830104613c7a868683613f5c565b15613c8a57806001019250613c8e565b8091505b50613c61565b80613ca457600092505050610d57565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01949350505050565b6000818310613cdf575060006113d7565b6000838152607860208181526040808420888552600190810183528185205487865293835281852089865201909152909120546129d7613d1d61324b565b6128ed896138b9858763ffffffff6123de16565b613d39613fdc565b60405180606001604052806000815260200160008152602001600081525090506000608260009054906101000a90046001600160a01b03166001600160a01b0316635e2308d26040518163ffffffff1660e01b815260040160206040518083038186803b158015613da957600080fd5b505afa158015613dbd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250613de1919081019061445d565b90508215613ee357600081613df461324b565b0390506000613e92608260009054906101000a90046001600160a01b03166001600160a01b0316630d4955e36040518163ffffffff1660e01b815260040160206040518083038186803b158015613e4a57600080fd5b505afa158015613e5e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250613e82919081019061445d565b6128ed848863ffffffff613b1616565b90506000613eb3613ea161324b565b6128ed8987860163ffffffff613b1616565b9050613ed0613ec061324b565b6128ed898763ffffffff613b1616565b60208601819052900384525061232f9050565b613efe613eee61324b565b6128ed868463ffffffff613b1616565b60408301525092915050565b613f12613fdc565b6113d7613f1f8585613aa4565b83613aa4565b60008183613f465760405162461bcd60e51b8152600401610c7591906150da565b506000838581613f5257fe5b0495945050505050565b6001600160a01b038316600090815260746020908152604080832085845290915281206001015482108015906113d757506001600160a01b0384166000908152607460209081526040808320868452909152902060020154613fbd83613fc7565b1115949350505050565b60009081526078602052604090206007015490565b60405180606001604052806000815260200160008152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061403e57805160ff191683800117855561406b565b8280016001018555821561406b579182015b8281111561406b578251825591602001919060010190614050565b5061407792915061407b565b5090565b61180991905b808211156140775760008155600101614081565b8035610d5781615529565b8035610d578161553d565b8051610d578161553d565b60008083601f8401126140c857600080fd5b50813567ffffffffffffffff8111156140e057600080fd5b6020830191508360018202830111156140f857600080fd5b9250929050565b8035610d5781615546565b8051610d5781615546565b60006020828403121561412757600080fd5b60006113d78484614095565b6000806000806080858703121561414957600080fd5b60006141558787614095565b945050602061416687828801614095565b9350506040614177878288016140a0565b9250506060614188878288016140ff565b91505092959194509250565b600080604083850312156141a757600080fd5b60006141b38585614095565b92505060206141c4858286016140ff565b9150509250929050565b60008060008060008060008060006101008a8c0312156141ed57600080fd5b60006141f98c8c614095565b995050602061420a8c828d016140ff565b98505060408a013567ffffffffffffffff81111561422757600080fd5b6142338c828d016140b6565b975097505060606142468c828d016140ff565b95505060806142578c828d016140ff565b94505060a06142688c828d016140ff565b93505060c06142798c828d016140ff565b92505060e061428a8c828d016140ff565b9150509295985092959850929598565b6000806000606084860312156142af57600080fd5b60006142bb8686614095565b93505060206142cc868287016140ff565b92505060406142dd868287016140ff565b9150509250925092565b600080600080608085870312156142fd57600080fd5b60006143098787614095565b945050602061431a878288016140ff565b9350506040614177878288016140ff565b60008060008060008060008060006101208a8c03121561434a57600080fd5b60006143568c8c614095565b99505060206143678c828d016140ff565b98505060406143788c828d016140ff565b97505060606143898c828d016140ff565b965050608061439a8c828d016140ff565b95505060a06143ab8c828d016140ff565b94505060c06143bc8c828d016140ff565b93505060e06143cd8c828d016140ff565b92505061010061428a8c828d016140ff565b6000602082840312156143f157600080fd5b60006113d784846140ab565b6000806020838503121561441057600080fd5b823567ffffffffffffffff81111561442757600080fd5b614433858286016140b6565b92509250509250929050565b60006020828403121561445157600080fd5b60006113d784846140ff565b60006020828403121561446f57600080fd5b60006113d7848461410a565b6000806040838503121561448e57600080fd5b600061449a85856140ff565b92505060206141c4858286016140a0565b600080604083850312156144be57600080fd5b60006141b385856140ff565b6000806000606084860312156144df57600080fd5b60006142bb86866140ff565b60006144f78383614fd5565b505060600190565b600061450b8383615008565b505060200190565b61451c816154b4565b82525050565b61451c81615498565b60006145368261548b565b614540818561548f565b935061454b83615479565b8060005b8381101561457957815161456388826144eb565b975061456e83615479565b92505060010161454f565b509495945050505050565b600061458f8261548b565b614599818561548f565b93506145a483615479565b8060005b838110156145795781516145bc88826144ff565b97506145c783615479565b9250506001016145a8565b61451c816154a3565b60006145e68261548b565b6145f08185611379565b93506146008185602086016154d5565b9290920192915050565b60006146158261548b565b61461f818561548f565b935061462f8185602086016154d5565b61463881615501565b9093019392505050565b60008154600181166000811461465f57600181146146a3576146e2565b607f6002830416614670818761548f565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00841681529550506020850192506146e2565b600282046146b1818761548f565b95506146bc8561547f565b60005b828110156146db578154888201526001909101906020016146bf565b8701945050505b505092915050565b61451c816154bf565b600061470060178361548f565b7f76616c696461746f722069736e277420736c6173686564000000000000000000815260200192915050565b6000614739600b8361548f565b7f7a65726f20616d6f756e74000000000000000000000000000000000000000000815260200192915050565b600061477260268361548f565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206181527f6464726573730000000000000000000000000000000000000000000000000000602082015260400192915050565b60006147d160118361548f565b7f616c7265616479206c6f636b6564207570000000000000000000000000000000815260200192915050565b600061480a601b8361548f565b7f536166654d6174683a206164646974696f6e206f766572666c6f770000000000815260200192915050565b600061484360298361548f565b7f63616c6c6572206973206e6f7420746865204e6f64654472697665724175746881527f20636f6e74726163740000000000000000000000000000000000000000000000602082015260400192915050565b60006148a260108361548f565b7f6e6f7468696e6720746f20737461736800000000000000000000000000000000815260200192915050565b60006148db60128361548f565b7f696e636f7272656374206475726174696f6e0000000000000000000000000000815260200192915050565b600061491460168361548f565b7f76616c696461746f722069736e27742061637469766500000000000000000000815260200192915050565b600061494d600c8361548f565b7f77726f6e67207374617475730000000000000000000000000000000000000000815260200192915050565b600061498660168361548f565b7f6e6f7420656e6f7567682074696d652070617373656400000000000000000000815260200192915050565b60006149bf60188361548f565b7f76616c696461746f7220616c7265616479206578697374730000000000000000815260200192915050565b60006149f8600d8361548f565b7f6e6f74206c6f636b656420757000000000000000000000000000000000000000815260200192915050565b6000614a31601a8361548f565b7f6f75747374616e64696e67207356495452412062616c616e6365000000000000815260200192915050565b6000614a6a60218361548f565b7f536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f81527f7700000000000000000000000000000000000000000000000000000000000000602082015260400192915050565b6000614ac9600c8361548f565b7f7a65726f20726577617264730000000000000000000000000000000000000000815260200192915050565b6000614b02601f8361548f565b7f6c6f636b7570206475726174696f6e2063616e6e6f7420646563726561736500815260200192915050565b6000614b3b60138361548f565b7f7772494420616c72656164792065786973747300000000000000000000000000815260200192915050565b6000614b7460208361548f565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572815260200192915050565b6000614bad60218361548f565b7f6d757374206265206c657373207468616e206f7220657175616c20746f20312e81527f3000000000000000000000000000000000000000000000000000000000000000602082015260400192915050565b6000614c0c60198361548f565b7f6e6f7420656e6f75676820756e6c6f636b6564207374616b6500000000000000815260200192915050565b6000614c45600c8361548f565b7f656d707479207075626b65790000000000000000000000000000000000000000815260200192915050565b6000614c7e60158361548f565b7f7265717565737420646f65736e27742065786973740000000000000000000000815260200192915050565b6000614cb760288361548f565b7f76616c696461746f72206c6f636b757020706572696f642077696c6c20656e6481527f206561726c696572000000000000000000000000000000000000000000000000602082015260400192915050565b6000614d1660178361548f565b7f6e6f7420656e6f756768206c6f636b6564207374616b65000000000000000000815260200192915050565b6000610d57600083611379565b6000614d5c60108361548f565b7f6e6f7420656e6f756768207374616b6500000000000000000000000000000000815260200192915050565b6000614d9560298361548f565b7f76616c696461746f7227732064656c65676174696f6e73206c696d697420697381527f2065786365656465640000000000000000000000000000000000000000000000602082015260400192915050565b6000614df4601b8361548f565b7f676f7620766f746573207265636f756e74696e67206661696c65640000000000815260200192915050565b6000614e2d60178361548f565b7f76616c696461746f7220646f65736e2774206578697374000000000000000000815260200192915050565b6000614e6660188361548f565b7f6e6f7420656e6f7567682065706f636873207061737365640000000000000000815260200192915050565b6000614e9f60148361548f565b7f6e6f74207356495452412066696e616c697a6572000000000000000000000000815260200192915050565b6000614ed860178361548f565b7f696e73756666696369656e742073656c662d7374616b65000000000000000000815260200192915050565b6000614f1160168361548f565b7f7374616b652069732066756c6c7920736c617368656400000000000000000000815260200192915050565b6000614f4a60148361548f565b7f4661696c656420746f2073656e64205649545241000000000000000000000000815260200192915050565b6000614f83602c8361548f565b7f6c6f636b6564207374616b652069732067726561746572207468616e2074686581527f2077686f6c65207374616b650000000000000000000000000000000000000000602082015260400192915050565b80516060830190614fe68482615008565b506020820151614ff96020850182615008565b506040820151611c2260408501825b61451c81611809565b600061390a82846145db565b6000610d5782614d42565b60208101610d578284614522565b608081016150448287614513565b6150516020830186614522565b61505e6040830185615008565b61506b6060830184615008565b95945050505050565b604081016150828285614522565b61390a6020830184614522565b6040810161509d8285614522565b61390a6020830184615008565b60208082528101610d54818461452b565b60208082528101610d548184614584565b60208101610d5782846145d2565b60208082528101610d54818461460a565b60208082528101610d57816146f3565b60208082528101610d578161472c565b60208082528101610d5781614765565b60208082528101610d57816147c4565b60208082528101610d57816147fd565b60208082528101610d5781614836565b60208082528101610d5781614895565b60208082528101610d57816148ce565b60208082528101610d5781614907565b60208082528101610d5781614940565b60208082528101610d5781614979565b60208082528101610d57816149b2565b60208082528101610d57816149eb565b60208082528101610d5781614a24565b60208082528101610d5781614a5d565b60208082528101610d5781614abc565b60208082528101610d5781614af5565b60208082528101610d5781614b2e565b60208082528101610d5781614b67565b60208082528101610d5781614ba0565b60208082528101610d5781614bff565b60208082528101610d5781614c38565b60208082528101610d5781614c71565b60208082528101610d5781614caa565b60208082528101610d5781614d09565b60208082528101610d5781614d4f565b60208082528101610d5781614d88565b60208082528101610d5781614de7565b60208082528101610d5781614e20565b60208082528101610d5781614e59565b60208082528101610d5781614e92565b60208082528101610d5781614ecb565b60208082528101610d5781614f04565b60208082528101610d5781614f3d565b60208082528101610d5781614f76565b60208101610d578284615008565b604081016153378285615008565b81810360208301526113d78184614642565b604081016153578285615008565b61390a60208301846146ea565b6040810161509d8285615008565b606081016153808286615008565b61538d6020830185615008565b6113d76040830184615008565b608081016153a88287615008565b6150516020830186615008565b60e081016153c3828a615008565b6153d06020830189615008565b6153dd6040830188615008565b6153ea6060830187615008565b6153f76080830186615008565b61540460a0830185615008565b61541160c0830184614522565b98975050505050505050565b60e0810161542b828a615008565b6154386020830189615008565b6154456040830188615008565b6154526060830187615008565b61545f6080830186615008565b61546c60a0830185615008565b61541160c0830184615008565b60200190565b60009081526020902090565b5190565b90815260200190565b6000610d57826154a8565b151590565b6001600160a01b031690565b6000610d57826154ca565b6000610d5782611809565b6000610d5782615498565b60005b838110156154f05781810151838201526020016154d8565b83811115611c225750506000910152565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01690565b61553281615498565b811461183957600080fd5b615532816154a3565b6155328161180956fea365627a7a723158205b69aa0e5aa05624737f0852357b99c9bf5a285c499fc3be09a88b2945f31aeb6c6578706572696d656e74616cf564736f6c63430005110040"
