package main

import (
	"fmt"
	"time"
)

func double(playerhandvalue int, bankhandvalue int, chips int, bets int) (int, int, int) {
	cardx := random()
	if playerhandvalue >= 11 && cardx == 11 {
		cardx = 1
	}
	playerhandvalue = playerhandvalue + cardx
	chips = chips - bets
	fmt.Printf("\nYou have now bet %d chips", bets*2)
	fmt.Printf("\nYou have now %d chips\n", chips)
	fmt.Printf("\nYou are hitting a %d card...\n", cardx)
	time.Sleep(1 * time.Second)
	fmt.Printf("Your hand value : %d\n", playerhandvalue)
	fmt.Printf("The dealer's hand value : %d\n", bankhandvalue)
	return chips, playerhandvalue, bets
}

func doubleendgame(playerhandvalue int, bankhandvalue int, action string, bets int, chips int) {
	if playerhandvalue > 21 {
		fmt.Printf("\nYour hand value : %d\n", playerhandvalue)
		fmt.Printf("The dealer's hand value : %d\n", bankhandvalue)
		fmt.Printf("You have bust and lose %d chips. :-(\n", bets*2)
		fmt.Printf("You still have %d chips !\n", chips)
		again(action, chips)
	}
}

func doublestand(bankhandvalue int, playerhandvalue int) int {
	fmt.Printf("\nYour hand value : %d\n", playerhandvalue)
	card2 := random()
	bankhandvalue = bankhandvalue + card2
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

func doublestandendgame(bankhandvalue int, playerhandvalue int, action string, bets int, chips int) int {
	if bankhandvalue > 21 {
		fmt.Printf("\nYour hand value : %d\n", playerhandvalue)
		fmt.Printf("The dealer's hand value : %d", bankhandvalue)
		fmt.Printf("\nYou won %d chips ! :-)\n", bets*2+bets*2)
		chips = chips + bets*2 + bets*2
		fmt.Printf("You have now %d chips ! \n", chips)
		again(action, chips)
	}
	return (chips)
}

func youwin2(action string, chips int, bets int) int {
	fmt.Printf("\nYou won %d chips ! :-)", bets*2+bets*2)
	chips = chips + bets*2 + bets*2
	fmt.Printf("\nYou have now %d chips !\n", chips)
	again(action, chips)
	return (chips)
}

func draw2(action string, chips int, bets int) int {
	fmt.Println("It's a draw, no one lose ! ;-)")
	chips = chips + bets*2
	fmt.Printf("You have always %d chips !\n", chips)
	again(action, chips)
	return (chips)
}

func youlose2(action string, chips int, bets int) int {
	fmt.Printf("You lose %d chips ! :-(\n", bets*2)
	fmt.Printf("You have still %d chips !\n", chips)
	again(action, chips)
	return (chips)
}

func doublethewinneris(bankhandvalue int, playerhandvalue int, chips int, bets int) int {
	var action string
	fmt.Printf("\n\nYour hand value : %d\n", playerhandvalue)
	fmt.Printf("The dealer's hand value : %d\n", bankhandvalue)
	if bankhandvalue < playerhandvalue {
		chips = youwin2(action, chips, bets)
	} else if bankhandvalue == playerhandvalue {
		chips = draw2(action, chips, bets)
	} else {
		chips = youlose2(action, chips, bets)
	}
	return (chips)
}
