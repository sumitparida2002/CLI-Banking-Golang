package main

import (
	"flag"
	"fmt"

	"github.com/sumitparida2002/CLI-Banking-Golang/pkg/model"
	// "os"
)

func main() {

	model.Init()
	// model.CreateAccount()
	// model.GetAllAccounts()
	// model.GetAccountByID("5")
	// model.DeleteAccount("1")

	rolePtr := flag.String("role", "user", "Defined Role")

	if *rolePtr == "user" {
		fmt.Println("Enter Your Pin:")

		var pin int
		fmt.Scanln(&pin)

	}

	flag.Parse()

	fmt.Println("Role:", *rolePtr)

	// fmt.Println(argsWithoutProg)
}
