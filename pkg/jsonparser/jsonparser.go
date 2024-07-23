package jsonparser

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	delimiter = ".._=.dl.=_.."
	Slice     = "_sl1c3_"
)

func Unmarshal(input string) (map[string]string, error) {
	var data interface{}

	result := make(map[string]string)
	if err := json.Unmarshal([]byte(input), &data); err != nil {
		return nil, err
	}
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
		parseJSON(value[0], prefix+Slice, result)
		result[Slice] = result[Slice] + strings.TrimSuffix(prefix, delimiter) + delimiter + delimiter
	default:
		// result[prefix[:len(prefix)-len(delimiter)]] = fmt.Sprintf("%v", value)
		result[strings.TrimSuffix(prefix, delimiter)] = fmt.Sprintf("%v", value)
	}
}

func Marshal(data map[string]string) ([]byte, error) {
	jsonData := make(map[string]interface{})

	for key, value := range data {
		if key == Slice {
			continue
		}
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
		if strings.Contains(keys[len(keys)-1], Slice) {
			var tmpLst []map[string]interface{}
			k := strings.Split(keys[len(keys)-1], Slice)[1]
			if _, ok := temp[Slice]; !ok {
				tmpLst = make([]map[string]interface{}, 1)
				tmpLst[0] = make(map[string]interface{}, 1)
			} else {
				tmpLst = temp[Slice].([]map[string]interface{})
			}

			tmpLst[0][k] = value
			temp[Slice] = tmpLst

			if k == "" {
				tmpLst := make([]interface{}, 1)
				tmpLst[0] = value
				temp[Slice] = tmpLst
			}

		} else {
			temp[keys[len(keys)-1]] = value
		}
	}

	var jd interface{}
	for _, v := range strings.Split(data[Slice], delimiter+delimiter) {
		if v == "" {
			jd = jsonData[Slice]
			continue
		}
		var tPr interface{}
		temp := jsonData
		keys := strings.Split(v, delimiter)
		for _, v := range keys {
			tPr = temp
			temp = temp[v].(map[string]interface{})
		}
		tPr.(map[string]interface{})[keys[len(keys)-1]] = temp[Slice]
		jd = jsonData
	}

	if jd == nil {
		jd = jsonData
	}

	d, err := json.Marshal(jd)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func Param2Str(param string) string {
	return strings.ReplaceAll(param, delimiter, ".")
}
