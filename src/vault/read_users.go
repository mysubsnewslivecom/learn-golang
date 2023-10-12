package vault

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type User struct {
	Name       string
	Occupation string
}

func ReadUser() {

	yfile, err := os.ReadFile("./src/vault/users.yaml")
	if err != nil {
		log.Fatal(err)
	}

	data := make(map[string]User)
	err2 := yaml.Unmarshal(yfile, &data)
	if err2 != nil {
		log.Fatal(err2)
	}

	for k, v := range data {
		log.Printf("%s: %s\n", k, v)
	}
}
