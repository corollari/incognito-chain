package blockchain

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

//Network fixed params
const (
	// SHARD_BLOCK_VERSION is the current latest supported block version.
	VERSION                       = 1
	RANDOM_NUMBER                 = 3
	SHARD_BLOCK_VERSION           = 1
	DefaultMaxBlkReqPerPeer       = 900
	DefaultMaxBlkReqPerTime       = 900
	MinCommitteeSize              = 3                // min size to run bft
	DefaultBroadcastStateTime     = 6 * time.Second  // in second
	DefaultStateUpdateTime        = 8 * time.Second  // in second
	DefaultMaxBlockSyncTime       = 30 * time.Second // in second
	DefaultCacheCleanupTime       = 40 * time.Second // in second
	WorkerNumber                  = 5
	MAX_S2B_BLOCK                 = 30
	MAX_BEACON_BLOCK              = 5
	LowerBoundPercentForIncDAO    = 3
	UpperBoundPercentForIncDAO    = 10
	GetValidBlock                 = 20
	TestRandom                    = true
	NumberOfFixedBlockValidators  = 4
	BEACON_ID                     = -1         // CommitteeID of beacon chain, used for highway
	ValidateTimeForSpamRequestTxs = 1581565837 // GMT: Thursday, February 13, 2020 3:50:37 AM. From this time, block will be checked spam request-reward tx
	TransactionBatchSize          = 30
	SpareTime                     = 1000 // in mili-second
)

// burning addresses
const (
	burningAddress  = "15pABFiJVeh9D5uiQEhQX4SVibGGbdAVipQxBdxkmDqAJaoG1EdFKHBrNfs"
	burningAddress2 = "12RxahVABnAVCGP3LGwCn8jkQxgw7z1x14wztHzn455TTVpi1wBq9YGwkRMQg3J4e657AbAnCvYCJSdA9czBUNuCKwGSRQt55Xwz8WA"
)

// CONSTANT for network MAINNET
const (
	// ------------- Mainnet ---------------------------------------------
	Mainnet                 = 0x01
	MainetName              = "mainnet"
	MainnetDefaultPort      = "9333"
	MainnetGenesisBlockTime = "2019-10-29T00:00:00.000Z"
	MainnetEpoch            = 350
	MainnetRandomTime       = 175
	MainnetOffset           = 4
	MainnetSwapOffset       = 4
	MainnetAssignOffset     = 8

	MainNetShardCommitteeSize     = 32
	MainNetMinShardCommitteeSize  = 22
	MainNetBeaconCommitteeSize    = 32
	MainNetMinBeaconCommitteeSize = 7
	MainNetActiveShards           = 8
	MainNetStakingAmountShard     = 1750000000000 // 1750 PRV = 1750 * 10^9 nano PRV

	MainnetMinBeaconBlkInterval = 40 * time.Second //second
	MainnetMaxBeaconBlkCreation = 10 * time.Second //second
	MainnetMinShardBlkInterval  = 40 * time.Second //second
	MainnetMaxShardBlkCreation  = 10 * time.Second //second

	//board and proposal parameters
	MainnetBasicReward = 1386666000 //1.386666 PRV
	//MainETHContractAddressStr = "0x0261DB5AfF8E5eC99fBc8FBBA5D4B9f8EcD44ec7" // v2-main - mainnet, branch master-temp-B-deploy, support erc20 with decimals > 18
	//MainETHContractAddressStr               = "0x3c8ec94213f09A1575f773470830124dfb40042e"                                                              // v3-main - mainnet
	MainETHContractAddressStr               = "0x6CC3873C3ca91cf5500DaD8B1A2c620B4f20507c"                                                              // v4-main - mainnet
	MainnetIncognitoDAOAddress              = "12S32fSyF4h8VxFHt4HfHvU1m9KHvBQsab5zp4TpQctmMdWuveXFH9KYWNemo7DRKvaBEvMgqm4XAuq1a1R4cNk2kfUfvXR3DdxCho3" // community fund
	MainnetCentralizedWebsitePaymentAddress = "12Rvjw6J3FWY3YZ1eDZ5uTy6DTPjFeLhCK7SXgppjivg9ShX2RRq3s8pdoapnH8AMoqvUSqZm1Gqzw7rrKsNzRJwSK2kWbWf1ogy885"

	// relaying header chain
	MainnetBNBChainID        = "Binance-Chain-Tigris"
	MainnetBTCChainID        = "Bitcoin-Mainnet"
	MainnetBTCDataFolderName = "btcrelayingv7"

	// BNB fullnode
	MainnetBNBFullNodeHost     = "dataseed1.ninicoin.io"
	MainnetBNBFullNodeProtocol = "https"
	MainnetBNBFullNodePort     = "443"

	MainnetPortalFeeder = "12RwJVcDx4SM4PvjwwPrCRPZMMRT9g6QrnQUHD54EbtDb6AQbe26ciV6JXKyt4WRuFQVqLKqUUbb7VbWxR5V6KaG9HyFbKf6CrRxhSm"
	// ------------- end Mainnet --------------------------------------
)

