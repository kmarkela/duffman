package pcollection

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func CollFromJson(colF, envF string) (*Collection, error) {

	return parseJSONs(colF, envF)
}

func parseJSONs(colF, envF string) (*Collection, error) {

	// Open the JSON file
	jsonC, err := os.Open(colF)
	if err != nil {
		return nil, fmt.Errorf("cannot open Collection file. Err: %s", err)
	}
	defer jsonC.Close()

	byteV, _ := io.ReadAll(jsonC)
	var rawCollection RawCollection
	if err := json.Unmarshal(byteV, &rawCollection); err != nil {
		return nil, fmt.Errorf("cannot unmarshal collection. Err: %s", err)
	}

	collection, err := rawCollection.filter()
	if err != nil {
		return nil, fmt.Errorf("cannot process Collection. Err: %s", err)
	}

	if envF == "" {
		return collection, nil
	}

	jsonE, err := os.Open(envF)
	if err != nil {
		return nil, fmt.Errorf("cannot open Enviroment file. Err: %s", err)
	}
	defer jsonE.Close()

	byteE, _ := io.ReadAll(jsonE)
	var env *Enviroment
	if err := json.Unmarshal(byteE, &env); err != nil {
		return nil, fmt.Errorf("cannot unmarshal env. Err: %s", err)
	}

	collection.Env = env.Values

	return collection, nil
}

// Simplify/Filter sub-layers of the collection
func (rc *RawCollection) filter() (*Collection, error) {

	var col *Collection = &Collection{}
	col.Variables = rc.Variable

	for _, v := range rc.Items {
		tr, err := v.filter()
		if err != nil {
			return nil, err
		}
		// fmt.Println(tr)
		col.Requests = append(col.Requests, tr...)
	}
	return col, nil
}

func (c *Collection) ResolveVars() {}
