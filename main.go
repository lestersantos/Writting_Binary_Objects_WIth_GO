package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
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

var fileName string  = "./newDir2/test.txt" 

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
			fmt.Println("Create a default file called testFIle.txt")

			createDirectory(fileName)
			createFIle(fileName)
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

func createFIle(path string){

	//name := filepath.Base(path)
	name := path
	if ifPathExist(name) == false{
		f, err := os.Create(name)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		fmt.Println(f.Name())
		fmt.Printf("File %s created succesfully\n",f.Name())	
	}
}

func ifPathExist(path string) bool{
	exist := false

	_, err := os.Stat(path)

	fmt.Printf("%T",err)
	fmt.Println(err)

	if err == nil {
		fmt.Println("File or directory exists.")
		exist = true
	} else if os.IsNotExist(err){
		fmt.Println("File or directory does not exist")
	} else {
		fmt.Println("Error occurred:", err)
		log.Fatal(err)
	}
	return exist
}

func createDirectory(path string) {

	dirPath := filepath.Dir(path)
	
	if ifPathExist(dirPath) == false{
		err := os.MkdirAll(dirPath, 0777)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Directory %s created succesfully\n",dirPath)	
	}
}