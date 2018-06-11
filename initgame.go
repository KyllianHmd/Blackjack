package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

func random() int {
	cards := [14]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10, 11}
	rand.Seed(time.Now().UnixNano())
	choosen := cards[rand.Intn(len(cards)-1)]
	return (choosen)
}

func checknegativebets(bets int, chips int) int {
	var neg int
	neg = 1
	for neg == 1 {
		if bets > chips {
			color.Set(color.FgRed, color.Bold)
			fmt.Println("\nSorry, you haven't enought chips.")
			color.Unset()
			fmt.Print("\nPlease, place your bet.\n> ")
			fmt.Scanf("%d\n", &bets)
			neg = 1
		} else {
			neg = 0
		}
	}
	return (bets)
}

func betamount(chips int) (int, int) {
	var bets int
	color.Set(color.FgWhite, color.Bold)
	fmt.Print("\nTo start, place your bet.\n> ")
	fmt.Scanf("%d\n", &bets)
	bets = checknegativebets(bets, chips)
	chips = chips - bets
	fmt.Printf("\nYou have bet %d chips.", bets)
	fmt.Printf("\nYou have now %d chips", chips)
	color.Unset()
	return chips, bets
}
func placeyourbets() int {
	var chips int
	fmt.Println()
	color.Set(color.FgWhite, color.Bold)
	fmt.Print("\nPlease, tell us your chips amount you want.\n> ")
	fmt.Scanf("%d\n", &chips)
	fmt.Printf("\nYou have %d chips", chips)
	color.Unset()
	fmt.Println()
	return chips
}

func initplayers() int {
	card1 := random()
	card2 := random()
	playerhandvalue := card1 + card2
	if playerhandvalue == 22 {
		playerhandvalue = 13
	}
	fmt.Println()
	color.Set(color.FgWhite, color.Bold)
	fmt.Printf("\nYour hand value : %d", playerhandvalue)
	color.Unset()
	return (playerhandvalue)
}

func initbank() int {
	card1 := random()
	bankhandvalue := card1
	fmt.Println()
	color.Set(color.FgWhite, color.Bold)
	fmt.Printf("The dealer's hand value : %d\n", bankhandvalue)
	color.Unset()
	return (bankhandvalue)
}

func decision() {
	color.Set(color.FgCyan, color.Bold)
	fmt.Println()
	fmt.Println("Type \"hit\" to hit a card")
	fmt.Println("Type \"stand\" to stand with your cards")
	fmt.Println("Type \"double\" your bets and pick up only one card")
	fmt.Println("Type \"surrender\" to surrend your hand")
	fmt.Println("Type \"chips\" to see your chips amount")
	fmt.Println("Type \"restart\" to restart the game")
	fmt.Println("Type \"quit\" to exit the game")
	fmt.Println()
	color.Unset()
}
