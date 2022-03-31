// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"fmt"
	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
	"github.com/racingtimegames/rt-smart-contracts/cadence/examples/go/lib"
	"github.com/racingtimegames/rt-smart-contracts/cadence/examples/go/marsart"
)

var mintArt = fmt.Sprintf(`
import NonFungibleToken from %s
import Marsart from %s
transaction(recipient: Address,
            name: String,
            artist: String,
            artistIntroduction: String,
            artworkIntroduction: String,
            typeId: UInt64,
            type: String,
            description: String,
            ipfsLink: String,
            MD5Hash: String,
            serialNumber: UInt32,
            totalNumber: UInt32 ) {
    let minter: &Marsart.NFTMinter

    prepare(signer: AuthAccount) {
        self.minter = signer.borrow<&Marsart.NFTMinter>(from: Marsart.MinterStoragePath)
            ?? panic("Could not borrow a reference to the NFT minter")
    }
    execute {
        let recipient = getAccount(recipient)
        let receiver = recipient
            .getCapability(Marsart.CollectionPublicPath)!
            .borrow<&{NonFungibleToken.CollectionPublic}>()
            ?? panic("Could not get receiver reference to the NFT Collection")
        self.minter.mintNFT(recipient: receiver,
                            name: name,
                            artist: artist,
                            artistIntroduction: artistIntroduction,
                            artworkIntroduction: artworkIntroduction,
                            typeId: typeId,
                            type: type,
                            description: description,
                            ipfsLink: ipfsLink,
                            MD5Hash: MD5Hash,
                            serialNumber: serialNumber,
                            totalNumber: totalNumber )
    }
}`, marsart.NonFungibleTokenAddress, marsart.ContractOwnAddress)

type NFTInfo struct {
	Name                string
	Artist              string
	ArtistIntroduction  string
	ArtworkIntroduction string
	TypeId              uint64
	Type                string
	Description         string
	IpfsLink            string
	MD5Hash             string
	SerialNumber        uint32
	TotalNumber         uint32
}

func MintNFT(nft *NFTInfo) *flow.TransactionResult {

	referenceBlock, err := marsart.FlowClient.GetLatestBlock(marsart.Ctx, false)
	if err != nil {
		panic(err)
	}
	acctAddress, acctKey, signer := lib.ServiceAccount(marsart.FlowClient, marsart.SigAlgo, marsart.HashAlgo, marsart.KeyIndex, marsart.Address, marsart.PrivateKey)
	tx := flow.NewTransaction().
		SetScript([]byte(mintArt)).
		SetGasLimit(200).
		SetProposalKey(acctAddress, acctKey.Index, acctKey.SequenceNumber).
		SetReferenceBlockID(referenceBlock.ID).
		SetPayer(acctAddress).
		AddAuthorizer(acctAddress)

	if err := tx.AddArgument(cadence.NewAddress(flow.HexToAddress(marsart.Address))); err != nil {
		panic(err)
	}
	if err := tx.AddArgument(cadence.String(nft.Name)); err != nil {
		panic(err)
	}
	if err := tx.AddArgument(cadence.String(nft.Artist)); err != nil {
		panic(err)
	}
	if err := tx.AddArgument(cadence.String(nft.ArtistIntroduction)); err != nil {
		panic(err)
	}
	if err := tx.AddArgument(cadence.String(nft.ArtworkIntroduction)); err != nil {
		panic(err)
	}
	if err := tx.AddArgument(cadence.UInt64(nft.TypeId)); err != nil {
		panic(err)
	}
	if err := tx.AddArgument(cadence.String(nft.Type)); err != nil {
		panic(err)
	}
	if err := tx.AddArgument(cadence.String(nft.Description)); err != nil {
		panic(err)
	}
	if err := tx.AddArgument(cadence.String(nft.IpfsLink)); err != nil {
		panic(err)
	}
	if err := tx.AddArgument(cadence.String(nft.MD5Hash)); err != nil {
		panic(err)
	}
	if err := tx.AddArgument(cadence.UInt32(nft.SerialNumber)); err != nil {
		panic(err)
	}
	if err := tx.AddArgument(cadence.UInt32(nft.TotalNumber)); err != nil {
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
