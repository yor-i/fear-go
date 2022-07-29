package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open and load file into string
	if len(os.Args) < 2 {
		fmt.Println("One argument (file path) must be supplied.")
		return
	}

	f, e := os.ReadFile(os.Args[1])
	if e != nil {
		fmt.Println("Couldn't open file.")
		return
	}

	str := string(f)

	// Takes user input
	reader := bufio.NewReader(os.Stdin)

	// The only variable in Fear
	var v byte = ' '

	// Looping through the file
	for i := 0; i < len(str); i++ {
		// Switch executing the current instruction
		switch str[i] {
		case 'i':
			v, e = reader.ReadByte()
			if e != nil {
				fmt.Println("Couldn't read byte.")
				return
			}
			break
		case 'o':
			fmt.Print(string(v))
			break
		case '+':
			v++
			break
		case '-':
			v--
			break
		case 'j':
			i = -1 // at the end of the cycle it becomes 0
			break
		case 'r':
			v = ' '
			break
		}
	}
}
