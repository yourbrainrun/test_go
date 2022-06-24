package roc

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"strconv"
	"time"
)

type Client struct {
	AuthKey   string
	SecretKey string
}

type Token struct {
	Token     string `json:"token"`
	Timestamp int64  `json:"timestamp"`
}

func (client *Client) GetToken(channelName string, userId string, expireTime int64) (authToken string, err error) {
	bodyStr := "authkey" + client.AuthKey + "channame" + channelName + "timestamp" + strconv.Itoa(int(expireTime)) + "userid" + userId
	bodyStr = client.AuthKey + bodyStr

	// md5(authKey + bodyStr)
	bodyStrByte := []byte(bodyStr)
	bodyMd5 := md5.New()
	bodyMd5.Write(bodyStrByte)
	bodyEncodeStr := hex.EncodeToString(bodyMd5.Sum(nil))

	// md5(secretKey)
	secretMd5 := md5.New()
	secretByte := []byte(client.SecretKey)
	secretMd5.Write(secretByte)
	secretEncodeStr := hex.EncodeToString(secretMd5.Sum(nil))

	// md5(body+secret)
	allStrByte := []byte(bodyEncodeStr + secretEncodeStr)
	allMd5 := md5.New()
	allMd5.Write(allStrByte)
	allEncodeStr := hex.EncodeToString(allMd5.Sum(nil))

	token := &Token{
		Token:     allEncodeStr,
		Timestamp: expireTime,
	}
	return token.GetBase64Str()
}

func (token *Token) GetBase64Str() (string, error) {
	val, err := json.Marshal(token)
	if err != nil {
		return "", err
	}
	baseStr := base64.StdEncoding.EncodeToString(val) + GetRandomStr(16)
	return baseStr, nil
}

func GetRandomStr(n int64) string {
	rand.Seed(time.Now().UnixNano())
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
