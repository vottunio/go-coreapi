package apiwrapper

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/vottunio/go-coreapi/decoder"
	"github.com/vottunio/log"
)

const (
	ErrorParsingJson         string = "ERROR_PARSING_JSON"
	ErrorUnauthorized        string = "ERROR_UNAUTHORIZED"
	ErrorHttpStatus          string = "ERROR_HTTP_STATUS_%d"
	ErrorApiWrapperUrlNotSet string = "ERROR_API_WRAPPER_URL_NOT_SET"
)

type ErrorDTO struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type SetReqHeaders func(req *http.Request, tokenAuth string, appID string)

type RequestApiEndpointInfo struct {
	EndpointUrl    string
	RequestData    interface{}
	ResponseData   interface{}
	HttpMethod     string
	TokenAuth      string
	AppID          string
	ResponseStatus int
}

func RequestApiEndpoint(r *RequestApiEndpointInfo, setReqHeaders SetReqHeaders, parseRequest, parseResponse bool) error {
	var req *http.Request
	var res *http.Response
	var statuscode int = 0
	var requestDataBuffer *bytes.Buffer

	if _, err := url.Parse(r.EndpointUrl); err == nil {
		if parseRequest {
			b, err := json.Marshal(r.RequestData)
			if err != nil {
				log.Printf("An error was raised marshalling request data. %v", err)
				return err
			}
			requestDataBuffer = bytes.NewBuffer(b)
		} else {
			requestDataBuffer = &bytes.Buffer{}
		}

		// log.Printf("Sending post request to validate token and app for token %s, customer id %s and app secret %s", jti, customerID, appSecret)

		if req, err = http.NewRequest(r.HttpMethod, r.EndpointUrl, requestDataBuffer); err == nil {
			setReqHeaders(req, r.TokenAuth, r.AppID)
			client := &http.Client{
				Timeout: 30 * time.Second,
			}

			res, err = client.Do(req)
			if err == nil {
				defer res.Body.Close()
				body, _ := io.ReadAll(res.Body)
				statuscode = res.StatusCode
				log.Tracef("Received statuscode %d", statuscode)
				switch statuscode {
				case http.StatusOK, http.StatusCreated, http.StatusAccepted:
					if res.Header.Get("Content-Type") == "image/png" {
						kk := make([]byte, len(body))
						copy(kk, body)
						r.ResponseData = &kk

					} else if parseResponse {
						err = decoder.JsonNumberDecode(body, &r.ResponseData)
						if err != nil {
							log.Printf("Error unmarshaling token information received from api: %+v", err)
							return fmt.Errorf("%s - Error unmarshaling token information received from api: %+v", ErrorParsingJson, err)
						}
					}
					r.ResponseStatus = statuscode
					return nil

				case http.StatusUnauthorized:
					return errors.New(ErrorUnauthorized + " - The token used in not authorized to perform the requested operation")

				default:
					errorMsg := ErrorDTO{}
					err := json.Unmarshal(body, &errorMsg)
					if err != nil {
						log.Printf("Error unmarshaling token information received from api: %+v", err)
						return fmt.Errorf(ErrorHttpStatus, statuscode)
					}
					return errors.New(errorMsg.Code + " - " + errorMsg.Message)
				}
			} else {
				log.Printf("error executing request with error %+v", err)
				return err
			}
		} else {
			log.Printf("error creating request to send to server %+v", err)
			return err
		}
	} else {
		log.Printf("Invalid url or not set")
		return errors.New(ErrorApiWrapperUrlNotSet + " - The url sent is not correct")
	}
}
