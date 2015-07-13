package shrutigo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path"
)

type Provider struct {
	Id          int    `json:"id, omitempty" db:"id"`
	Name        string `json:"name" db:"name"`
	DisplayName string `json:"display_name" db:"display_name"`
	Description string `json:"description, omitempty" db:"description"`
	WebURL      string `json:"web_url, omitempty" db:"web_url"`
	IconURL     string `json:"icon_url, omitempty" db:"icon_url"`
	Active      bool   `json:"active, omitempty" db:"active"`
}

func (client *Client) GetAllProviders() (providers []*Provider, err error) {

	url := client.Protocol + path.Join(client.Host, "providers")

	response, err := http.Get(url)
	if err != nil {
		return
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	providers = make([]*Provider, 0)
	err = json.Unmarshal(contents, &providers)
	if err != nil {
		return
	}

	return
}

func (client *Client) GetSingleProvider(providerName string) (p *Provider, err error) {

	url := client.Protocol + path.Join(client.Host, "providers", providerName)

	response, err := http.Get(url)
	if err != nil {
		return
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	p = new(Provider)
	err = json.Unmarshal(contents, p)
	if err != nil {
		return
	}

	return
}
