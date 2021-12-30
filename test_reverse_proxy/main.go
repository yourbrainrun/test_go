package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

func main() {
	router := gin.Default()
	router.GET("/channel", func(c *gin.Context) {
		proxy := reverseProxy(c)

		delete(c.Request.Header, "If-None-Match")
		proxy.ServeHTTP(c.Writer, c.Request)
	})

	_ = router.Run(":999")
}

func reverseProxy(c *gin.Context) *httputil.ReverseProxy {

	director := func(req *http.Request) {
		req.URL.Scheme = "https"
		fmt.Println(c.Request.URL.Path, "proxy")
		req.URL.Host = "bayer-elp-console-api-demo.cloudhub.vip"
		req.Host = "bayer-elp-console-api-demo.cloudhub.vip"
		req.URL.Path = "/channel" //c.Request.URL.Path
	}

	modifyFunc := func(res *http.Response) error {
		if res.StatusCode == http.StatusOK {
			proxyBody, err := ioutil.ReadAll(res.Body)
			if err != nil {
				return err
			}

			newBody := []byte(string(proxyBody) + " test ")
			res.Body = ioutil.NopCloser(bytes.NewBuffer(newBody))
			res.ContentLength = int64(len(newBody))
			res.Header.Set("Content-Length", fmt.Sprint(len(newBody)))
			return nil
		} else {
			return errors.New("error statusCode" + fmt.Sprintf("%d", res.StatusCode))
		}
	}

	errorHandler := func(res http.ResponseWriter, req *http.Request, err error) {
		res.Write([]byte(err.Error()))
	}

	return &httputil.ReverseProxy{Director: director, ModifyResponse: modifyFunc, ErrorHandler: errorHandler}
}
