package vault

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"net/http"
	"time"
)

/*
token := "s.Cz26eGSOO8O854IyCzlVPGtN"
	vaultAddr := "http://127.0.0.1:8200"
*/
var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

type Provider struct {
	path   string
	client *api.Logical
	result map[string]interface{}
}

func New(token, addr, path string) (*Provider, error) {
	client, clientError := api.NewClient(&api.Config{Address: addr})
	if clientError != nil {
		return nil, fmt.Errorf("Vault client error")
	}
	client.SetToken(token)
	return &Provider{
		path:   path,
		client: client.Logical(),
		result: map[string]interface{}{},
	}, nil
}
func (p Provider) VaultSecretWrite(path string, data map[string]interface{}) error {
	if _, err := p.client.Write(p.path+"/"+path, data); err != nil {
		return err
	}
	return nil
}
