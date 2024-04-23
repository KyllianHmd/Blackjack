package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

func split(playerCard1 int, playerCard2 int, bankhandvalue int, chips int, bets int) (int, int, int) {

	color.Set(color.FgWhite, color.Bold)
	chips = chips - bets
	fmt.Printf("You have bets now %d chips", bets*2)
	fmt.Printf("\nYou have now %d chips\n", chips)
	fmt.Println()
	// handlefirsthand()
	// fmt.Println("You are hitting a card for your first hand...")
	// hand1 := playerCard1 + random()
	// fmt.Printf("Your first hand value : %d\n", hand1)

	// handlesecondhand()
	// fmt.Println("You are hitting a card for your second hand...")
	// hand2 := playerCard2 + random()
	// fmt.Printf("Your first hand value : %d\n", hand1)

	time.Sleep(1 * time.Second)
	fmt.Printf("The dealer's hand value : %d\n", bankhandvalue)
	color.Unset()

	os.Exit(0)
	// card3 := random()
	// if playerhandvalue == 11 && card3 == 11 {
	// 	playerhandvalue = 12
	// } else {
	// 	playerhandvalue = playerhandvalue + card3
	// }
	// fmt.Println()
	// color.Set(color.FgWhite, color.Bold)
	// chips = chips - bets
	// fmt.Printf("You have bets now %d chips", bets*2)
	// fmt.Printf("\nYou have now %d chips\n", chips)
	// fmt.Println()
	// fmt.Println("You are hitting a card...")
	// time.Sleep(1 * time.Second)
	// fmt.Printf("Your hand value : %d\n", playerhandvalue)
	// fmt.Printf("The dealer's hand value : %d\n", bankhandvalue)
	// color.Unset()
	// return chips, playerhandvalue, bets
	return 1, 1, 1
}
