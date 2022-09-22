package utils

import "fmt"

type FaimlyAccount struct {
	// declare fields name & type
	// declare a field for user input option
	opt string
	// declare a field to control if break the for loop
	loop bool
	// balance
	balance float64
	// deposit/withdraw money
	money float64
	// transaction details
	notes string
	// flag to indicate if there's any transactions
	flag bool
	// declare a detail string to print out the transaction details
	details string
}

// factory mode constructor, return a *FamilyAccount instance
func NewFamilyAccount() *FaimlyAccount {
	return &FaimlyAccount{
		opt:     "",
		loop:    true,
		balance: 100.0,
		money:   0.0,
		notes:   "",
		flag:    false,
		details: "Flag\t\tBalance\t\tMoney\t\tNotes",
	}
}

// option 1 Show details
func (account *FaimlyAccount) showDetails() {
	fmt.Println("-------------- Show FamilyAccount Details--------------")
	if account.flag {
		fmt.Println(account.details)
	} else {
		fmt.Println("Cannot find any transaction records")
	}
}

// option 2 Deposit
func (account *FaimlyAccount) deposit() {
	fmt.Println("Enter deposit $")
	fmt.Scanln(&account.money)
	account.balance += account.money
	fmt.Println("Enter notes for deposit transaction")
	fmt.Scanln(&account.notes)
	account.details += fmt.Sprintf("\nDeposit\t\t%v\t\t%v\t\t%v", account.balance, account.money, account.notes)
	account.flag = true
}

// option 3 Withdraw
func (account *FaimlyAccount) withdraw() {
	fmt.Println("Enter withdraw $")
	fmt.Scanln(&account.money)
	if account.money > account.balance {
		fmt.Printf("Your balance is not sufficient: $%v\n", account.balance)
		account.MainMenu()
	}
	account.balance -= account.money
	fmt.Println("Enter notes for withdraw transaction")
	fmt.Scanln(&account.notes)
	account.details += fmt.Sprintf("\nWithdraw\t\t%v\t\t%v\t\t%v", account.balance, account.money, account.notes)
	account.flag = true
}

// option 4 Quit
func (account *FaimlyAccount) quit() {
	fmt.Println("Are you sure to exit? Enter y/n")
	choice := ""
	fmt.Scanln(&choice)
	for {
		if choice == "y" || choice == "n" {
			break
		} else {
			fmt.Println("Enter something wrong")
		}
	}
	if choice == "y" {
		account.loop = false
	}
}

// Show main menu
func (account *FaimlyAccount) MainMenu() {
	for {
		fmt.Println("-------------- Family Income Software--------------")
		fmt.Println("1. Income Details")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Quit")
		fmt.Println("Please select the option (1-4)")
		fmt.Scanln(&account.opt)
		switch account.opt {
		case "1":
			account.showDetails()
		case "2":
			account.deposit()
		case "3":
			account.withdraw()
		case "4":
			account.quit()
		default:
			fmt.Println("The option number is not valid")
		}
		if !account.loop {
			break
		}
	}
	fmt.Println("You have been logged out")
}
