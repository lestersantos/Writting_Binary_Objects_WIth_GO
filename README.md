# Writting Binary objects WIth GO

Converting objects into byte streams that can be stored, shared, and reconstructed.

## Go version

At the time of writting we are working with `go version go1.21.6 linux/amd64`.

Check if the path exist or not.
If exist create it
Creat a new file
Not override the file created for first time
Show menu
Writte in file
Read and recover the file content and print it out 
exit the application


# Homework #1 

Create a Go Object with the following attributes:

1. Object Course
- type: Boolean
- Id: int
- Code: int
- Name: char[16]

# Packages in GO

[Go Code organization](https://go.dev/doc/code)

Go programs are organized into packages. A `package` is a collection of source files in the same directory that are compiled together. 

The first statement in a Go source file must be `package name`. Executable commands must always use package main.

# Discover packges in GO

[Search packages or symbols](https://pkg.go.dev/)

# Creating a simple file with Go

[Creating a simple file with Go](https://gosamples.dev/create-file/#:~:text=To%20create%20a%20new%20empty,removed%20without%20deleting%20the%20file.)

Using the `os.Create`library we can create single file in the current directory.

If we enter something different like a path with a folder that doesn't exist, we get the follwoing error

```sh
open /newDIr/testFIle.txt: no such file or directory
exit status 1
```

## Creating a file with Go and check if the path exist or not

## About Absolute and Relative paths

[Absolute and Relative paths](https://www.linkedin.com/pulse/difference-between-absolute-path-relative-linux-waqas-muazam#:~:text=Absolute%20paths%20always%20start%20with,relative%20to%20the%20current%20directory.)

An absolute path is a complete path to a file or directory from the `root directory`.

A relative path is a path to a file or directory that is relative to the `current directory`. 

## The `filepath` package

[Package filepath](https://pkg.go.dev/path/filepath@go1.21.6)

Package filepath help us manipulate filename paths.

## Cleaning path string with `filepath.Dir function`

[Go package filepath function Dir](https://pkg.go.dev/path/filepath#Dir)

Dir function drops the final element and returns all the parents, basically will return the entire relative path. E.g.

```go
	fmt.Println(filepath.Dir("/foo/bar/baz.js"))
	fmt.Println(filepath.Dir("/foo/bar/baz"))
	fmt.Println(filepath.Dir("/foo/bar/baz/"))
	fmt.Println(filepath.Dir("/dirty//path///"))
	fmt.Println(filepath.Dir("dev.txt"))
```

```bash
Output:

/foo/bar
/foo/bar
/foo/bar/baz
/dirty/path
.
```
# The `os` package

[Pacakge os](https://pkg.go.dev/os)

Package `os` provides a platform-ndependent interface to operating system functionality.

## Check directory or file existens with `os.Stat` function

[os.Stat function](https://pkg.go.dev/os#Stat)

Stat returns a FileInfo describing the named file. If there is an error, it will
be of type *PathError

```go
_, err := os.Stat(path)
```

```bash
>>Error type
*fs.PathErrorstat 

>>Error info send it to `os.IsNotExist(err)`
./newDir2/test.txt: no such file or directory
```

### About Underscore variable in Go

[Bland identifier (underscore)](https://www.geeksforgeeks.org/what-is-blank-identifierunderscore-in-golang/)

`_` Underscore in Go is known as th `Blank Identifier`, is a feature of Go to define unused variables, Go doesn't allow the programmer to define an unused variable if you
do sucn, then the compiler throw an error.

When a funciton returns multiple values, but we need only a few values and we want to
discard some values, underscore thells the compiler to ignore the value without any error. 

```bash
./prog.go:15:7: <var> declared and not used
```

## The `os.IsNotExist`function

IsNotExist returns a boolean indicating whether the err is known to report that a file
or directory `does not exist`. E.g


# Create a directory

[os.MkdirAll create a directory](https://pkg.go.dev/os#MkdirAll)

MKdirAll creates a directory named path, and any necessary parents, and returns nil,
or else returns an error, the specified permissions are used for all directories that
MkdirAll creates. If path is already a directory, it won't do anything.

# Linux Permissions Considerations

[Chmod permissions](https://en.wikipedia.org/wiki/Chmod)

The main parts of the chmod permissions:

Class: User - Group - Other
Mode: r (read), w(write), x(execute)

- By group:  rwxr-x--- groups of 3 characters
- Numerical: 0754 four digits for 3 groups and 1 special flag.
- Symbolic: reference class [operator] [mode]

If you create a file without a prior check if the file exist it will override it and create a new one, deleting all the content
of the named file.

# Structs

[Go Structs](https://go.dev/tour/moretypes/2)
A struct is a collection of fields.


# OpenFIle

## Constants for file permissions

https://pkg.go.dev/os#pkg-constants

## Visual studio code Select coincidence

ctrl + shift + L

# Pointers 

[Pointers in C](https://www.w3schools.com/c/c_pointers.php)

A `pointer` is a variables that stores the memory address for another variable as its value, and is created with the `*` when declaring a variable.

`int* ptr = &myAge`

The `reference & operator` is used to get the `memory address` of a variable.

`&myAge`

`Deference * operator` when not used in declaration you used to get the value of the
variable the pointer points to. 

`printf("%d\n", *ptr)`

```c
int myAge = 43;     // An int variable
int* ptr = &myAge; // A pointer variable, with the name ptr, that stores the address of myAge
```
Pointers are importan in `C`, because they allow us to manipulate the data in the computer's memory
This can reduce the code and improve the performance. You use pointers with `Data structures eg. lists, trees and graphs` and also when working with `files`.

## Pointers in Go

[Pointers in Go](https://go.dev/tour/moretypes/1)

Go `pointers`, e.g

```go
var p *int
```

`Reference & operator` e.g

```go
i := 42
p = &i
```
`Deference * operator` also know as `indirecting` in Go.

```go
fmt.Println(*p)
*p = 21
```

We can run the following snippet in a go console program and check the results.
```go
	i, j := 42, 2701


	p := &i         // point to i
	fmt.Println("i:",i)
	fmt.Println("&i:",&i)
	fmt.Println("&p:",&p)
	fmt.Println("p:",p)
	fmt.Println("*p",*p) // read i through the pointer

	fmt.Println("Update i=42, throught *p")
	fmt.Println("*p=21")
	*p = 21         // set i through the pointer
	fmt.Println("i:",i)  // see the new value of i
	fmt.Println("p:",p)
	fmt.Println("*p",*p) // read i through the pointer

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
```