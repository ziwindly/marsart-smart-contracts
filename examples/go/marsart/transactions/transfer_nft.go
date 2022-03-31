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
	transferNft = fmt.Sprintf(`
	import NonFungibleToken from %s
	import Marsart from %s

	transaction(recipient: Address, withdrawID: UInt64) {
    prepare(signer: AuthAccount) {
        let recipient = getAccount(recipient)
        let collectionRef = signer.borrow<&Marsart.Collection>(from: Marsart.CollectionStoragePath)
            ?? panic("Could not borrow a reference to the owner's collection")
        let depositRef = recipient.getCapability(Marsart.CollectionPublicPath)!.borrow<&{NonFungibleToken.CollectionPublic}>()!
        let nft <- collectionRef.withdraw(withdrawID: withdrawID)
        depositRef.deposit(token: <-nft)
    }
}`, marsart.NonFungibleTokenAddress, marsart.ContractOwnAddress)
)

// TransferNFT This transaction transfers a NFT from one account to another.
func TransferNFT(toAddress string, nftId uint64) *flow.TransactionResult {
	referenceBlock, err := marsart.FlowClient.GetLatestBlock(marsart.Ctx, false)
	if err != nil {
		panic(err)
	}

	acctAddress, acctKey, signer := lib.ServiceAccount(marsart.FlowClient, marsart.SigAlgo, marsart.HashAlgo, marsart.KeyIndex, marsart.Address, marsart.PrivateKey)

	tx := flow.NewTransaction().
		SetScript([]byte(transferNft)).
		SetGasLimit(100).
		SetProposalKey(acctAddress, acctKey.Index, acctKey.SequenceNumber).
		SetReferenceBlockID(referenceBlock.ID).
		SetPayer(acctAddress).
		AddAuthorizer(acctAddress)

	if err := tx.AddArgument(cadence.NewAddress(flow.HexToAddress(toAddress))); err != nil {
		panic(err)
	}

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
