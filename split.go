package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func split(playerCard1 int, playerCard2 int, bankhandvalue int, chips int, bets int) int {

	color.Set(color.FgWhite, color.Bold)
	chips = chips - bets
	fmt.Printf("You have now bet %d chips", bets*2)
	fmt.Printf("\nYou have now %d chips\n", chips)
	fmt.Println()

	hand1 := splitHandleFirstHand(playerCard1, chips, bets)
	hand2 := splitHandleSecondHand(playerCard2, chips, bets)

	bankhandvalue = splitDealerHand(bankhandvalue)
	chips = splitthewinneris(bankhandvalue, hand1, hand2, chips, bets)
	return chips
}

func splitHandleFirstHand(playerCard1 int, chips int, bets int) int {
	color.Set(color.FgWhite, color.Bold)
	cardx := random()
	hand1 := playerCard1 + cardx
	fmt.Printf("You are hitting a %d card for your first hand...\n", cardx)
	fmt.Printf("Your first hand value : %d (%d + %d)\n", hand1, playerCard1, (hand1 - playerCard1))
	hand1 = continueSplitGameFirstHand(hand1, chips, bets)
	color.Unset()
	return (hand1)
}

func splitHandleSecondHand(playerCard2 int, chips int, bets int) int {
	color.Set(color.FgWhite, color.Bold)

	cardx := random()
	hand2 := playerCard2 + cardx
	fmt.Printf("You are hitting a %d card for your second hand...\n", cardx)
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
	cardx := random()
	if playerhandvalue >= 11 && cardx == 11 {
		cardx = 1
	}
	playerhandvalue = playerhandvalue + cardx
	fmt.Println()
	color.Set(color.FgWhite, color.Bold)
	fmt.Printf("You are hitting a %d card for your %s hand...\n", cardx, handNumber)
	time.Sleep(1 * time.Second)
	fmt.Printf("Your %s hand value : %d\n", handNumber, playerhandvalue)
	color.Unset()
	return (playerhandvalue)
}

