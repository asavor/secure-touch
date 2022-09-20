package request

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	http "github.com/saucesteals/fhttp"
	"io"
	"net/url"
	"regexp"
	"secure-touch/module/asos/header"
)

type secureTouchToken struct {
	PingVersion  string `json:"pingVersion"`
	AppId        string `json:"appId"`
	AppSessionId string `json:"appSessionId"`
}

func (a *Asos) GenerateSecureTouchToken() error {
	tokenPayload := &secureTouchToken{
		PingVersion:  "1.3.0p",
		AppId:        a.AppID,
		AppSessionId: a.AppSessionId,
	}
	jsonPayload, err := json.Marshal(tokenPayload)
	if err != nil {
		return errors.New("failed to marshal token payload")
	}
	baseURL := base64.StdEncoding.EncodeToString(jsonPayload)
	encodeUrl := url.QueryEscape(baseURL)

	request, _ := http.NewRequest("GET", "https://st-static.asos.com/sdk/pong.js?body="+encodeUrl, nil)

	request.Header = header.SecureTouchTokenHeader(a.UserAgent, a.Sch)
	request.Close = true

	response, err := a.Do(request)
	if err != nil {
		return err
	}

	responseByte, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	SecureTouchToken := regexp.MustCompile(`(?m) = '(.*)'`)
	SecureTouchTokenGroup := SecureTouchToken.FindSubmatch(responseByte)
	if len(SecureTouchTokenGroup) != 2 {
		return errors.New("can not find secureTouchToken")
	}
	a.StToken = string(SecureTouchTokenGroup[1])

	return nil

}
