package coreapi

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type CoreApi struct {
	RootUrl   string
	tokenAuth string
	appID     string
}

func New(tokenAuth, appID, rootUrl string) CoreApi {
	if !strings.HasSuffix(rootUrl, "/") {
		rootUrl += "/"
	}

	return CoreApi{tokenAuth: tokenAuth, appID: appID, RootUrl: rootUrl}
}

// Creates a new custodied wallet for a new user
func (c *CoreApi) CreateNewCustodiedWallet(requestDto *NewWalletRequestDTO) (*AccountZeroResponseDTO, error) {

	var responseDto *AccountZeroResponseDTO

	err := c.sendCoreTransaction(NewWalletUrl, http.MethodPost, &requestDto, &responseDto)

	if err != nil {
		log.Printf("An error has raised calling core api Create New Custodied Wallet. %+v", err)
		return nil, err
	}

	return responseDto, nil
}

// Returns the user account zero address. It is obtained creating the wallet from the seed prhase ang deriving the account 0.
func (c *CoreApi) AccountZeroAddress(requestDto *AccountZeroRequestDTO) (*AccountZeroResponseDTO, error) {

	var responseDto *AccountZeroResponseDTO

	err := c.sendCoreTransaction(AccountZeroAddressUrl, http.MethodGet, &requestDto, &responseDto)

	if err != nil {
		log.Printf("An error has raised calling core api AccountZeroAddress. %+v", err)
		return nil, err
	}

	return responseDto, nil
}

// Returns all the customer users paginated by offset and rows
func (c *CoreApi) ListUserInfo(offset, rows uint32) (response *[]UserInfotDTO, err error) {

	url := fmt.Sprintf("%s?o=%d&n=%d", UsersRegisteredWithCustodiedWalletUrl, offset, rows)
	err = c.sendCoreTransaction(url, http.MethodGet, nil, &response)
	if err != nil {
		log.Printf("An error has raised listing customer users. %+v", err)
		return nil, err
	}

	return response, nil
}

// Returns the User seed prhase encrypted with the customer RSA public key
func (c *CoreApi) UserMnemonic(requestDto *UserWalletSeedRequestDTO) (*UserWalletSeedResponseDTO, error) {

	var responseDto *UserWalletSeedResponseDTO

	err := c.sendCoreTransaction(WalletSeedPhraseUrl, http.MethodGet, &requestDto, &responseDto)

	if err != nil {
		log.Printf("An error has raised calling core api UserMnemonic. %+v", err)
		return nil, err
	}

	return responseDto, nil
}

// Sends a custodied wallet mutable transaction
func (c *CoreApi) SendCustodiedWalletMutableTransaction(requestDto *AbiMutableRequestDTO) (*AbiMutableResponseDTO, error) {
	var responseDto *AbiMutableResponseDTO

	err := c.sendCoreTransaction(CustodiedWalletCoreMutableUrl, http.MethodPost, &requestDto, &responseDto)

	if err != nil {
		log.Printf("An error has raised calling core api SendCustodiedWalletMutableTransaction. %+v", err)
		return nil, err
	}

	return responseDto, nil
}

// Sends a Mutable transaction
func (c *CoreApi) SendMutableTransaction(requestDto *AbiMutableRequestDTO) (*AbiMutableResponseDTO, error) {
	var responseDto *AbiMutableResponseDTO

	err := c.sendCoreTransaction(CoreMutableUrl, http.MethodPost, &requestDto, &responseDto)

	if err != nil {
		log.Printf("An error has raised calling core api sending mutable transaction. %+v", err)
		return nil, err
	}

	return responseDto, nil
}

func (c *CoreApi) SendViewTransaction(requestDto *AbiViewOptionsDTO) ([]interface{}, error) {
	var responseDto []interface{}

	err := c.sendCoreTransaction(CoreViewUrl, http.MethodGet, &requestDto, &responseDto)

	if err != nil {
		log.Printf("An error has raised calling core api SendViewTransaction. %+v", err)
		return nil, err
	}

	return responseDto, nil
}

func (c *CoreApi) TransferNetworkNativeCrypto(requestDto *AbiMutableRequestDTO) (*AbiMutableResponseDTO, error) {
	var responseDto *AbiMutableResponseDTO

	err := c.sendCoreTransaction(TransferNativeNetworkCryptoUrl, http.MethodPost, &requestDto, &responseDto)

	if err != nil {
		log.Printf("An error has raised calling core api to transfer networknative crypto. %+v", err)
		return nil, err
	}

	return responseDto, nil
}
