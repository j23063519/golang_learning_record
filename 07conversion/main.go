package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcom := "Welcome to user input"
	fmt.Println(welcom)

	// os.Stdin
	// standard input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")

	// comma ok / err ok
	// after inputing something you can get what you input and error
	// if you don't want get error you can use _ to ignore it
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of rating is %T \n", input)

	// numRating, err := strconv.ParseFloat(input, 64)
	// output:
	// strconv.ParseFloat: parsing "01\n": invalid syntax

	// strings.TrimSpace:
	// remove space from strings
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}
}
