#### php openssl_encrypt 生成加密串 用golang 解析问题

```php
$data = "123456";
$key = "lOLWf2ubswSDrUSt6lfwLQ";

$encode = openssl_encrypt($data, 'AES-128-ECB', $key, 0);
var_dump($encode);

//$decode = "bRpR18AplFNCRq4fD1Nj9BKyTanzItbJPQCBu0UCmH9qMVtejCoQM0OZTeF1bS4Z";
$decode = openssl_decrypt($encode, 'AES-128-ECB', $key, 0);
var_dump($decode);

```

golang 使用库 github.com/forgoer/openssl v0.0.0-20210828150411-6c5378b5b719
> https://github.com/forgoer/openssl 库说明

AES-128-ECB / AES-256-ECB

* key 长度问题 php 对长或者短的 key 做了截断 或 末尾补 \0 处理
* php 加密结果会自动base64 encode ，解密前会自动base64 decode

```golang

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
```

```golang
package main

import (
	"encoding/base64"
	"fmt"
	"github.com/yourbrainrun/test_go/test_openssl/crypt"
)

func main() {
	key := "testkey"

	svc1 := crypt.NewOpensslAES128ECB(key)
	enStr := "xquWr9JWshGZQSmcNuiuiKsHsI1SrJ27qV78RBBvXplqMVtejCoQM0OZTeF1bS4Z"
	decodeString, err := base64.StdEncoding.DecodeString(enStr)
	if err != nil {
		return
	}
	fmt.Println(string(decodeString))
	fmt.Println(len(key))
	str1 := svc1.Decrypt(string(decodeString))
	fmt.Println(str1)

	return
}

```

https://blog.csdn.net/Jum_Summer/article/details/111989021

```
golang里面封装有加解密库，安装好后，使用时直接加上

go get -u github.com/thinkoner/openssl

OpenSSL库的功能包装，用于对称和非对称加密和解密。
AES-ECB
AES-CBC
DES-ECB
DES-CBC
3DES-ECB
3DES-CBC

```

#### 其他参考

> https://blog.csdn.net/a_lzq/article/details/108392654 java 的缺位/过长处理方式

> https://www.php.net/manual/en/function.openssl-encrypt.php php对key 处理

Important: The key should have exactly the same length as the cipher you are using. For example, if you use AES-256 then
you should provide a $key that is 32 bytes long (256 bits == 32 bytes). Any additional bytes in $key will be truncated
and not used at all

passphrase（key） The passphrase. If the passphrase is shorter than expected, it is silently padded with NUL characters;
if the passphrase is longer than expected, it is silently truncated.

hex2bin 把十六进制值的字符串转换为 ASCII 字符。

