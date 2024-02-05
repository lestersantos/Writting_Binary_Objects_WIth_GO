package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	fp "path/filepath"
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
			courseRegister()
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

func courseRegister(){
	
	if ifFilePathExist(fileName) == false {
		createDirectory(fileName)
		createFile(fileName)
	}

	err := readFile(fileName)

	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return 
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

func createFile(filePath string){

	//name := filefilePath.Base(filePath)
	name := filePath

	file, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println(file.Name())
	fmt.Printf("File %s created succesfully\n",file.Name())	
}

func ifFilePathExist(filePath string) bool{
	exist := false

	_, err := os.Stat(filePath)

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

func createDirectory(filePath string) {

	dirfilePath := fp.Dir(filePath)

	err := os.MkdirAll(dirfilePath, 0777)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Directory %s created succesfully\n",dirfilePath)	
}

func openFile(filePath string)(*os.File, error){
	// I believe that the perm 0755 isn't necessary unless we're creating the file
	file, err := os.OpenFile(filePath, os.O_RDWR, 0755)
	if err != nil {
		fmt.Println("Error opening file ",err)
		return nil, err
	}
	return file, nil
}

func readFile(filePath string) error{
	
	file, err := openFile(filePath)

	if err != nil {
		return err
	}
	
	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Iterate through each line in the file
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}