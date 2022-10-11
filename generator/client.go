package generator

import (
	"fmt"
	"github.com/google/uuid"
	http "github.com/saucesteals/fhttp"
	"github.com/saucesteals/fhttp/cookiejar"
	"github.com/saucesteals/mimic"
	"net/url"
	"secure-touch/generator/device"
)

type Client struct {
	*http.Client
	UserAgent     string
	Sch           string
	DeviceID      string
	ClientVersion string
	AuthToken     string
	DeviceType    string
	AppID         string
	Website       string
	Instanceuuid  string
	StToken       string
	AppSessionId  string
	LoginUrl      string
	Device        *device.OneDevice
	StartTime     int64
}

func MakeClient(proxy *url.URL) (*Client, error) {

	m, _ := mimic.Chromium(mimic.BrandChrome, `106.0.0.0`)

	jar, _ := cookiejar.New(nil)

	client := &http.Client{Transport: m.ConfigureTransport(&http.Transport{Proxy: http.ProxyURL(proxy)}), Jar: jar}

	c := &Client{Client: client}

	c.Sch = `"Chromium";v="106", "Google Chrome";v="106", "Not;A=Brand";v="99"`
	c.UserAgent = fmt.Sprintf("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", m.Version())
	c.DeviceID = "Id-" + uuid.New().String()
	c.ClientVersion = "3.13.2w"
	c.DeviceType = "Chrome(106.0.0.0)-Windows(10)"
	c.AppID = "asos"
	c.AuthToken = "YjIxMzVjdDIxSnVsVnlP"
	c.Website = "https://my.asos.com/"
	c.Device = device.GetRandomDevice()
	c.Instanceuuid = uuid.New().String()

	return c, nil

}
