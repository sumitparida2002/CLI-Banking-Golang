package main

import (
	"flag"
	"fmt"
	"os"

	// "fmt"
	"github.com/sumitparida2002/CLI-Banking-Golang/pkg/model"
	// "os"
)

func main() {

	rolePtr := flag.String("role", "user", "Defined Role")

	flag.Parse()

	model.Init(*rolePtr)

	if *rolePtr == "user" {
		fmt.Println("Enter Your Account No.")
		var accountNo string
		fmt.Scan(&accountNo)

		result := model.GetAccountByID(accountNo)

		var password string
		fmt.Scan(&password)

		if password != result.Password {
			print("Invalid Password")
			os.Exit(1)
		}

		print("Welcome")
		print("Would you like to \n 1.Check Balance\n 2.Withdraw \n 3.Transfer")
		var option int

		switch option {
		case 1:
			print(result.Balance)

		case 2:
			print("Not Done Yet lilbro")

		case 3:
			print("Not Done Yet lilbro")

		default:
			print("Please select the option")
		}

		os.Exit(0)

	}

	print("Welcome")
	print("Would you like to \n 1.Get All Accounts\n 2.Get Specific Account \n 3.Update Account Information \n 4.Delete Account \n 5.Deactivate Account  ")
	var option int

	switch option {
	case 1:
		print("Not done")

	case 2:
		print("Not Done Yet lilbro")

	case 3:
		print("Not Done Yet lilbro")

	default:
		print("Please select the option")
	}

	os.Exit(0)

}
