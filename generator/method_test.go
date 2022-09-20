package generator

import (
	"fmt"
	"log"
	"net/url"
	"testing"
)

func TestStarterPost(t *testing.T) {

	formatProxy, _ := url.Parse("http://localhost:8888")

	c, err := MakeClient(formatProxy)
	if err != nil {
		log.Println("Failed to create client ")
	}
	payload, err := c.GenerateStarterPayload()
	if err != nil {
		log.Println("failed to create starter payload ")
	}
	post, err := c.SecureTouchPost("starter", payload, true, false)
	if err != nil {
		return
	}
	fmt.Println(post)

}
