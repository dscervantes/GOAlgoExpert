/*Go or Golan is a statically typed language, programmer must specify the type
of variable, typically results in less run-time errors. Type checking is done at compile time

- Go is a high performance service side language - mostly used in the backend
for creating performance web apps
- Created in 2012 at Google - performance language that is easy to read and write
- Built in features for concurrency

- Save your file with the .go extension
- Compile the code then run just like java: go run introgo.go will compile and run your code
this will delete the compiled file after
- Compile the code into an executable binary file, it will create a .exe file in your
repository, this is useful when you want to run that file into another machine with the
same operating system: go build
- To then run that exe file in linux you would do: ./introgo */

package main

import "fmt" // Allows us to output text

func main() { //Starting code, must have a function to run
	fmt.Println("Hello World!")
}

//-------------------------------------------------------------------------

//FYI on Comments on code

//One Single line comment

/*
Mutiline comment
*/

//--------------------------------------------------------------------------

//Questions:

/*

1. GO is a statically typed language as well as a strongly typed language
2. Every executable go program must contain a single package and function named main.

*/
