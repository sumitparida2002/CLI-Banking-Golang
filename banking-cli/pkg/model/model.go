package model

import (
	// "errors"
	"fmt"

	"github.com/sumitparida2002/CLI-Banking-Golang/pkg/config"
	"gorm.io/gorm"
)

var access string

var adminKey string = "yip"

type AccountDetails struct {
	gorm.Model
	AccountHolder string
	AccountNo     string
	Password      string
	Balance       int
}

var db *gorm.DB

func Init(role string) (adminKey string) {
	db = config.Connect()
	access = role
	db.AutoMigrate(&AccountDetails{})
	if access == "admin" {
		return adminKey
	}
	return
}

func CreateAccount() {
	if access == "user" {
		println("Permission Denied")
		return
	}
	print("Creating Account")
	db.Create(&AccountDetails{AccountHolder: "Sumit", AccountNo: "5", Password: "420", Balance: 10})
}
func GetAllAccounts() {
	if access == "user" {
		println("Permission Denied")
		return
	}
	var bankAccounts []AccountDetails
	db.Find(&bankAccounts)

	for i := 0; i < len(bankAccounts); i++ {
		fmt.Printf("%s \n", bankAccounts[i].AccountHolder)
	}

}

func GetAccountByID(accountNo string) (result AccountDetails) {
	var bankAccount AccountDetails
	db.Where("account_no").Find(&bankAccount)
	return bankAccount

}

func DeleteAccount(accountNo string) {
	if access == "user" {
		println("Permission Denied")
		return
	}
	db.Where("account_no = ?", accountNo).Delete(&AccountDetails{})
}

//Wrong PIN Deactivate
