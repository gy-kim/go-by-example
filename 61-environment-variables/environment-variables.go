package main

/*

https://gobyexample.com/environment-variables

Environment variables are a universal mechanism for
conveying configuration information to Unix programs.
Let's look at how to set, get, and list environment
variables.

*/

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	/*
		To set a key/value pair, use os.Setenv. To get a value for a
		key, use os.Getenv. This will return an empty string if the
		key isn't present in the environment.
	*/
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	/*
		Use os.Environ to list all key/value pairs in the
		environment. This returns a slice of strings in the form
		KEY=value. You can strings.Split them to get the key
		and value. Here we print all the keys.
	*/
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair)
	}

	/*
		$ go run environment-variables.go
		Running the program show that we pick up the value for
		FOO that we set in the program, but that BAR is empty.

		$ BAR=2 go run environment-variables.go
		If we set BAR in the environment first, the running
		program picks that value up.
	*/
}
