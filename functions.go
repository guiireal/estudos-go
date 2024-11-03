package main

import "fmt"

func printData(name string, age int) {
	fmt.Printf("%s tem %d anos.", name, age)
}

func sum(num1, num2 int) int {
	return num1 + num2
}

func main3() {
	printData("João", 30)
	fmt.Println(sum(10, 20))
}
