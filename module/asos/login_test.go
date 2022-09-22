package asos

import (
	"fmt"
	"secure-touch/module/asos/request"
	"testing"
)

func TestLogin(t *testing.T) {

	asos, _ := request.Init()

	err := asos.Home()
	if err != nil {
		return
	}
	err = asos.GetAkamai()
	if err != nil {
		return
	}

	err = asos.GenerateSecureTouchToken()
	if err != nil {
		return
	}

	starterPayload, err := asos.GenerateStarterPayload()
	if err != nil {
		return
	}
	_, err = asos.SecureTouchPost("starter", starterPayload, true, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	count := 0
	interactions, err := asos.GenerateInteractionPayload(false, false, count)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = asos.SecureTouchPost("interactions", interactions, true, true)
	if err != nil {
		fmt.Println(err)
		return
	}

	metaData, err := asos.GenerateMetadata()
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = asos.SecureTouchPost("metadata", metaData, true, true)
	if err != nil {
		fmt.Println(err)
		return
	}

	interactions1, err := asos.GenerateInteractionPayload(true, true, count)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = asos.SecureTouchPost("interactions", interactions1, true, true)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = asos.Login()
	if err != nil {
		return
	}

}
