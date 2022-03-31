// SPDX-License-Identifier: Apache-2.0

package t_test

import (
	"github.com/onflow/flow-go-sdk/crypto"
	"github.com/racingtimegames/rt-smart-contracts/cadence/examples/go/lib"
	"github.com/racingtimegames/rt-smart-contracts/cadence/examples/go/marsart"
	"github.com/racingtimegames/rt-smart-contracts/cadence/examples/go/marsart/scripts"
	"github.com/racingtimegames/rt-smart-contracts/cadence/examples/go/marsart/transactions"
	"testing"
)

func TestMarsartDeployContracts(t *testing.T) {
	marsart.Address = ""
	marsart.PrivateKey = ""
	marsart.SigAlgo = crypto.ECDSA_P256
	marsart.HashAlgo = crypto.SHA3_256
	marsart.KeyIndex = 0

	_result := transactions.DeployContract()
	if _result.Error != nil {
		println(_result.Error.Error())
	}
	println(_result.Status)
}

func TestCreateAccount(t *testing.T) {
	seed := ""

	marsart.Address = ""
	marsart.PrivateKey = ""
	marsart.SigAlgo = crypto.ECDSA_P256
	marsart.HashAlgo = crypto.SHA3_256
	marsart.KeyIndex = 0

	acc, prv := transactions.CreateAddress(seed, crypto.ECDSA_secp256k1, crypto.SHA3_256)

	println("New Account Address:", acc)
	println("New Account Signature Algorithm:", crypto.ECDSA_secp256k1.String())
	println("New Account Hash Algorithm:", crypto.SHA3_256.String())
	println("New Account Private Key:", prv)
}

func TestTransferFlow(t *testing.T) {
	marsart.Address = ""
	marsart.PrivateKey = ""
	marsart.SigAlgo = crypto.ECDSA_P256
	marsart.HashAlgo = crypto.SHA3_256
	marsart.KeyIndex = 0

	ToAddress := ""
	Flow := "10.000"
	_result := transactions.TransferFlow(ToAddress, Flow)
	if _result.Error != nil {
		println(_result.Error.Error())
	}
	println(_result)
}

func TestSetupAccount(t *testing.T) {
	marsart.Address = ""
	marsart.PrivateKey = ""
	marsart.SigAlgo = crypto.ECDSA_secp256k1
	marsart.HashAlgo = crypto.SHA3_256
	marsart.KeyIndex = 0

	_result := transactions.SetupAccount()
	if _result.Error != nil {
		println(_result.Error.Error())
	}
	println(_result.Status)
}

func TestMarsartUpdateContracts(t *testing.T) {
	marsart.Address = ""
	marsart.PrivateKey = ""
	marsart.SigAlgo = crypto.ECDSA_secp256k1
	marsart.HashAlgo = crypto.SHA3_256
	marsart.KeyIndex = 0

	_result := transactions.UpdateContract()
	if _result.Error != nil {
		println(_result.Error.Error())
	}
	println(_result.Status)
}

func TestMintNFT(t *testing.T) {
	marsart.Address = ""
	marsart.PrivateKey = ""
	marsart.SigAlgo = crypto.ECDSA_P256
	marsart.HashAlgo = crypto.SHA3_256
	marsart.KeyIndex = 0

	//test nft info
	nft := &transactions.NFTInfo{
		Name:                "",
		Artist:              "",
		ArtistIntroduction:  "",
		ArtworkIntroduction: "",
		TypeId:              1,
		Type:                "",
		Description:         "",
		IpfsLink:            "",
		MD5Hash:             "",
		SerialNumber:        4,
		TotalNumber:         10}

	_result := transactions.MintNFT(nft)
	if _result.Error != nil {
		println(_result.Error.Error())
	}
	println(_result)
}

func TestAccountNFT(t *testing.T) {
	searchAddress := ""

	_result := scripts.AccountInfo(searchAddress)

	println(lib.CadenceValueToJsonString(_result))
}

func TestTransferNFT(t *testing.T) {
	marsart.Address = ""
	marsart.PrivateKey = ""
	marsart.SigAlgo = crypto.ECDSA_P256
	marsart.HashAlgo = crypto.SHA3_256
	marsart.KeyIndex = 0

	ToAddress := ""
	NFTId := uint64(1)

	_result := transactions.TransferNFT(ToAddress, NFTId)
	if _result.Error != nil {
		println(_result.Error.Error())
	}
	println(_result)
}
