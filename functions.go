package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func addExpense() error {
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

func deleteExpense() error {
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
