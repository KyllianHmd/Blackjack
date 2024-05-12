package main

import (
	"fmt"
	"time"
)

func split(playerCard1 int, playerCard2 int, bankhandvalue int, chips int, bets int) int {
	chips = chips - bets
	fmt.Printf("You have now bet %d chips", bets*2)
	fmt.Printf("\nYou have now %d chips\n\n", chips)

	hand1, chips, bets := splitHandleFirstHand(playerCard1, chips, bets)
	hand2, chips, bets := splitHandleSecondHand(playerCard2, chips, bets)

	bankhandvalue = splitDealerHand(bankhandvalue)
	chips = splitthewinneris(bankhandvalue, hand1, hand2, chips, bets)
	return chips
}

func splitHandleFirstHand(playerCard1 int, chips int, bets int) (int, int, int) {
	cardx := random()
	hand1 := playerCard1 + cardx
	fmt.Printf("You are hitting a %d card for your first hand...\n", cardx)
	fmt.Printf("Your first hand value : %d (%d + %d)\n", hand1, playerCard1, (hand1 - playerCard1))
	hand1, chips, bets = continueSplitGameFirstHand(hand1, chips, bets)
	return hand1, chips, bets
}

func splitHandleSecondHand(playerCard2 int, chips int, bets int) (int, int, int) {
	cardx := random()
	hand2 := playerCard2 + cardx
	fmt.Printf("You are hitting a %d card for your second hand...\n", cardx)
	fmt.Printf("Your second hand value : %d (%d + %d)\n", hand2, playerCard2, (hand2 - playerCard2))
	hand2, chips, bets = continueSplitGameSecondHand(hand2, chips, bets)
	return hand2, chips, bets
}
func continueSplitGameFirstHand(hand1 int, chips int, bets int) (int, int, int) {
	fmt.Print("\n> ")
	for {
		var action string
		fmt.Scanf("%s", &action)
		if action == "hit" {
			hand1 = hitSplitHand(hand1, "first")
			if hand1 > 21 {
				return hand1, chips, bets
			}
		} else if action == "stand" {
			fmt.Printf("\nYour first hand value : %d\n", hand1)
			return hand1, chips, bets
		} else if action == "quit" {
			quit()
		} else if action == "restart" {
			restart()
		} else if action == "chips" {
			howmanychips(chips)
		} else if action == "surrender" {
			return 0, chips, bets
		} else if action == "double" {
			if chips-bets < 0 {
				fmt.Println("\nSorry, you haven't enought money to double your hand.")
			} else {
				chips, hand1, bets = splitDoubleDownHand(hand1, chips, bets)
				return hand1, chips, bets
			}
		} else if action == "split" {
			fmt.Println("\nSorry, you can not split your hand again because you've already done it once.")
		} else {
			fmt.Println("\nPlease, enter a real decision.")
		}
		fmt.Print("\n> ")
	}
}

func continueSplitGameSecondHand(hand2 int, chips int, bets int) (int, int, int) {
	fmt.Print("\n> ")
	for {
		var action string
		fmt.Scanf("%s", &action)
		if action == "hit" {
			hand2 = hitSplitHand(hand2, "second")
			if hand2 > 21 {
				return hand2, chips, bets
			}
		} else if action == "stand" {
			fmt.Printf("\nYour second hand value : %d\n", hand2)
			return hand2, chips, bets
		} else if action == "quit" {
			quit()
		} else if action == "restart" {
			restart()
		} else if action == "chips" {
			howmanychips(chips)
		} else if action == "surrender" {
			return 0, chips, bets
		} else if action == "double" {
			if chips-bets < 0 {
				fmt.Println("\nSorry, you haven't enought money to double your hand.")
			} else {
				chips, hand2, bets = splitDoubleDownHand(hand2, chips, bets)
				return hand2, chips, bets
			}
		} else if action == "split" {
			fmt.Println("\nSorry, you can not split your hand again because you've already done it once.")
		} else {
			fmt.Println("\nPlease, enter a real decision.")
		}
		fmt.Print("\n> ")
	}
}

