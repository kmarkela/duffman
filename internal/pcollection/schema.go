package pcollection

type NodeList []Node

type Schema struct {
	Nodes       NodeList
	Name        string
	Description string
	Schema      string // postman schema (e.g. https://schema.getpostman.com/json/collection/v2.1.0/collection.json)
}

type Node struct {
	Name string
	Node []Node
	Req  *Req
}
