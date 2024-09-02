package pcollection

type NodeList []Node

type Schema struct {
	Nodes NodeList
}

type Node struct {
	Name string
	Node NodeList
	Req  *Req
}

// func createSchema(rc *RawCollection) (Schema, error) {

// 	var sc Schema = Schema{}
// 	// var n Node = Node{}

// 	for _, v := range rc.Items {

// 		fmt.Printf("%s - {}: %v", v.Name, v.Item)
// 	}

// 	return sc, nil
// }
