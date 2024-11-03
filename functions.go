package main

import "fmt"

func printData(name string, age int) {
	fmt.Printf("%s tem %d anos.", name, age)
}

func main() {
	printData("João", 20)
}
