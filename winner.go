package main

import (
	"fmt"

	"github.com/fatih/color"
)

func youwin(action string, chips int, bets int) int {
	color.Set(color.FgGreen, color.Bold)
	fmt.Printf("You won %d chips ! :-)", bets*2)
	chips = chips + bets*2
	fmt.Printf("\nYou have now %d chips !", chips)
	color.Unset()
	color.Set(color.FgCyan, color.Bold)
	again(action, chips)
	return (chips)
}

func draw(action string, chips int, bets int) int {
	color.Set(color.FgMagenta, color.Bold)
	fmt.Println("It's a draw, no one lose ! ;-)")
	chips = chips + bets
	fmt.Printf("You have always %d chips !\n", chips)
	color.Unset()
	color.Set(color.FgCyan, color.Bold)
	again(action, chips)
	return (chips)
}

func youlose(action string, chips int, bets int) int {
	color.Set(color.FgRed, color.Bold)
	fmt.Printf("You lose %d chips ! :-(\n", bets)
	fmt.Printf("You have still %d chips !\n", chips)
	color.Unset()
	color.Set(color.FgCyan, color.Bold)
	again(action, chips)
	return (chips)
}

func thewinneris(bankhandvalue int, playerhandvalue int, chips int, bets int) int {
	var action string
	fmt.Println()
	color.Set(color.FgWhite, color.Bold)
	fmt.Printf("\nYour hand value : %d\n", playerhandvalue)
	fmt.Printf("The dealer's hand value : %d\n", bankhandvalue)
	color.Unset()
	if bankhandvalue < playerhandvalue {
		chips = youwin(action, chips, bets)
	} else if bankhandvalue == playerhandvalue {
		chips = draw(action, chips, bets)
	} else {
		chips = youlose(action, chips, bets)
	}
	return (chips)
}
