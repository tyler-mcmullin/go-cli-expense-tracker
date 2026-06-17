package expense

type Expense struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Date        string  `json:"date"`
}

func getNextID(expenses []Expense) (int, error) {
	maxID := 0

	for _, expense := range expenses {
		if expense.ID > maxID {
			maxID = expense.ID
		}
	}

	return maxID + 1, nil
}
