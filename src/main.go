package main

import (
	// "github.com/mysubsnewslivecom/learn-golang/src/apis"
	// "github.com/mysubsnewslivecom/learn-golang/src/greetings"
	// "github.com/mysubsnewslivecom/learn-golang/src/menu"
	"log"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"
	"github.com/mysubsnewslivecom/learn-golang/src/apis"
	"github.com/mysubsnewslivecom/learn-golang/src/utils"
	"github.com/spf13/viper"
)

func init() {

	viper.SetConfigFile("/workspaces/learn-golang/.env")

	// viper.AddConfigPath("")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	log.Println(".env file loaded")

	var env utils.Environment
	// parse & validate environment variables
	_, err := flags.Parse(&env)
	if err != nil {
		if flags.WroteHelp(err) {
			os.Exit(0)
		}
		log.Fatalf("unable to parse environment variables: %v", err)
	}

	log.Printf("VaultApproleSecretIDFile: %v", env.VaultApproleSecretIDFile)

	// var hcpClientId, hcpClientSecret = os.Getenv("HCP_CLIENT_ID"), os.Getenv("HCP_CLIENT_SECRET")

	// var hcpAPIToken string
	// if hcpAPIToken, envErr := os.LookupEnv("HCP_API_TOKEN"); !envErr {
	// 	hcpAPIToken, err = vault.GetClientToken(hcpClientId, hcpClientSecret)
	// }
	// _ = os.Setenv("HCP_API_TOKEN", hcpAPIToken)

}

func main() {

	// var config utils.Config
	appConfigPath := "/workspaces/learn-golang/env/"
	config, err := utils.LoadConfig(appConfigPath)
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	var sb strings.Builder

	sb.WriteString(config.DBProtocol)
	sb.WriteString("://")
	sb.WriteString(config.DBUser)
	sb.WriteString(":")
	sb.WriteString(config.DBPassword)
	sb.WriteString("@")
	sb.WriteString(config.DBHost)
	sb.WriteString(":")
	sb.WriteString(config.DBPort)
	sb.WriteString("/")
	sb.WriteString(config.DBSchema)
	sb.WriteString("?")
	sb.WriteString(config.DBSSL)

	log.Printf("DBDriver: %v", config.DBDriver)
	log.Printf("ServerAddress: %v", config.ServerAddress)
	log.Printf("DBSource: %v", sb.String())

	apis.GetAnything()

	// organization.PhoneBook()
	// greetings.Greetings("Linux")
	// apis.BoredApi()
	// menu.MenuStart()
	// log.Println(env.DatabaseHostname)
	// vault.Unseal()
	// vault.GetSecret()
	// vault.ReadConfigVault()
	// vault.ReadUser()
	// vault.UsingToken()

	utils.Main()

}
