package generator

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	http "github.com/saucesteals/fhttp"
	"io"
	"log"
	"secure-touch/generator/header"
	"secure-touch/generator/types"
	"strconv"
	"strings"
	"time"
)

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }

func (c *Client) SecureTouchPost(method, payload string, gzipData, encryptData bool) (string, error) {
	var secureTouchEndpoint strings.Builder
	secureTouchEndpoint.WriteString("https://st.asos.com/SecuredTouch/rest/services/v2/")
	secureTouchEndpoint.WriteString(method)
	secureTouchEndpoint.WriteString("/")
	secureTouchEndpoint.WriteString(c.AppID)

	request, err := http.NewRequest("POST", secureTouchEndpoint.String(), nil)
	request.Close = true

	if err != nil {
		log.Println(err)
		return "", err
	}

	clientePoch := time.Now().UTC().UnixNano() / 1e6

	request.Header = header.SecureTouchHeader(c.UserAgent, c.Sch, strconv.FormatInt(clientePoch, 10), c.Instanceuuid, c.AuthToken, c.ClientVersion, c.Website)

	var postPayload bytes.Buffer

	if gzipData {
		w := gzip.NewWriter(&postPayload)
		w.Write([]byte(payload))
		w.Close()
		request.Header.Set("content-encoding", "gzip")

		request.Header.Set("content-length", strconv.Itoa(len(postPayload.Bytes())))

		if !encryptData {
			request.Body = nopCloser{bytes.NewReader(postPayload.Bytes())}
		}
	}

	if encryptData {
		key := "eG9yLWVuY3J5cHRpb24"
		request.Header.Set("encrypted", "1")
		request.Header.Set("content-length", strconv.Itoa(len(postPayload.Bytes())))
		request.Body = nopCloser{bytes.NewReader(EncryptByte(postPayload.Bytes(), key))}
	}

	response, err := c.Do(request)
	if err != nil {

		return "", err
	}

	var jsonResponse types.StaterResponse

	err = json.NewDecoder(response.Body).Decode(&jsonResponse)
	defer response.Body.Close()

	if err != nil {
		log.Println("could not parse starter header")
		return "", err
	}
	if method == "starter" {
		c.AuthToken = jsonResponse.Token
	}

	return "", err

}
