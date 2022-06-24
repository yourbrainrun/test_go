package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	tt := GetClientPullType("{\"decoding_capacity\":3}")
	fmt.Println(tt)
}

func GetClientPullType(options string) string {

	mapOption := make(map[string]interface{}, 0)
	err := json.Unmarshal([]byte(options), &mapOption)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	if val, ok := mapOption["decoding_capacity"]; ok {
		fmt.Println("ok decode", val)
		dd := fmt.Sprintf("%#v", val)
		fmt.Println(dd)
		fmt.Println(reflect.TypeOf(val))
		if typeInt, assertOk := val.(float64); assertOk {
			fmt.Println("is int", typeInt)
			switch typeInt {
			case 3:
				return "3"

			case 4:
				return "4"

			case 5:
				return "5"
			default:
				return "default"
			}
		}
	}

	return "ret"
}
