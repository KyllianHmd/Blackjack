package main

import (
	"fmt"
	"time"
)

func stand(bankhandvalue int, playerhandvalue int) int {
	fmt.Printf("\nYour hand value : %d\n", playerhandvalue)
	card2 := random()
	if bankhandvalue == 11 && card2 == 11 {
		bankhandvalue = 12
	} else {
		bankhandvalue = bankhandvalue + card2
	}
	fmt.Printf("The dealer hit a %d card...\n", card2)
	time.Sleep(1 * time.Second)
	fmt.Printf("The dealer's hand value : %d\n", bankhandvalue)
	var i int64
	i = 5
	for bankhandvalue < 17 {
		cardx := random()
		i = i + 1
		bankhandvalue = bankhandvalue + cardx
		fmt.Printf("The dealer hit a %d card...\n", cardx)
		time.Sleep(1 * time.Second)
		fmt.Printf("The dealer's hand value : %d\n", bankhandvalue)
	}
	return (bankhandvalue)
}

func standendgame(bankhandvalue int, playerhandvalue int, action string, bets int, chips int) int {
	if bankhandvalue > 21 {
		fmt.Printf("\nYour hand value : %d\n", playerhandvalue)
		fmt.Printf("The dealer's hand value : %d", bankhandvalue)
		fmt.Printf("\nYou won %d chips ! :-)\n", bets*2)
		chips = chips + bets*2
		fmt.Printf("You have now %d chips ! \n", chips)
		again(action, chips)
	}
	return (chips)
}
