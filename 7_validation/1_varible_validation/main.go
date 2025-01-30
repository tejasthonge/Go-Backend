package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
)

func validateEmail(email string) error {
	err := validator.New().Var(email, "required,email")
	return err
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the email")
	email, _ := reader.ReadString('\n')
	trimEmail := strings.TrimSpace(email)
	err := validateEmail(trimEmail)
	if err != nil {
		fmt.Println(err.Error())

	} else {
		fmt.Println("email ", trimEmail, "is validate")
	}

}
