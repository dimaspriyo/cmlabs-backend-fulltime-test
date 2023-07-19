package main

import "fmt"

func main() {
	fmt.Println()
	fmt.Println(3 % 3)
	for i := 1; i <= 100; i++ {
		fmt.Println("\n\n====== Angka  ", i, "======")
		var isFizz bool
		var isBuzz bool

		if i%3 == 0 {
			isFizz = true
		}

		if i%5 == 0 {
			isBuzz = true
		}

		if isFizz == true && isBuzz == true {
			fmt.Println("FizzBuzz")
		}

		if isFizz {
			fmt.Println("Fizz")
		} else {
			fmt.Println("Buzz")
		}

	}
}
