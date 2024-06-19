package pcollection

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func New(colF, envF string) (Collection, error) {

	var collection = Collection{}

	jsonC, err := os.Open(colF)
	if err != nil {
		return collection, fmt.Errorf("cannot open Collection file. Err: %s", err)
	}
	defer jsonC.Close()

	byteV, _ := io.ReadAll(jsonC)
	var rawCollection RawCollection
	if err := json.Unmarshal(byteV, &rawCollection); err != nil {
		return collection, fmt.Errorf("cannot unmarshal collection. Err: %s", err)
	}

	collection.Variables = rawCollection.Variable

	reqLt, err := getReqLt(&rawCollection)
	if err != nil {
		return collection, fmt.Errorf("cannot process Collection. Err: %s", err)
	}

	collection.Requests = reqLt

	if envF == "" {
		return collection, nil
	}

	jsonE, err := os.Open(envF)
	if err != nil {
		return collection, fmt.Errorf("cannot open Enviroment file. Err: %s", err)
	}
	defer jsonE.Close()

	byteE, _ := io.ReadAll(jsonE)
	var env Enviroment
	if err := json.Unmarshal(byteE, &env); err != nil {
		return collection, fmt.Errorf("cannot unmarshal env. Err: %s", err)
	}

	collection.Env = env.Values

	return collection, nil
}

func getReqLt(rc *RawCollection) ([]Req, error) {

	var rlt []Req

	for _, v := range rc.Items {
		tr, err := v.i2ReqLt()
		if err != nil {
			return nil, err
		}
		rlt = append(rlt, tr...)
	}

	return rlt, nil

}

func (c *Collection) ResolveVars() {}
