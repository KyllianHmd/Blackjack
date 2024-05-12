package main

import (
	"fmt"
	"time"
)

func hit(playerhandvalue int, bankhandvalue int) int {
	cardx := random()
	if playerhandvalue >= 11 && cardx == 11 {
		cardx = 1
	}
	playerhandvalue = playerhandvalue + cardx
	fmt.Printf("\nYou are hitting a %d card...\n", cardx)
	time.Sleep(1 * time.Second)
	fmt.Printf("Your hand value : %d\n", playerhandvalue)
	fmt.Printf("The dealer's hand value : %d\n", bankhandvalue)
	return (playerhandvalue)
}
func hitendgame(playerhandvalue int, bankhandvalue int, action string, bets int, chips int) {
	if playerhandvalue > 21 {
		fmt.Printf("\nYour hand value : %d\n", playerhandvalue)
		fmt.Printf("The dealer's hand value : %d\n", bankhandvalue)
		fmt.Printf("You have bust and lose %d chips. :-(\n", bets)
		fmt.Printf("You still have %d chips !\n", chips)
		again(action, chips)
	}
}
