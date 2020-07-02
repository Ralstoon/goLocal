package main

import "fmt"

func main() {
	naturals := make(chan int)
	squareer := make(chan int)
	//counter
	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()
	// squareer
	go func() {
		// for {
		// 	x, ok := <-naturals
		// 	if !ok {
		// 		break
		// 	}
		// 	squareer <- x * x
		// }
		for x := range naturals {
			squareer <- x * x
		}
		close(squareer)
	}()
	// printer
	for x := range squareer {
		fmt.Println(x)
	}
	// close(squareer)
}
