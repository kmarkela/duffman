package interactive

import (
	"encoding/json"
	"log"

	// "log"

	"github.com/kmarkela/duffman/internal/pcollection"
)

func buildReqStr(r pcollection.Req) string {
	marshaled, err := json.MarshalIndent(r, "", "   ")
	if err != nil {
		log.Fatalf("marshaling error: %s", err)
	}
	return string(marshaled)

}

func buildVarStr(col pcollection.Collection) string {

	marshaled, err := json.MarshalIndent(col, "", "   ")
	if err != nil {
		log.Fatalf("marshaling error: %s", err)
	}
	// logger.Logger.Info(string(marshaled))
	return string(marshaled)

}

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// )

// type Teacher struct {
// 	ID        string  `json:"id"`
// 	Firstname string  `json:"firstname"`
// 	Lastname  string  `json:"lastname"`
// 	TT        Test123 `json:"tt"`
// }

// type Test123 struct {
// 	T1 string `json:"T1"`
// 	T2 string `json:"t2"`
// }

// func main() {

// 	tt := Test123{
// 		T1: "asdasd",
// 		T2: "m,kiuyhnmju",
// 	}
// 	john := Teacher{
// 		ID:        "678930",
// 		Firstname: "John",
// 		Lastname:  "Doe",
// 		TT:        tt,
// 	}
// 	marshaled, err := json.MarshalIndent(john, "", "   ")
// 	if err != nil {
// 		log.Fatalf("marshaling error: %s", err)
// 	}
// 	fmt.Println(string(marshaled))
// }
