package committeestate

import (
	"errors"

	"github.com/incognitochain/incognito-chain/common"
	"github.com/incognitochain/incognito-chain/dataaccessobject/statedb"
	"github.com/incognitochain/incognito-chain/incognitokey"
	"github.com/incognitochain/incognito-chain/instruction"
)

//processUnstakeInstruction : process unstake instruction from beacon block

func (b *BeaconCommitteeStateV1) processUnstakeInstruction(
	unstakeInstruction *instruction.UnstakeInstruction,
	env *BeaconCommitteeStateEnvironment,
	committeeChange *CommitteeChange,
) (*CommitteeChange, [][]string, error) {
	newCommitteeChange := committeeChange
	incurredInstructions := [][]string{}
	returnStakerInfoPublicKeys := make(map[byte][]string)
	stakingTxs := make(map[byte][]string)
	percentReturns := make(map[byte][]uint)
	indexNextEpochShardCandidate := make(map[string]int)

	for i, v := range b.nextEpochShardCandidate {
		key, err := v.ToBase58()
		if err != nil {
			return newCommitteeChange, nil, err
		}
		indexNextEpochShardCandidate[key] = i
	}

	for _, committeePublicKey := range unstakeInstruction.CommitteePublicKeys {
		if common.IndexOfStr(committeePublicKey, env.unassignedCommonPool) == -1 {
			// TODO: @tin is it possible for one key to be not in any list?
			if common.IndexOfStr(committeePublicKey, env.allSubstituteCommittees) != -1 {
				// if found in committee list then turn off auto staking
				if _, ok := b.autoStake[committeePublicKey]; ok {
					b.autoStake[committeePublicKey] = false
					newCommitteeChange.Unstake = append(newCommitteeChange.Unstake, committeePublicKey)
				}
			}
		} else {
			if _, ok := b.autoStake[committeePublicKey]; ok {
				delete(b.autoStake, committeePublicKey)
			}

			if _, ok := b.stakingTx[committeePublicKey]; ok {
				delete(b.stakingTx, committeePublicKey)
			}
			committeePublicKeyStruct := incognitokey.CommitteePublicKey{}
			err := committeePublicKeyStruct.FromBase58(committeePublicKey)
			if err != nil {
				return newCommitteeChange, nil, err
			}
			rewardReceiverKey := committeePublicKeyStruct.GetIncKeyBase58()
			if _, ok := b.rewardReceiver[rewardReceiverKey]; ok {
				delete(b.rewardReceiver, rewardReceiverKey)
			}

			indexCandidate := indexNextEpochShardCandidate[committeePublicKey]
			b.nextEpochShardCandidate = append(b.nextEpochShardCandidate[:indexCandidate], b.nextEpochShardCandidate[indexCandidate+1:]...)
			stakerInfo, has, err := statedb.GetStakerInfo(env.ConsensusStateDB, committeePublicKey)
			if err != nil {
				return newCommitteeChange, nil, err
			}
			if !has {
				return newCommitteeChange, nil, errors.New("Can't find staker info")
			}

			newCommitteeChange.NextEpochShardCandidateRemoved =
				append(newCommitteeChange.NextEpochShardCandidateRemoved, committeePublicKeyStruct)

			returnStakerInfoPublicKeys[stakerInfo.ShardID()] =
				append(returnStakerInfoPublicKeys[stakerInfo.ShardID()], committeePublicKey)
			percentReturns[stakerInfo.ShardID()] =
				append(percentReturns[stakerInfo.ShardID()], 100)
			stakingTxs[stakerInfo.ShardID()] =
				append(stakingTxs[stakerInfo.ShardID()], stakerInfo.TxStakingID().String())
		}
	}

	for i, v := range returnStakerInfoPublicKeys {
		if v != nil {
			returnStakingIns := instruction.NewReturnStakeInsWithValue(
				v,
				i,
				stakingTxs[i],
				percentReturns[i],
			)
			incurredInstructions = append(incurredInstructions, returnStakingIns.ToString())
		}
	}
	return newCommitteeChange, incurredInstructions, nil
}

func (b *BeaconCommitteeStateV1) getSubstituteCandidates() ([]string, error) {
	commonPoolValidators := []string{}

	candidateBeaconWaitingForNextRandom := b.nextEpochBeaconCandidate
	candidateBeaconWaitingForNextRandomStr, err := incognitokey.CommitteeKeyListToString(candidateBeaconWaitingForNextRandom)
	if err != nil {
		return nil, err
	}
	commonPoolValidators = append(commonPoolValidators, candidateBeaconWaitingForNextRandomStr...)
	candidateShardWaitingForNextRandom := b.nextEpochShardCandidate
	candidateShardWaitingForNextRandomStr, err := incognitokey.CommitteeKeyListToString(candidateShardWaitingForNextRandom)
	if err != nil {
		return nil, err
	}
	commonPoolValidators = append(commonPoolValidators, candidateShardWaitingForNextRandomStr...)

	return commonPoolValidators, nil
}

func (b *BeaconCommitteeStateV1) getAllSubstituteCommittees() ([]string, error) {
	validators := []string{}

	for _, committee := range b.shardCommittee {
		committeeStr, err := incognitokey.CommitteeKeyListToString(committee)
		if err != nil {
			return nil, err
		}
		validators = append(validators, committeeStr...)
	}
	for _, substitute := range b.shardSubstitute {
		substituteStr, err := incognitokey.CommitteeKeyListToString(substitute)
		if err != nil {
			return nil, err
		}
		validators = append(validators, substituteStr...)
	}

	beaconCommittee := b.beaconCommittee
	beaconCommitteeStr, err := incognitokey.CommitteeKeyListToString(beaconCommittee)
	if err != nil {
		return nil, err
	}
	validators = append(validators, beaconCommitteeStr...)
	beaconSubstitute := b.beaconSubstitute
	beaconSubstituteStr, err := incognitokey.CommitteeKeyListToString(beaconSubstitute)
	if err != nil {
		return nil, err
	}
	validators = append(validators, beaconSubstituteStr...)

	candidateBeaconWaitingForCurrentRandom := b.currentEpochBeaconCandidate
	candidateBeaconWaitingForCurrentRandomStr, err := incognitokey.CommitteeKeyListToString(candidateBeaconWaitingForCurrentRandom)
	if err != nil {
		return nil, err
	}
	validators = append(validators, candidateBeaconWaitingForCurrentRandomStr...)
	candidateShardWaitingForCurrentRandom := b.currentEpochShardCandidate
	candidateShardWaitingForCurrentRandomStr, err := incognitokey.CommitteeKeyListToString(candidateShardWaitingForCurrentRandom)
	if err != nil {
		return nil, err
	}
	validators = append(validators, candidateShardWaitingForCurrentRandomStr...)

	return validators, nil
}

func (b *BeaconCommitteeStateV1) processUnstakeChange(committeeChange *CommitteeChange, env *BeaconCommitteeStateEnvironment) (*CommitteeChange, error) {

	newCommitteeChange := committeeChange

	unstakingIncognitoKey, err := incognitokey.CommitteeBase58KeyListToStruct(newCommitteeChange.Unstake)
	if err != nil {
		return newCommitteeChange, err
	}
	err = statedb.StoreStakerInfoV1(
		env.ConsensusStateDB,
		unstakingIncognitoKey,
		b.rewardReceiver,
		b.autoStake,
		b.stakingTx,
	)

	return newCommitteeChange, err
}