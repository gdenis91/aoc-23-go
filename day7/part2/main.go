package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"

	"github.com/gdenis91/aoc-23-go/util"
)

//go:embed input.txt
var input string

func main() {
	hands := mustParseInput(input)
	slices.SortFunc(hands, func(a hand, b hand) int {
		if a.beats(b) {
			return 1
		}
		return -1
	})
	fmt.Println("Ranked Hands:")
	var winnings int
	for i, h := range hands {
		fmt.Printf("Hand %d: %s - %s\n", i, h, h.handType())
		winnings += (i + 1) * h.bid
	}
	fmt.Println("Winnings:", winnings)
}

// 32T3K 765
func mustParseInput(input string) []hand {
	lines := strings.Split(input, "\n")
	hands := make([]hand, 0, len(lines))
	for _, l := range lines {
		fields := strings.Fields(l)
		h := hand{
			bid:        util.MustAtoi(fields[1]),
			cards:      make([]card, 0, 5),
			cardCounts: make(map[card]int, 5),
			numWilds:   0,
		}
		for _, r := range fields[0] {
			c := cardFromRune(r)
			h.cards = append(h.cards, c)
			if c == CardJ {
				h.numWilds++
			}
			h.cardCounts[c]++
		}
		hands = append(hands, h)
	}
	return hands
}

type card int

const (
	CardJ card = iota
	Card2
	Card3
	Card4
	Card5
	Card6
	Card7
	Card8
	Card9
	CardT
	CardQ
	CardK
	CardA
)

func cardFromRune(r rune) card {
	switch r {
	case '2':
		return Card2
	case '3':
		return Card3
	case '4':
		return Card4
	case '5':
		return Card5
	case '6':
		return Card6
	case '7':
		return Card7
	case '8':
		return Card8
	case '9':
		return Card9
	case 'T':
		return CardT
	case 'J':
		return CardJ
	case 'Q':
		return CardQ
	case 'K':
		return CardK
	case 'A':
		return CardA
	}
	panic(fmt.Errorf("unrecognized card %s", string(r)))
}

func (c card) String() string {
	switch c {
	case Card2:
		return "2"
	case Card3:
		return "3"
	case Card4:
		return "4"
	case Card5:
		return "5"
	case Card6:
		return "6"
	case Card7:
		return "7"
	case Card8:
		return "8"
	case Card9:
		return "9"
	case CardT:
		return "T"
	case CardJ:
		return "J"
	case CardQ:
		return "Q"
	case CardK:
		return "K"
	case CardA:
		return "A"
	}
	return "-"
}

type handType int

const (
	HighCard    handType = iota // 5 card types 2 3 4 5 6
	OnePair                     // 4 card types 2 2 3 4 5
	TwoPair                     // 3 card types 2 2 3 3 4
	ThreeOfKind                 // 3 card types 2 2 2 3 4
	FullHouse                   // 2 card types 2 2 2 3 3
	FourOfKind                  // 2 card types 2 2 2 2 3
	FiveOfKind                  // 1 card types 2 2 2 2 2
)

func (t handType) String() string {
	switch t {
	case HighCard:
		return "HighCard"
	case OnePair:
		return "OnePair"
	case TwoPair:
		return "TwoPair"
	case ThreeOfKind:
		return "ThreeOfKind"
	case FullHouse:
		return "FullHouse"
	case FourOfKind:
		return "FourOfKind"
	case FiveOfKind:
		return "FiveOfKind"
	}
	return "-"
}

type hand struct {
	cards      []card
	cardCounts map[card]int
	bid        int
	numWilds   int
}

func (h hand) String() string {
	cards := make([]string, 0, len(h.cards))
	for _, v := range h.cards {
		cards = append(cards, v.String())
	}
	return strings.Join(cards, " ") + " NumWilds: " + fmt.Sprint(h.numWilds)
}

func (h hand) handType() handType {
	fmt.Println("wilds: ", h.numWilds)
	switch len(h.cardCounts) {
	case 1:
		return FiveOfKind
	case 2:
		for _, v := range h.cardCounts {
			if v == 4 || v == 1 {
				if h.numWilds == 1 || h.numWilds == 4 {
					return FiveOfKind
				}
				return FourOfKind
			}
			if h.numWilds == 2 || h.numWilds == 3 {
				return FiveOfKind
			}
			return FullHouse
		}
	case 3:
		for _, v := range h.cardCounts {
			if v == 3 {
				if h.numWilds == 1 || h.numWilds == 3 {
					return FourOfKind
				} else if h.numWilds == 2 {
					return FiveOfKind
				}
				return ThreeOfKind
			}
		}
		if h.numWilds == 1 {
			return FullHouse
		} else if h.numWilds == 2 {
			return FourOfKind
		}
		return TwoPair
	case 4:
		if h.numWilds == 1 || h.numWilds == 2 {
			return ThreeOfKind
		} else if h.numWilds == 3 {
			return FiveOfKind
		}
		return OnePair
	case 5:
		if h.numWilds == 1 {
			return OnePair
		} else if h.numWilds == 2 {
			return ThreeOfKind
		} else if h.numWilds == 3 {
			return FourOfKind
		} else if h.numWilds == 4 {
			return FiveOfKind
		}
		return HighCard
	}
	panic(fmt.Errorf("unknown hand type: %v", h.cardCounts))
}

func (h hand) beats(other hand) bool {
	thisHand := h.handType()
	otherHand := other.handType()
	if thisHand == otherHand {
		for i := 0; i < len(h.cards); i++ {
			if h.cards[i] == other.cards[i] {
				continue
			}
			return h.cards[i] > other.cards[i]
		}
	}
	return thisHand > otherHand
}
