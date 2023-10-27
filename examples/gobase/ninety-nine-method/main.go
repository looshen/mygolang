package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			value := j * i
			space := ""
			if value < 10 {
				space = " "
			}
			fmt.Printf(fmt.Sprintf("%d*%d=%d  %s", j, i, value, space))
		}
		fmt.Println()
	}
}
