package coreapi

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/vottunio/go-coreapi/apiwrapper"
	"github.com/vottunio/log"
)

type SendCoreTransactionModel struct {
	Url            string
	HttpMethod     string
	RequestDto     interface{}
	ResponseDto    interface{}
	TokenAuth      string
	AppID          string
	ResponseStatus int
	ParseRequest   bool
	ParseResponse  bool
}

func (c *CoreApi) sendCoreTransaction(s *SendCoreTransactionModel) error {

	kk := &apiwrapper.RequestApiEndpointInfo{
		EndpointUrl:  c.createAbsoluteCoreApiUrl(s.Url),
		RequestData:  s.RequestDto,
		ResponseData: s.ResponseDto,
		HttpMethod:   s.HttpMethod,
		TokenAuth:    s.TokenAuth,
		AppID:        s.AppID,
	}
	err := apiwrapper.RequestApiEndpoint(
		kk,
		setReqHeaders,
		s.ParseRequest,
		s.ParseResponse,
	)

	if err != nil {
		log.Errorf("An error has raised calling core auth endpoint. %+v", err)
		log.Errorf("%+v", *s)
		return err
	}

	s.ResponseDto = kk.ResponseData
	s.ResponseStatus = kk.ResponseStatus
	return err
}

func (c *CoreApi) createAbsoluteCoreApiUrl(relativePath string) string {

	return c.RootUrl + relativePath
}

func setReqHeaders(req *http.Request, tokenAuth, appID string) {

	req.Header.Set(CONTENT_TYPE, MIME_TYPE_JSON)

	req.Header.Add(AUTHORIZATION, fmt.Sprintf("Bearer %s", tokenAuth))
	req.Header.Add(AUTH_APP_ID, appID)

}

func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}
