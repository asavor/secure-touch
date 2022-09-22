package request

import (
	"fmt"
	"github.com/justhyped/OrderedForm"
	http "github.com/saucesteals/fhttp"
	"io"
	"secure-touch/module/asos/header"
	"strings"
)

func (a *Asos) Login() error {

	form := new(OrderedForm.OrderedForm)
	form.Set("idsrv.xsrf", a.idsrvXsrf)
	form.Set("SecuredTouchToken", a.StToken)
	form.Set("Username", "dddsa23e@outlook.com")
	form.Set("Password", "dsaewe43!232")

	request, err := http.NewRequest("POST", fmt.Sprintf(`https://my.asos.com/identity/login?signin=%v&checkout=False&showAllOptions=False`, a.LoginUrl), strings.NewReader(form.URLEncode()))
	if err != nil {
		return err
	}
	request.Header = header.PostLoginHeader(a.UserAgent, a.Sch)
	request.Close = true

	response, err := a.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	responseByte, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	fmt.Println(responseByte)

	return nil

}
