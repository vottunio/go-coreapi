package coreapi

import (
	"math/big"

	"github.com/vottun-com/utils/types"
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

// DTO used to get the wallet first account (the first one derived using seed)
type AccountZeroRequestDTO struct {
	// The user email
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

type MessageWeb3DTO struct {
	Domain         string         `json:"domain"`
	WalletAddress  string         `json:"walletAddress"`
	ChainID        uint16         `json:"chainId"`
	Statement      string         `json:"statement"`
	Uri            string         `json:"uri"`
	IssuedAt       int64          `json:"issuedAt"`
	Version        *string        `json:"version:omitempty"`
	Nonce          *types.SqlUuid `json:"Nonce:omitempty"`
	ExpirationTime *int64         `json:"expirationTime:omitempty"`
}

type MessageWeb3ResponseDTO struct {
	Domain string `json:"message"`
	Nonce  string `json:"nonce"`
}
