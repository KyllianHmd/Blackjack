package main

import (
	"fmt"
	"io/ioutil"

	"github.com/fatih/color"
)

func infinitegame(playerhandvalue int, bankhandvalue int, chips int, bets int) {
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
				color.Set(color.FgRed, color.Bold)
				fmt.Println("\nSorry, you haven't enought money to double your hand.")
				color.Unset()
			} else {
				chips, playerhandvalue, bets = double(playerhandvalue, bankhandvalue, chips, bets)
				doubleendgame(playerhandvalue, bankhandvalue, action, bets, chips)
				bankhandvalue = doublestand(bankhandvalue, playerhandvalue)
				chips = doublestandendgame(bankhandvalue, playerhandvalue, action, bets, chips)
				chips = doublethewinneris(bankhandvalue, playerhandvalue, chips, bets)
			}
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

func luckyhand(playerhandvalue int, bankhandvalue int, chips int, bets int) int {
	if playerhandvalue == 21 && bankhandvalue != 11 {
		color.Set(color.FgGreen, color.Bold)
		fmt.Println("\nB L A C K J A C K !")
		fmt.Println("You won x1.5 your bets !")
		color.Unset()
		chips = chips + bets*2 + bets/2
		blackjack2(chips)
	}
	return chips
}

func blackjack() {
	chips := placeyourbets()
	chips, bets := betamount(chips)
	playerhandvalue := initplayers()
	bankhandvalue := initbank()
	chips = luckyhand(playerhandvalue, bankhandvalue, chips, bets)
	decision()
	color.Set(color.FgWhite, color.Bold)
	fmt.Print("> ")
	color.Unset()
	infinitegame(playerhandvalue, bankhandvalue, chips, bets)
}

func blackjack2(chips int) {
	color.Set(color.FgWhite, color.Bold)
	fmt.Printf("\nYou have %d chips !", chips)
	if chips == 0 {
		fmt.Println("\nYou have lost all of your chips !")
		fmt.Println("We are restarting you a new game.")
		restart()
	}
	color.Unset()
	chips, bets := betamount(chips)
	playerhandvalue := initplayers()
	bankhandvalue := initbank()
	chips = luckyhand(playerhandvalue, bankhandvalue, chips, bets)
	decision()
	color.Set(color.FgWhite, color.Bold)
	fmt.Print("> ")
	color.Unset()
	infinitegame(playerhandvalue, bankhandvalue, chips, bets)
}

func main() {
	color.Set(color.FgCyan, color.Bold)
	b, err := ioutil.ReadFile("bj.txt")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(string(b))
	color.Unset()
	blackjack()
}
