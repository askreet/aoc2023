package day7

type Hand struct {
	Cards     []byte
	Bid       int
	Freq      map[byte]int8
	Value     int
	JokerRule bool
	TypeName  string
}

func NewHand(in []byte, bid int, jokers bool) Hand {
	h := Hand{
		Cards:     in,
		Bid:       bid,
		JokerRule: jokers,
	}
	h.freqAnalysis()
	h.computeValue()
	return h
}

const (
	HighCard = 1 + iota
	OnePair
	TwoPair
	ThreeofaKind
	FullHouse
	FourofaKind
	FiveofaKind
)

// Value converts a hand into a unique value allowing it to be sorted among its peers.
func (h *Hand) computeValue() {
	// T 11 22 33 44 55
	// \--[ Type of hand, highest wins.
	//   \\--[ Value of first card, breaks ties.
	//      \\--[ Value of second card, breaks ties.
	//         \\ etc.
	v := 0

	switch {
	case h.isFiveOfAKind():
		h.TypeName = "Five-of-a-Kind"
		v = 7_00_00_00_00_00
	case h.isFourOfAKind():
		h.TypeName = "Four-of-a-Kind"
		v = 6_00_00_00_00_00
	case h.isFullHouse():
		h.TypeName = "Full House"
		v = 5_00_00_00_00_00
	case h.isThreeOfAKind():
		h.TypeName = "Three-of-a-Kind"
		v = 4_00_00_00_00_00
	case h.isTwoPair():
		h.TypeName = "Two Pair"
		v = 3_00_00_00_00_00
	case h.isOnePair():
		h.TypeName = "One Pair"
		v = 2_00_00_00_00_00
	case h.isHighCard():
		h.TypeName = "High Card"
		v = 1_00_00_00_00_00
	}

	v += h.valueOfCard(h.Cards[0]) * 1_00_00_00_00
	v += h.valueOfCard(h.Cards[1]) * 1_00_00_00
	v += h.valueOfCard(h.Cards[2]) * 1_00_00
	v += h.valueOfCard(h.Cards[3]) * 1_00
	v += h.valueOfCard(h.Cards[4]) * 1

	h.Value = v
}

var CardValues = map[byte]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}
var CardValuesWithJokerRule = map[byte]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1,
}

func (h *Hand) Rank() int {
	return h.Value / 1_00_00_00_00_00
}

func (h *Hand) valueOfCard(b byte) int {
	var cardValues = CardValues
	if h.JokerRule {
		cardValues = CardValuesWithJokerRule
	}

	if v, ok := cardValues[b]; ok {
		return v
	} else {
		panic("unexpected card byte " + string(b))
	}
}

func (h *Hand) nJokers() int8 {
	if !h.JokerRule {
		return 0
	}

	if n, ok := h.Freq['J']; ok {
		return n
	} else {
		return 0
	}
}

func (h *Hand) freqAnalysis() {
	h.Freq = make(map[byte]int8)

	for _, b := range h.Cards {
		if _, ok := h.Freq[b]; ok {
			h.Freq[b] += 1
		} else {
			h.Freq[b] = 1
		}
	}
}

func (h *Hand) isFiveOfAKind() bool {
	jokersAvailable := h.nJokers()

	for b, v := range h.Freq {
		if b == 'J' && v == 5 {
			return true
		} else if b == 'J' {
			continue
		}

		if v+jokersAvailable == 5 {
			return true
		}
	}

	return false
}

func (h *Hand) isFourOfAKind() bool {
	jokersAvailable := h.nJokers()

	for b, v := range h.Freq {
		if b == 'J' {
			continue
		}

		if v+jokersAvailable == 4 {
			return true
		}
	}

	return false
}

func (h *Hand) isFullHouse() bool {
	jokersAvailable := h.nJokers()
	foundThreeOfAKind := false
	foundTwoOfAKind := false

	for b, n := range h.Freq {
		if b == 'J' {
			continue
		}

		if n+jokersAvailable >= 3 {
			foundThreeOfAKind = true
			jokersAvailable -= 3 - n
			continue
		}

		if n+jokersAvailable >= 2 {
			foundTwoOfAKind = true
			jokersAvailable -= 2 - n
			continue
		}
	}

	return foundTwoOfAKind && foundThreeOfAKind
}

func (h *Hand) isThreeOfAKind() bool {
	for _, v := range h.Freq {
		if v+h.nJokers() == 3 {
			return true
		}
	}

	return false
}

func (h *Hand) isTwoPair() bool {
	nPairs := 0
	jokersAvailable := h.nJokers()

	for _, n := range h.Freq {
		if n+jokersAvailable >= 2 {
			nPairs++
			jokersAvailable -= 2 - n
		}
	}

	return nPairs == 2
}

func (h *Hand) isOnePair() bool {
	nPairs := 0
	jokersAvailable := h.nJokers()

	for _, n := range h.Freq {
		if n+jokersAvailable >= 2 {
			nPairs++
			jokersAvailable -= 2 - n
		}
	}

	return nPairs == 1
}

func (h *Hand) isHighCard() bool { return true }
