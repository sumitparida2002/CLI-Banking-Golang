package main

import (
	"flag"
	"fmt"

	"github.com/sumitparida2002/CLI-Banking-Golang/pkg/config"
	// "os"
)

func main() {

	config.Connect()

	rolePtr := flag.String("role", "user", "Defined Role")

	flag.Parse()

	fmt.Println("Role:", *rolePtr)

	// fmt.Println(argsWithoutProg)
}
