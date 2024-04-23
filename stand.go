package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func stand(bankhandvalue int, playerhandvalue int) int {
	color.Set(color.FgWhite, color.Bold)
	fmt.Printf("\nYour hand value : %d\n", playerhandvalue)
	color.Unset()
	card2 := random()
	if bankhandvalue == 11 && card2 == 11 {
		bankhandvalue = 12
	} else {
		bankhandvalue = bankhandvalue + card2
	}
	bankhandvalue = bankhandvalue + card2
	color.Set(color.FgWhite, color.Bold)
	fmt.Println("The dealer hit a card...")
	time.Sleep(1 * time.Second)
	fmt.Printf("The dealer's hand value : %d\n", bankhandvalue)
	color.Unset()
	var i int64
	i = 5
	for bankhandvalue < 17 {
		cardx := random()
		i = i + 1
		bankhandvalue = bankhandvalue + cardx
		color.Set(color.FgWhite, color.Bold)
		fmt.Println("The dealer hit a card...")
		time.Sleep(1 * time.Second)
		fmt.Printf("The dealer's hand value : %d\n", bankhandvalue)
		color.Unset()
	}
	return (bankhandvalue)
}

func standendgame(bankhandvalue int, playerhandvalue int, action string, bets int, chips int) int {
	if bankhandvalue > 21 {
		color.Set(color.FgWhite, color.Bold)
		fmt.Printf("\nYour hand value : %d\n", playerhandvalue)
		fmt.Printf("The dealer's hand value : %d", bankhandvalue)
		color.Unset()
		color.Set(color.FgGreen, color.Bold)
		fmt.Printf("\nYou won %d chips ! :-)\n", bets*2)
		chips = chips + bets*2
		fmt.Printf("You have now %d chips ! \n", chips)
		color.Unset()
		again(action, chips)
	}
	return (chips)
}
