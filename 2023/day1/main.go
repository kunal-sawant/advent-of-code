package main

import (
	"fmt"
	"os"
)

func main() {
	filecontent, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	var lf byte = 10
	digit1Asigned := false
	digit2Asigned := false
	var digit1, digit2, sum int
	for _, v := range filecontent {
		if v == lf {
			if !digit2Asigned {
				digit2 = digit1
			}
			// fmt.Printf("val: %d", 10*digit1+digit2)

			sum += 10*digit1 + digit2
			digit1Asigned = false
			digit2Asigned = false
		}
		if v >= 48 && v <= 57 {
			if !digit1Asigned {
				digit1 = int(v - 48)
				digit1Asigned = true
			} else {
				digit2 = int(v - 48)
				digit2Asigned = true
			}
		}
	}

	fmt.Printf("sum: %d", sum)
}
