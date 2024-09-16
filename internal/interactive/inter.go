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

func (i *Inter) Run(col *pcollection.Collection) {
	i.RenderList(col)
}
