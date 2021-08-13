package request

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

type BaseService struct {
	BaseUrl string
}

func NewBaseService(baseUrl string) *BaseService {
	return &BaseService{BaseUrl: baseUrl}
}

func (svc *BaseService) Request(method string, path string, data map[string]string) (string, error) {
	// 创建 client
	client := &http.Client{}

	// 创建 request
	url := svc.BaseUrl
	if path[0:1] == "/" {
		url = url + path
	} else {
		url = url + "/" + path
	}

	// 拼凑参数
	postData := bytes.NewBuffer(nil)
	dataWriter := multipart.NewWriter(postData)
	defer dataWriter.Close()

	if method == http.MethodPost && data != nil {
		for k, v := range data {
			_ = dataWriter.WriteField(k, v)
		}
	}

	// 创建请求
	req, err := http.NewRequest(method, url, postData)
	if err != nil {
		return "", err
	}

	// 如果是 post 需要添加 form-data 的格式声明
	if method == http.MethodPost {
		req.Header.Set("Content-Type", dataWriter.FormDataContentType())
	}

	// 如果是其他类型请求拼凑参数
	if method != http.MethodPost && data != nil {
		query := req.URL.Query()
		for k, v := range data {
			query.Add(k, v)
		}
		req.URL.RawQuery = query.Encode()
	}

	// 开始请求
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	bytesData, _ := ioutil.ReadAll(resp.Body)
	buf := bytes.NewBuffer(bytesData)

	return buf.String(), nil
}