// VARIABLE for mainnet
var PreSelectBeaconNodeMainnetSerializedPubkey = []string{}
var PreSelectBeaconNodeMainnetSerializedPaymentAddress = []string{}
var PreSelectShardNodeMainnetSerializedPubkey = []string{}
var PreSelectShardNodeMainnetSerializedPaymentAddress = []string{}

var SelectBeaconNodeMainnetSerializedPubkeyV2 = make(map[uint64][]string)
var SelectBeaconNodeMainnetSerializedPaymentAddressV2 = make(map[uint64][]string)
var SelectShardNodeMainnetSerializedPubkeyV2 = make(map[uint64][]string)
var SelectShardNodeMainnetSerializedPaymentAddressV2 = make(map[uint64][]string)
var MainnetReplaceCommitteeEpoch = []uint64{}

// END CONSTANT for network MAINNET

// CONSTANT for network TESTNET
const (
	Testnet                 = 0x16
	TestnetName             = "testnet"
	TestnetDefaultPort      = "9444"
	TestnetGenesisBlockTime = "2019-11-29T00:00:00.000Z"
	TestnetEpoch            = 100
	TestnetRandomTime       = 50
	TestnetOffset           = 1
	TestnetSwapOffset       = 1
	TestnetAssignOffset     = 2

	TestNetShardCommitteeSize     = 32
	TestNetMinShardCommitteeSize  = 4
	TestNetBeaconCommitteeSize    = 4
	TestNetMinBeaconCommitteeSize = 4
	TestNetActiveShards           = 8
	TestNetStakingAmountShard     = 1750000000000 // 1750 PRV = 1750 * 10^9 nano PRV

	TestNetMinBeaconBlkInterval = 10 * time.Second //second
	TestNetMaxBeaconBlkCreation = 8 * time.Second  //second
	TestNetMinShardBlkInterval  = 10 * time.Second //second
	TestNetMaxShardBlkCreation  = 6 * time.Second  //second

	//board and proposal parameters
	TestnetBasicReward                      = 400000000 //40 mili PRV
	TestnetETHContractAddressStr            = "0x98cD12B85Df38c621d37cD6222E262835a7E87B9"
	TestnetIncognitoDAOAddress              = "12S5Lrs1XeQLbqN4ySyKtjAjd2d7sBP2tjFijzmp6avrrkQCNFMpkXm3FPzj2Wcu2ZNqJEmh9JriVuRErVwhuQnLmWSaggobEWsBEci" // community fund
	TestnetCentralizedWebsitePaymentAddress = "12S5Lrs1XeQLbqN4ySyKtjAjd2d7sBP2tjFijzmp6avrrkQCNFMpkXm3FPzj2Wcu2ZNqJEmh9JriVuRErVwhuQnLmWSaggobEWsBEci"

	// relaying header chain
	TestnetBNBChainID        = "Binance-Chain-Ganges"
	TestnetBTCChainID        = "Bitcoin-Testnet"
	TestnetBTCDataFolderName = "btcrelayingv8"

	// BNB fullnode
	TestnetBNBFullNodeHost     = "data-seed-pre-0-s3.binance.org"
	TestnetBNBFullNodeProtocol = "https"
	TestnetBNBFullNodePort     = "443"
	TestnetPortalFeeder        = "12S2ciPBja9XCnEVEcsPvmCLeQH44vF8DMwSqgkH7wFETem5FiqiEpFfimETcNqDkARfht1Zpph9u5eQkjEnWsmZ5GB5vhc928EoNYH"
)

