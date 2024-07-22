package jsonparser

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	delimiter = ".._=.dl.=_.."
	slice     = "_sl1c3_"
)

func Unmarshal(input string) (map[string]string, error) {
	var data interface{}
	if err := json.Unmarshal([]byte(input), &data); err != nil {
		return nil, err
	}
	result := make(map[string]string)
	parseJSON(data, "", result)
	return result, nil
}

func parseJSON(data interface{}, prefix string, result map[string]string) {
	switch value := data.(type) {
	case map[string]interface{}:
		for k, v := range value {
			parseJSON(v, prefix+k+delimiter, result)
		}
	// parse first element of a slice
	case []interface{}:
		parseJSON(value[0], prefix+slice, result)
		// result[prefix[:len(prefix)-len(delimiter)]+slice] = ""
	default:
		result[prefix[:len(prefix)-len(delimiter)]] = fmt.Sprintf("%v", value)
	}
}

func Marshal(data map[string]string) ([]byte, error) {
	jsonData := make(map[string]interface{})

	for key, value := range data {
		// Split the key into parts
		keys := strings.Split(key, delimiter)

		// Traverse the keys to set the value in jsonData
		temp := jsonData
		for i := 0; i < len(keys)-1; i++ {

			if _, ok := temp[keys[i]]; !ok {
				temp[keys[i]] = make(map[string]interface{})

			}
			temp = temp[keys[i]].(map[string]interface{})

		}
		if strings.Contains(keys[len(keys)-1], slice) {
			var tmpLst []map[string]interface{}
			k := strings.Split(keys[len(keys)-1], slice)[1]
			if _, ok := temp["test"]; !ok {
				tmpLst = make([]map[string]interface{}, 1)
				tmpLst[0] = make(map[string]interface{}, 1)
			} else {
				tmpLst = temp["test"].([]map[string]interface{})
			}
			tmpLst[0][k] = value
			temp["test"] = tmpLst
			fmt.Println(temp)

		} else {
			temp[keys[len(keys)-1]] = value
			fmt.Println(temp)
		}
	}

	d, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func Param2Srt(param string) string {
	return strings.ReplaceAll(param, delimiter, ".")
}
