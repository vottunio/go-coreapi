package coreapi

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vottunio/sdk-core-go/apiwrapper"
)

const (
	NewWalletUrl                          string = "evm/wallet/custodied/new"
	AccountZeroAddressUrl                 string = "evm/wallet/custodied/address"
	UsersRegisteredWithCustodiedWalletUrl string = "evm/wallet/custodied/list"
	WalletSeedPhraseUrl                   string = "evm/wallet/custodied/user/wallet"
	CustodiedWalletCoreMutableUrl         string = "evm/wallet/custodied/transact/mutable"
	CoreMutableUrl                        string = "evm/transact/mutable"
	CoreViewUrl                           string = "evm/transact/view"
	CONTENT_TYPE                          string = "Content-Type"
	AUTH_APP_ID                           string = "x-application-vkn"
	AUTHORIZATION                         string = "Authorization"
	MIME_TYPE_JSON                        string = "application/json; charset=UTF-8"
)

type CoreApi struct {
	RootUrl   string
	tokenAuth string
	appID     string
}

func New(tokenAuth, appID, rootUrl string) CoreApi {
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
		log.Printf("An error has raised calling core api Create New Custodied Wallet. %+v", err)
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
		log.Printf("An error has raised calling core api Create New Custodied Wallet. %+v", err)
		return nil, err
	}

	return responseDto, nil
}

// Sends a custodied wallet mutable transaction
func (c *CoreApi) SendCustodiedWalleetMutableTransaction(requestDto *AbiMutableRequestDTO) (*AbiMutableResponseDTO, error) {
	var responseDto *AbiMutableResponseDTO

	err := c.sendCoreTransaction(CustodiedWalletCoreMutableUrl, http.MethodPost, &requestDto, &responseDto)

	if err != nil {
		log.Printf("An error has raised calling core api Create New Custodied Wallet. %+v", err)
		return nil, err
	}

	return responseDto, nil
}

// Sends a Mutable transaction
func (c *CoreApi) SendMutableTransaction(requestDto *AbiMutableRequestDTO) (*AbiMutableResponseDTO, error) {
	var responseDto *AbiMutableResponseDTO

	err := c.sendCoreTransaction(CoreMutableUrl, http.MethodPost, &requestDto, &responseDto)

	if err != nil {
		log.Printf("An error has raised calling core api Create New Custodied Wallet. %+v", err)
		return nil, err
	}

	return responseDto, nil
}
func (c *CoreApi) SendViewTransaction(requestDto *AbiViewOptionsDTO) ([]interface{}, error) {
	var responseDto []interface{}

	err := c.sendCoreTransaction(CoreViewUrl, http.MethodGet, &requestDto, &responseDto)

	if err != nil {
		log.Printf("An error has raised calling core api Create New Custodied Wallet. %+v", err)
		return nil, err
	}

	return responseDto, nil
}
func (c *CoreApi) sendCoreTransaction(url, httpMethod string, requestDto, responseDto interface{}) error {
	return apiwrapper.RequestApiEndpoint(
		&apiwrapper.RequestApiEndpointInfo{
			EndpointUrl:  c.createAbsoluteCoreApiUrl(url),
			RequestData:  requestDto,
			ResponseData: &responseDto,
			HttpMethod:   httpMethod,
			TokenAuth:    c.tokenAuth,
			AppID:        c.appID,
		},
		setReqHeaders,
	)
}

func (c *CoreApi) createAbsoluteCoreApiUrl(relativePath string) string {

	return c.RootUrl + relativePath
}

func setReqHeaders(req *http.Request, tokenAuth, appID string) {

	req.Header.Set(CONTENT_TYPE, MIME_TYPE_JSON)

	req.Header.Add(AUTHORIZATION, fmt.Sprintf("Bearer %s", tokenAuth))
	req.Header.Add(AUTH_APP_ID, appID)

}
