package main

import (
	"fmt"
	"marunk20/cli-ropc/adfs"
	"os"
)

func main() {


	initialChecks()

	if os.Args[1] == "login" {
		userFullName := adfs.LoginAndGetUserFullName()
		fmt.Println("Welcome ", userFullName)
	} else {
		fmt.Println("This is a POC CLI. Only login command will work as of now")
	}

}

func initialChecks() {

	if len(os.Args) != 2 {
		fmt.Println("This is a POC CLI. Only login command will work as of now")
		os.Exit(1)
	}

}
