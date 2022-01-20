package vault

import (
	"api/pkg/config"
	"fmt"
	"github.com/hashicorp/vault/api"
)

/*
token := "s.Cz26eGSOO8O854IyCzlVPGtN"
	vaultAddr := "http://127.0.0.1:8200"
*/

type Provider struct {
	path    string
	client  *api.Logical
	results map[string]interface{}
}

var token string = config.GetEnvironment("VAULT_TOKEN", config.STRING).(string)
var addr string = config.GetEnvironment("VAULT_ADDRESS", config.STRING).(string)
var path string = config.GetEnvironment("VAULT_PATH", config.STRING).(string)

func New() (*Provider, error) {

	client, clientError := api.NewClient(&api.Config{Address: addr})
	if clientError != nil {
		return nil, fmt.Errorf("Vault client error")
	}
	client.SetToken(token)
	return &Provider{
		path:    path,
		client:  client.Logical(),
		results: map[string]interface{}{},
	}, nil
}

func (p Provider) VaultSecretWrite(path string, data map[string]interface{}) error {
	inputData := map[string]interface{}{
		"data": map[string]interface{}{
			"first": "ankit",
		},
	}
	if _, err := p.client.Write(p.path+"/"+path, inputData); err != nil {
		return err
	}
	return nil
}

func (p *Provider) Get(v string) {
	read, err := p.client.Read(p.path + "/" + v)
	if err != nil {

	}
	fmt.Println(read.Data)
}
