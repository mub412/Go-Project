package main

import "fmt"

func calculate() (result int) {
	fmt.Println("First Result")

	show := func() {
		result = result + 10
		fmt.Println("defer", result)
	}
	defer show()
	result = 5
	fmt.Println("Second", result)

	return
}

func calc() int {
	result := 0
	fmt.Println("First", result)

	show := func() {
		result = result + 10
		fmt.Println("defer", result)
	}
	defer show()
	result = 5
	fmt.Println("Second", result)

	return result
}

func main() {
	a := calculate()
	fmt.Println("Main first", a)

	b := calc()
	fmt.Println("Main second", b)
}
