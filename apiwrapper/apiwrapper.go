package apiwrapper

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	ErrorParsingJson         string = "ERROR_PARSING_JSON"
	ErrorUnauthorized        string = "ERROR_UNAUTHORIZED"
	ErrorHttpStatus          string = "ERROR_HTTP_STATUS_%d"
	ErrorApiWrapperUrlNotSet string = "ERROR_API_WRAPPER_URL_NOT_SET"
)

type SetReqHeaders func(req *http.Request, tokenAuth string, appID string)

type RequestApiEndpointInfo struct {
	EndpointUrl  string
	RequestData  interface{}
	ResponseData interface{}
	HttpMethod   string
	TokenAuth    string
	AppID        string
}

func RequestApiEndpoint(r *RequestApiEndpointInfo, setReqHeaders SetReqHeaders) error {
	var req *http.Request
	var res *http.Response
	var statuscode int = 0

	if _, err := url.Parse(r.EndpointUrl); err == nil {
		b, err := json.Marshal(r.RequestData)
		if err != nil {
			log.Printf("An error was raised marshalling request data. %v", err)
			return err
		}

		// log.Printf("Sending post request to validate token and app for token %s, customer id %s and app secret %s", jti, customerID, appSecret)

		if req, err = http.NewRequest(r.HttpMethod, r.EndpointUrl, bytes.NewBuffer(b)); err == nil {
			setReqHeaders(req, r.TokenAuth, r.AppID)
			client := &http.Client{
				Timeout: 30 * time.Second,
			}

			res, err = client.Do(req)
			if err == nil {
				defer res.Body.Close()
				body, _ := io.ReadAll(res.Body)
				statuscode = res.StatusCode
				log.Printf("Received statuscode %d", statuscode)
				switch statuscode {
				case http.StatusOK, http.StatusCreated:
					err = json.Unmarshal(body, &r.ResponseData)
					if err != nil {
						log.Printf("Error unmarshaling token information received from api: %+v", err)
						return errors.New(ErrorParsingJson)
					}
					return nil

				case http.StatusUnauthorized:
					return errors.New(ErrorUnauthorized)

				default:
					return fmt.Errorf(ErrorHttpStatus, statuscode)
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
		return errors.New(ErrorApiWrapperUrlNotSet)
	}
}
