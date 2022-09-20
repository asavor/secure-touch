package request

import (
	http "github.com/saucesteals/fhttp"
	"secure-touch/module/asos/header"
)

func (a *Asos) GetAkamai() error {
	request, err := http.NewRequest("GET", "https://my.asos.com/K_MOcUaqn6qh4/4vRZP4Ur/Q6rIn4/OfYpcL9pJa/Lh8ZKF85/FW/krMSQOJgYC", nil)
	if err != nil {
		return err
	}
	request.Header = header.HomeHeader(a.UserAgent, a.Sch)
	request.Close = true

	_, err = a.Do(request)
	return nil

}
