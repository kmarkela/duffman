package interactive

import (
	"net/http"
	"net/url"

	"github.com/kmarkela/duffman/internal/pcollection"
)

type Inter struct {
	tr *http.Transport
}

func New(proxy string) (Inter, error) {

	var inter = Inter{}
	var tr = http.Transport{}

	if proxy != "" {
		// parse proxy
		proxyUrl, err := url.Parse(proxy)
		if err != nil {
			return inter, err
		}

		tr = http.Transport{Proxy: http.ProxyURL(proxyUrl)}

	}

	inter.tr = &tr

	return inter, nil

}

// func printSh(tab int, nl pcollection.NodeList) {
// 	for _, v := range nl {

// 		if v.Node == nil {
// 			fmt.Printf("%s%s : %s\n", strings.Repeat("-", tab), v.Name, v.Req.URL)
// 			continue
// 		}

// 		fmt.Printf("%s %s\n", strings.Repeat("-", tab), v.Name)

// 		printSh(tab+2, v.Node)

// 	}
// }

func (i *Inter) Run(col *pcollection.Collection) {

	// for _, v := range col.Schema.Nodes {
	// printSh(0, col.Schema.Nodes)
	// }

	RenderList(col.Schema)
}
