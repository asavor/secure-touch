package generator

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"testing"
)

func TestEncryptString(t *testing.T) {
	key := "eG9yLWVuY3J5cHRpb24"
	payload := `dqwdqwdqw`
	fmt.Println(EncryptString(payload, key))
}

func TestEncryptByte(t *testing.T) {

	payload := `{"hello":"DSAdas"}`
	key := "eG9yLWVuY3J5cHRpb24"
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	w.Write([]byte(payload))
	w.Close()

	fmt.Println(EncryptByte(b.Bytes(), key))

}
