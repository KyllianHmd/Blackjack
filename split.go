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
	fmt.Printf("You have now bets %d chips", bets*2)
	fmt.Printf("\nYou have now %d chips\n", chips)
	fmt.Println()
	hand1 := splitHandleFirstHand(playerCard1, bankhandvalue, chips, bets)
	if hand1 > 21 {
		color.Set(color.FgWhite, color.Bold)
		fmt.Printf("\nYour first hand has bust and you lost %d chips. :-(\n", bets)
		color.Unset()
	}
	fmt.Println("Passage hand2")
	os.Exit(0)

	// _ = splitHandleSecondHand(playerCard2, bankhandvalue, chips, bets)

	// time.Sleep(1 * time.Second)
	// fmt.Printf("The dealer's hand value : %d\n", bankhandvalue)
	// color.Unset()

	// card3 := random()
	// if playerhandvalue == 11 && card3 == 11 {
	// 	playerhandvalue = 12
	// } else {
	// 	playerhandvalue = playerhandvalue + card3
	// }
	// fmt.Println()
	// color.Set(color.FgWhite, color.Bold)
	// chips = chips - bets
	// fmt.Printf("You have now bets %d chips", bets*2)
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

func splitHandleFirstHand(playerCard1 int, bankhandvalue int, chips int, bets int) int {
	color.Set(color.FgWhite, color.Bold)
	fmt.Println("You are hitting a card for your first hand...")
	hand1 := playerCard1 + random()
	fmt.Printf("Your first hand value : %d (%d + %d)\n", hand1, playerCard1, (hand1 - playerCard1))
	hand1 = continueSplitGameFirstHand(hand1, bankhandvalue, chips, bets)
	color.Unset()
	return (hand1)
}

func splitHandleSecondHand(playerCard2 int, bankhandvalue int, chips int, bets int) int {
	color.Set(color.FgWhite, color.Bold)
	fmt.Println("\nYou are hitting a card for your second hand...")
	hand2 := playerCard2 + random()
	fmt.Printf("Your second hand value : %d (%d + %d)\n", hand2, playerCard2, (hand2 - playerCard2))
	continueSplitGameSecondHand(hand2, bankhandvalue, chips, bets)
	color.Unset()
	return (hand2)
}
func continueSplitGameFirstHand(hand1 int, bankhandvalue int, chips int, bets int) int {
	fmt.Print("\n> ")
	for {
		var action string
		fmt.Scanf("%s", &action)
		if action == "hit" {
			fmt.Println("Hit First Hand")
			hand1 = hitSplitHand(hand1, "first")
			return hand1

			// checkBustFirstHand(hand, bankhandvalue, action, bets, chips)
		} else if action == "stand" {
			color.Set(color.FgWhite, color.Bold)
			fmt.Printf("\nYour first hand value : %d\n", hand1)
			color.Unset()
			return hand1
		} else if action == "quit" {
			quit()
		} else if action == "restart" {
			restart()
		} else if action == "chips" {
			howmanychips(chips)
		} else if action == "surrender" {
			fmt.Println("Surrender First Hand")
			// chips = surrender(chips, bets)
			// blackjack2(chips)
		} else if action == "double" {
			if chips-bets < 0 {
				color.Set(color.FgRed, color.Bold)
				fmt.Println("\nSorry, you haven't enought money to double your hand.")
				color.Unset()
			} else {
				fmt.Println("Double First Hand")
				// chips, hand1, bets = double(hand1, bankhandvalue, chips, bets)
				// doubleendgame(hand, bankhandvalue, action, bets, chips)
				// bankhandvalue = doublestand(bankhandvalue, hand)
				// chips = doublestandendgame(bankhandvalue, hand, action, bets, chips)
				// chips = doublethewinneris(bankhandvalue, hand, chips, bets)
			}
		} else if action == "split" {
			fmt.Println("\nSorry, you can not split your hand again because you've already done it once.")
		} else {
			color.Set(color.FgRed, color.Bold)
			fmt.Println("\nPlease, enter a real decision.")
			color.Unset()
		}
		color.Set(color.FgWhite, color.Bold)
		fmt.Print("\n> ")
		color.Unset()
	}
}

func continueSplitGameSecondHand(hand2 int, bankhandvalue int, chips int, bets int) {
	fmt.Print("\n> ")
	for {
		var action string
		fmt.Scanf("%s", &action)
		if action == "hit" {
			fmt.Println("Hit First Hand")
			// hand = hit(hand, bankhandvalue)
			// hitendgame(hand, bankhandvalue, action, bets, chips)
		} else if action == "stand" {
			splitStandFirstHand(hand2)
			return
		} else if action == "quit" {
			quit()
		} else if action == "restart" {
			restart()
		} else if action == "chips" {
			howmanychips(chips)
		} else if action == "surrender" {
			fmt.Println("Surrender First Hand")
			// chips = surrender(chips, bets)
			// blackjack2(chips)
		} else if action == "double" {
			if chips-bets < 0 {
				color.Set(color.FgRed, color.Bold)
				fmt.Println("\nSorry, you haven't enought money to double your hand.")
				color.Unset()
			} else {
				fmt.Println("Double First Hand")
				// chips, hand1, bets = double(hand1, bankhandvalue, chips, bets)
				// doubleendgame(hand, bankhandvalue, action, bets, chips)
				// bankhandvalue = doublestand(bankhandvalue, hand)
				// chips = doublestandendgame(bankhandvalue, hand, action, bets, chips)
				// chips = doublethewinneris(bankhandvalue, hand, chips, bets)
			}
		} else if action == "split" {
			fmt.Println("\nSorry, you can not split your hand again because you've already done it once.")
		} else {
			color.Set(color.FgRed, color.Bold)
			fmt.Println("\nPlease, enter a real decision.")
			color.Unset()
		}
		color.Set(color.FgWhite, color.Bold)
		fmt.Print("\n> ")
		color.Unset()
	}
}

func splitStandFirstHand(playerhandvalue int) {
	color.Set(color.FgWhite, color.Bold)
	fmt.Printf("\nYour first hand value : %d\n", playerhandvalue)
	color.Unset()
}

func hitSplitHand(playerhandvalue int, handNumber string) int {
	card3 := random()
	if playerhandvalue == 11 && card3 == 11 {
		playerhandvalue = 12
	} else {
		playerhandvalue = playerhandvalue + card3
	}
	fmt.Println()
	color.Set(color.FgWhite, color.Bold)
	fmt.Printf("You are hitting a card for your %s hand...\n", handNumber)
	time.Sleep(1 * time.Second)
	fmt.Printf("Your hand value : %d\n", playerhandvalue)
	color.Unset()
	return (playerhandvalue)
}
