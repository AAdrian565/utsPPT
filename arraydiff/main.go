package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scaner := bufio.NewReader(os.Stdin)

	//input array1
	fmt.Print("input array 1:")
	line, err := scaner.ReadString('\n')
	array1 := strings.Fields(line)

	//input array2
	fmt.Print("input array 2:")
	line, err = scaner.ReadString('\n')
	array2 := strings.Fields(line)

	//Check for errors
	if err != nil {
		fmt.Println("Error reading value!")
	}

	// fmt.Println(array1)
	// fmt.Println(array2)

	//find highest lenght
	loop := 0
	if len(array1) > len(array1) {
		loop = len(array1)
	} else {
		loop = len(array2)
	}

	//loop based on the highest lenght
	for i := 0; i < loop; i++ {
		if array1[i] != array2[i] {
			fmt.Println("index ke", i, "berbeda")
		}
	}
}
