package coreapi

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type CoreApi struct {
	RootUrl   string
	tokenAuth string
	appID     string
}

func New(tokenAuth, appID, rootUrl string) *CoreApi {
	if !strings.HasSuffix(rootUrl, "/") {
		rootUrl += "/"
	}

	return &CoreApi{tokenAuth: tokenAuth, appID: appID, RootUrl: rootUrl}
}

// Deploys a new contract
func (c *CoreApi) DeployNewContract(requestDto *ContractDeployRequestDTO) (*ContractDeployResponseDTO, error) {

	var responseDto *ContractDeployResponseDTO

	err := c.sendCoreTransaction(
		&SendCoreTransactionModel{
			Url:           ContractDeployUrl,
			HttpMethod:    http.MethodPost,
			RequestDto:    &requestDto,
			ResponseDto:   &responseDto,
			TokenAuth:     c.tokenAuth,
			AppID:         c.appID,
			ParseRequest:  true,
			ParseResponse: true,
		},
	)

	if err != nil {
		log.Printf("An error has raised calling core api Create New Custodied Wallet. %+v", err)
		return nil, err
	}

	return responseDto, nil
}

// Creates a new custodied wallet for a new user
func (c *CoreApi) CreateNewCustodiedWallet(requestDto *NewWalletRequestDTO) (*AccountZeroResponseDTO, error) {

	var responseDto *AccountZeroResponseDTO

	err := c.sendCoreTransaction(
		&SendCoreTransactionModel{
			Url:           NewWalletUrl,
			HttpMethod:    http.MethodPost,
			RequestDto:    &requestDto,
			ResponseDto:   &responseDto,
			TokenAuth:     c.tokenAuth,
			AppID:         c.appID,
			ParseRequest:  true,
			ParseResponse: true,
		},
	)
	if err != nil {
		log.Printf("An error has raised calling core api Create New Custodied Wallet. %+v", err)
		return nil, err
	}

	return responseDto, nil
}

// Returns the user account zero address. It is obtained creating the wallet from the seed prhase ang deriving the account 0.
func (c *CoreApi) AccountZeroAddress(requestDto *AccountZeroRequestDTO) (*AccountZeroResponseDTO, error) {

	var responseDto *AccountZeroResponseDTO

	err := c.sendCoreTransaction(
		&SendCoreTransactionModel{
			Url:           AccountZeroAddressUrl,
			HttpMethod:    http.MethodGet,
			RequestDto:    &requestDto,
			ResponseDto:   &responseDto,
			TokenAuth:     c.tokenAuth,
			AppID:         c.appID,
			ParseRequest:  true,
			ParseResponse: true,
		},
	)
	if err != nil {
		log.Printf("An error has raised calling core api AccountZeroAddress. %+v", err)
		return nil, err
	}

	return responseDto, nil
}

// Returns all the customer users paginated by offset and rows
func (c *CoreApi) ListUserInfo(offset, rows uint32) (response *[]UserInfotDTO, err error) {

	url := fmt.Sprintf("%s?o=%d&n=%d", UsersRegisteredWithCustodiedWalletUrl, offset, rows)

	err = c.sendCoreTransaction(
		&SendCoreTransactionModel{
			Url:           url,
			HttpMethod:    http.MethodGet,
			RequestDto:    nil,
			ResponseDto:   &response,
			TokenAuth:     c.tokenAuth,
			AppID:         c.appID,
			ParseRequest:  false,
			ParseResponse: true,
		},
	)
	if err != nil {
		log.Printf("An error has raised listing customer users. %+v", err)
		return nil, err
	}

	return response, nil
}

// Returns the User seed prhase encrypted with the customer RSA public key
func (c *CoreApi) UserMnemonic(requestDto *UserWalletSeedRequestDTO) (*UserWalletSeedResponseDTO, error) {

	var responseDto *UserWalletSeedResponseDTO

	err := c.sendCoreTransaction(
		&SendCoreTransactionModel{
			Url:           WalletSeedPhraseUrl,
			HttpMethod:    http.MethodGet,
			RequestDto:    &requestDto,
			ResponseDto:   &responseDto,
			TokenAuth:     c.tokenAuth,
			AppID:         c.appID,
			ParseRequest:  true,
			ParseResponse: true,
		},
	)
	if err != nil {
		log.Printf("An error has raised calling core api UserMnemonic. %+v", err)
		return nil, err
	}

	return responseDto, nil
}

