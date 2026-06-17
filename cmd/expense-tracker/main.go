package main

import (
	"fmt"
	"os"

	"github.com/tyler-mcmullin/go-cli-expense-tracker/expense"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Command requires arguments - Try \"expense-tracker --help\"")
		return
	}

	switch os.Args[1] {

	case "add":
		err := expense.AddExpense()

		if err != nil {
			fmt.Println("Could not add expense:", err)
			return
		}
	case "delete":
		err := expense.DeleteExpense()

		if err != nil {
			fmt.Println("Could not delete expense: ", err)
			return
		}
	case "list":
		err := expense.ListExpenses()

		if err != nil {
			fmt.Println("Could not list expenses: ", err)
			return
		}
	case "summary":
		err := expense.ShowSummary()

		if err != nil {
			fmt.Println("Could not show summary: ", err)
			return
		}
	case "help":
		err := expense.ShowHelp()

		if err != nil {
			fmt.Println("Could not show help: ", err)
			return
		}
	default:
		fmt.Println("Unknown command: ", os.Args[1], " - Try \"expense-tracker help\"")
	}

}
