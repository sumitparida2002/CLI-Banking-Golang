package model

import (
	"fmt"

	"github.com/sumitparida2002/CLI-Banking-Golang/pkg/config"
	"gorm.io/gorm"
)

type BankAccount struct {
	gorm.Model
	AccountNo string `gorm:"primaryKey"`
	Name      string
	Balance   int
}

var db *gorm.DB

func Init() {
	db = config.Connect()
	db.AutoMigrate(&BankAccount{})
}

func CreateAccount() {
	db.Create(&BankAccount{AccountNo: "5", Name: "Sumit"})
}
func GetAllAccounts() {
	var bankAccounts []BankAccount
	db.Find(&bankAccounts)

	for i := 0; i < len(bankAccounts); i++ {
		fmt.Printf("%s \n", bankAccounts[i].Name)
	}

}

func GetAccountByID(accountNo string) {
	var bankAccount BankAccount
	db.Where("account_no").Find(&bankAccount)
	fmt.Printf("%s %d\n", bankAccount.Name, bankAccount.ID)
}

func DeleteAccount(accountNo string) {
	db.Where("account_no = ?", accountNo).Delete(&BankAccount{})
}
