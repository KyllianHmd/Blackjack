package main

import (
	"fmt"
	"os"
)

func infinitegame(playerCard1 int, playerCard2 int, bankhandvalue int, chips int, bets int) {
	playerhandvalue := playerCard1 + playerCard2
	for {
		var action string
		fmt.Scanf("%s", &action)
		if action == "hit" {
			playerhandvalue = hit(playerhandvalue, bankhandvalue)
			hitendgame(playerhandvalue, bankhandvalue, action, bets, chips)
		} else if action == "stand" {
			bankhandvalue = stand(bankhandvalue, playerhandvalue)
			chips = standendgame(bankhandvalue, playerhandvalue, action, bets, chips)
			chips = thewinneris(bankhandvalue, playerhandvalue, chips, bets)
		} else if action == "quit" {
			quit()
		} else if action == "restart" {
			restart()
		} else if action == "chips" {
			howmanychips(chips)
		} else if action == "surrender" {
			chips = surrender(chips, bets)
			blackjack2(chips)
		} else if action == "double" {
			if chips-bets < 0 {
				fmt.Println("\nSorry, you haven't enought money to double your hand.")
			} else {
				chips, playerhandvalue, bets = double(playerhandvalue, bankhandvalue, chips, bets)
				doubleendgame(playerhandvalue, bankhandvalue, action, bets, chips)
				bankhandvalue = doublestand(bankhandvalue, playerhandvalue)
				chips = doublestandendgame(bankhandvalue, playerhandvalue, action, bets, chips)
				chips = doublethewinneris(bankhandvalue, playerhandvalue, chips, bets)
			}
		} else if action == "split" {
			if chips-bets < 0 {
				fmt.Println("\nSorry, you haven't enought money to split your hand.")
			} else if playerCard1 != playerCard2 {
				fmt.Println("\nSorry, you can not split your hand as your two cards are differents.")
			} else {
				chips = split(playerCard1, playerCard2, bankhandvalue, chips, bets)
			}
		} else {
			fmt.Println("\nPlease, enter a real decision.")
		}
		fmt.Print("\n> ")
	}
}

func luckyhand(playerhandvalue int, bankhandvalue int, chips int, bets int) int {
	if playerhandvalue == 21 && bankhandvalue != 11 {
		fmt.Println("\nB L A C K J A C K !")
		fmt.Println("You won x1.5 your bets !")
		chips = chips + bets*2 + bets/2
		blackjack2(chips)
	}
	return chips
}

func blackjack() {
	chips := placeyourbets()
	chips, bets := betamount(chips)
	playerCard1, playerCard2 := initplayers()
	bankhandvalue := initbank()
	chips = luckyhand(playerCard1+playerCard2, bankhandvalue, chips, bets)
	decision()
	fmt.Print("> ")
	infinitegame(playerCard1, playerCard2, bankhandvalue, chips, bets)
}

func blackjack2(chips int) {
	fmt.Printf("\nYou have %d chips !", chips)
	if chips == 0 {
		fmt.Println("\nYou have lost all of your chips !")
		fmt.Println("We are restarting you a new game.")
		restart()
	}
	chips, bets := betamount(chips)
	playerCard1, playerCard2 := initplayers()
	bankhandvalue := initbank()
	chips = luckyhand(playerCard1+playerCard2, bankhandvalue, chips, bets)
	decision()
	fmt.Print("> ")
	infinitegame(playerCard1, playerCard2, bankhandvalue, chips, bets)
}

func main() {
	b, err := os.ReadFile("bj.txt")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(string(b))
	blackjack()
}
