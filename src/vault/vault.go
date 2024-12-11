package vault

import (
	"context"
	"log"

	vault "github.com/hashicorp/vault/api"
)

// This is the accompanying code for the Developer Quick Start.
// WARNING: Using root tokens is insecure and should never be done in production!
func UsingToken() {
	// config := vault.DefaultConfig()

	// // config.Address = "http://vault-server:8200"
	// config.Address = os.Getenv("VAULT_ADDR")

	// config.HttpClient = &http.Client{
	// 	Timeout: 10 * time.Second,
	// }

	// client, err := vault.NewClient(config)
	// if err != nil {
	// 	log.Fatalf("unable to initialize Vault client: %v", err)
	// }

	// // Authenticate
	// vaultToken, ok := os.LookupEnv("VAULT_TOKEN")
	// if !ok {
	// 	log.Fatal("VAULT_TOKEN is not set")
	// }
	// client.SetToken(vaultToken)

	client, _ := newVaultClient()

	secretData := map[string]interface{}{
		"password": "Hashi123",
	}
	// Write a secret
	_, err := client.KVv2("secret/data").Put(context.Background(), "my-secret-password", secretData)
	if err != nil {
		log.Fatalf("unable to write secret: %v", err)
	}

	log.Printf("[INFO] Secret written successfully.")

	// Read a secret from the default mount path for KV v2 in dev mode, "secret"
	secret, err := client.KVv2("secret/data").Get(context.Background(), "my-secret-password")
	if err != nil {
		log.Fatalf("unable to read secret: %v", err)
	}

	value, ok := secret.Data["password"].(string)
	if !ok {
		log.Fatalf("value type assertion failed: %T %#v", secret.Data["password"], secret.Data["password"])
	}

	if value != "Hashi123" {
		log.Fatalf("unexpected password value %q retrieved from vault", value)
	}

	log.Printf("[INFO] Access granted!")
}

// // newVaultClient creates a new client for connecting to vault.
// func newVaultClient() (*vault.Client, error) {
// 	log.Printf("[INFO] Creating vault/api client")

// 	config := vault.DefaultConfig()

// 	config.Address = os.Getenv("VAULT_ADDR")
// 	config.HttpClient = &http.Client{
// 		Timeout: 10 * time.Second,
// 	}

// 	client, err := vault.NewClient(config)
// 	if err != nil {
// 		log.Fatalf("unable to initialize Vault client: %v", err)
// 	}

// 	if config.Address != "" {
// 		log.Printf("[DEBUG] Setting vault address to %s", config.Address)
// 	}

// 	// Authenticate
// 	vaultToken, ok := os.LookupEnv("VAULT_TOKEN")
// 	if !ok {
// 		log.Fatal("VAULT_TOKEN is not set")
// 	}
// 	client.SetToken(vaultToken)
// 	// status, _ := client.Sys().SealStatus()

// 	// log.Printf("%v", status)

// 	return client, nil
// }

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func unsealVault(vc *vault.Client, initResponse *vault.InitResponse) string {
	log.Printf("Unseal vault")
	_, err := vc.Sys().Unseal(initResponse.Keys[0])
	fatal(err)
	return initResponse.RootToken
}

func initializeNewVault(vc *vault.Client) *vault.InitResponse {
	log.Printf("Initialize fresh vault")
	vaultInit := &vault.InitRequest{
		SecretShares:    1,
		SecretThreshold: 1,
	}
	initResponse, err := vc.Sys().Init(vaultInit)
	fatal(err)

	return initResponse

}

func InitializeVault() string {
	vc, _ := newVaultClient()
	initResponse := initializeNewVault(vc)
	return unsealVault(vc, initResponse)
}
