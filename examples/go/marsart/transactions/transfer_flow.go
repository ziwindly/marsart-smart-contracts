// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"context"
	"fmt"
	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
	"github.com/racingtimegames/rt-smart-contracts/cadence/examples/go/lib"
	"github.com/racingtimegames/rt-smart-contracts/cadence/examples/go/marsart"
)

var transferFLowScript string = fmt.Sprintf(`
import FungibleToken from %s
import FlowToken from %s

transaction(amount: UFix64, recipient: Address) {
  let sentVault: @FungibleToken.Vault
  prepare(signer: AuthAccount) {
    let vaultRef = signer.borrow<&FlowToken.Vault>(from: /storage/flowTokenVault)
      ?? panic("failed to borrow reference to sender vault")

    self.sentVault <- vaultRef.withdraw(amount: amount)
  }

  execute {
    let receiverRef =  getAccount(recipient)
      .getCapability(/public/flowTokenReceiver)
      .borrow<&{FungibleToken.Receiver}>()
        ?? panic("failed to borrow reference to recipient vault")

    receiverRef.deposit(from: <-self.sentVault)
  }
}`, marsart.FungibleTokenAddress, marsart.FlowTokenAddress)

func TransferFlow(tAddress string, f string) *flow.TransactionResult {
	referenceBlock, err := marsart.FlowClient.GetLatestBlock(context.Background(), false)
	if err != nil {
		panic(err)
	}

	fmt.Println("referenceBlock.Height --- ", referenceBlock.Height)

	acctAddress, acctKey, signer := lib.ServiceAccount(marsart.FlowClient, marsart.SigAlgo, marsart.HashAlgo, marsart.KeyIndex, marsart.Address, marsart.PrivateKey)
	tx := flow.NewTransaction().
		SetScript([]byte(transferFLowScript)).
		SetGasLimit(100).
		SetProposalKey(acctAddress, acctKey.Index, acctKey.SequenceNumber).
		SetReferenceBlockID(referenceBlock.ID).
		SetPayer(acctAddress).
		AddAuthorizer(acctAddress)

	toAmount, err := cadence.NewUFix64(f)
	if err != nil {
		panic(err)
	}

	toAddress := cadence.NewAddress(flow.HexToAddress(tAddress))

	if err := tx.AddArgument(toAmount); err != nil {
		panic(err)
	}

	if err := tx.AddArgument(toAddress); err != nil {
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
