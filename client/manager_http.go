package client

import (
	"net/http"
	"net/url"

	"github.com/ory-am/fosite"
	"github.com/ory-am/hydra/pkg"
)

type HTTPManager struct {
	Client   *http.Client
	Endpoint *url.URL
}

func (m *HTTPManager) GetClient(id string) (fosite.Client, error) {
	var c fosite.DefaultClient
	var r = pkg.NewSuperAgent(pkg.JoinURL(m.Endpoint, id).String())
	r.Client = m.Client
	if err := r.Get(&c); err != nil {
		return nil, err
	}

	return &c, nil
}

func (m *HTTPManager) CreateClient(c *fosite.DefaultClient) error {
	var r = pkg.NewSuperAgent(m.Endpoint.String())
	r.Client = m.Client
	return r.Create(c)
}

func (m *HTTPManager) DeleteClient(id string) error {
	var r = pkg.NewSuperAgent(pkg.JoinURL(m.Endpoint, id).String())
	r.Client = m.Client
	return r.Delete()
}

func (m *HTTPManager) GetClients() (map[string]*fosite.DefaultClient, error) {
	cs := make(map[string]*fosite.DefaultClient)
	var r = pkg.NewSuperAgent(m.Endpoint.String())
	r.Client = m.Client
	if err := r.Get(&cs); err != nil {
		return nil, err
	}

	return cs, nil
}
