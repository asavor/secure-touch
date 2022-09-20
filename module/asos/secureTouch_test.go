package asos

import (
	"fmt"
	"secure-touch/module/asos/request"
	"testing"
)

func TestAsos_GenerateSecureTouchToken(t *testing.T) {

	asos, _ := request.Init()

	err := asos.Home()
	if err != nil {
		return
	}

	err = asos.GenerateSecureTouchToken()
	fmt.Println(asos.StToken)
	err = asos.Login()
}
