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

type Fork struct {
	previousVersion [4]byte `ssz:"bytes4"` // Previous fork version
	currentVersion  [4]byte `ssz:"bytes4"` // Current fork version
	epoch           uint64  `ssz:"uint64"` // Fork epoch number
}

type Crosslink struct {
	epoch             uint64   `ssz:"uint64"`  // Epoch number
	crosslinkDataRoot [32]byte `ssz:"bytes32"` // Shard data since the previous crosslink
}

type Eth1Data struct {
	depositRoot [32]byte `ssz:"bytes32"` // Root of the deposit tree
	blockHash   [32]byte `ssz:"bytes32"` // Block hash
}

type Eth1DataVote struct {
	eth1Data  Eth1Data `ssz:"container"` // Data being voted for
	voteCount uint64   `ssz:"uint64"`    // Vote count
}

type AttestationData struct {
	// LMD GHOST vote
	slot            uint64   `ssz:"uint64"`
	beaconBlockRoot [32]byte `ssz:"bytes32"`

	// FFG vote
	sourceEpoch uint64   `ssz:"uint64"`
	sourceRoot  [32]byte `ssz:"bytes32"`
	targetRoot  [32]byte `ssz:"bytes32"`

	// Crosslink vote
	shard             uint64    `ssz:"uint64"`
	previousCrosslink Crosslink `ssz:"container"`
	crosslinkDataRoot [32]byte  `ssz:"bytes32"`
}

type AttestationDataAndCustodyBit struct {
	data       AttestationData `ssz:"container"` // Attestation data
	custodyBit bool            `ssz:"bool"`      // Custody bit
}

type SlashableAttestation struct {
	validatorIndices   []uint64        `ssz:"list:uint64"` // Validator indices
	data               AttestationData `ssz:"container"`   // Attestation data
	custodyBitfield    []byte          `ssz:"bytes"`       // Custody bitfield
	aggregateSignature [96]byte        `ssz:"bytes96"`     // Aggregate signature
}

type DepositInput struct {
	pubkey                [48]byte `ssz:"bytes48"` // BLS pubkey
	withdrawalCredentials [32]byte `ssz:"bytes32"` // Withdrawal credentials
	proofOfPossession     [96]byte `ssz:"bytes96"` // A BLS signature of this `DepositInput`
}

type DepositData struct {
	amount       uint64       `ssz:"uint64"`    // Amount in Gwei
	timestamp    uint64       `ssz:"uint64"`    // Timestamp from deposit contract
	depositInput DepositInput `ssz:"container"` // Deposit input
}

type BeaconBlockHeader struct {
	slot                uint64   `ssz:"uint64"`
	previous_block_root [32]byte `ssz:"bytes32"`
	state_root          [32]byte `ssz:"bytes32"`
	blockBodyRoot       [32]byte `ssz:"bytes32"`
	signature           [96]byte `ssz:"bytes96"`
}

type Validator struct {
	pubkey                [48]byte `ssz:"bytes48"` // BLS public key
	withdrawalCredentials [32]byte `ssz:"bytes32"` // Withdrawal credentials
	activationEpoch       uint64   `ssz:"uint64"`  // Epoch when validator activated
	exitEpoch             uint64   `ssz:"uint64"`  // Epoch when validator exited
	withdrawableEpoch     uint64   `ssz:"uint64"`  // Epoch when validator is eligible to withdraw
	initiatedExit         bool     `ssz:"bool"`    // Did the validator initiate an exit
	slashed               bool     `ssz:"bool"`    // Was the validator slashed
}

type PendingAttestation struct {
	aggregationBitfield []byte          `ssz:"bytes"`     // Attester aggregation bitfield
	data                AttestationData `ssz:"container"` // Attestation data
	custodyBitfield     bytes           `ssz:"bytes"`     // Custody bitfield
	inclusionSlot       uint64          `ssz:"uint64"`    // Inclusion slot
}
type root [32]byte //`ssz:"bytes32"`
type HistoricalBatch struct {
	blockRoots [SlotsPerHistoricalRoot]root // Block roots
	stateRoots [SlotsPerHistoricalRoot]root // State roots
}
