package jsonparser

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	delimiter = ".._=.dl.=_.."
	slice     = "|_!$+"
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
			parseJSON(v, fmt.Sprintf("%s%d%s", prefix, i, delimiter), result)
		}
	default:
		result[prefix[:len(prefix)-1]] = interface2str(value)
	}
}

type convertible interface {
	~int | ~float64 | ~string | ~bool | interface{}
}

func interface2str[T convertible](value T) string {
	switch v := any(value).(type) {
	case string:
		return v
	case int:
		return fmt.Sprintf("%d", v)
	case float64:
		return fmt.Sprintf("%f", v)
	case bool:
		return fmt.Sprintf("%t", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// func Unmarshal(s string) (map[string]string, error) {

// 	var data map[string]interface{} = make(map[string]interface{})
// 	var dataList []map[string]interface{}
// 	var result map[string]string

// 	switch {
// 	case strings.HasPrefix(s, "["):
// 		if err := json.Unmarshal([]byte(s), &dataList); err != nil {
// 			return result, fmt.Errorf("error parsing JSON: %s", err.Error())
// 		}
// 		// TODO: Fuzzing only firs element in the list. update to fuzz all
// 		data = dataList[0]
// 		data["DM-data-in-slice"] = struct{}{}
// 	default:
// 		if err := json.Unmarshal([]byte(s), &data); err != nil {
// 			return result, fmt.Errorf("error parsing JSON: %s", err.Error())
// 		}
// 	}

// 	result = parseJSONBody(data, "")

// 	return result, nil
// }

// func parseSlice(data []interface{}, prefix string) map[string]string {
// 	for i, v := range data {

// 	}
// }

func parseJSONBody(data map[string]interface{}, prefix string) map[string]string {
	result := make(map[string]string)

	for key, value := range data {
		switch v := value.(type) {
		case map[string]interface{}:
			// Nested object, recurse
			nestedResult := parseJSONBody(v, prefix+key+".")
			for nestedKey, nestedValue := range nestedResult {
				result[nestedKey] = nestedValue
			}
		// TODO: add parse of list
		// case []interface{}:
		default:
			// Leaf node, add to the result map
			result[prefix+key] = fmt.Sprintf("%v", value)
		}
	}

	return result
}

func MarshalJSONBody(data map[string]string) []byte {
	var list bool
	if _, ok := data["DM-data-in-slice"]; ok {
		delete(data, "DM-data-in-slice")
		list = true
	}
	jsonData := make(map[string]interface{})

	for key, value := range data {
		// Split the key into parts
		keys := strings.Split(key, ".")

		// Traverse the keys to set the value in jsonData
		temp := jsonData
		for i := 0; i < len(keys)-1; i++ {
			if _, ok := temp[keys[i]]; !ok {
				temp[keys[i]] = make(map[string]interface{})
			}
			temp = temp[keys[i]].(map[string]interface{})
		}
		temp[keys[len(keys)-1]] = value
	}

	var d []byte
	var err error
	if list {
		ljd := make([]map[string]interface{}, 1)
		ljd[0] = jsonData

		d, err = json.Marshal(ljd)
		if err != nil {
			// TODO: log it in verbose
			return nil
		}

	} else {
		d, err = json.Marshal(jsonData)
		if err != nil {
			// TODO: log it in verbose
			return nil
		}
	}
	return d
}
