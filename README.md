# Expense Tracker CLI

A command-line expense tracker built in Go that allows users to manage, view, and summarize personal expenses. The application stores data locally as JSON and provides a simple CLI interface for adding, deleting, listing, and summarizing expenses by month or all-time.

## Features

- Add expenses with:
  - Description
  - Amount
  - Automatically recorded date
  - Automatically generated unique expense ID

- List expenses:
  - View recent expenses
  - Display all expenses
  - Control output length

- Delete expenses by ID

- Generate summaries:
  - View total spending
  - Filter expenses by month

- Persistent local storage using JSON

- Global CLI installation support

## Demo

Example usage:

```bash
expense-tracker add --description "Dinner" --amount 20.50
```

Output:

```
Expense added successfully (ID: 1 )
```

List expenses:

```bash
expense-tracker list
```

Example:

```
ID    Date         Description             Amount
----------------------------------------------------
1     2026-06-17   Lunch                   $12.50
2     2026-06-17   Gas                     $40.00

2 expenses shown
```

Generate a summary:

```bash
expense-tracker summary --month 6
```

Output:

```
ID    Date         Description                Amount
----------------------------------------------------
1     2026-06-17   Dinner                     $14.00
3     2026-06-17   Food                       $50.00

Total expenses for June: $64.00
```

Or view total expenditure:

```bash
expense-tracker summary
```

Output:

```
Total expenses: $64.00
```

## Installation

### Requirements

- Go 1.20+

### Install from source

Clone the repository:

```bash
git clone https://github.com/tyler-mcmullin/go-cli-expense-tracker.git
cd go-cli-expense-tracker
```

Install the CLI:

```bash
go install ./cmd/expense-tracker
```

Verify installation:

```bash
expense-tracker help
```

The executable will be installed to:

```
~/go/bin/expense-tracker
```

Make sure this directory is included in your PATH.

## Usage

### Add an expense

```bash
expense-tracker add --description "Coffee" --amount 5.75
```

### List expenses

Show expenses (this defaults to 30):

```bash
expense-tracker list
```

Change the number displayed:

```bash
expense-tracker list --len 10
```

Show all expenses:

```bash
expense-tracker list --all
```

### Delete an expense

```bash
expense-tracker delete --id 3
```

### View spending summary

Show total of all expenses:

```bash
expense-tracker summary
```

Show expenditure summary for a specific month:

```bash
expense-tracker summary --month 6
```

### Help

Shows list of commands, flags, and associated values/defaults

```bash
expense-tracker help
```

## Storage

Expenses are stored locally in:

```
~/.expense-tracker/expenses.json
```

Example data:

```json
[
  {
    "id": 1,
    "description": "Lunch",
    "amount": 12.5,
    "date": "2026-06-17"
  }
]
```

## Project Structure

```
.
├── cmd
│   └── expense-tracker
│       └── main.go
│
├── expense
│   ├── constants.go
│   ├── expense.go
│   ├── functions.go
│   └── storage.go
│
├── go.mod
└── README.md
```

## Technologies Used

- Go
- JSON
- Standard Library
  - flag
  - encoding/json
  - os
  - time

## Design Notes

The project follows a package-based Go structure:

- `cmd/expense-tracker`
  - Handles CLI input and command routing

- `expense`
  - Handles expense type structure
  - Manages ID numbers

- `constants`
  - Stores months as key:value pairs

- `storage`
  - Helper functions related to saving new JSON data and loading existing JSON data
  - Keeps data local to /.expense-tracker

- `functions`
  - Main command logic routed to by expense-tracker
  - Contains logic for add, delete, list, summary, and help commands

Storage and command logic are separated to keep the code modular and maintainable.
