package request

import (
	"net/url"
	"secure-touch/generator"
)

type Asos struct {
	*generator.Client
	idsrvXsrf string
}

func Init() (*Asos, error) {
	proxy, _ := url.Parse("http://localhost:8888")

	c, err := generator.MakeClient(proxy)
	if err != nil {
		return nil, err
	}

	a := &Asos{Client: c}

	return a, nil

}