func hitSplitHand(playerhandvalue int, handNumber string) int {
	cardx := random()
	if playerhandvalue >= 11 && cardx == 11 {
		cardx = 1
	}
	playerhandvalue = playerhandvalue + cardx
	fmt.Printf("\nYou are hitting a %d card for your %s hand...\n", cardx, handNumber)
	time.Sleep(1 * time.Second)
	fmt.Printf("Your %s hand value : %d\n", handNumber, playerhandvalue)
	return (playerhandvalue)
}

func splitDealerHand(bankhandvalue int) int {
	cardx := random()
	bankhandvalue = bankhandvalue + cardx
	fmt.Printf("The dealer hit a %d card...\n", cardx)
	time.Sleep(1 * time.Second)
	fmt.Printf("The dealer's hand value : %d\n", bankhandvalue)
	for bankhandvalue < 17 {
		cardx := random()
		bankhandvalue = bankhandvalue + cardx
		fmt.Printf("The dealer hit a %d card...\n", cardx)
		time.Sleep(1 * time.Second)
		fmt.Printf("The dealer's hand value : %d\n", bankhandvalue)
	}
	return bankhandvalue
}

func splitthewinneris(bankhandvalue int, hand1 int, hand2 int, chips int, bets int) int {
	var action string
	fmt.Printf("\nYour first hand value : %d\n", hand1)
	fmt.Printf("Your second hand value : %d\n", hand2)
	fmt.Printf("The dealer's hand value : %d\n", bankhandvalue)

	if hand1 > 21 && hand2 > 21 {
		chips = h1loseh2lose(action, chips, bets)
		return chips
	} else if hand1 <= 21 && hand2 > 21 && bankhandvalue <= 21 {
		if hand1 > bankhandvalue {
			chips = h1winh2bust(action, chips, bets)
		} else if hand1 == bankhandvalue {
			chips = h1drawh2bust(action, chips, bets)
		} else {
			chips = h1loseh2bust(action, chips, bets)
		}
	} else if hand1 > 21 && hand2 <= 21 && bankhandvalue <= 21 {
		if hand2 > bankhandvalue {
			chips = h1busth2win(action, chips, bets)
		} else if hand2 == bankhandvalue {
			chips = h1busth2draw(action, chips, bets)
		} else {
			chips = h1busth2lose(action, chips, bets)
		}
	} else if hand1 <= 21 && hand2 <= 21 && bankhandvalue > 21 {
		chips = h1winh2win(action, chips, bets)
	} else if hand1 <= 21 && hand2 > 21 && bankhandvalue > 21 {
		chips = h1winh2lose(action, chips, bets)
	} else {
		if bankhandvalue <= 21 && hand1 <= 21 && hand2 <= 21 && hand1 != 0 && hand2 != 0 {
			if (bankhandvalue < hand1 && bankhandvalue < hand2) && hand1 <= 21 && hand2 <= 21 {
				chips = h1winh2win(action, chips, bets)
			} else if bankhandvalue < hand1 && (bankhandvalue > hand2 || hand2 > 21) {
				chips = h1winh2lose(action, chips, bets)
			} else if bankhandvalue > hand1 && (bankhandvalue < hand2 && hand2 <= 21) {
				chips = h1loseh2win(action, chips, bets)
			} else if bankhandvalue > hand1 && (bankhandvalue > hand2 || hand2 > 21) {
				chips = h1loseh2lose(action, chips, bets)
			} else if bankhandvalue > hand1 && bankhandvalue == hand2 {
				chips = h1loseh2draw(action, chips, bets)
			} else if bankhandvalue == hand1 && (bankhandvalue > hand2 || hand2 > 21) {
				chips = h1drawh2lose(action, chips, bets)
			} else if bankhandvalue == hand1 && bankhandvalue == hand2 {
				chips = h1drawh2draw(action, chips)
			} else if bankhandvalue == hand1 && bankhandvalue < hand2 {
				chips = h1drawh2win(action, chips, bets)
			} else if bankhandvalue < hand1 && bankhandvalue == hand2 {
				chips = h1winh2draw(action, chips, bets)
			} else {
				fmt.Println("AUTRE CAS !")
			}
		} else {
			fmt.Println("SURRENDER CASES")
			if hand1 == 0 && hand2 > bankhandvalue && hand2 != 0 {
				fmt.Println("h1ffh2win")
				chips = h1ffh2win(action, chips, bets)
			} else if hand1 == 0 && hand2 == bankhandvalue {
				fmt.Println("h1ffh2draw")
				chips = h1ffh2draw(action, chips, bets)
			} else if hand1 == 0 && hand2 < bankhandvalue && hand2 != 0 {
				fmt.Println("h1ffh2lose")
				chips = h1ffh2lose(action, chips, bets)
			} else if hand1 == 0 && hand2 == 0 {
				fmt.Println("h1ffh2ff")
				chips = h1ffh2ff(action, chips, bets)
			} else if hand2 == 0 && hand1 > bankhandvalue && hand1 != 0 {
				fmt.Println("h1winh2ff")
				chips = h1winh2ff(action, chips, bets)
			} else if hand2 == 0 && hand1 == bankhandvalue {
				fmt.Println("h1drawh2ff")
				chips = h1drawh2ff(action, chips, bets)
			} else if hand2 == 0 && hand1 < bankhandvalue && hand1 != 0 {
				fmt.Println("h1loseh2ff")
				chips = h1loseh2ff(action, chips, bets)
			}
		}
	}
	return (chips)
}

