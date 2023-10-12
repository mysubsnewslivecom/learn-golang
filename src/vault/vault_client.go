package vault

import (
	"log"
	"net/http"
	"os"
	"time"

	vault "github.com/hashicorp/vault/api"
)

// newVaultClient creates a new client for connecting to vault.
func newVaultClient() (*vault.Client, error) {
	log.Printf("[INFO] Creating vault/api client")

	config := vault.DefaultConfig()

	config.Address = os.Getenv("VAULT_ADDR")
	config.HttpClient = &http.Client{
		Timeout: 10 * time.Second,
	}

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("unable to initialize Vault client: %v", err)
	}

	if config.Address != "" {
		log.Printf("[DEBUG] Setting vault address to %s", config.Address)
	}

	// Authenticate
	vaultToken, ok := os.LookupEnv("VAULT_TOKEN")
	if !ok {
		log.Fatal("VAULT_TOKEN is not set")
	}
	client.SetToken(vaultToken)
	// status, _ := client.Sys().SealStatus()

	// log.Printf("%v", status)

	return client, nil
}
