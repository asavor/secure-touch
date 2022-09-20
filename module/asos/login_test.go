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
	interactions, err := asos.GenerateInteractionPayload(true, false)
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

	interactions3, err := asos.GenerateInteractionPayload(true, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = asos.SecureTouchPost("interactions", interactions3, true, true)
	if err != nil {
		fmt.Println(err)
		return
	}

	interactions2, err := asos.GenerateInteractionPayload(true, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(interactions2)

	_, err = asos.SecureTouchPost("interactions", interactions2, true, true)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = asos.Login()
	if err != nil {
		return
	}

}
