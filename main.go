package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	passwordGenerator()
}

func passwordGenerator() {
	lowLet := "qwertyuiopasdfghjklzxcvbnm"
	upLet := "QWERTYUIOPASDFGHJKLZXCVBNM"
	nums := "0123456789"
	symbols := "!?@#$^&*()_-[]"
	password := ""

	var passwordLength int

	rand.Seed(time.Now().UnixNano())

	fmt.Print("What is length of your password?\n")
	fmt.Scan(&passwordLength)

	if passwordLength <= 0 {
		fmt.Print("Incorrect input. Try again!\n")
		passwordGenerator()
	} else {
		for i := 0; i < passwordLength; i++ {
			randLowLet := lowLet[rand.Intn(len(lowLet))]
			randUpLet := upLet[rand.Intn(len(upLet))]
			randNums := nums[rand.Intn(len(nums))]
			randSymbols := symbols[rand.Intn(len(symbols))]
			password += string(randLowLet) + string(randUpLet) + string(randNums) + string(randSymbols)
		}
		correctPassword := password[:passwordLength]
		fmt.Printf("Your new generated password: %s\n", string(correctPassword))
		fmt.Printf("Length: %d\n", len(correctPassword))
		repeat()
	}
}

func repeat() {
	var choice int

	fmt.Print("Do you want to create one more password?\n[0][1] <- ?\t")
	fmt.Scan(&choice)

	if choice == 0 {
		fmt.Print("Thanks for using my password generator!\n")
	} else if choice == 1 {
		passwordGenerator()
	} else {
		panic("incorrect input...")
	}
}
