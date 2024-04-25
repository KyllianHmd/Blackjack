package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func hit(playerhandvalue int, bankhandvalue int) int {
	card3 := random()
	if playerhandvalue == 11 && card3 == 11 {
		playerhandvalue = 12
	} else {
		playerhandvalue = playerhandvalue + card3
	}
	fmt.Println()
	color.Set(color.FgWhite, color.Bold)
	fmt.Println("You are hitting a card...")
	time.Sleep(1 * time.Second)
	fmt.Printf("Your hand value : %d\n", playerhandvalue)
	fmt.Printf("The dealer's hand value : %d\n", bankhandvalue)
	color.Unset()
	return (playerhandvalue)
}
func hitendgame(playerhandvalue int, bankhandvalue int, action string, bets int, chips int) {
	if playerhandvalue > 21 {
		color.Set(color.FgWhite, color.Bold)
		fmt.Printf("\nYour hand value : %d\n", playerhandvalue)
		fmt.Printf("The dealer's hand value : %d", bankhandvalue)
		color.Unset()
		color.Set(color.FgRed, color.Bold)
		fmt.Printf("\nYou have bust and lose %d chips. :-(\n", bets)
		fmt.Printf("You still have %d chips !\n", chips)
		color.Unset()
		color.Set(color.FgCyan, color.Bold)
		again(action, chips)
	}
}