// Sends a custodied wallet mutable transaction
func (c *CoreApi) SendCustodiedWalletMutableTransaction(requestDto *AbiMutableRequestDTO) (*AbiMutableResponseDTO, error) {
	var responseDto *AbiMutableResponseDTO

	err := c.sendCoreTransaction(
		&SendCoreTransactionModel{
			Url:           CustodiedWalletCoreMutableUrl,
			HttpMethod:    http.MethodPost,
			RequestDto:    &requestDto,
			ResponseDto:   &responseDto,
			TokenAuth:     c.tokenAuth,
			AppID:         c.appID,
			ParseRequest:  true,
			ParseResponse: true,
		},
	)
	if err != nil {
		log.Printf("An error has raised calling core api SendCustodiedWalletMutableTransaction. %+v", err)
		return nil, err
	}

	return responseDto, nil
}

// Sends a Mutable transaction
func (c *CoreApi) SendMutableTransaction(requestDto *AbiMutableRequestDTO) (*AbiMutableResponseDTO, error) {
	var responseDto *AbiMutableResponseDTO

	err := c.sendCoreTransaction(
		&SendCoreTransactionModel{
			Url:           CoreMutableUrl,
			HttpMethod:    http.MethodPost,
			RequestDto:    &requestDto,
			ResponseDto:   &responseDto,
			TokenAuth:     c.tokenAuth,
			AppID:         c.appID,
			ParseRequest:  true,
			ParseResponse: true,
		},
	)

	if err != nil {
		log.Printf("An error has raised calling core api sending mutable transaction. %+v", err)
		return nil, err
	}

	return responseDto, nil
}

func (c *CoreApi) SendViewTransaction(requestDto *AbiViewOptionsDTO) ([]interface{}, error) {
	var responseDto []interface{}

	err := c.sendCoreTransaction(
		&SendCoreTransactionModel{
			Url:           CoreViewUrl,
			HttpMethod:    http.MethodGet,
			RequestDto:    &requestDto,
			ResponseDto:   &responseDto,
			TokenAuth:     c.tokenAuth,
			AppID:         c.appID,
			ParseRequest:  true,
			ParseResponse: true,
		},
	)

	if err != nil {
		log.Printf("An error has raised calling core api SendViewTransaction. %+v", err)
		return nil, err
	}

	return responseDto, nil
}

func (c *CoreApi) GetChainBalance(account common.Address, networkID uint64) (*GetChainBalanceResponseDTO, error) {
	responseDto := &GetChainBalanceResponseDTO{}

	err := c.sendCoreTransaction(
		&SendCoreTransactionModel{
			Url:           fmt.Sprintf(GetChainBalanceUrl, account, networkID),
			HttpMethod:    http.MethodGet,
			ResponseDto:   responseDto,
			TokenAuth:     c.tokenAuth,
			AppID:         c.appID,
			ParseRequest:  false,
			ParseResponse: true,
		},
	)

	if err != nil {
		log.Printf("An error has raised calling core api Chain Balance for account {%s} and network {%d}. %+v", account, networkID, err)
		return nil, err
	}

	return responseDto, nil
}
func (c *CoreApi) TransferNetworkNativeCrypto(requestDto *AbiMutableRequestDTO) (*AbiMutableResponseDTO, error) {
	var responseDto *AbiMutableResponseDTO

	err := c.sendCoreTransaction(
		&SendCoreTransactionModel{
			Url:           TransferNativeNetworkCryptoUrl,
			HttpMethod:    http.MethodPost,
			RequestDto:    &requestDto,
			ResponseDto:   &responseDto,
			TokenAuth:     c.tokenAuth,
			AppID:         c.appID,
			ParseRequest:  true,
			ParseResponse: true,
		},
	)

	if err != nil {
		log.Printf("An error has raised calling core api to transfer networknative crypto. %+v", err)
		return nil, err
	}

	return responseDto, nil
}

func (c *CoreApi) GetBlockchains() (*ChainsListDTO, error) {

	responseDto := &ChainsListDTO{}

	err := c.sendCoreTransaction(
		&SendCoreTransactionModel{
			Url:           GetBlockchains,
			HttpMethod:    http.MethodGet,
			ResponseDto:   responseDto,
			TokenAuth:     c.tokenAuth,
			AppID:         c.appID,
			ParseRequest:  false,
			ParseResponse: true,
		},
	)

	if err != nil {
		log.Printf("An error has raised getting blockchains. %+v", err)
		return nil, err
	}

	return responseDto, nil
}

func (c *CoreApi) GetGasPrice(gasprice *GasPriceRequestDTO) (*GasPriceResponseDTO, error) {

	responseDto := &GasPriceResponseDTO{}

	err := c.sendCoreTransaction(
		&SendCoreTransactionModel{
			Url:           fmt.Sprintf(GetGasPrice, gasprice.Network),
			HttpMethod:    http.MethodGet,
			ResponseDto:   &responseDto,
			TokenAuth:     c.tokenAuth,
			AppID:         c.appID,
			ParseRequest:  false,
			ParseResponse: true,
		},
	)

	if err != nil {
		log.Printf("An error has raised getting gasprice. %+v", err)
		return nil, err
	}

	return responseDto, nil
}
