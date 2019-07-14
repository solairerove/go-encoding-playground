package crypt

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// AwesomeCrypt tbd
func AwesomeCrypt() {
	fmt.Println("\nAwesomeCrypt")

	s := "password1234"

	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

	pass := "password1234"

	err = bcrypt.CompareHashAndPassword(bs, []byte(pass))
	if err != nil {
		fmt.Println("You can't log in")
		return
	}
	fmt.Println("You logged in")
}