func h1winh2win(action string, chips int, bets int) int {
	fmt.Printf("\nYou won %d chips ! :-)", bets*2+bets*2)
	chips = chips + bets*2 + bets*2
	fmt.Printf("\nYou have now %d chips !\n", chips)
	again(action, chips)
	return (chips)
}

func h1winh2lose(action string, chips int, bets int) int {
	fmt.Printf("\nYou won %d chips ! :-)", bets)
	chips = chips + bets
	fmt.Printf("\nYou have now %d chips !\n", chips)
	again(action, chips)
	return (chips)
}

func h1loseh2win(action string, chips int, bets int) int {
	fmt.Printf("\nYou won %d chips ! :-)", bets)
	chips = chips + bets
	fmt.Printf("\nYou have now %d chips !\n", chips)
	again(action, chips)
	return (chips)
}

func h1loseh2lose(action string, chips int, bets int) int {
	fmt.Printf("\nYou lost %d chips ! :-)", bets*2)
	fmt.Printf("\nYou have now %d chips !\n", chips)
	again(action, chips)
	return (chips)
}

func h1loseh2draw(action string, chips int, bets int) int {
	fmt.Printf("\nYou lost %d chips ! :-)", bets)
	chips = chips + bets
	fmt.Printf("\nYou have now %d chips !\n", chips)
	again(action, chips)
	return (chips)
}

func h1drawh2lose(action string, chips int, bets int) int {
	fmt.Printf("\nYou lost %d chips ! :-)", bets)
	chips = chips + bets
	fmt.Printf("\nYou have now %d chips !\n", chips)
	again(action, chips)
	return (chips)
}

func h1drawh2draw(action string, chips int) int {
	fmt.Println("It's a draw, no one lose ! ;-)")
	fmt.Printf("You have always %d chips !\n", chips)
	again(action, chips)
	return (chips)
}

func h1drawh2win(action string, chips int, bets int) int {
	fmt.Printf("\nYou won %d chips ! :-)", bets*2)
	chips = chips + bets*2 + bets
	fmt.Printf("\nYou have now %d chips !\n", chips)
	again(action, chips)
	return (chips)
}

func h1winh2draw(action string, chips int, bets int) int {
	fmt.Printf("\nYou won %d chips ! :-)", bets*2)
	chips = chips + bets*2
	fmt.Printf("\nYou have now %d chips !\n", chips)
	again(action, chips)
	return (chips)
}

func h1busth2win(action string, chips int, bets int) int {
	fmt.Printf("\nYou won %d chips ! :-)", bets*2)
	chips = chips + bets
	fmt.Printf("\nYou have now %d chips !\n", chips)
	again(action, chips)
	return (chips)
}

