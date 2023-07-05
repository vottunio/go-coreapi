package coreapi

import (
	"fmt"
	"net/http"

	"github.com/vottunio/sdk-core-go/apiwrapper"
)

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
