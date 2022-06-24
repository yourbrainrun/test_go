package tests

import (
	"github.com/yourbrainrun/test_go/test_sign/roc"
	"testing"
)

func TestGetToken(t *testing.T) {

	var (
		authKey     = "testAuthKey"
		secretKey   = "testSecretKey"
		channelName = "testChannelName"
		userId      = "testUserId"
		expiredTime = 1594704891
	)

	client := roc.Client{
		AuthKey:   authKey,
		SecretKey: secretKey,
	}
	token, err := client.GetToken(channelName, userId, int64(expiredTime))
	if err == nil {
		token = token[0 : len(token)-16]
	}

	if token != "eyJ0b2tlbiI6IjBiNDkzZTVjN2FjMjM2OThkOWNkYTJhZmMzNGUyMTQyIiwidGltZXN0YW1wIjoxNTk0NzA0ODkxfQ==" {
		t.Errorf("TestGetToken@GetToken test failed token=%s ", token)
	}
}
