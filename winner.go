package main

import (
	"fmt"
)

func youwin(action string, chips int, bets int) int {
	fmt.Printf("You won %d chips ! :-)", bets*2)
	chips = chips + bets*2
	fmt.Printf("\nYou have now %d chips !", chips)
	again(action, chips)
	return (chips)
}

func draw(action string, chips int, bets int) int {
	fmt.Println("It's a draw, no one lose ! ;-)")
	chips = chips + bets
	fmt.Printf("You have always %d chips !\n", chips)
	again(action, chips)
	return (chips)
}

func youlose(action string, chips int, bets int) int {
	fmt.Printf("You lose %d chips ! :-(\n", bets)
	fmt.Printf("You have still %d chips !\n", chips)
	again(action, chips)
	return (chips)
}

func thewinneris(bankhandvalue int, playerhandvalue int, chips int, bets int) int {
	var action string
	fmt.Printf("\n\nYour hand value : %d\n", playerhandvalue)
	fmt.Printf("The dealer's hand value : %d\n", bankhandvalue)
	if bankhandvalue < playerhandvalue {
		chips = youwin(action, chips, bets)
	} else if bankhandvalue == playerhandvalue {
		chips = draw(action, chips, bets)
	} else {
		chips = youlose(action, chips, bets)
	}
	return (chips)
}
