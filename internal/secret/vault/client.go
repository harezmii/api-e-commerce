package vault

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"net/http"
	"time"
)

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func VaultSecretWrite() bool {
	token := "s.Cz26eGSOO8O854IyCzlVPGtN"
	vaultAddr := "http://127.0.0.1:8200"

	client, clientError := api.NewClient(&api.Config{Address: vaultAddr, HttpClient: httpClient})
	if clientError != nil {
		return false
	}
	client.SetToken(token)

	inputData := map[string]interface{}{
		"data": map[string]interface{}{
			"first": "ankit",
		},
	}
	output, err := client.Logical().Write("secret/data/abd", inputData)
	fmt.Println(output)
	if err != nil {
		return false
	}
	return true
}

func VaultSecretRead() map[string]interface{} {
	token := "s.Cz26eGSOO8O854IyCzlVPGtN"
	vaultAddr := "http://127.0.0.1:8200"

	client, clientError := api.NewClient(&api.Config{Address: vaultAddr, HttpClient: httpClient})
	if clientError != nil {
		return nil
	}
	client.SetToken(token)

	output, err := client.Logical().Read("secret/data/abd")
	fmt.Println(output.Data)
	if err != nil {
		return nil
	}
	return output.Data
}
