package main

import (
	"fmt"
	"math/rand"
	"time"
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
			fmt.Println("\nSorry, you haven't enought chips.")
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
	fmt.Print("\nTo start, place your bet.\n> ")
	fmt.Scanf("%d\n", &bets)
	bets = checknegativebets(bets, chips)
	chips = chips - bets
	fmt.Printf("\nYou have bet %d chips.", bets)
	fmt.Printf("\nYou have now %d chips", chips)
	return chips, bets
}
func placeyourbets() int {
	var chips int
	fmt.Print("\n\nPlease, tell us your chips amount you want.\n> ")
	fmt.Scanf("%d\n", &chips)
	fmt.Printf("\nYou have %d chips\n", chips)
	return chips
}

func initplayers() (int, int) {
	// playerCard1 := random()
	// playerCard2 := random()
	// playerhandvalue := playerCard1 + playerCard2
	// if playerhandvalue == 22 {
	// 	playerhandvalue = 12
	// }

	// FOR THE TEST ---
	playerCard1 := 2
	playerCard2 := 2
	playerhandvalue := playerCard1 + playerCard2
	// FOR THE TEST ---

	fmt.Printf("\n\nYour hand value : %d (%d + %d)", playerhandvalue, playerCard1, playerCard2)
	return playerCard1, playerCard2
}

func initbank() int {
	card1 := random()
	bankhandvalue := card1
	fmt.Printf("\nThe dealer's hand value : %d\n", bankhandvalue)
	return (bankhandvalue)
}

func decision() {
	fmt.Println("\nType \"hit\" to hit a card")
	fmt.Println("Type \"stand\" to stand with your cards")
	fmt.Println("Type \"double\" your bets and pick up only one card")
	fmt.Println("Type \"split\" and play with double game cards")
	fmt.Println("Type \"surrender\" to surrend your hand")
	fmt.Println("Type \"chips\" to see your chips amount")
	fmt.Println("Type \"restart\" to restart the game")
	fmt.Println("Type \"quit\" to exit the game\n")
}
