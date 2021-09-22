package crypt

import (
	"fmt"
	"github.com/forgoer/openssl"
)

type OpensslAES128ECB struct {
	secret string
	key    []byte
}

func NewOpensslAES128ECB(key string) *OpensslAES128ECB {
	obj := OpensslAES128ECB{
		secret: key,
	}
	obj.getKey()
	return &obj
}

func (svc *OpensslAES128ECB) Encrypt(data string) string {
	dst, _ := openssl.AesECBEncrypt([]byte(data), svc.key, openssl.ZEROS_PADDING)
	return fmt.Sprintf("%s", dst)
}

func (svc *OpensslAES128ECB) Decrypt(data string) string {
	dst, _ := openssl.AesECBDecrypt([]byte(data), svc.key, openssl.ZEROS_PADDING)
	return string(dst)
}

func (svc *OpensslAES128ECB) getKey() {
	key := make([]byte, 16)
	copy(key, svc.secret)
	svc.key = key
}
