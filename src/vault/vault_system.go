package vault

import (
	"log"
	"os"

	"github.com/hashicorp/vault/api"
)

func Unseal() {
	client, _ := newVaultClient()
	status, err := client.Sys().SealStatus()
	if err != nil {
		log.Fatalf("checking seal status: %v", err)
	}
	log.Printf(status.BuildDate)
	keyPath, _ := os.LookupEnv("VAULT_KEY_FILE")
	log.Printf("%v", keyPath)
	_ = readConfig(keyPath)
	if !status.Sealed {
		var resp *api.SealStatusResponse
		resp, err = client.Sys().Unseal(client.Token())
		if err != nil {
			log.Printf("using unseal key on %v: %v", client.Address(), err)
		}
		log.Printf(resp.ClusterID)

	}
}
