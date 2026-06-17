package expense

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func AddExpense() error {
	addCmd := flag.NewFlagSet("add", flag.ContinueOnError)

	description := addCmd.String("description", "", "expense description")
	amount := addCmd.Float64("amount", 0, "expense amount")

	err := addCmd.Parse(os.Args[2:])
	if err != nil {
		return err
	}

	expenses, err := loadExpenses()
	if err != nil {
		return err
	}

	id, err := getNextID(expenses)
	if err != nil {
		return err
	}

	newExpense := Expense{
		ID:          id,
		Description: *description,
		Amount:      *amount,
		Date:        time.Now().Format("2006-01-02"),
	}

	expenses = append(expenses, newExpense)

	err = saveExpenses(expenses)
	if err != nil {
		return err
	}

	fmt.Println("Expense added successfully (ID: ", id, ")")

	return nil
}

func DeleteExpense() error {
	delCmd := flag.NewFlagSet("delete", flag.ContinueOnError)

	id := delCmd.Int("id", 0, "expense ID to delete")

	err := delCmd.Parse(os.Args[2:])
	if err != nil {
		return err
	}

	expenses, err := loadExpenses()
	if err != nil {
		return err
	}

	found := false

	for i, expense := range expenses {
		if expense.ID == *id {
			expenses = append(expenses[:i], expenses[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("Expense with id %d not found", *id)
	}

	err = saveExpenses(expenses)
	if err != nil {
		return fmt.Errorf("Could not save expenses: %w", err)
	}

	fmt.Println("Expense deleted successfully (ID: ", *id, ")")

	return nil
}

func ListExpenses() error {
	listCmd := flag.NewFlagSet("list", flag.ContinueOnError)

	listLen := listCmd.Int("len", 30, "length of list to output")
	all := listCmd.Bool("all", false, "list all expenses")

	err := listCmd.Parse(os.Args[2:])
	if err != nil {
		return err
	}

	expenses, err := loadExpenses()
	if err != nil {
		return err
	}

	var expensesToShow []Expense

	if *all {
		expensesToShow = expenses
	} else {
		limit := *listLen
		if limit > len(expenses) {
			limit = len(expenses)
		}
		expensesToShow = expenses[:limit]
	}

	total := 0

	fmt.Printf("%-5s %-12s %-20s %12s\n",
		"ID",
		"Date",
		"Description",
		"Amount",
	)

	fmt.Println("----------------------------------------------------")

	for _, expense := range expensesToShow {
		fmt.Printf("%-5d %-12s %-20s %12s\n",
			expense.ID,
			expense.Date,
			expense.Description,
			fmt.Sprintf("$%.2f", expense.Amount),
		)
		total++
	}
	fmt.Println(total, " expenses shown")

	return nil
}

func ShowSummary() error {
	sumCmd := flag.NewFlagSet("summary", flag.ContinueOnError)

	month := sumCmd.Int("month", 0, "month to show total for, will show all if field empty")

	err := sumCmd.Parse(os.Args[2:])
	if err != nil {
		return err
	}

	if *month < 0 || *month > 12 {
		return fmt.Errorf("Invalid month\n")
	}

	expenses, err := loadExpenses()
	if err != nil {
		return err
	}

	total := 0.0

	if *month == 0 {
		for _, expense := range expenses {
			total += expense.Amount
		}

		fmt.Printf("Total expenses: $%.2f\n", total)
		return nil
	}

	fmt.Printf("%-5s %-12s %-20s %12s\n",
		"ID",
		"Date",
		"Description",
		"Amount",
	)

	fmt.Println("----------------------------------------------------")

	for _, expense := range expenses {

		date, err := time.Parse("2006-01-02", expense.Date)
		if err != nil {
			return err
		}

		if int(date.Month()) == *month {
			fmt.Printf(
				"%-5d %-12s %-20s %12s\n",
				expense.ID,
				expense.Date,
				expense.Description,
				fmt.Sprintf("$%.2f", expense.Amount),
			)

			total += expense.Amount
		}
	}

	fmt.Printf("Total expenses for the month of %s: $%.2f\n", months[*month], total)

	return nil
}

func ShowHelp() error {
	fmt.Println("")
	fmt.Println("Command format: expense-tracker command --flag value")
	fmt.Println("Commands with flags surrounded by \"*\" are optional and command have default behaviors")
	fmt.Println("")
	fmt.Println("Command\tflags(values)")
	fmt.Println("add\tdescription(string), amount(float)")
	fmt.Println("delete\tid(integer)")
	fmt.Println("list\t*len(integer(default is 30))*, *all*")
	fmt.Println("summary\t*month*(1 -> 12(0 is default and will total all expenses)")
	return nil
}
