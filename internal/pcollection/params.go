package pcollection

import (
	"encoding/json"
	"fmt"
	"strings"
)

func UnmarshalJSONBody(s string) (map[string]string, error) {

	var data map[string]interface{} = make(map[string]interface{})
	var dataList []map[string]interface{}
	var result map[string]string

	switch {
	case strings.HasPrefix(s, "["):
		if err := json.Unmarshal([]byte(s), &dataList); err != nil {
			return result, fmt.Errorf("error parsing JSON: %s", err.Error())
		}
		// TODO: Fuzzing only firs element in the list. update to fuzz all
		data = dataList[0]
		data["DM-data-in-slice"] = struct{}{}
	default:
		if err := json.Unmarshal([]byte(s), &data); err != nil {
			return result, fmt.Errorf("error parsing JSON: %s", err.Error())
		}
	}

	result = parseJSONBody(data, "")

	return result, nil
}

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

// TODO:
// func UnmarshalXMLBody() {}
// func MarshalXMLBody() {}
