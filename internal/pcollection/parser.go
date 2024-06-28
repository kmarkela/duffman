package pcollection

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func New(colF, envF string, vars []string) (Collection, error) {

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
	vl, err := pVars(vars)
	if err != nil {
		return collection, fmt.Errorf("cannot parse provided variables %v. Err: %s", vars, err)
	}

	for k, v := range vl {
		for _, j := range collection.Variables {
			if strings.EqualFold(j.Key, k) {
				continue
			}
			j.Value = v
		}

	}

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
		return collection, fmt.Errorf("cannot open Environment file. Err: %s", err)
	}
	defer jsonE.Close()

	byteE, _ := io.ReadAll(jsonE)
	var env Environment
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

func pVars(vars []string) (map[string]string, error) {
	rh := make(map[string]string)

	for _, h := range vars {
		p := strings.Split(h, ":")
		if len(p) < 2 {
			return nil, fmt.Errorf("%s is wrong header format", h)
		}
		rh[strings.TrimSpace(p[0])] = p[1]
	}

	return rh, nil
}

func (c *Collection) ResolveVars() {}
