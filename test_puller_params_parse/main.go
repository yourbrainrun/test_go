package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	str := `{
  "reqParamsStr": "{\"app_id\":\"Y25hNTVFNzhvcUI\",\"nonce\":\"635ada86-6a1f-467a-b84b-2f73c44e66e6\",\"organization_id\":\"1651071864\",\"project_type\":\"cdn_live\",\"puller_type\":\"subscribe_scheduling\",\"region\":\"cna\",\"rtmp_pull_url\":\"rtmp://play.cloudatory.cn/live/qgstestccccccccccccccccccc\",\"sign\":\"ea68a508fc62033bca08d0f9f1f982e596941bace781331663d9b622a5e7598a\",\"stream_url\":\"rtn://play.cloudatory.cn/live/qgstestccccccccccccccccccc\",\"timestamp\":\"1656046487308\"}",
  "pullerSessionId": "b672abc693eb1fa74dfb1ad8bc0f0820",
  "reqTs": 1656039287314,
  "agentID": "ms_cna_bjbgph-ali_172.16.0.213",
  "isRescuer": false
}`

	data := NewBootManager(str)
	fmt.Println(data.ReqParams["app_id"])
	fmt.Println(data.ReqParams["stream_url"])
}

type BootParams struct {
	ReqParams       map[string]string
	PullerSessionId string `json:"pullerSessionId"`
	ReqTs           int64  `json:"reqTs"`
	AgentId         string `json:"agentID"`
	IsRescuer       bool   `json:"isRescuer"`
	ReqParamsString string `json:"reqParamsStr"`
}

func NewBootManager(paramsString string) *BootParams {
	var lastParams BootParams
	buf := bytes.NewBuffer(nil)
	buf.WriteString(paramsString)
	err := json.Unmarshal(buf.Bytes(), &lastParams)
	if err != nil {
		fmt.Println("action", "puller cron parse params floor 1", "message", "parse error"+err.Error())
		return nil
	}

	buf.Reset()
	buf.WriteString(lastParams.ReqParamsString)
	params := make(map[string]string, 0)
	err = json.Unmarshal(buf.Bytes(), &params)
	if err != nil {
		fmt.Println("action", "puller cron parse params floor 2", "message", "parse error"+err.Error())
		return nil
	}

	lastParams.ReqParams = params
	return &lastParams
}
