package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

func surrender(chips int, bets int) int {
	chips = chips + bets/2
	color.Set(color.FgRed, color.Bold)
	fmt.Println("\nYou have surrender !")
	color.Unset()
	return (chips)
}

func howmanychips(chips int) {
	fmt.Printf("You still have %d chips", chips)
}

func quit() {
	color.Set(color.FgCyan, color.Bold)
	b, err := os.ReadFile("cya.txt")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(string(b))
	color.Unset()
	os.Exit(0)
}

func restart() {
	color.Set(color.FgYellow, color.Bold)
	fmt.Print("\nRestarting the game... ")
	color.Unset()
	time.Sleep(2 * time.Second)
	fmt.Println("\n_____________________________________________________")
	blackjack()
}

func again(action string, chips int) {
	color.Set(color.FgCyan, color.Bold)
	fmt.Println("\nAgain ? Type Y or N")
	color.Unset()
	color.Set(color.FgWhite, color.Bold)
	fmt.Print("> ")
	color.Unset()
	fmt.Scanf("%s", &action)
	if action == "Y" || action == "y" {
		color.Set(color.FgYellow, color.Bold)
		fmt.Print("\nStarting a new game...")
		color.Unset()
		time.Sleep(2 * time.Second)
		fmt.Println("\n_____________________________________________________")
		blackjack2(chips)
	} else {
		color.Set(color.FgCyan, color.Bold)
		b, err := os.ReadFile("cya.txt")
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(string(b))
		color.Unset()
		os.Exit(0)
	}
}
