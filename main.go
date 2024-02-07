package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	fp "path/filepath"
	"strconv"
	"strings"
)

// 1. Object Course
// - type: Boolean
// - Id: int
// - Code: int
// - Name: char[16]

type Course struct{
	Type bool
	Id int32
	Code int32
	Name [16]byte
}

// Application menu
// 1. registro curso
// 2. Ver registros
// 3. salir

var fileName string  = "./ht/test.txt" 
var currentOffset int64
var currentId int32

func main(){
	if ifFilePathExist(fileName) == false {
		createDirectory(fileName)
		createFile(fileName)
		currentOffset = 0
		currentId = 0
	}else{
		file, err := openFile(fileName)
		if err != nil {
			return
		}
		
		updateWritePointer(file)
		defer file.Close()	
	}
	
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
			courseRegister()
		case 2:
			fmt.Println("Showing database...")
			showRegisters()
		case 3:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid option. Please choose again.")
		}
	}
}

func showRegisters(){
	// Open bin file
	file, err := openFile(fileName)
	if err != nil {
		return
	}
	defer file.Close()
	
	readAll(file)
}

func courseRegister(){

	var newCourse Course

	// Prompt the user for input
	fmt.Println("REGISTRO DE CURSO")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Name: ")
	scanner.Scan()
	copy(newCourse.Name[:],scanner.Bytes())

	// fmt.Print("Enter ID: ")
	// scanner.Scan()
	// i64,_:= strconv.ParseInt(scanner.Text(),10,32)
	// newCourse.Id = int32(i64)
	newCourse.Id = currentId + 1
	currentId += 1;

	fmt.Print("Enter Code: ")
	scanner.Scan()
	i64,_:= strconv.ParseInt(scanner.Text(),10,32)
	newCourse.Code = int32(i64)

	fmt.Print("Enter Type: ")
	scanner.Scan()
	newCourse.Type,_ = strconv.ParseBool(scanner.Text())

	fmt.Println("Object to write:",newCourse)
	
	file, err := openFile(fileName)
	if err != nil {
		return
	}

	// Write object in bin file
	if err := writeObjectToFile(file,newCourse,currentOffset); err != nil {
		return
	}
	currentOffset += int64(binary.Size(newCourse))

	readAll(file)
	// var tempCourse Course

	// if err := ReadObject(file, &tempCourse, int64(0*binary.Size(tempCourse) + 0)); err != nil {
	// 	return
	// }

	// printObject(tempCourse)

	// if err := ReadObject(file, &tempCourse, int64(1*binary.Size(tempCourse) + 0)); err != nil {
	// 	return
	// }

	// printObject(tempCourse)
	// // Writing 10 objects to the file
	// for i := 0; i < 10; i++ {
	// 	// Set values to newData
	// 	copy(newData.Name[:], "Sergie")
	// 	newData.Id = int32(i+1)

	// 	// Write object in bin file
	// 	if err := writeObjectToFile(file,newData,int64(i * binary.Size(newData))); err != nil {
	// 		return
	// 	}
	// }

	// // Read 10 objects from bin file
	// for i := 0; i < 10; i++ {
	// 	var TempData Course
	// 	// Read object from bin file
	// 	if err := ReadObject(file, &TempData, int64(i * binary.Size(TempData) + 0)); err != nil {
	// 		return
	// 	}

	// 	// Print object
	// 	PrintData(TempData)
	// }
	

	// Close bin file
	defer file.Close()
}

func writeObjectToFile(file *os.File, data interface{}, offset  int64) error {
	tempOffset, _ := file.Seek(offset, 0)
	fmt.Println("Last Offset: ",tempOffset)

	err := binary.Write(file, binary.LittleEndian, data)
	if err != nil {
		fmt.Println("Err WriteObject==",err)
		return err
	}
	return nil
}

func ReadObject(file *os.File, data interface{}, offset  int64) error {
	file.Seek(offset, 0)
	err := binary.Read(file, binary.LittleEndian, data)
	if err != nil {
		fmt.Println("Err ReadObject==",err)
		return err
	}
	return nil
}

func readAll(file *os.File) error {
	var tempCourse Course

	tempOffset, _ := file.Seek(0, 0)

	fmt.Println(strings.Repeat("-",60))
	fmt.Println("\t\t   COURSE'S TABLE ")
	fmt.Println(strings.Repeat("-",60))
	fmt.Printf("| %5s%-5s | %12s%-8s | %8s%-5s | %5s%-5s \n", "ID","", "NAME","", "CODE","","TYPE","")

	for {
		
		err := binary.Read(file, binary.LittleEndian, &tempCourse)
		if err != nil {
			log.Print(err)
			fmt.Println("Err ReadObject==",err)
			break
		}else {
			printObject(tempCourse)

			tempOffset,_= file.Seek(tempOffset+int64(binary.Size(tempCourse)),0)
		}

	}
	return nil
}

func updateWritePointer(file *os.File) error {
	var tempCourse Course

	tempOffset, _ := file.Seek(0, 0)

	for {
		
		err := binary.Read(file, binary.LittleEndian, &tempCourse)
		if err != nil {
			fmt.Printf("Last id: %d next id: %d\n",tempCourse.Id,tempCourse.Id+1)
			fmt.Println("Err ReadObject==",err)
			currentOffset = tempOffset
			currentId = tempCourse.Id
			break
		}else {
			tempOffset,_= file.Seek(tempOffset+int64(binary.Size(tempCourse)),0)
			currentOffset = tempOffset
			fmt.Println("Updating position: ",tempOffset)
		}

	}
	return nil
}

func printObject(data Course){
	// fmt.Println(fmt.Sprintf("Name: %s, id: %d", string(data.Name[:]), data.Id))

	//fmt.Printf("%d %s %v %d \n",data.Id,string(data.Name[:]),data.Type, data.Code)

	//fmt.Printf("| %5d%-5s | %12s%-13s | %8d%-5s | %5v%-5s \n",data.Id,"", string(data.Name[:]),"",data.Code,"",data.Type,"")

	fmt.Printf("| %d \t     |\t %s %s \t|\t %d \t|\t %v \t \n",data.Id, string(data.Name[:])," ",data.Code,data.Type)
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
	defer file.Close()

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