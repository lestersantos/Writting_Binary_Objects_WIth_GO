# Writting Binary objects WIth GO

Converting objects into byte streams that can be stored, shared, and reconstructed.

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

[Go COde organization](https://go.dev/doc/code)

Go programs are organized into packages. A `package` is a collection of source files in the same directory that are compiled together. 

The first statement in a Go source file must be `pacage name`. Executable commands must always use package main.

# Discover packges in GO

[Search packages or symbols](https://pkg.go.dev/)

# Creating a simple file with Go

[Creating a simple file with Go](https://gosamples.dev/create-file/#:~:text=To%20create%20a%20new%20empty,removed%20without%20deleting%20the%20file.)

Using the `os.Create`library we can create single file in the current directory.

If we enter something different like a path, we get the follwoing error

```sh
open /newDIr/testFIle.txt: no such file or directory
exit status 1
```

## Creating a file with Go and check if the path exist or not


# Structs

[Go Structs](https://go.dev/tour/moretypes/2)
A struct is a collection of fields.

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
