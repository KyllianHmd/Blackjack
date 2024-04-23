package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func double(playerhandvalue int, bankhandvalue int, chips int, bets int) (int, int, int) {
	card3 := random()
	if playerhandvalue == 11 && card3 == 11 {
		playerhandvalue = 12
	} else {
		playerhandvalue = playerhandvalue + card3
	}
	fmt.Println()
	color.Set(color.FgWhite, color.Bold)
	chips = chips - bets
	fmt.Printf("You have bets now %d chips", bets*2)
	fmt.Printf("\nYou have now %d chips\n", chips)
	fmt.Println()
	fmt.Println("You are hitting a card...")
	time.Sleep(1 * time.Second)
	fmt.Printf("Your hand value : %d\n", playerhandvalue)
	fmt.Printf("The dealer's hand value : %d\n", bankhandvalue)
	color.Unset()
	return chips, playerhandvalue, bets
}

func doubleendgame(playerhandvalue int, bankhandvalue int, action string, bets int, chips int) {
	if playerhandvalue > 21 {
		color.Set(color.FgWhite, color.Bold)
		fmt.Printf("\nYour hand value : %d\n", playerhandvalue)
		fmt.Printf("The dealer's hand value : %d", bankhandvalue)
		color.Unset()
		color.Set(color.FgRed, color.Bold)
		fmt.Printf("\nYou have bust and lose %d chips. :-(\n", bets*2)
		fmt.Printf("You still have %d chips !\n", chips)
		color.Unset()
		color.Set(color.FgCyan, color.Bold)
		again(action, chips)
	}
}

func doublestand(bankhandvalue int, playerhandvalue int) int {
	color.Set(color.FgWhite, color.Bold)
	fmt.Printf("\nYour hand value : %d\n", playerhandvalue)
	color.Unset()
	card2 := random()
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

func doublestandendgame(bankhandvalue int, playerhandvalue int, action string, bets int, chips int) int {
	if bankhandvalue > 21 {
		color.Set(color.FgWhite, color.Bold)
		fmt.Printf("\nYour hand value : %d\n", playerhandvalue)
		fmt.Printf("The dealer's hand value : %d", bankhandvalue)
		color.Unset()
		color.Set(color.FgGreen, color.Bold)
		fmt.Printf("\nYou won %d chips ! :-)\n", bets*2+bets*2)
		chips = chips + bets*2 + bets*2
		fmt.Printf("You have now %d chips ! \n", chips)
		color.Unset()
		again(action, chips)
	}
	return (chips)
}

func youwin2(action string, chips int, bets int) int {
	color.Set(color.FgGreen, color.Bold)
	fmt.Printf("\nYou won %d chips ! :-)", bets*2+bets*2)
	chips = chips + bets*2 + bets*2
	fmt.Printf("\nYou have now %d chips !\n", chips)
	color.Unset()
	color.Set(color.FgCyan, color.Bold)
	again(action, chips)
	return (chips)
}

func draw2(action string, chips int, bets int) int {
	color.Set(color.FgMagenta, color.Bold)
	fmt.Println("It's a draw, no one lose ! ;-)")
	chips = chips + bets*2
	fmt.Printf("You have always %d chips !\n", chips)
	color.Unset()
	color.Set(color.FgCyan, color.Bold)
	again(action, chips)
	return (chips)
}

func youlose2(action string, chips int, bets int) int {
	color.Set(color.FgRed, color.Bold)
	fmt.Printf("You lose %d chips ! :-(\n", bets*2)
	fmt.Printf("You have still %d chips !\n", chips)
	color.Unset()
	color.Set(color.FgCyan, color.Bold)
	again(action, chips)
	return (chips)
}

func doublethewinneris(bankhandvalue int, playerhandvalue int, chips int, bets int) int {
	var action string
	fmt.Println()
	color.Set(color.FgWhite, color.Bold)
	fmt.Printf("\nYour hand value : %d\n", playerhandvalue)
	fmt.Printf("The dealer's hand value : %d\n", bankhandvalue)
	color.Unset()
	if bankhandvalue < playerhandvalue {
		chips = youwin2(action, chips, bets)
	} else if bankhandvalue == playerhandvalue {
		chips = draw2(action, chips, bets)
	} else {
		chips = youlose2(action, chips, bets)
	}
	return (chips)
}
