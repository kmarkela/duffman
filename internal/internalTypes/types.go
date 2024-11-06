package internalTypes

// moving types here to prevent circular reference

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
