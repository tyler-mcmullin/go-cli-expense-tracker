package expense

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func loadExpenses() ([]Expense, error) {
	path, err := getExpenseFilePath()
	data, err := os.ReadFile(path)

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
	path, err := getExpenseFilePath()

	data, err := json.MarshalIndent(
		expenses,
		"",
		"  ",
	)

	if err != nil {
		return err
	}

	return os.WriteFile(
		path,
		data,
		0644,
	)
}

func getExpenseFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not get home directory: %w\n", err)
	}

	dir := filepath.Join(home, ".expense-tracker")

	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return "", fmt.Errorf("could not create expense directory: %w\n", err)
	}

	return filepath.Join(dir, "expenses.json"), nil
}
