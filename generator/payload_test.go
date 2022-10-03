package generator

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"log"
	"net/url"
	"secure-touch/generator/device"
	"testing"
)

func TestClient_GenerateStarterPayload(t *testing.T) {
	formatProxy, _ := url.Parse("http://localhost:8888")

	c, err := MakeClient(formatProxy)
	if err != nil {
		log.Println("Failed to create client ")
	}
	payload, err := c.GenerateStarterPayload()
	if err != nil {
		log.Println("failed to create starter payload ")
	}
	fmt.Println(payload)

}

func TestClient_GenerateStarterPayloadWithZip(t *testing.T) {

	formatProxy, _ := url.Parse("http://localhost:8888")

	c, err := MakeClient(formatProxy)
	if err != nil {
		log.Println("Failed to create client ")
	}
	payload, err := c.GenerateStarterPayload()
	if err != nil {
		log.Println("failed to create starter payload ")
	}

	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	w.Write([]byte(payload))
	w.Close()

	fmt.Println(b.Bytes())

}

func TestClient_GenerateStarterPayloadWithZipThenByteEncode(t *testing.T) {

	formatProxy, _ := url.Parse("http://localhost:8888")

	c, err := MakeClient(formatProxy)
	if err != nil {
		log.Println("Failed to create client ")
	}
	payload, err := c.GenerateStarterPayload()
	if err != nil {
		log.Println("failed to create starter payload ")
	}

	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	w.Write([]byte(payload))
	w.Close()
	key := "eG9yLWVuY3J5cHRpb24"

	byte := EncryptByte(b.Bytes(), key)

	fmt.Println(byte)

}

func TestClient_GenerateInteractionPayload(t *testing.T) {

	formatProxy, _ := url.Parse("http://localhost:8888")

	c, err := MakeClient(formatProxy)
	if err != nil {
		log.Println("Failed to create client ")
	}
	payload, err := c.GenerateInteractionPayload(true, true, 1)
	if err != nil {
		log.Println("failed to create starter payload ")
	}
	fmt.Println(payload)

}

func TestClient_GenerateMetadata(t *testing.T) {

	device.GetDevices()

	formatProxy, _ := url.Parse("http://localhost:8888")

	c, err := MakeClient(formatProxy)
	if err != nil {
		log.Println("Failed to create client ")
	}
	payload, err := c.GenerateMetadata()
	if err != nil {
		log.Println("failed to create starter payload ")
	}
	fmt.Println(payload)

}
