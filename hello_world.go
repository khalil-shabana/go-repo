package main

import "fmt"

func main() {
	fmt.Println(greeting_recipirnt("Khalil"))
}

func greeting() string {
	msg := "Hello, World!"
	return msg
}

func greeting_recipirnt(name string) string {
	msg := fmt.Sprintf("Hello, %s!", name)
	return msg
}
