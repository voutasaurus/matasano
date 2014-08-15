package main

import (
	"unicode"
	"strings"
)

var normLow = map[rune]float64{
	'a': 1.0, 'b': 1.0, 'c': 1.0, 'd': 1.0, 'e': 1.0, 'f': 1.0, 'g': 1.0, 'h': 1.0, 'i': 1.0, 'j': 1.0,
	'k': 1.0, 'l': 1.0, 'm': 1.0, 'n': 1.0, 'o': 1.0, 'p': 1.0, 'q': 1.0, 'r': 1.0, 's': 1.0, 't': 1.0,
	'u': 1.0, 'v': 1.0, 'w': 1.0, 'x': 1.0, 'y': 1.0, 'z': 1.0,
}

var normUp = map[rune]float64{
	'A': 1.0, 'B': 1.0, 'C': 1.0, 'D': 1.0, 'E': 1.0, 'F': 1.0, 'G': 1.0, 'H': 1.0, 'I': 1.0, 'J': 1.0,
	'K': 1.0, 'L': 1.0, 'M': 1.0, 'N': 1.0, 'O': 1.0, 'P': 1.0, 'Q': 1.0, 'R': 1.0, 'S': 1.0, 'T': 1.0,
	'U': 1.0, 'V': 1.0, 'W': 1.0, 'X': 1.0, 'Y': 1.0, 'Z': 1.0,
}

var	other = rune(0)
var	space = ' '
var	punc = ','

var normMisc = map[rune]float64{
	space: 1.0, punc: 1.0, other: 1.0,
}

var norm = map[rune]float64{}

func score(msg string) float64 {
	// soft filter on nonstandard characters
	// plaintext will likely have a very small number of nonstandard characters
	lmod := 1.0
	for _, s := range msg {
		if !unicode.In(s, unicode.Letter, unicode.Punct, unicode.Space) {
			lmod *= 0.1
		}
	}

	if lmod < 0.01 {
		return lmod		
	}

	// calculate normality measure for character frequency
	// plaintext will likely have similar character frequency to natural language
	fmod := 1.0
	appear := make(map[rune]int)
	countAppearance(msg, appear)
	fmod = normality(appear)

	// average word length and standard deviation
	// plaintext will likely not have really long words
	wmod := 1.0
	words := strings.Split(msg, " ")
	total := 0
	for _, word := range words {
		total += len(word)
	}
	average := float64(total) / float64(len(words))
	_ = average

	return lmod * fmod * wmod
}

func normality(appear map[rune]int) float64 {
	// distance from appear to norm?
	// invert for similarity?
	// normalise

	// Similarity matrix?
	return 1.0
}

func countAppearance(msg string, appear map[rune]int) {
	for _, c := range msg {
		if ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') {
			appear[c]++
		} else if c < 128 && unicode.In(c, unicode.Punct) {
			appear[punc]++
		} else if unicode.In(c, unicode.Space) {
			appear[space]++
		} else {
			appear[other]++
		}
	}
	return
}


/*
		// find deviation from normal character frequencies
		appear := make(map[rune]int)
		for _, c := range msg {
			appear[c]++
		}

		freq := make(map[rune]float64)
		msgLen := float64(len(msg))
		for c := range appear {
			freq[c] = float64(appear[c]) / msgLen
		}

		//norm[]
*/
