package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	str := "sadfasdfasdfas"
	str = "app_idapp_id_valuenoncenonce_valuestream_urlhttp://test.com/alive/test123456"
	buf := bytes.NewBuffer(nil)
	buf.WriteString(str)
	fmt.Println(fmt.Sprintf("%x", md5.Sum(buf.Bytes())))

	w := md5.New()
	io.WriteString(w, str)
	//将str写入到w中
	md5str2 := fmt.Sprintf("%x", w.Sum(nil))
	fmt.Println(md5str2)
}
