package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type LastParams struct {
	ReqParamsString string            `json:"reqParamsStr"`
	McuSessionId    string            `json:"mcuSessionId"`
	ReqTs           int64             `json:"reqTs"`
	AgentId         string            `json:"agentID"`
	IsRescuer       bool              `json:"isRescuer"`
	ReqParams       map[string]string `json:"-"`
}

func main() {
	str := `{
  "reqParamsStr": "{\"app_id\":\"Y25hMXFsTE03Tw\",\"client_id\":\"randomuid_0.36712070929511254\",\"client_info\":\"qgstestrtcdnpubclientinfo\",\"client_ip\":\"115.171.217.69\",\"nonce\":\"5eb5d083-7610-427f-b174-285388110a4e\",\"options_str\":\"{}\",\"organization_id\":\"82065966\",\"project_type\":\"mcu\",\"region\":\"cna\",\"request_id\":\"1f9b8571a740f3fa29cd80a890e5cd9a\",\"sign\":\"1d3ef6dd6c4144b1a4201a9fbcf4bf4880a97c334ebd7445361a878607420fe3\",\"stream_url\":\"rtn://Y25hMXFsTE03Tw.streaming.cloudhub.vip/live/qgstest\",\"terminal_id\":\"jsTestTerminalId_968265357464565900\",\"timestamp\":\"1642074409011\"}",
  "mixingSessionId": "1f9b8571a740f3fa29cd80a890e5cd9a",
  "reqTs": 1642067209019,
  "agentID": "ms_cna_bjbgph-ali_172.16.0.213",
  "isRescuer": false
}`

	var lastParams LastParams
	buf := bytes.NewBuffer(nil)
	buf.WriteString(str)
	err := json.Unmarshal(buf.Bytes(), &lastParams)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		buf.Reset()
		buf.WriteString(lastParams.ReqParamsString)

		params := make(map[string]string, 0)
		err := json.Unmarshal(buf.Bytes(), &params)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		lastParams.ReqParams = params
		fmt.Println(lastParams)
		fmt.Println(lastParams.ReqParams["app_id"])
	}
}
