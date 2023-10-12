package vault

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// type Config struct {
// 	tokens string `yaml:"tokens"`
// }

func readConfig(path string) error {
	// Person struct represents the person key in YAML.
	type Person struct {
		Name  string `yaml:"name"`
		Age   int    `yaml:"age"`
		Email string `yaml:"email"`
	}

	var err error
	// var fi os.FileInfo
	// var tokens Config

	// if path != "" {
	// 	_, err = os.Stat(path)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	data, err := os.ReadFile(path)

	// 	if err != nil {
	// 		return err
	// 	}

	// 	if err := yaml.Unmarshal(data, &tokens); err != nil {
	// 		panic(err)
	// 	}
	// }
	// log.Printf("tokens: %v", tokens.tokens)

	// if len(tokens.tokens) < 1 {
	// 	return errors.New("no tokens found in config")
	// }
	// read the output.yaml file

	var data []byte
	data, err = os.ReadFile("output.yaml")

	if err != nil {
		log.Fatal(err)
	}

	// create a person struct and deserialize the data into that struct
	var person Person

	if err := yaml.Unmarshal(data, &person); err != nil {
		panic(err)
	}

	log.Println(person)

	// print the fields to the console
	log.Printf("Name: %s\n", person.Name)
	log.Printf("Age: %d\n", person.Age)
	log.Printf("Email: %s\n", person.Email)
	return nil
}
