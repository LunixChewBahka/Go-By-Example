package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		// should execute 3x since i := 1; if i := 0 then 4x
		fmt.Println(i)
		i = i + 1
		// or i += 1
	}

	// this chunk will execute 3x
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		// will print once
		fmt.Println("loop")
		// after printing, it will break out of the for loop
		break
	}

	// will loop 6x
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		// but will only print the ones that are not even
		// in this case 1, 3, 5
		fmt.Println(n)
	}
}
