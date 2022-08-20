package blackjack

var Cards = map[string]int{
    "ace": 11,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
    "ten": 10,
    "jack": 10,
    "queen": 10,
    "king": 10,
    "other": 0,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    return Cards[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	cardSum := ParseCard(card1) + ParseCard(card2)
    var move string

    switch {
    case (card1 == "ace") && (card2 == "ace"):
    	move = "P"
    case (cardSum == 21):
    	if ParseCard(dealerCard) >= 10 {
        	move = "S"    
        } else {
        	move = "W"
        }        
	case (cardSum >= 17) && (cardSum <= 20):
		move = "S"
    case (cardSum >= 12) && (cardSum <= 16):
		if ParseCard(dealerCard) >= 7 {
            move = "H"
        } else {
        	move = "S"
        }
    case cardSum <= 11:
		move = "H"
    }
	return move
}
