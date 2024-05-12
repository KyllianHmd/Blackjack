package main

import (
	"fmt"
	"os"
	"time"
)

func surrender(chips int, bets int) int {
	chips = chips + bets/2
	fmt.Println("\nYou have surrender !")
	return (chips)
}

func howmanychips(chips int) {
	fmt.Printf("You still have %d chips", chips)
}

func quit() {
	b, err := os.ReadFile("cya.txt")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(string(b))
	os.Exit(0)
}

func restart() {
	fmt.Print("\nRestarting the game... ")
	time.Sleep(2 * time.Second)
	fmt.Println("\n_____________________________________________________")
	blackjack()
}

func again(action string, chips int) {
	fmt.Println("\nAgain ? Type Y or N")
	fmt.Print("> ")
	fmt.Scanf("%s", &action)
	if action == "Y" || action == "y" {
		fmt.Print("\nStarting a new game...")
		time.Sleep(2 * time.Second)
		fmt.Println("\n_____________________________________________________")
		blackjack2(chips)
	} else {
		b, err := os.ReadFile("cya.txt")
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(string(b))
		os.Exit(0)
	}
}
