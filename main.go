package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 1. Object Course
// - type: Boolean
// - Id: int
// - Code: int
// - Name: char[16]

type Course struct{
	id int
	code int
}

// Application menu
// 1. registro curso
// 2. Ver registros
// 3. salir

func main(){
	menu()
}

func menu() {
	for {
		fmt.Println("Select an option:")
		fmt.Println("1. Registro curso")
		fmt.Println("2. Ver registros")
		fmt.Println("3. Salir")

		option := getUserInput("Enter the option number: ")

		switch option {
		case 1:
			fmt.Println("You chose Option 1")
		case 2:
			fmt.Println("You chose Option 2")
		case 3:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid option. Please choose again.")
		}
	}
}

func getUserInput(prompt string) int {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return getUserInput(prompt)
	}
	return input
}