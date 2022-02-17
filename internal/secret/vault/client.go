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

func New() (*Provider, error) {
	cfg := config.GetConf()
	client, clientError := api.NewClient(&api.Config{Address: cfg.Vault.VaultAddress})
	if clientError != nil {
		return nil, fmt.Errorf("Vault client error")
	}
	client.SetToken(cfg.Vault.VaultToken)
	return &Provider{
		path:    cfg.Vault.VaultPath,
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