func h1drawh2bust(action string, chips int, bets int) int {
	fmt.Printf("\nYou lose %d chips ! :-)", bets*2)
	chips = chips - bets
	fmt.Printf("\nYou have now %d chips !\n", chips)
	again(action, chips)
	return (chips)
}

func h1busth2lose(action string, chips int, bets int) int {
	fmt.Printf("\nYou lose %d chips ! :-)", bets*2)
	fmt.Printf("\nYou have now %d chips !\n", chips)
	again(action, chips)
	return (chips)
}

func h1winh2bust(action string, chips int, bets int) int {
	fmt.Printf("\nYou won %d chips ! :-)", bets*2)
	chips = chips + bets
	fmt.Printf("\nYou have now %d chips !\n", chips)
	again(action, chips)
	return (chips)
}

func h1busth2draw(action string, chips int, bets int) int {
	fmt.Printf("\nYou lose %d chips ! :-)", bets)
	chips = chips + bets
	fmt.Printf("\nYou have now %d chips !\n", chips)
	again(action, chips)
	return (chips)
}

func h1loseh2bust(action string, chips int, bets int) int {
	fmt.Printf("\nYou lose %d chips ! :-)", bets*2)
	fmt.Printf("\nYou have now %d chips !\n", chips)
	again(action, chips)
	return (chips)
}
func h1ffh2win(action string, chips int, bets int) int {
	fmt.Printf("\nYou have surrender your first hand and won %d chips with your second one!\n", bets*2)
	chips = chips + bets/2 + bets*2
	fmt.Printf("You have now %d chips !\n", chips)
	again(action, chips)
	return (chips)
}

func h1ffh2draw(action string, chips int, bets int) int {
	fmt.Printf("\nYou have surrender your first hand and draw with your second one. You won %d chips.\n", bets)
	chips = chips + bets/2 + bets
	fmt.Printf("You have now %d chips !\n", chips)
	again(action, chips)
	return (chips)
}
func h1ffh2lose(action string, chips int, bets int) int {
	fmt.Printf("\nYou have surrender your first hand and lose %d chips with your second one. \n", bets)
	chips = chips + bets/2 - bets
	fmt.Printf("You have now %d chips !\n", chips)
	again(action, chips)
	return (chips)
}
func h1ffh2ff(action string, chips int, bets int) int {
	fmt.Printf("\nYou have surrender your first and your second hand. We've refund you %d chips.\n", bets/2+bets/2)
	chips = chips + bets/2 + bets/2
	fmt.Printf("You have now %d chips !\n", chips)
	again(action, chips)
	return (chips)
}

func h1winh2ff(action string, chips int, bets int) int {
	fmt.Printf("\nYou have won %d chips with your first hand and you have surrender your second one!\n", bets*2)
	chips = chips + bets/2 + bets*2
	fmt.Printf("You have now %d chips !\n", chips)
	again(action, chips)
	return (chips)
}

func h1drawh2ff(action string, chips int, bets int) int {
	fmt.Printf("\nYou draw with your first one and surrender your second hand. You won %d chips.\n", bets)
	chips = chips + bets/2 + bets
	fmt.Printf("You have now %d chips !\n", chips)
	again(action, chips)
	return (chips)
}

func h1loseh2ff(action string, chips int, bets int) int {
	fmt.Printf("\nYou have lost %d chips with your first one and surrender your second hand. \n", bets)
	chips = chips + bets/2 - bets
	fmt.Printf("You have now %d chips !\n", chips)
	again(action, chips)
	return (chips)
}

func splitDoubleDownHand(playerhandvalue int, chips int, bets int) (int, int, int) {
	cardx := random()
	if playerhandvalue >= 11 && cardx == 11 {
		cardx = 1
	}
	playerhandvalue = playerhandvalue + cardx
	chips = chips - bets
	fmt.Printf("\nYou have now bet %d chips", (bets*2 + bets))
	fmt.Printf("\nYou have now %d chips\n", chips)
	fmt.Printf("\nYou are hitting a %d card...\n", cardx)
	time.Sleep(1 * time.Second)
	fmt.Printf("Your hand value : %d\n", playerhandvalue)
	return chips, playerhandvalue, bets
}
