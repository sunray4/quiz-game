package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string
	var age uint  
	var ans string
	var ans2 string

	fmt.Println("welcome to my quiz game :)")
	fmt.Print("what is your name? ")

	fmt.Scan(&name)

	fmt.Printf("Hello %v, welcome to the game!\n", name)

	fmt.Print("what is your age? ")
	fmt.Scan(&age)

	if (age >= 10) {
		fmt.Println("yay you can play!")
	} else {
		fmt.Println("you can't play...")
		return
	}

	fmt.Printf("what is better, rtx 3080 or rtx 3090? ")
	fmt.Scan(&ans, &ans2)
	ansLC := strings.ToLower(ans)

	if (ansLC + " " + ans2 == "rtx 3080") {
		fmt.Println("correct!")
	} else {
		fmt.Println("incorrect")
	}
	


} 