package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func split(playerCard1 int, playerCard2 int, bankhandvalue int, chips int, bets int) int {

	color.Set(color.FgWhite, color.Bold)
	chips = chips - bets
	fmt.Printf("You have now bets %d chips", bets*2)
	fmt.Printf("\nYou have now %d chips\n", chips)
	fmt.Println()

	hand1 := splitHandleFirstHand(playerCard1, chips, bets)
	bustedHand1 := false
	bustedHand2 := false

	hand2 := splitHandleSecondHand(playerCard2, chips, bets)

	if hand1 > 21 {
		bustedHand1 = true
	}
	if hand2 > 21 {
		bustedHand2 = true
	}
	if bustedHand1 && bustedHand2 {
		color.Set(color.FgRed, color.Bold)
		fmt.Printf("\nYou have bust and lose %d chips. :-(\n", bets*2)
		fmt.Printf("You still have %d chips !\n", chips)
		color.Unset()
	}

	bankhandvalue = splitDealerHand(bankhandvalue)
	chips = splitthewinneris(bankhandvalue, hand1, hand2, chips, bets)
	return chips
}

func splitHandleFirstHand(playerCard1 int, chips int, bets int) int {
	color.Set(color.FgWhite, color.Bold)
	fmt.Println("You are hitting a card for your first hand...")
	hand1 := playerCard1 + random()
	fmt.Printf("Your first hand value : %d (%d + %d)\n", hand1, playerCard1, (hand1 - playerCard1))
	hand1 = continueSplitGameFirstHand(hand1, chips, bets)
	color.Unset()
	return (hand1)
}

