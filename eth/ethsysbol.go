package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

const (
	// HexPrefix 16进制前缀
	HexPrefix = "0x"

	// ERC20MethodTransfer erc20方法transfer的16进制字符串
	ERC20MethodTransfer = "a9059cbb"
	// ERC20TransferLength erc20方法transfer的input长度
	ERC20TransferLength = 130

	// ERC20MethodBalanceOf erc20方法balanceOf的16进制字符串
	ERC20MethodBalanceOf = "70a08231"

	// ERC20Name erc20 name 的16进制字符串
	ERC20Name = "0x06fdde03"
	// ERC20Symbol erc20 symbol 的16进制字符串
	ERC20Symbol = "0x95d89b41"
	// ERC20Decimals erc20 decimals 的16进制字符串
	ERC20Decimals = "0x313ce567"
	// ERC20AbiDefaultLength ERC20 Abi字符串 Default Length
	ERC20AbiDefaultLength = 194

	// ConfirmedNum confirmed num till to 12
	ConfirmedNum = int64(12)

	// EthGasLimit eth gas limit transaction gas limit
	EthGasLimit = int64(21000)
	// DefaultErc20GasPrice default gas price for erc20
	DefaultErc20GasPrice = int64(60000)

	// DefaultErc20Icon default erc20 icon
	DefaultErc20Icon = "https://ops.58wallet.io/home/img/avatar-DEFAULT@2x.png"

	//ReceiptStatusSuccessful Receipt Status Successful
	ReceiptStatusSuccessful = "0x1"

	// StatusPass status pass
	StatusPass = 0
	// StatusFailed status failed
	StatusFailed = 1
	// StatusPending status pending
	StatusPending = 2

	// EtherscanURL EtherscanURL
	EtherscanURL = "https://etherscan.io/tx/"
	// RopstenURL RopstenURL
	RopstenURL = "https://ropsten.etherscan.io/tx/"

	// EtcURL EtcURL
	EtcURL = "https://gastracker.io/tx/"

	// RopstenTestNode Ropsten test node
	RopstenTestNode = "http://116.62.148.208:8545"

	// EventTransferHash hash of Event Transfer
	EventTransferHash = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	// AddressInHashIndex Address in hash  from index
	AddressInHashIndex = 26

	// MyEtherWalletEndpointTestNet MyEtherWalletEndpointTestNet
	MyEtherWalletEndpointTestNet = "https://api.myetherapi.com/rop"
	// MyEtherWalletEndpoint MyEtherWalletEndpoint
	MyEtherWalletEndpoint = "https://api.myetherapi.com/eth"
)

func convertToASCII(abiString string) (string, error) {
	if len(abiString) < ERC20AbiDefaultLength {
		return "", nil
	}
	trimedAbi := strings.TrimRight(abiString[66+64:], "0")
	if len(trimedAbi)%2 != 0 {
		trimedAbi += "0"
	}
	res, err := hex.DecodeString(trimedAbi)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(res), nil
}

func main() {
	// constrast addr 0x937ea0eed76dc56a1862d473f6b5dbf6ea1df191
	abi := "0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000034f4d470000000000000000000000000000000000000000000000000000000000"

	fmt.Println(convertToASCII(abi))
}
