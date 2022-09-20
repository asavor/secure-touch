package request

import (
	"errors"
	http "github.com/saucesteals/fhttp"
	"io"
	"regexp"
	"secure-touch/module/asos/header"
)

func (a *Asos) Home() error {
	request, err := http.NewRequest("GET", "https://my.asos.com/identity/login", nil)
	if err != nil {
		return err
	}
	request.Header = header.HomeHeader(a.UserAgent, a.Sch)
	request.Close = true

	response, err := a.Do(request)
	if err != nil {
		return err
	}
	ByteResponse, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	sessionIdRegex := regexp.MustCompile(`(?m)=asos&sessionId=(.*?)"`)
	SessionIDGroup := sessionIdRegex.FindSubmatch(ByteResponse)
	if len(SessionIDGroup) != 2 {
		return errors.New("can not find sessionID")
	}
	a.AppSessionId = string(SessionIDGroup[1])

	idsrvXsrfGroupTokenRegex := regexp.MustCompile(`(?m)idsrv.xsrf" type="hidden" value="(.*?)"`)
	idsrvXsrfGroup := idsrvXsrfGroupTokenRegex.FindSubmatch(ByteResponse)
	if len(idsrvXsrfGroup) != 2 {
		return errors.New("can not find idsrvXsrf token")
	}
	a.idsrvXsrf = string(idsrvXsrfGroup[1])

	a.LoginUrl = response.Request.URL.Query().Get("signin")
	return nil

}
