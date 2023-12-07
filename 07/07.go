package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	cards []Card
	htype HandType
	bid   int64
}

type HandType int

const (
	HighCard = iota
	OnePair
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func (h HandType) Val() int {
	return int(h)
}

func (h HandType) String() string {
	return [...]string{"High Card", "One Pair", "Two Pairs", "Three of a Kind", "Full House", "Four of a Kind", "Five of a Kind"}[h]
}

type Card int

const (
	Two = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

func (c Card) Val() int {
	return int(c)
}

func (c Card) String() string {
	return [...]string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}[c]
}

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	var data []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		curr_row := scanner.Text()
		if err != nil {
			fmt.Println(err)
		}
		data = append(data, curr_row)
	}

	var hands []Hand
	for _, d := range data {
		h := strings.Split(d, " ")[0]
		b := strings.Split(d, " ")[1]
		cards := parse_hand(h)
		handtype := get_hand_type(cards)
		bet, _ := strconv.ParseInt(b, 10, 64)
		hand := Hand{cards: cards, htype: handtype, bid: bet}
		// fmt.Printf("%v:\n cards: %v\n type: %v\n bid: %v\n", d, cards, handtype.String(), bet)
		hands = append(hands, hand)
	}

	slices.SortFunc(hands, func(a, b Hand) int {
		if handval := cmp.Compare(a.htype, b.htype); handval != 0 {
			return handval
		} else {
			for i := 0; i < 5; i++ {
				if c := cmp.Compare(a.cards[i], b.cards[i]); c != 0 {
					return c
				}
			}
		}
		return 0
	})

	part1 := 0
	for i, hand := range hands {
		part1 += (i + 1) * int(hand.bid)
	}
	fmt.Printf("part1: %v\n", part1)
	// fmt.Printf("%v: %T\n", strings.Split("test", "")[0], strings.Split("test", "")[0])
}

func parse_hand(s string) []Card {
	cards := strings.Split(s, "")
	var hand []Card
	for _, c := range cards {
		card := get_card(c)
		hand = append(hand, card)
	}
	return hand
}

func get_card(s string) Card {
	switch s {
	case "2":
		return Two
	case "3":
		return Three
	case "4":
		return Four
	case "5":
		return Five
	case "6":
		return Six
	case "7":
		return Seven
	case "8":
		return Eight
	case "9":
		return Nine
	case "T":
		return Ten
	case "J":
		return Jack
	case "Q":
		return Queen
	case "K":
		return King
	case "A":
		return Ace
	}
	return -1
}

func get_hand_type(h []Card) HandType {
	c_counts := make(map[Card]int)
	for _, c := range h {
		_, exists := c_counts[c]
		if exists {
			c_counts[c]++
		} else {
			c_counts[c] = 1
		}
	}
	if l := len(c_counts); l == 1 {
		return FiveOfAKind
	} else if l == 5 {
		return HighCard
	} else if l == 4 {
		return OnePair
	} else if l == 2 { // FH / 4ofakind
		for _, v := range c_counts {
			if v == 4 || v == 1 {
				return FourOfAKind
			} else {
				return FullHouse
			}
		}
	} else if l == 3 {
		for _, v := range c_counts {
			if v == 3 {
				return ThreeOfAKind
			} else if v == 2 {
				return TwoPairs
			}
		}
	}
	return -1
}
