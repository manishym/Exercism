package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	m := map[string]int{
		"ace":   11,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"ten":   10,
		"jack":  10,
		"queen": 10,
		"king":  10,
	}
	if val, ok := m[card]; ok {
		return val
	} else {
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	// panic("Please implement the FirstTurn function")
	c1 := ParseCard(card1)
	c2 := ParseCard(card2)
	d := ParseCard(dealerCard)
	if c1 < c2 {
		c1, c2 = c2, c1
	}
	switch {
	case c1 == 11 && c2 == 11:
		return "P"
	case c1 == 11 && c2 == 10:
		if d < 10 {
			return "W"
		} else {
			return "S"
		}
	case c1+c2 >= 17:
		return "S"

	case c1+c2 <= 11:
		return "H"

	case d >= 7:
		return "H"
	default:
		return "S"
	}

}
