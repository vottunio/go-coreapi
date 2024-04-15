package coreapi

import (
	"math/big"

	"github.com/vottun-com/ethereum/v2"
	"github.com/vottunio/go-coreapi/types"
)

type NewWalletRequestDTO struct {
	User string `json:"userEmail"`
	Pin  string `json:"pin"`
}

type AccountZeroResponseDTO struct {
	AccountAddress string `json:"accountAddress"`
}

type ContractDeployRequestDTO struct {
	ContractSpecsID uint64        `json:"contractSpecsId"`
	WalletAddress   string        `json:"sender,omitempty"`
	Network         uint64        `json:"blockchainNetwork"`
	GasLimit        uint64        `json:"gasLimit"`
	GasPrice        *big.Int      `json:"gasPrice,omitempty"`
	Nonce           *uint64       `json:"nonce,omitempty"`
	Alias           *string       `json:"alias,omitempty"`
	Params          []interface{} `json:"params,omitempty"`
}

type ContractDeployResponseDTO struct {
	ContractAddress string `json:"contractAddress"`
	TxHash          string `json:"txHash"`
}

type AbiMutableRequestDTO struct {
	ContractAddress *string       `json:"contractAddress,omitempty"`
	Sender          string        `json:"sender"`
	Recipient       *string       `json:"recipient,omitempty"`
	Method          *string       `json:"method,omitempty"`
	Nonce           *uint64       `json:"nonce,omitempty"`
	Network         uint64        `json:"blockchainNetwork"`
	Gas             *uint64       `json:"gas,omitempty"`
	GasPrice        *big.Int      `json:"gasPrice,omitempty"`
	Value           *big.Int      `json:"value,omitempty"`
	Params          []interface{} `json:"params,omitempty"`
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
	SpecsID         *uint64       `json:"contractSpecsId,omitempty"`
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
	Message string `json:"message"`
	Nonce   string `json:"nonce"`
}

type GetChainBalanceResponseDTO struct {
	Balance float64 `json:"balance"`
}
type BlockchainClientDTO struct {
	ID                    uint64  `json:"id"`
	Name                  string  `json:"name"`
	TokenSymbol           string  `json:"symbol"`
	Mainnet               bool    `json:"isMainnet"`
	Explorer              *string `json:"explorer"`
	TestnetFaucet         *string `json:"testnetFaucet"`
	TypeID                uint64  `json:"typeId"`
	ActiveForTransactions bool    `json:"activeForTransactions"`
	TypeName              string  `json:"typeName"`
}

type ChainsListDTO struct {
	Mainnet []BlockchainClientDTO `json:"mainnetNetworks"`
	Testnet []BlockchainClientDTO `json:"testnetNetworks"`
}

type GasPriceRequestDTO struct {
	Network uint64 `json:"blockchainNetwork"`
}

type GasPriceResponseDTO struct {
	GasPrice float64 `json:"gasPriceGwei"`
}

type BlockchainTransactionDTO struct {
	Network   uint64                       `json:"network"`
	Tx        *ethereum.Transaction        `json:"transaction"`
	Receipt   *ethereum.TransactionReceipt `json:"receipt"`
	Fees      *TransactionFees             `json:"transactionFees,omitempty"`
	Error     bool                         `json:"error"`
	ErrorInfo ErrorDTO                     `json:"errorInfo"`
}

type TransactionFees struct {
	Currency string `json:"currency"`
	// TodayValue  float64 `json:"todayPrice"`
	// ValueAtDate float64 `json:"priceAtTxDate"`
	VottunFee float64 `json:"fee"`
}

type ErrorDTO struct {
	Code string `json:"code"`

	Message string `json:"message"`
}
