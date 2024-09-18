package client

import (
	"net/http"
	"net/url"

	"github.com/kmarkela/duffman/internal/pcollection"
)

type Client struct {
	tr *http.Transport
}

func New(proxy string) (Client, error) {

	var inter = Client{}
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

func (c *Client) Run(col *pcollection.Collection) {
	c.RenderList(col)
}
