// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"fmt"
	"github.com/onflow/flow-go-sdk"
	"github.com/racingtimegames/rt-smart-contracts/cadence/examples/go/lib"
	"github.com/racingtimegames/rt-smart-contracts/cadence/examples/go/marsart"
)

var setupAccountContracts = fmt.Sprintf(`
import NonFungibleToken from %s
import Marsart from %s
transaction {
    prepare(signer: AuthAccount) {
        if signer.borrow<&Marsart.Collection>(from: Marsart.CollectionStoragePath) == nil {
            let collection <- Marsart.createEmptyCollection()
            
            signer.save(<-collection, to: Marsart.CollectionStoragePath)
            signer.link<&Marsart.Collection{NonFungibleToken.CollectionPublic, Marsart.MarsartCollectionPublic}>(Marsart.CollectionPublicPath, target: Marsart.CollectionStoragePath)
        }
    }
}`, marsart.NonFungibleTokenAddress, marsart.ContractOwnAddress)

func SetupAccount() *flow.TransactionResult {
	referenceBlock, err := marsart.FlowClient.GetLatestBlock(marsart.Ctx, false)
	if err != nil {
		panic(err)
	}

	acctAddress, acctKey, signer := lib.ServiceAccount(marsart.FlowClient, marsart.SigAlgo, marsart.HashAlgo, marsart.KeyIndex, marsart.Address, marsart.PrivateKey)

	tx := flow.NewTransaction().
		SetScript([]byte(setupAccountContracts)).
		SetGasLimit(100).
		SetProposalKey(acctAddress, acctKey.Index, acctKey.SequenceNumber).
		SetReferenceBlockID(referenceBlock.ID).
		SetPayer(acctAddress).
		AddAuthorizer(acctAddress)

	if err := tx.SignEnvelope(acctAddress, acctKey.Index, signer); err != nil {
		panic(err)
	}

	if err := marsart.FlowClient.SendTransaction(marsart.Ctx, *tx); err != nil {
		panic(err)
	}

	return lib.WaitForSeal(marsart.Ctx, marsart.FlowClient, tx.ID())
}
