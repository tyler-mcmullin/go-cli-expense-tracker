package main

import (
	"encoding/json"
	"os"
)

func loadExpenses() ([]Expense, error) {
	data, err := os.ReadFile("expenses.json")

	if err != nil {
		if os.IsNotExist(err) {
			return []Expense{}, nil
		}

		return nil, err
	}

	var expenses []Expense

	err = json.Unmarshal(data, &expenses)
	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func saveExpenses(expenses []Expense) error {
	data, err := json.MarshalIndent(
		expenses,
		"",
		" ",
	)

	if err != nil {
		return err
	}

	return os.WriteFile(
		"expenses.json",
		data,
		0644,
	)
}
