package blackjack

import (
	_ "strconv"
)
// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch {
	case card == "ace":
		return 11
	case card == "one":
		return 1
	case card == "two":
		return 2
	case card == "three":
		return 3
	case card == "four":
		return 4
	case card == "five":
		return 5
	case card == "six":
		return 6
	case card == "seven":
		return 7
	case card == "eight":
		return 8
	case card == "nine"
		return 9
	case card == "ten" || card == "jack" || card == "queen" || card == "king":
		return 10
	default:
		return 0
	}
	}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	switch {
	case ParseCard(card1) + ParseCard(card2) == 21:
		return true
		default:
		return false
}
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	switch {
	case isBlackjack && dealerScore < 10:
		return "W"
	case !isBlackjack:
		return "P"
	default:
		return "S"
	
	}
}
// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	switch {
	case handScore >= 17:
		return "S"
	case handScore <= 11:
		return "H"
	case (handScore <=12 || >=16) && dealerScore >=7:
		return "P"
	default:
		return "S"

	}
}
