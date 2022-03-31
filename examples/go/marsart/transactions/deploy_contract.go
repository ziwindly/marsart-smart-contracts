// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"context"
	"fmt"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/templates"
	"github.com/racingtimegames/rt-smart-contracts/cadence/examples/go/lib"
	"github.com/racingtimegames/rt-smart-contracts/cadence/examples/go/marsart"
	"io/ioutil"
)

func DeployContract() *flow.TransactionResult {

	referenceBlock, err := marsart.FlowClient.GetLatestBlock(context.Background(), false)
	if err != nil {
		panic(err)
	}

	serviceAcctAddr, serviceAcctKey, singer := lib.ServiceAccount(marsart.FlowClient, marsart.SigAlgo, marsart.HashAlgo, marsart.KeyIndex, marsart.Address, marsart.PrivateKey)

	contractPath := fmt.Sprintf("../../contracts/%s.cdc", marsart.ContractsName)

	code, err := ioutil.ReadFile(contractPath)
	if err != nil {
		panic(err)
	}

	tx := templates.AddAccountContract(serviceAcctAddr, templates.Contract{
		Name:   marsart.ContractsName,
		Source: string(code),
	})
	tx.SetProposalKey(
		serviceAcctAddr,
		serviceAcctKey.Index,
		serviceAcctKey.SequenceNumber,
	)
	tx.SetReferenceBlockID(referenceBlock.ID)
	tx.SetPayer(serviceAcctAddr)
	tx.SetGasLimit(9999)
	if err := tx.SignEnvelope(serviceAcctAddr, serviceAcctKey.Index, singer); err != nil {
		panic(err)
	}

	if err := marsart.FlowClient.SendTransaction(marsart.Ctx, *tx); err != nil {
		panic(err)
	}
	println(tx.ID().String())
	return lib.WaitForSeal(marsart.Ctx, marsart.FlowClient, tx.ID())
}
