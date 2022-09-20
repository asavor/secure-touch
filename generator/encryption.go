package generator

import (
	"strings"
)

//(t.prototype.encryptionString = function (e, t) {
//for (var a = [], r = 0; r < e.length; r++) {
//var n = e.charCodeAt(r) ^ t.charCodeAt(r % t.length);
//a.push(String.fromCharCode(n));
//}
//return a.join("");
//}),
//(t.prototype.encryptionBytes = function (e, t) {
//for (var a = new Uint8Array(e.length), r = 0; r < e.length; r++)
//a[r] = e[r] ^ t.charCodeAt(r % t.length);
//return a;
//}),

func EncryptString(payload, key string) string {
	var a []string
	for i := 0; i < len(payload); i++ {
		n := []rune(payload)[i] ^ []rune(key)[i%len(key)]
		a = append(a, string(n))
	}
	return strings.Join(a, "")
}

func EncryptByte(payload []byte, key string) []byte {

	a := make([]uint8, len(payload))

	for r := 0; r < len(payload); r++ {
		a[r] = uint8(payload[r]) ^ uint8([]rune(key)[r%len(key)])
	}
	bytePayload := make([]byte, len(a))

	for i, v := range a {
		bytePayload[i] = v
	}

	return bytePayload
}