func splitHandleSecondHand(playerCard2 int, chips int, bets int) int {
	color.Set(color.FgWhite, color.Bold)
	fmt.Println("\nYou are hitting a card for your second hand...")
	hand2 := playerCard2 + random()
	fmt.Printf("Your second hand value : %d (%d + %d)\n", hand2, playerCard2, (hand2 - playerCard2))
	hand2 = continueSplitGameSecondHand(hand2, chips, bets)
	color.Unset()
	return (hand2)
}
func continueSplitGameFirstHand(hand1 int, chips int, bets int) int {
	fmt.Print("\n> ")
	for {
		var action string
		fmt.Scanf("%s", &action)
		if action == "hit" {
			hand1 = hitSplitHand(hand1, "first")
			if hand1 > 21 {
				return hand1
			}
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
		} else if action == "double" {
			if chips-bets < 0 {
				color.Set(color.FgRed, color.Bold)
				fmt.Println("\nSorry, you haven't enought money to double your hand.")
				color.Unset()
			} else {
				fmt.Println("Double First Hand")
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

func continueSplitGameSecondHand(hand2 int, chips int, bets int) int {
	fmt.Print("\n> ")
	for {
		var action string
		fmt.Scanf("%s", &action)
		if action == "hit" {
			hand2 = hitSplitHand(hand2, "second")
			if hand2 > 21 {
				return hand2
			}
		} else if action == "stand" {
			color.Set(color.FgWhite, color.Bold)
			fmt.Printf("\nYour second hand value : %d\n", hand2)
			color.Unset()
			return hand2
		} else if action == "quit" {
			quit()
		} else if action == "restart" {
			restart()
		} else if action == "chips" {
			howmanychips(chips)
		} else if action == "surrender" {
			fmt.Println("Surrender Second Hand")
		} else if action == "double" {
			if chips-bets < 0 {
				color.Set(color.FgRed, color.Bold)
				fmt.Println("\nSorry, you haven't enought money to double your hand.")
				color.Unset()
			} else {
				fmt.Println("Double Second Hand")
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

func hitSplitHand(playerhandvalue int, handNumber string) int {
	card3 := random()
	if playerhandvalue >= 11 && card3 == 11 {
		card3 = 1
	}
	playerhandvalue = playerhandvalue + card3

	fmt.Println()
	color.Set(color.FgWhite, color.Bold)
	fmt.Printf("You are hitting a card for your %s hand...\n", handNumber)
	time.Sleep(1 * time.Second)
	fmt.Printf("Your %s hand value : %d\n", handNumber, playerhandvalue)
	color.Unset()
	return (playerhandvalue)
}

func splitDealerHand(bankhandvalue int) int {
	color.Set(color.FgWhite, color.Bold)
	fmt.Println("The dealer is hitting a card...")
	time.Sleep(1 * time.Second)
	bankhandvalue = bankhandvalue + random()
	fmt.Printf("The dealer's hand value : %d\n", bankhandvalue)
	for bankhandvalue < 17 {
		cardx := random()
		bankhandvalue = bankhandvalue + cardx
		color.Set(color.FgWhite, color.Bold)
		fmt.Printf("The dealer hit a %d card...\n", cardx)
		time.Sleep(1 * time.Second)
		fmt.Printf("The dealer's hand value : %d\n", bankhandvalue)
		color.Unset()
	}
	color.Unset()
	return bankhandvalue
}

func splitthewinneris(bankhandvalue int, hand1 int, hand2 int, chips int, bets int) int {
	var action string
	fmt.Println()
	color.Set(color.FgWhite, color.Bold)
	fmt.Printf("\nYour first hand value : %d\n", hand1)
	fmt.Printf("Your second hand value : %d\n", hand2)
	fmt.Printf("The dealer's hand value : %d\n", bankhandvalue)
	color.Unset()
	if (bankhandvalue < hand1 && bankhandvalue < hand2) && hand1 <= 21 && hand2 <= 21 {
		fmt.Println("H1 win && H2 win")
		chips = h1winh2win(action, chips, bets)
	} else if bankhandvalue < hand1 && (bankhandvalue > hand2 || hand2 > 21) {
		fmt.Println("H1 win && H2 lose")
		chips = h1winh2lose(action, chips, bets)
	} else if (bankhandvalue > hand1 || hand1 > 21) && (bankhandvalue < hand2 && hand2 <= 21) {
		fmt.Println("H1 lose && H2 win")
		chips = h1loseh2win(action, chips, bets)
	} else if (bankhandvalue > hand1 || hand1 > 21) && (bankhandvalue > hand2 || hand2 > 21) {
		fmt.Println("H1 lose && H2 lose")
		chips = h1loseh2lose(action, chips, bets)
	} else if (bankhandvalue > hand1 || hand1 > 21) && bankhandvalue == hand2 {
		fmt.Println("H1 lose && H2 draw")
		chips = h1loseh2draw(action, chips, bets)
	} else if bankhandvalue == hand1 && (bankhandvalue > hand2 || hand2 > 21) {
		fmt.Println("H1 draw && H2 lose")
		chips = h1drawh2lose(action, chips, bets)
	} else if bankhandvalue == hand1 && bankhandvalue == hand2 {
		fmt.Println("H1 draw && H2 draw")
		chips = h1drawh2draw(action, chips, bets)
	} else if bankhandvalue == hand1 && bankhandvalue < hand2 {
		fmt.Println("H1 draw && H2 win")
		chips = h1drawh2win(action, chips, bets)
	} else if bankhandvalue < hand1 && bankhandvalue == hand2 {
		fmt.Println("H1 win && H2 draw")
		chips = h1winh2draw(action, chips, bets)
	} else {
		// chips = youlose2(action, chips, bets)
		fmt.Println("AUTRE CAS !!")
	}
	return (chips)
}

func h1winh2win(action string, chips int, bets int) int {
	color.Set(color.FgGreen, color.Bold)
	fmt.Printf("\nYou won %d chips ! :-)", bets*2+bets*2)
	chips = chips + bets*2 + bets*2
	fmt.Printf("\nYou have now %d chips !\n", chips)
	color.Unset()
	color.Set(color.FgCyan, color.Bold)
	again(action, chips)
	return (chips)
}

func h1winh2lose(action string, chips int, bets int) int {
	color.Set(color.FgGreen, color.Bold)
	fmt.Printf("\nYou won %d chips ! :-)", bets*2)
	chips = chips + bets*2
	fmt.Printf("\nYou have now %d chips !\n", chips)
	color.Unset()
	color.Set(color.FgCyan, color.Bold)
	again(action, chips)
	return (chips)
}

func h1loseh2win(action string, chips int, bets int) int {
	color.Set(color.FgGreen, color.Bold)
	fmt.Printf("\nYou won %d chips ! :-)", bets*2)
	chips = chips + bets*2
	fmt.Printf("\nYou have now %d chips !\n", chips)
	color.Unset()
	color.Set(color.FgCyan, color.Bold)
	again(action, chips)
	return (chips)
}

func h1loseh2lose(action string, chips int, bets int) int {
	color.Set(color.FgRed, color.Bold)
	fmt.Printf("\nYou lost %d chips ! :-)", bets*2)
	chips = chips - bets
	fmt.Printf("\nYou have now %d chips !\n", chips)
	color.Unset()
	color.Set(color.FgCyan, color.Bold)
	again(action, chips)
	return (chips)
}

func h1loseh2draw(action string, chips int, bets int) int {
	color.Set(color.FgRed, color.Bold)
	fmt.Printf("\nYou lost %d chips ! :-)", bets*2)
	chips = chips - bets
	fmt.Printf("\nYou have now %d chips !\n", chips)
	color.Unset()
	color.Set(color.FgCyan, color.Bold)
	again(action, chips)
	return (chips)
}

func h1drawh2lose(action string, chips int, bets int) int {
	color.Set(color.FgGreen, color.Bold)
	fmt.Printf("\nYou lost %d chips ! :-)", bets*2)
	chips = chips - bets
	fmt.Printf("\nYou have now %d chips !\n", chips)
	color.Unset()
	color.Set(color.FgCyan, color.Bold)
	again(action, chips)
	return (chips)
}

func h1drawh2draw(action string, chips int, bets int) int {
	color.Set(color.FgWhite, color.Bold)
	fmt.Println("It's a draw, no one lose ! ;-)")
	fmt.Printf("You have always %d chips !\n", chips)
	color.Unset()
	color.Set(color.FgCyan, color.Bold)
	again(action, chips)
	return (chips)
}

func h1drawh2win(action string, chips int, bets int) int {
	color.Set(color.FgGreen, color.Bold)
	fmt.Printf("\nYou won %d chips ! :-)", bets*2)
	chips = chips + bets*2
	fmt.Printf("\nYou have now %d chips !\n", chips)
	color.Unset()
	color.Set(color.FgCyan, color.Bold)
	again(action, chips)
	return (chips)
}

func h1winh2draw(action string, chips int, bets int) int {
	color.Set(color.FgGreen, color.Bold)
	fmt.Printf("\nYou won %d chips ! :-)", bets*2)
	chips = chips + bets*2
	fmt.Printf("\nYou have now %d chips !\n", chips)
	color.Unset()
	color.Set(color.FgCyan, color.Bold)
	again(action, chips)
	return (chips)
}
