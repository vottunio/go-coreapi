package coreapi

import (
	"math/big"
)

type NewWalletRequestDTO struct {
	User string `json:"userEmail"`
	Pin  string `json:"pin"`
}

type AccountZeroResponseDTO struct {
	AccountAddress string `json:"accountAddress"`
}

type AbiMutableRequestDTO struct {
	ContractAddress string        `json:"contractAddress"`
	Sender          string        `json:"sender"`
	Recipient       *string       `json:"recipient,omitempty"`
	Method          string        `json:"method"`
	Nonce           *uint64       `json:"nonce"`
	Network         uint64        `json:"blockchainNetwork"`
	Gas             uint64        `json:"gas"`
	GasPrice        *big.Int      `json:"gasPrice"`
	Value           uint64        `json:"value"`
	Params          []interface{} `json:"params"`
	SpecsID         *uint64       `json:"contractSpecsId,omitempty"`
	Pin             *string       `json:"pin,omitempty"`
}

type AbiMutableResponseDTO struct {
	TxHash string `json:"txHash"`
	Nonce  uint64 `json:"nonce"`
}

type AccountZeroRequestDTO struct {
	Email string `json:"userEmail"`
}

type UserWalletSeedRequestDTO struct {
	UserEmail string `json:"userEmail"`
	Pin       string `json:"pin"`
}

type UserWalletSeedResponseDTO struct {
	WalletSeedPhrase string `json:"walletSeedPhrase"`
}

type UserInfotDTO struct {
	ID                string `json:"id"`
	UserEmail         string `json:"userEmail"`
	AccountHash       string `json:"accountHash"`
	CreationTimestamp int64  `json:"creationTimestamp"`
}

type AbiViewOptionsDTO struct {
	ContractAddress string        `json:"contractAddress"`
	Method          string        `json:"method"`
	Network         uint64        `json:"blockchainNetwork"`
	Params          []interface{} `json:"params"`
}
