package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
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
	test3()
	//test2()
	//test1()
}
func test3() {
	mixingConfigStr := `{"videoWidth":640,"videoHeight":360,"videoFrameRate":15,"videoBandWidth":400,"backgroundColor":0,"audioBandWidth":48,"audioChannels":1,"mixingExtraInfo":"","mixingSources":[{"streamUrl":"rtn://Y25hMjlZeGtjdTY.streaming.cloudhub.vip/live/qgstestmcu1","x":0,"y":0,"width":320,"height":180,"zOrder":0,"fit":true},{"streamUrl":"rtn://Y25hMjlZeGtjdTY.streaming.cloudhub.vip/live/qgstestmcu2","x":320,"y":180,"width":320,"height":180,"zOrder":0,"fit":true}]}`

	outer := make(map[string]interface{}, 0)
	buf := bytes.NewBuffer(nil)
	buf.WriteString(mixingConfigStr)

	err := json.Unmarshal(buf.Bytes(), &outer)
	if err != nil {
		return
	}

	dealMixingSource(outer)
	//dealMixingExtraInfo(outer)

	marshal, err := json.Marshal(outer)
	if err != nil {
		return
	}
	fmt.Println("============end ===========")
	fmt.Println(string(marshal))

}

func dealMixingSource(outer map[string]interface{}) {
	if value, ok := outer["mixingSources"]; ok {
		if MixingSourcesSli, assertOk := value.([]interface{}); assertOk {
			for _, mixSourceMap := range MixingSourcesSli {
				if mixSource, assertOk := mixSourceMap.(map[string]interface{}); assertOk {
					if streamUrl, ok := mixSource["streamUrl"]; ok {
						if streamUrlStr, assertStr := streamUrl.(string); assertStr {
							urlDetail, err := url.Parse(streamUrlStr)
							if err != nil {
								continue
							}
							mixSource["appId"] = urlDetail.Host
						}
					}
				}
			}
		} else {
			fmt.Println("========== assert ========")
			fmt.Println("action", outer["mixingSources"])
		}

	} else {
		fmt.Println("========== exist ========")
		fmt.Println(outer["mixingSources"])
		return
	}
}

func test2() {
	str := `{
  "reqParamsStr": "[{\"app_id\":\"Y25hMXFsTE03Tw\"}]",
  "mixingSessionId": "1f9b8571a740f3fa29cd80a890e5cd9a",
  "reqTs": 1642067209019,
  "agentID": "ms_cna_bjbgph-ali_172.16.0.213",
  "isRescuer": false
}`

	outer := make(map[string]interface{}, 0)
	buf := bytes.NewBuffer(nil)
	buf.WriteString(str)
	err := json.Unmarshal(buf.Bytes(), &outer)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if value, ok := outer["reqParamsStr"]; ok {
			fmt.Println(value)
			if val, ok := value.(string); ok {
				buf.Reset()
				buf.WriteString(val)

				floor2 := make([]map[string]interface{}, 0)
				fmt.Println(val)
				err = json.Unmarshal(buf.Bytes(), &floor2)
				if err != nil {
					return
				} else {
					for index, map1 := range floor2 {
						map1["122"] = 3333
						fmt.Println(index, map1)

					}
				}

				fmt.Println(floor2, "floor2")

				fmt.Println(outer, "outer")
				outer["reqParamsStr"] = floor2
				outerAdd(outer)
				fmt.Println(outer, "outer ====")
			}

		} else {
			fmt.Println("outer not found key")
		}
	}

}

func outerAdd(outer map[string]interface{}) {
	outer["outer_add"] = "outer_add"
}

func test1() {
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
