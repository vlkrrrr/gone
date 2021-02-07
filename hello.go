package main

import "fmt"

const pynispynis = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "wurld"
	}
	return pynispynis + name
}

func main() {
	fmt.Println(Hello("Volker"))
}
