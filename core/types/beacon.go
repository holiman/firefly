//Copyright (c) 2019, The go-ethereum Authors. All rights reserved.
// This file is part of the firefly library.
//
//Redistribution and use in source and binary forms, with or without
//modification, are permitted provided that the following conditions are met:
//* Redistributions of source code must retain the above copyright
//notice, this list of conditions and the following disclaimer.
//* Redistributions in binary form must reproduce the above copyright
//notice, this list of conditions and the following disclaimer in the
//documentation and/or other materials provided with the distribution.
//* Neither the name of The go-ethereum Authors nor the
//names of its contributors may be used to endorse or promote products
//derived from this software without specific prior written permission.
//
// Package types contains data types related to Ethereum 2.0 consensus.
// See https://github.com/ethereum/eth2.0-specs/blob/dev/specs/core/0_beacon-chain.md//data-structures
package types

// Beacon transactions

type ProposerSlashing struct {
	proposerIndex uint64            `ssz:"uint64"`    // Proposer index
	header1       BeaconBlockHeader `ssz:"container"` // First block header
	header2       BeaconBlockHeader `ssz:"container"` // Second block header
}

type AttesterSlashing struct {
	slashableAttestation1 SlashableAttestation `ssz:"container"` // First slashable attestation
	slashableAttestation2 SlashableAttestation `ssz:"container"` // Second slashable attestation
}

type Attestation struct {
	aggregationBitfield []byte          `ssz:"bytes"`     // Attester aggregation bitfield
	data                AttestationData `ssz:"container"` // Attestation data
	custodyBitfield     bytes           `ssz:"bytes"`     // Custody bitfield
	aggregateSignature  [96]byte        `ssz:"bytes96"`   // BLS aggregate signature
}

type Deposit struct {
	proof       [DepositContractTreeDepth][32]byte `ssz:"list32:bytes32"` // Branch in the deposit tree
	index       uint64                             `ssz:"uint64"`         // Index in the deposit tree
	depositData DepositData                        `ssz:"container"`      // Data
}

type VoluntaryExit struct {
	epoch          uint64   `ssz:"uint64"`  // Minimum epoch for processing exit
	validatorIndex uint64   `ssz:"uint64"`  // Index of the exiting validator
	signature      [96]byte `ssz:"bytes96"` // Validator signature
}

type Transfer struct {
	sender    uint64   `ssz:"uint64"`  // Sender index
	recipient uint64   `ssz:"uint64"`  // Recipient index
	amount    uint64   `ssz:"uint64"`  // Amount in Gwei
	fee       uint64   `ssz:"uint64"`  // Fee in Gwei for block proposer
	slot      uint64   `ssz:"uint64"`  // Inclusion slot
	pubkey    [48]byte `ssz:"bytes48"` // Sender withdrawal pubkey
	signature [96]byte `ssz:"bytes96"` // Sender signature
}

// Beacon blocks

type BeaconBlockBody struct {
	randaoReveal      [96]byte           `ssz:"bytes96"`
	eth1Data          Eth1Data           `ssz:"container"`
	proposerSlashings []ProposerSlashing `ssz:"list:container"`
	attesterSlashings []AttesterSlashing `ssz:"list:container"`
	attestations      []Attestation      `ssz:"list:container"`
	deposits          []Deposit          `ssz:"list:container"`
	voluntary_exits   []VoluntaryExit    `ssz:"list:container"`
	transfers         []Transfer         `ssz:"list:container"`
}

type BeaconBlock struct {
	// Header
	slot              uint64          `ssz:"uint64"`
	previousBlockRoot root            `ssz:"bytes32"`
	stateRoot         root            `ssz:"bytes32"`
	body              BeaconBlockBody `ssz:"container"`
	signature         [96]byte        `ssz:"bytes96"`
}

//   Beacon state

type BeaconState struct {
	// Misc
	slot        uint64 `ssz:"uint64"`
	genesisTime uint64 `ssz:"uint64"`
	fork        Fork   `ssz:"container"` // For versioning hard forks

	// Validator registry
	validatorRegistry            []Validator `ssz:"list:container"`
	validatorBalances            []uint64    `ssz:"uint64"`
	validatorRegistryUpdateEpoch uint64      `ssz:"uint64"`

	// Randomness and committees
	latestRandaoMixes           [LatestRandaoMixesLength][32]byte `ssz:"bytes32"`
	previousShufflingStartShard uint64                            `ssz:"uint64"`
	currentShufflingStartShard  uint64                            `ssz:"uint64"`
	previousShufflingEpoch      uint64                            `ssz:"uint64"`
	currentShufflingEpoch       uint64                            `ssz:"uint64"`
	previousShufflingSeed       [32]byte                          `ssz:"bytes32"`
	currentShufflingSeed        [32]byte                          `ssz:"bytes32"`

	// Finality
	previousEpochAttestations []PendingAttestation `ssz:"list:container"`
	currentEpochAttestations  []PendingAttestation `ssz:"list:container"`
	previousJustifiedEpoch    uint64               `ssz:"uint64"`
	currentJustifiedEpoch     uint64               `ssz:"uint64"`
	previousJustifiedRoot     root                 `ssz:"bytes32"`
	currentJustifiedRoot      root                 `ssz:"bytes32"`
	justificationBitfield     uint64               `ssz:"uint64"`
	finalizedEpoch            uint64               `ssz:"uint64"`
	finalizedRoot             [32]byte             `ssz:"bytes32"`

	// Recent state
	latestCrosslinks [ShardCount]Crosslink

	latestBlockRoots       [SlotsPerHistoricalRoot]root       `ssz:"list:container,8192"`
	latestStateRoots       [SlotsPerHistoricalRoot]root       `ssz:"list:container,8192"`
	latestActiveIndexRoots [LatestActiveIndexRootsLength]root `ssz:"list:container,8192"`
	// Balances slashed at every withdrawal period
	latestSlashedBalances [LatestSlashedExitLength]uint64 `ssz:"uint64"`
	// `latest_block_header.state_root == ZERO_HASH` temporarily
	latestBlockHeader BeaconBlockHeader `ssz:"container"`
	historicalRoots   []root            `ssz:"list:bytes32"`

	// Ethereum 1.0 chain data
	latestEth1Data Eth1Data       `ssz:"container"`
	eth1DataVotes  []Eth1DataVote `ssz:"list:container"`
	depositIndex   uint64         `ssz:"uint64"`
}
