package oauth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/AnuragCheela/go-utils/resterrors"
	"github.com/mercadolibre/golang-restclient/rest"
)

const (
	paramAccessToken = "access_token"
	headerXClientID  = "X-Client-Id"
	headerXUserID    = "X-User-Id"
)

var (
	oauthRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8080",
		Timeout: 200 * time.Millisecond,
	}
)

type accessToken struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	ClientID string `json:"client_id"`
}

// AuthenticateRequest Function
func AuthenticateRequest(request *http.Request) resterrors.RestErr {
	if request == nil {
		return nil
	}
	cleanRequest(request)

	accessTokenID := strings.TrimSpace(request.URL.Query().Get(paramAccessToken))
	if accessTokenID == "" {
		return resterrors.NewBadRequestError("Access Token missing")
	}
	at, err := getAccessToken(accessTokenID)
	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil
		}
		return err
	}

	request.Header.Add(headerXClientID, fmt.Sprintf("%v", at.ClientID))
	request.Header.Add(headerXUserID, fmt.Sprintf("%v", at.UserID))
	return nil
}

func cleanRequest(request *http.Request) {
	if request == nil {
		return
	}
	request.Header.Del(headerXClientID)
	request.Header.Del(headerXUserID)
}

func getAccessToken(accessTokenID string) (*accessToken, resterrors.RestErr) {
	response := oauthRestClient.Get(fmt.Sprintf("/oauth/access_token/%s", accessTokenID))
	if response == nil || response.Response == nil {
		return nil, resterrors.NewInternalServerError("invalid restclient response when trying to get access token",
			errors.New("network timeout"))
	}
	if response.StatusCode > 299 {
		restErr, err := resterrors.NewRestErrorFromBytes(response.Bytes())
		if err != nil {
			return nil, resterrors.NewInternalServerError("invalid error interface when trying to get access token", err)
		}
		return nil, restErr
	}
	var at accessToken
	if err := json.Unmarshal(response.Bytes(), &at); err != nil {
		return nil, resterrors.NewInternalServerError("error when trying to unmarshal access token response",
			errors.New("error processing json"))
	}
	return &at, nil
}

// GetUserID function
func GetUserID(request *http.Request) string {
	if request == nil {
		return ""
	}
	userID := request.Header.Get(headerXUserID)
	return userID
}

// GetClientID function
func GetClientID(request *http.Request) string {
	if request == nil {
		return ""
	}
	clientID := request.Header.Get(headerXClientID)
	return clientID
}