func splitDealerHand(bankhandvalue int) int {
	color.Set(color.FgWhite, color.Bold)
	cardx := random()
	bankhandvalue = bankhandvalue + cardx
	fmt.Printf("The dealer hit a %d card...\n", cardx)
	time.Sleep(1 * time.Second)
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

	if hand1 > 21 && hand2 > 21 {
		chips = h1loseh2lose(action, chips, bets)
		return chips
	} else if hand1 <= 21 && hand2 > 21 && bankhandvalue <= 21 {
		// fmt.Println("H1 can win, H2 bust, bank can win")
		if hand1 > bankhandvalue {
			chips = h1winh2bust(action, chips, bets)
		} else if hand1 == bankhandvalue {
			chips = h1drawh2bust(action, chips, bets)
		} else {
			chips = h1loseh2bust(action, chips, bets)
		}
	} else if hand1 > 21 && hand2 <= 21 && bankhandvalue <= 21 {
		// fmt.Println("H1 bust, H2 can win, bank can win")
		if hand2 > bankhandvalue {
			chips = h1busth2win(action, chips, bets)
		} else if hand2 == bankhandvalue {
			chips = h1busth2draw(action, chips, bets)
		} else {
			chips = h1busth2lose(action, chips, bets)
		}
	} else if hand1 <= 21 && hand2 <= 21 && bankhandvalue > 21 {
		// fmt.Println("H1 win, H2 win, bank bust")
		chips = h1winh2win(action, chips, bets)
	} else if hand1 <= 21 && hand2 > 21 && bankhandvalue > 21 {
		// fmt.Println("H1 win, H2 bust, bank bust")
		chips = h1winh2lose(action, chips, bets)
	} else {
		if bankhandvalue <= 21 && hand1 <= 21 && hand2 <= 21 {
			if (bankhandvalue < hand1 && bankhandvalue < hand2) && hand1 <= 21 && hand2 <= 21 {
				// fmt.Println("H1 win && H2 win")
				chips = h1winh2win(action, chips, bets)
			} else if bankhandvalue < hand1 && (bankhandvalue > hand2 || hand2 > 21) {
				// fmt.Println("H1 win && H2 lose")
				chips = h1winh2lose(action, chips, bets)
			} else if bankhandvalue > hand1 && (bankhandvalue < hand2 && hand2 <= 21) {
				// fmt.Println("H1 lose && H2 win")
				chips = h1loseh2win(action, chips, bets)
			} else if bankhandvalue > hand1 && (bankhandvalue > hand2 || hand2 > 21) {
				// fmt.Println("H1 lose && H2 lose")
				chips = h1loseh2lose(action, chips, bets)
			} else if bankhandvalue > hand1 && bankhandvalue == hand2 {
				// fmt.Println("H1 lose && H2 draw")
				chips = h1loseh2draw(action, chips, bets)
			} else if bankhandvalue == hand1 && (bankhandvalue > hand2 || hand2 > 21) {
				chips = h1drawh2lose(action, chips, bets)
			} else if bankhandvalue == hand1 && bankhandvalue == hand2 {
				// fmt.Println("H1 draw && H2 draw")
				chips = h1drawh2draw(action, chips, bets)
			} else if bankhandvalue == hand1 && bankhandvalue < hand2 {
				// fmt.Println("H1 draw && H2 win")
				chips = h1drawh2win(action, chips, bets)
			} else if bankhandvalue < hand1 && bankhandvalue == hand2 {
				// fmt.Println("H1 win && H2 draw")
				chips = h1winh2draw(action, chips, bets)
			} else {
				// chips = youlose2(action, chips, bets)
				fmt.Println("AUTRE CAS !")
			}
		}
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
	fmt.Printf("\nYou won %d chips ! :-)", bets)
	chips = chips + bets
	fmt.Printf("\nYou have now %d chips !\n", chips)
	color.Unset()
	color.Set(color.FgCyan, color.Bold)
	again(action, chips)
	return (chips)
}

func h1loseh2win(action string, chips int, bets int) int {
	color.Set(color.FgGreen, color.Bold)
	fmt.Printf("\nYou won %d chips ! :-)", bets)
	chips = chips + bets
	fmt.Printf("\nYou have now %d chips !\n", chips)
	color.Unset()
	color.Set(color.FgCyan, color.Bold)
	again(action, chips)
	return (chips)
}

func h1loseh2lose(action string, chips int, bets int) int {
	color.Set(color.FgRed, color.Bold)
	fmt.Printf("\nYou lost %d chips ! :-)", bets*2)
	fmt.Printf("\nYou have now %d chips !\n", chips)
	color.Unset()
	color.Set(color.FgCyan, color.Bold)
	again(action, chips)
	return (chips)
}

func h1loseh2draw(action string, chips int, bets int) int {
	color.Set(color.FgRed, color.Bold)
	fmt.Printf("\nYou lost %d chips ! :-)", bets)
	chips = chips + bets
	fmt.Printf("\nYou have now %d chips !\n", chips)
	color.Unset()
	color.Set(color.FgCyan, color.Bold)
	again(action, chips)
	return (chips)
}

func h1drawh2lose(action string, chips int, bets int) int {
	color.Set(color.FgRed, color.Bold)
	fmt.Printf("\nYou lost %d chips ! :-)", bets)
	chips = chips + bets
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
	chips = chips + bets*2 + bets
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

func h1busth2win(action string, chips int, bets int) int {
	color.Set(color.FgGreen, color.Bold)
	fmt.Printf("\nYou won %d chips ! :-)", bets*2)
	chips = chips + bets
	fmt.Printf("\nYou have now %d chips !\n", chips)
	color.Unset()
	color.Set(color.FgCyan, color.Bold)
	again(action, chips)
	return (chips)
}

func h1drawh2bust(action string, chips int, bets int) int {
	color.Set(color.FgRed, color.Bold)
	fmt.Printf("\nYou lose %d chips ! :-)", bets*2)
	chips = chips - bets
	fmt.Printf("\nYou have now %d chips !\n", chips)
	color.Unset()
	color.Set(color.FgCyan, color.Bold)
	again(action, chips)
	return (chips)
}

func h1busth2lose(action string, chips int, bets int) int {
	color.Set(color.FgRed, color.Bold)
	fmt.Printf("\nYou lose %d chips ! :-)", bets*2)
	fmt.Printf("\nYou have now %d chips !\n", chips)
	color.Unset()
	color.Set(color.FgCyan, color.Bold)
	again(action, chips)
	return (chips)
}

func h1winh2bust(action string, chips int, bets int) int {
	color.Set(color.FgGreen, color.Bold)
	fmt.Printf("\nYou won %d chips ! :-)", bets*2)
	chips = chips + bets
	fmt.Printf("\nYou have now %d chips !\n", chips)
	color.Unset()
	color.Set(color.FgCyan, color.Bold)
	again(action, chips)
	return (chips)
}

func h1busth2draw(action string, chips int, bets int) int {
	color.Set(color.FgRed, color.Bold)
	fmt.Printf("\nYou lose %d chips ! :-)", bets)
	chips = chips + bets
	fmt.Printf("\nYou have now %d chips !\n", chips)
	color.Unset()
	color.Set(color.FgCyan, color.Bold)
	again(action, chips)
	return (chips)
}

func h1loseh2bust(action string, chips int, bets int) int {
	color.Set(color.FgRed, color.Bold)
	fmt.Printf("\nYou lose %d chips ! :-)", bets*2)
	fmt.Printf("\nYou have now %d chips !\n", chips)
	color.Unset()
	color.Set(color.FgCyan, color.Bold)
	again(action, chips)
	return (chips)
}
