package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("c!pher t00l\n")

	// input := "ABC NOP"

	fmt.Print("Sisesta TEXT: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	mapped := strings.Map(rot13, input)
	fmt.Println("ROT13 InPut: ", input)
	fmt.Println("ROT13 OutPut: ", mapped)

	reversed := strings.Map(reverseABC, input)
	fmt.Println("Reverse input: ", input)
	fmt.Println("Reverse output: ", reversed)

	//mapped := strings.Map(rot13, input)
	// fmt.Println(input)
	// fmt.Println(mapped)

}

func rot13(m rune) rune { // a b c d e f g h i j k l m n o p q r s t u v w x y z
	if m >= 'a' && m <= 'z' { // väikesed tähed - rotate13
		if m >= 'm' {
			return m - 13
		} else {
			return m + 13
		}
	} else if m >= 'A' && m <= 'Z' { // Suured tähed - rotate13
		if m >= 'M' {
			return m - 13
		} else {
			return m + 13
		}
	}
	return m
}

func reverseABC(r rune) rune {

	if r >= 'a' && r <= 'z' { // väikesed tähed
		return 'z' - (r - 'a')
	} else if r >= 'A' && r <= 'Z' { // Suured tähed
		return 'Z' - (r - 'A')
	} else {
		return r
	}

}