// CONSTANT for network TESTNET-2
const (
	Testnet2                 = 0x32
	Testnet2Name             = "testnet-2"
	Testnet2DefaultPort      = "9444"
	Testnet2GenesisBlockTime = "2020-08-11T00:00:00.000Z"
	Testnet2Epoch            = 100
	Testnet2RandomTime       = 50
	Testnet2Offset           = 1
	Testnet2SwapOffset       = 1
	Testnet2AssignOffset     = 2

	TestNet2ShardCommitteeSize     = 32
	TestNet2MinShardCommitteeSize  = 4
	TestNet2BeaconCommitteeSize    = 4
	TestNet2MinBeaconCommitteeSize = 4
	TestNet2ActiveShards           = 8
	TestNet2StakingAmountShard     = 1750000000000 // 1750 PRV = 1750 * 10^9 nano PRV

	TestNet2MinBeaconBlkInterval = 10 * time.Second //second
	TestNet2MaxBeaconBlkCreation = 8 * time.Second  //second
	TestNet2MinShardBlkInterval  = 10 * time.Second //second
	TestNet2MaxShardBlkCreation  = 6 * time.Second  //second

	//board and proposal parameters
	Testnet2BasicReward                      = 400000000 //40 mili PRV
	Testnet2ETHContractAddressStr            = "0x4A4cC6451D8B8645bDeE96e6a0967b471eC286DD"
	Testnet2IncognitoDAOAddress              = "12S5Lrs1XeQLbqN4ySyKtjAjd2d7sBP2tjFijzmp6avrrkQCNFMpkXm3FPzj2Wcu2ZNqJEmh9JriVuRErVwhuQnLmWSaggobEWsBEci" // community fund
	Testnet2CentralizedWebsitePaymentAddress = "12S5Lrs1XeQLbqN4ySyKtjAjd2d7sBP2tjFijzmp6avrrkQCNFMpkXm3FPzj2Wcu2ZNqJEmh9JriVuRErVwhuQnLmWSaggobEWsBEci"

	// relaying header chain
	Testnet2BNBChainID        = "Binance-Chain-Ganges"
	Testnet2BTCChainID        = "Bitcoin-Testnet-2"
	Testnet2BTCDataFolderName = "btcrelayingv9"

	// BNB fullnode
	Testnet2BNBFullNodeHost     = "data-seed-pre-0-s3.binance.org"
	Testnet2BNBFullNodeProtocol = "https"
	Testnet2BNBFullNodePort     = "443"
	Testnet2PortalFeeder        = "12S2ciPBja9XCnEVEcsPvmCLeQH44vF8DMwSqgkH7wFETem5FiqiEpFfimETcNqDkARfht1Zpph9u5eQkjEnWsmZ5GB5vhc928EoNYH"
)

// VARIABLE for testnet
var PreSelectBeaconNodeTestnetSerializedPubkey = []string{}
var PreSelectBeaconNodeTestnetSerializedPaymentAddress = []string{}
var PreSelectShardNodeTestnetSerializedPubkey = []string{}
var PreSelectShardNodeTestnetSerializedPaymentAddress = []string{}

// VARIABLE for testnet
var SelectBeaconNodeTestnetSerializedPubkeyV2 = make(map[uint64][]string)
var SelectBeaconNodeTestnetSerializedPaymentAddressV2 = make(map[uint64][]string)
var SelectShardNodeTestnetSerializedPubkeyV2 = make(map[uint64][]string)
var SelectShardNodeTestnetSerializedPaymentAddressV2 = make(map[uint64][]string)
var TestnetReplaceCommitteeEpoch = []uint64{}

var IsTestNet = true
var IsTestNet2 = false

