package main

/*

https://gobyexample.com/command-line-arguments

Command-line arguments are a common way to
parameterize execution of programs. For example, go run
hello.go uses run and hello.go arguments to the go
program.

*/

import (
	"fmt"
	"os"
)

func main() {
	/*
		os.Args provides access to raw command-line arguments.
		Note that the first value in this slice is the path to the
		program, and os.Args[1:] holds the arguments to the
		gram.
	*/
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// You can get individual args with normal indexing.
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)

	/*
		go build command-line-arguments.go
		./command-line-arguments.go a b c d

		To experiment with command-line arguments it's best to
		build a binary with go build first.
	*/
}
