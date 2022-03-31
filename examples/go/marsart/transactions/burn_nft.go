// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"fmt"
	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
	"github.com/racingtimegames/rt-smart-contracts/cadence/examples/go/lib"
	"github.com/racingtimegames/rt-smart-contracts/cadence/examples/go/marsart"
)

var (
	burn = fmt.Sprintf(`
	import NonFungibleToken from %s
	import RacingTime from %s

	transaction(burnID: UInt64) {
    prepare(signer: AuthAccount) {
		let collection <- signer.load<@RacingTime.Collection>(from: RacingTime.CollectionStoragePath)!
		
		let nft <- collection.withdraw(withdrawID: burnID)
		
		destroy nft
		
		signer.save(<-collection, to: RacingTime.CollectionStoragePath)
    }
}`, marsart.NonFungibleTokenAddress, marsart.ContractOwnAddress)
)

func BurnNFT(nftId uint64) *flow.TransactionResult {
	referenceBlock, err := marsart.FlowClient.GetLatestBlock(marsart.Ctx, false)
	if err != nil {
		panic(err)
	}

	acctAddress, acctKey, signer := lib.ServiceAccount(marsart.FlowClient, marsart.SigAlgo, marsart.HashAlgo, marsart.KeyIndex, marsart.Address, marsart.PrivateKey)

	tx := flow.NewTransaction().
		SetScript([]byte(burn)).
		SetGasLimit(100).
		SetProposalKey(acctAddress, acctKey.Index, acctKey.SequenceNumber).
		SetReferenceBlockID(referenceBlock.ID).
		SetPayer(acctAddress).
		AddAuthorizer(acctAddress)

	if err := tx.AddArgument(cadence.NewUInt64(nftId)); err != nil {
		panic(err)
	}

	if err := tx.SignEnvelope(acctAddress, acctKey.Index, signer); err != nil {
		panic(err)
	}

	if err := marsart.FlowClient.SendTransaction(marsart.Ctx, *tx); err != nil {
		panic(err)
	}

	return lib.WaitForSeal(marsart.Ctx, marsart.FlowClient, tx.ID())
}
