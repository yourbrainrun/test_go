package collectors

import (
	"bytes"
	"carina-exporter/configs"
	"encoding/json"
	"fmt"
	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/golang/protobuf/proto"
	"time"
)

func PutSlsLog(json string) {
	slsConfig := configs.GetSlsConfig()
	projectName := slsConfig.ProjectName
	storeName := slsConfig.LogStore

	buf := bytes.NewBuffer(nil)
	buf.WriteString(json)

	logs := []*sls.Log{}
	content, err := jsonToLogContent(json)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	log := &sls.Log{
		Time:     proto.Uint32(uint32(time.Now().Unix())),
		Contents: content,
	}
	logs = append(logs, log)

	logGroup := &sls.LogGroup{
		Topic:  proto.String(""),
		Source: proto.String("10.230.201.117"),
		Logs:   logs,
	}

	client := sls.CreateNormalInterface(slsConfig.EndPoint, slsConfig.AppKey, slsConfig.AppSecret, "")
	if err := client.PutLogs(projectName, storeName, logGroup); err != nil {
		fmt.Println(err.Error())
		return
	}
}

func jsonToLogContent(strJson string) ([]*sls.LogContent, error) {
	tempMap := make(map[string]interface{})
	buf := bytes.NewBuffer(nil)
	buf.WriteString(strJson)
	content := []*sls.LogContent{}

	if err := json.Unmarshal(buf.Bytes(), &tempMap); err == nil {
		for k, v := range tempMap {
			switch value := v.(type) {
			case map[string]interface{}:
				if tempStr, err := json.Marshal(value); err == nil {
					content = append(content, &sls.LogContent{
						Key:   proto.String(k),
						Value: proto.String(string(tempStr)),
					})
				} else {
					return []*sls.LogContent{}, err
				}

			default:
				content = append(content, &sls.LogContent{
					Key:   proto.String(k),
					Value: proto.String(fmt.Sprintf("%v", v)),
				})
			}
		}

		return content, nil
	} else {
		return []*sls.LogContent{}, err
	}
}