func init() {
	if len(os.Args) > 0 && (strings.Contains(os.Args[0], "test") || strings.Contains(os.Args[0], "Test")) {
		return
	}
	var keyData []byte
	var keyDataV2 []byte
	var err error

	keyData, err = ioutil.ReadFile("keylist.json")
	if err != nil {
		panic(err)
	}

	keyDataV2, err = ioutil.ReadFile("keylist-v2.json")
	if err != nil {
		panic(err)
	}

	type AccountKey struct {
		PrivateKey     string
		PaymentAddress string
		// PubKey     string
		CommitteePublicKey string
	}

	type KeyList struct {
		Shard  map[int][]AccountKey
		Beacon []AccountKey
	}
	type KeyListV2 struct {
		Epoch  uint64
		Shard  map[int][]AccountKey
		Beacon []AccountKey
	}

	keylist := KeyList{}
	keylistV2 := []KeyListV2{}

	err = json.Unmarshal(keyData, &keylist)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(keyDataV2, &keylistV2)
	if err != nil {
		panic(err)
	}
	if IsTestNet {
		for i := 0; i < TestNetMinBeaconCommitteeSize; i++ {
			PreSelectBeaconNodeTestnetSerializedPubkey = append(PreSelectBeaconNodeTestnetSerializedPubkey, keylist.Beacon[i].CommitteePublicKey)
			PreSelectBeaconNodeTestnetSerializedPaymentAddress = append(PreSelectBeaconNodeTestnetSerializedPaymentAddress, keylist.Beacon[i].PaymentAddress)
		}

		for i := 0; i < TestNetActiveShards; i++ {
			for j := 0; j < TestNetMinShardCommitteeSize; j++ {
				PreSelectShardNodeTestnetSerializedPubkey = append(PreSelectShardNodeTestnetSerializedPubkey, keylist.Shard[i][j].CommitteePublicKey)
				PreSelectShardNodeTestnetSerializedPaymentAddress = append(PreSelectShardNodeTestnetSerializedPaymentAddress, keylist.Shard[i][j].PaymentAddress)
			}
		}
		for _, v := range keylistV2 {
			epoch := v.Epoch
			TestnetReplaceCommitteeEpoch = append(TestnetReplaceCommitteeEpoch, epoch)
			for i := 0; i < TestNetMinBeaconCommitteeSize; i++ {
				SelectBeaconNodeTestnetSerializedPubkeyV2[epoch] = append(SelectBeaconNodeTestnetSerializedPubkeyV2[epoch], v.Beacon[i].CommitteePublicKey)
				SelectBeaconNodeTestnetSerializedPaymentAddressV2[epoch] = append(SelectBeaconNodeTestnetSerializedPaymentAddressV2[epoch], v.Beacon[i].PaymentAddress)
			}
			for i := 0; i < TestNetActiveShards; i++ {
				for j := 0; j < TestNetMinShardCommitteeSize; j++ {
					SelectShardNodeTestnetSerializedPubkeyV2[epoch] = append(SelectShardNodeTestnetSerializedPubkeyV2[epoch], v.Shard[i][j].CommitteePublicKey)
					SelectShardNodeTestnetSerializedPaymentAddressV2[epoch] = append(SelectShardNodeTestnetSerializedPaymentAddressV2[epoch], v.Shard[i][j].PaymentAddress)
				}
			}
		}
	} else {
		GenesisParam = genesisParamsMainnetNew
		for i := 0; i < MainNetMinBeaconCommitteeSize; i++ {
			PreSelectBeaconNodeMainnetSerializedPubkey = append(PreSelectBeaconNodeMainnetSerializedPubkey, keylist.Beacon[i].CommitteePublicKey)
			PreSelectBeaconNodeMainnetSerializedPaymentAddress = append(PreSelectBeaconNodeMainnetSerializedPaymentAddress, keylist.Beacon[i].PaymentAddress)
		}
		for i := 0; i < MainNetActiveShards; i++ {
			for j := 0; j < MainNetMinShardCommitteeSize; j++ {
				PreSelectShardNodeMainnetSerializedPubkey = append(PreSelectShardNodeMainnetSerializedPubkey, keylist.Shard[i][j].CommitteePublicKey)
				PreSelectShardNodeMainnetSerializedPaymentAddress = append(PreSelectShardNodeMainnetSerializedPaymentAddress, keylist.Shard[i][j].PaymentAddress)
			}
		}
		for _, v := range keylistV2 {
			epoch := v.Epoch
			MainnetReplaceCommitteeEpoch = append(MainnetReplaceCommitteeEpoch, epoch)
			for i := 0; i < MainNetMinBeaconCommitteeSize; i++ {
				SelectBeaconNodeMainnetSerializedPubkeyV2[epoch] = append(SelectBeaconNodeMainnetSerializedPubkeyV2[epoch], v.Beacon[i].CommitteePublicKey)
				SelectBeaconNodeMainnetSerializedPaymentAddressV2[epoch] = append(SelectBeaconNodeMainnetSerializedPaymentAddressV2[epoch], v.Beacon[i].PaymentAddress)
			}
			for i := 0; i < MainNetActiveShards; i++ {
				for j := 0; j < MainNetMinShardCommitteeSize; j++ {
					SelectShardNodeMainnetSerializedPubkeyV2[epoch] = append(SelectShardNodeMainnetSerializedPubkeyV2[epoch], v.Shard[i][j].CommitteePublicKey)
					SelectShardNodeMainnetSerializedPaymentAddressV2[epoch] = append(SelectShardNodeMainnetSerializedPaymentAddressV2[epoch], v.Shard[i][j].PaymentAddress)
				}
			}
		}
	}
}

// For shard
// public key

// END CONSTANT for network TESTNET

// -------------- FOR INSTRUCTION --------------
// Action for instruction
const (
	SetAction     = "set"
	SwapAction    = "swap"
	RandomAction  = "random"
	StakeAction   = "stake"
	AssignAction  = "assign"
	StopAutoStake = "stopautostake"
)
