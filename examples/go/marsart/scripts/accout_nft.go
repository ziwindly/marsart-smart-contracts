// SPDX-License-Identifier: Apache-2.0

package scripts

import (
	"context"
	"fmt"
	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
	"github.com/racingtimegames/rt-smart-contracts/cadence/examples/go/marsart"
)

var getNftInfo string = fmt.Sprintf(`
import Marsart from %s

pub struct MarsartNFTData {
	pub let id: UInt64
	pub let name: String
	pub let artist: String
	pub let artistIntroduction: String
	pub let artworkIntroduction: String
	pub let typeId: UInt64
	pub let type: String
	pub let description: String
	pub let ipfsLink: String
	pub let MD5Hash: String
	pub let serialNumber: UInt32
	pub let totalNumber: UInt32

	init(
		id: UInt64,
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
		totalNumber: UInt32){
		self.id=id
		self.name=name
		self.artist=artist
		self.artistIntroduction=artistIntroduction
		self.artworkIntroduction=artworkIntroduction
		self.typeId=typeId
		self.type=type
		self.description=description
		self.ipfsLink=ipfsLink
		self.MD5Hash=MD5Hash
		self.serialNumber=serialNumber
		self.totalNumber=totalNumber
	}
}

pub fun main(address:Address) : [MarsartNFTData] {
	var nfts: [MarsartNFTData] = []
    let account = getAccount(address)
	
	if let artCollection= account.getCapability(Marsart.CollectionPublicPath).borrow<&{Marsart.MarsartCollectionPublic}>()  {
		for id in artCollection.getIDs() {
			if let marsartFT = artCollection.borrowMarsart(id: id) {
				nfts.append(MarsartNFTData(id: id, 
						name: marsartFT.data.name,
						artist: marsartFT.data.artist,
						artistIntroduction: marsartFT.data.artistIntroduction,
						artworkIntroduction: marsartFT.data.artworkIntroduction,
						typeId: marsartFT.data.typeId,
						type: marsartFT.data.type,
						description: marsartFT.data.description,
						ipfsLink: marsartFT.data.ipfsLink,
						MD5Hash: marsartFT.data.MD5Hash,
						serialNumber: marsartFT.data.serialNumber,
						totalNumber: marsartFT.data.totalNumber ))           
			}
		}
	}
	
    return nfts
}`, marsart.ContractOwnAddress)

func AccountInfo(searchAddress string) cadence.Value {
	ctx := context.Background()
	result, err := marsart.FlowClient.ExecuteScriptAtLatestBlock(ctx, []byte(getNftInfo), []cadence.Value{cadence.NewAddress(flow.HexToAddress(searchAddress))})
	if err != nil {
		panic(err)
	}

	return result
}
