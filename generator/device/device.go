package device

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var GlobalDevices Devices

type Devices []struct {
	Navigator Navigator `json:"navigator"`
	Window    Window    `json:"window"`
	Screen    Screen    `json:"screen"`
	Wv        string    `json:"wv"`
	Wr        string    `json:"wr"`
}

type OneDevice struct {
	Navigator Navigator `json:"navigator"`
	Window    Window    `json:"window"`
	Screen    Screen    `json:"screen"`
	Wv        string    `json:"wv"`
	Wr        string    `json:"wr"`
}
type Navigator struct {
	UserAgent string   `json:"userAgent"`
	Language  string   `json:"language"`
	Languages []string `json:"languages"`
}
type Window struct {
	InnerHeight int `json:"innerHeight"`
	InnerWidth  int `json:"innerWidth"`
	OuterWidth  int `json:"outerWidth"`
}
type Screen struct {
	AvailWidth  int `json:"availWidth"`
	AvailHeight int `json:"availHeight"`
	Width       int `json:"width"`
	Height      int `json:"height"`
	PixelDepth  int `json:"pixelDepth"`
	ColorDepth  int `json:"colorDepth"`
}

func GetDevices() error {
	resp, err := http.Get("https://file.asavor.com/asos.json?key=735deefc-833e-403a-a5a8-d8bc7eaad7ad")
	if err != nil {
		log.Fatalln(err)
	}
	var DevicesJson Devices

	json.NewDecoder(resp.Body).Decode(&DevicesJson)
	//for _, k := range DevicesJson {
	//	GlobalDevices = append(GlobalDevices, k)
	//}

	GlobalDevices = DevicesJson
	return nil
}

func GetRandomDevice() *OneDevice {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(GlobalDevices)

	return (*OneDevice)(&GlobalDevices[rand.Intn(max-min+1)+min])

}
