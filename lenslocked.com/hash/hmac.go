package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"hash"
)

func NewHMAC(key string) HMAC {
	h := hmac.New(sha256.New, []byte(key))
	return HMAC{h}
}

type HMAC struct {
	hash.Hash
}

func (h HMAC) Bytes(input []byte) []byte {
	h.Reset()
	h.Write(input)
	return h.Sum(nil)
}

func (h HMAC) String(input string) string {
	b := h.Bytes([]byte(input))
	return base64.URLEncoding.EncodeToString(b)
}
