package main

import (
	"strings"
	"unicode"
	"math"
)


	// fmt.Println("Here's some scores for actual messages:")
	// fmt.Println(score("test and something"))
	// fmt.Println(score("Here's a message and things are great."))

	// fmt.Println("Here's some scores for random noise with well placed spaces:")
	// fmt.Println(score("adskf hgkwi wgfuh tbmeibv vriewrf iybvfj."))
	// fmt.Println(score("jicds fkcvbw qeigyw qedibf vndsv weiugrw."))

func score(msg string) float64 {
	// soft filter on nonstandard characters
	// plaintext will likely have a very small number of nonstandard characters
	lmod := 1.0
	for _, s := range msg {
		if !unicode.In(s, unicode.Letter, unicode.Punct, unicode.Space) {
			if (s < '0') || (s > '9') {
				if s != '\n' {
					lmod *= 0.1					
				}
			}
		}
	}

	if lmod < 0.01 {
		return lmod
	}

	// calculate regularity measure for character frequency
	// plaintext will likely have similar character frequency to natural language
	fmod := 1.0
	appear := make(map[rune]int)
	countAppearance(msg, appear)
	fmod = regularity(appear)

	// average word length and standard deviation
	// plaintext will likely not have really long words
	wmod := 1.0
	spaceMsg := strings.Replace(msg, "\n", " ", -1)
	spaceMsg = strings.Replace(spaceMsg, "\t", " ", -1)
	spaceMsg = strings.Trim(spaceMsg, " ")
	spaceMsg = strings.Replace(spaceMsg, "  ", " ", -1)
	words := strings.Split(spaceMsg, " ")
	total := 0
	for _, word := range words {
		total += len(word)
	}
	average := float64(total) / float64(len(words))
	wmod /= (math.Abs(average - 5.08)+1.0)

	// TODO: calculate standard deviation
	// regular sd = 0.4305961

	return lmod * fmod * wmod
}

func normalise_corpus(m map[rune]int) map[rune]float64 {
	n := make(map[rune]float64)
	n[space] = spacefreq // precalculated
	s := sum(m)
	normaliser := (1 - spacefreq) / float64(s)
	for key, val := range m {
		n[key] = float64(val) * normaliser
	}
	return n
}

func normalise(m map[rune]int) map[rune]float64 {
	n := make(map[rune]float64)
	s := sum(m)
	normaliser := 1.0 / float64(s)
	for key, val := range m {
		n[key] = float64(val) * normaliser
	}
	return n
}

func regularity(appear map[rune]int) float64 {
	// First, we normalise the 'appear' and 'regular' distributions
	// as they are currently given as number of appearances
	regularNorm := normalise_corpus(NYTcorpus)
	appearNorm := normalise(appear)

	// distance from appear to regular
	dist := d(regularNorm, appearNorm)

	// invert for similarity and normalise
	return 1.0/(dist+1.0)

	// This will yeild a value between
	// 1/(sqrt(n)+1) and 1
	// where n is the number of keys
}

func d(m1, m2 map[rune]float64) float64 {
	keys := make(map[rune]bool)
	
	for key := range m1 {
		keys[key] = true
	}
	for key := range m2 {
		keys[key] = true
	}

	squareSum := 0.0
	for key := range keys {
		squareSum += (m1[key] - m2[key])*(m1[key] - m2[key])
	}

	return math.Sqrt(squareSum)
}

func countAppearance(msg string, appear map[rune]int) {
	for _, c := range msg {
		if ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') {
			appear[c]++
		} else if c < 128 && unicode.In(c, unicode.Punct) {
			appear[punc]++
		} else if unicode.In(c, unicode.Space) || c == '\n' {
			appear[space]++
		} else if (c < '0') || (c > '9') {
			appear[digit]++
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
