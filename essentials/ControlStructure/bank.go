package main

import "fmt"

func main() {
	var balance float64 = 0
	fmt.Println("Welcome to Go Bank")
	fmt.Println("What you want to do?")
	fmt.Println("1.Deposit")
	fmt.Println("2.Withdraw")
	fmt.Println("3.Balance")
	fmt.Println("4.Exit")
	var choice int

	for choice != 4 {
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			var deposit float64
			fmt.Print("Enter the amount to deposit: ")
			fmt.Scan(&deposit)
			balance += deposit
			fmt.Println("Deposit successful. New balance:", balance)
		case 2:
			var withdraw float64
			fmt.Print("Enter the amount to withdraw: ")
			fmt.Scan(&withdraw)
			if withdraw > balance {
				fmt.Println("Insufficient funds")
			} else {
				balance -= withdraw
				fmt.Println("Withdrawal successful. New balance:", balance)
			}
		case 3:
			fmt.Println("Current balance:", balance)
		case 4:
			fmt.Println("Thank you for using Go Bank")
		default:
			fmt.Println("Invalid choice")
		}
	}
}
