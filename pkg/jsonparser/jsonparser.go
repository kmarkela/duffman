package jsonparser

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	delimiter = ".._=.dl.=_.."
	slice     = ".|_!$+_"
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
	case []interface{}:
		for i, v := range value {
			parseJSON(v, fmt.Sprintf("%s%d%s", prefix, i, slice), result)
		}
	default:
		result[prefix[:len(prefix)-len(delimiter)]] = fmt.Sprintf("%v", value)
	}
}

// func MarshalJSONBody(data map[string]string) []byte {
// 	var list bool
// 	if _, ok := data["DM-data-in-slice"]; ok {
// 		delete(data, "DM-data-in-slice")
// 		list = true
// 	}
// 	jsonData := make(map[string]interface{})

// 	for key, value := range data {
// 		// Split the key into parts
// 		keys := strings.Split(key, ".")

// 		// Traverse the keys to set the value in jsonData
// 		temp := jsonData
// 		for i := 0; i < len(keys)-1; i++ {
// 			if _, ok := temp[keys[i]]; !ok {
// 				temp[keys[i]] = make(map[string]interface{})
// 			}
// 			temp = temp[keys[i]].(map[string]interface{})
// 		}
// 		temp[keys[len(keys)-1]] = value
// 	}

// 	var d []byte
// 	var err error
// 	if list {
// 		ljd := make([]map[string]interface{}, 1)
// 		ljd[0] = jsonData

// 		d, err = json.Marshal(ljd)
// 		if err != nil {
// 			// TODO: log it in verbose
// 			return nil
// 		}

// 	} else {
// 		d, err = json.Marshal(jsonData)
// 		if err != nil {
// 			// TODO: log it in verbose
// 			return nil
// 		}
// 	}
// 	return d
// }

func Marshal(jMap map[string]string) ([]byte, error) {
	data := make(map[string]interface{})

	for key, value := range data {
		// Split the key into parts
		keys := strings.Split(key, delimiter)

		// Traverse the keys to set the value in jsonData
		temp := data
		for i := 0; i < len(keys)-1; i++ {
			if _, ok := temp[keys[i]]; !ok {
				temp[keys[i]] = make(map[string]interface{})
			}
			temp = temp[keys[i]].(map[string]interface{})
		}
		temp[keys[len(keys)-1]] = value
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
