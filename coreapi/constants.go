package coreapi

const (
	ContractDeployUrl                     string = "evm/contract/deploy"
	NewWalletUrl                          string = "evm/wallet/custodied/new" //New wallet url
	AccountZeroAddressUrl                 string = "evm/wallet/custodied/address"
	UsersRegisteredWithCustodiedWalletUrl string = "evm/wallet/custodied/list"
	WalletSeedPhraseUrl                   string = "evm/wallet/custodied/user/wallet"
	CustodiedWalletCoreMutableUrl         string = "evm/wallet/custodied/transact/mutable"
	CoreMutableUrl                        string = "evm/transact/mutable"
	TransferNativeNetworkCryptoUrl        string = "evm/chain/transfer"
	CoreViewUrl                           string = "evm/transact/view"
	PrepareWeb3MessageUrl                 string = "login/web3/prepare"
	ValidateWeb3SignatureUrl              string = "login/web3/validate?nonce=%s&signature=%s"
	GetChainBalanceUrl                    string = "evm/chain/%s/balance?network=%d"
	GetBlockchains                        string = "evm/info/chains"
	GetGasPrice                           string = "evm/network/gasprice?network=%d"
)

const (
	CONTENT_TYPE   string = "Content-Type"
	AUTH_APP_ID    string = "x-application-vkn"
	AUTHORIZATION  string = "Authorization"
	MIME_TYPE_JSON string = "application/json; charset=UTF-8"
)
