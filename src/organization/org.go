package organization

import (
	"bufio"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"

	"os"
	"strings"
	"syscall"
)

func Org() {
	company, username, password, hash, match, err := ReadUsersInputs()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("")
	fmt.Println("--------- You have provided the following information------------")
	log.Printf("CompanyName: %s \nUsername: %s \nPassword: %s \nHashedPassword: %s \nMatch:  %v \n", company, username, password, string(hash), match)
	log.Println("---------End------------")

}

func ReadUsersInputs() (string, string, string, []byte, bool, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Your FirstName: ")
	firstName, err := reader.ReadString('\n')
	if len(strings.TrimSpace(firstName)) == 0 {
		err = fmt.Errorf("your FirstName  can't be empty %v", firstName)
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if err != nil {
		return "", "", "", nil, false, err
	}

	fmt.Print("Enter Your LastName: ")
	lastName, err := reader.ReadString('\n')
	if len(strings.TrimSpace(lastName)) == 0 {
		err = fmt.Errorf("your LastName  can't be empty %v", lastName)
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if err != nil {
		return "", "", "", nil, false, err
	}

	fmt.Print("Enter Your OrganisationName: ")
	companyName, err := reader.ReadString('\n')
	if len(strings.TrimSpace(companyName)) == 0 {
		err = fmt.Errorf("your Company Name  can't be empty %v", companyName)
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if err != nil {
		return "", "", "", nil, false, err
	}
	// create username
	username := firstName[0:1] + lastName

	fmt.Print("Enter Password: ")
	bytePassword, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		return "", "", "", nil, false, nil
	}

	password := string(bytePassword)
	// hash the password
	hash, err := PasswordHash(bytePassword)
	if err != nil {
		return "", "", "", nil, false, err
	}
	//check if matches
	passwordMatch := HashPasswordCheck(bytePassword, hash)
	return strings.TrimSpace(companyName), strings.TrimSpace(username), strings.TrimSpace(password), hash, passwordMatch, nil
}
func PasswordHash(password []byte) ([]byte, error) {
	resBytes, err := bcrypt.GenerateFromPassword(password, 15)
	return resBytes, err
}
func HashPasswordCheck(password, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}
