package sign

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/url"
	"sort"
)

func paramsSort(params url.Values) []string {
	keys := make([]string, 0)
	for k, _ := range params {
		keys = append(keys, k)
	}

	sort.Strings(keys[:])
	return keys
}

func getSortedString(params url.Values) string {
	var signString string
	sliceSort := paramsSort(params)

	for _, value := range sliceSort {
		if value != "sign" {
			signString = fmt.Sprintf("%s%s%s", signString, value, params.Get(value))
		}
	}
	return signString
}

func GetSign(params url.Values, secret string) string {
	signStr := getSortedString(params)
	buf := bytes.NewBuffer(nil)
	buf.WriteString(signStr)
	buf.WriteString(secret)

	return fmt.Sprintf("%x", sha256.Sum256(buf.Bytes()))
}

func Params(c *gin.Context) {
	var err error
	if err = c.Request.ParseMultipartForm(10); err == nil {

	} else if err = c.Request.ParseForm(); err == nil {

	} else {
		fmt.Println(err.Error())
		return
	}

	data := c.Request.PostForm
	fmt.Println(data)
}
