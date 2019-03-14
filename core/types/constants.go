package types

const (

	// Misc
	ShardCount                 = 1024 // (= 1,024)
	TargetCommitteeSize        = 128  // (= 128)
	MaxBalanceChurnQuotient    = 32   // (= 32)
	MaxIndicesPerSlashableVote = 4096 // (= 4,096)
	MaxExitDequeusPerEpoch     = 4    //(= 4)
	ShuffleRoundCount          = 90

	// Deposit contract
	//  TBD: DepositContractAddress
	DepositContractTreeDepth = 32 // (= 32)

	// Gwei values
	MinDepositAmountGwei           = 1000000000  // (= 1,000,000,000) 	Gwei
	MaxDepositAmountGwei           = 32000000000 // (= 32,000,000,000) 	Gwei
	ForkChoiceBalanceIncrementGwei = 1000000000  // (= 1,000,000,000) 	Gwei
	EjectionBalanceGwei            = 16000000000 //(= 16,000,000,000) 	Gwei

	//Time parameters
	SecondPerSlot                    = 6    // 6 seconds
	MinAttestationInclusiondDelay    = 4    //(= 4) 	slots 	24 seconds
	SlotsPerEpoch                    = 64   // (= 64) 	slots 	6.4 minutes
	MinSeedLookahead                 = 1    // (= 1) 	epochs 	6.4 minutes
	ActivationExitDelay              = 4    // (= 4) 	epochs 	25.6 minutes
	EpochsPerEth1VotingPeriod        = 16   // (= 16) 	epochs 	~1.7 hours
	SlotsPerHistoricalRoot           = 8192 // (= 8,192) 	slots 	~13 hours
	MinValidatorWithdrawabilityDelay = 256  // (= 256) 	epochs 	~27 hours
	PersistentCommitteePeriod        = 2048 // (= 2,048) 	epochs 	9 days

	// State list lengths
	LatestRandaoMixesLength      = 8192 // epochs  ~36 days
	LatestActiveIndexRootsLength = 8192 // epochs  ~36 days
	LatestShasledExitLength      = 8192 // epochs  ~36 days

	//Reward and penalty quotients
	BaseRewardQuotient                 = 32        // 2**5 (= 32)
	WhistleblowerRewardQuotient        = 512       // 2**9 (= 512)
	AttestationInclusionRewardQuotient = 8         // 2**3 (= 8)
	InactivityPenaltyQuotient          = 0x1000000 // 2**24 (= 16,777,216)
	MinPenaltyQuotient                 = 32        // 2**5 (= 32)

	//Max transactions per block
	MaxProposesSlashings = 16  // 2**4 (= 16)
	MaxAttesterSlashings = 1   // 2**0 (= 1)
	MaxAttestations      = 128 // 2**7 (= 128)
	MaxDeposits          = 16  // 2**4 (= 16)
	MaxVoluntaryExits    = 16  // 2**4 (= 16)
	MaxTransfers         = 16  // 2**4 (= 16)
)
const (
	// Signature domains

	DomainBeaconBLock   = iota // 0
	DomainRandao        = iota // 1
	DomainAttestation   = iota // 2
	DomainDeposit       = iota // 3
	DomainVoluntaryExit = iota // 4
	DomainTransfer      = iota // 5
)

var (

	//Initial values
	GenesisForkVersion      = 0
	GenesisSlot             = 0x100000000 // 4294967296
	GenesisEpoch            = GenesisSlot / SlotsPerEpoch
	GenesisStartShard       = 0
	FarFutureEpoch          = 0xffffffffffffffff //2**64 - 1
	ZeroHash                [32]byte
	EmptySignature          [96]byte
	BlsWithdrawalPrefixByte = byte(0)
)
