package day7

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

const ExampleInput = "32T3K 765\nT55J5 684\nKK677 28\nKTJJT 220\nQQQJA 483\n"

func TestSolution_Part1(t *testing.T) {
	input := bytes.NewBufferString(ExampleInput)

	result := Solution{}.Part1(input)

	assert.Equal(t, 6440, result)
}

func TestSolution_Part2(t *testing.T) {
	input := bytes.NewBufferString(ExampleInput)

	result := Solution{}.Part2(input)

	assert.Equal(t, 5905, result)
}

func TestHand_JokerFullHouse(t *testing.T) {
	assert.Equal(t, FullHouse, jokerHand("AAJKK").Rank())
}

func TestHand_JokerCantBeUsedTwiceInFullHouse(t *testing.T) {
	assert.Equal(t, ThreeofaKind, jokerHand("AAJKQ").Rank())
}

func TestHand_JokerCantBeUsedTwiceInTwoPair(t *testing.T) {
	assert.Equal(t, OnePair, jokerHand("AJKQ9").Rank())
}

func TestHand_JWeaknessWithJokers(t *testing.T) {
	assert.True(t,
		jokerHand("JKKK2").Value < jokerHand("QQQQ2").Value)
	assert.True(t,
		jokerHand("JKKK2").Value < jokerHand("22223").Value)
}

func TestHand_AllJokers(t *testing.T) {
	assert.Equal(t, FiveofaKind, jokerHand("JJJJJ").Rank())
}

func TestHand_OnePairWithJoker(t *testing.T) {
	assert.Equal(t, OnePair, jokerHand("J2839").Rank())
}

func TestHand_ThreeOfAKind_WithTwoJokers(t *testing.T) {
	hand := jokerHand("7JJT3")
	assert.Equal(t, ThreeofaKind, hand.Rank())
}

func jokerHand(cards string) *Hand {
	h := NewHand([]byte(cards), 0, true)
	return &h
}
