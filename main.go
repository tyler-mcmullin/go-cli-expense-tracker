package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Command requires arguments - Try \"expense-tracker --help\"")
		return
	}

	switch os.Args[1] {

	case "add":
		err := addExpense()

		if err != nil {
			fmt.Println("Could not add expense:", err)
			return
		}
	case "delete":
		err := deleteExpense()

		if err != nil {
			fmt.Println("Could not delete expense: ", err)
			return
		}
	case "list":
	case "total":
	case "help":
	default:
		fmt.Println("Unknown command: ", os.Args[1], " - Try \"expense-tracker --help\"")
	}

}
