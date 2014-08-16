package main

import (
	"fmt"
	"sort"
)

// for normalised distributions

type prettymap map[rune]float64

func (p prettymap) String() string {
	s := "prettymap[\n"

	pairs := sortMapByValue(p)
	for _, pair := range pairs {
		s += fmt.Sprintf("%c", pair.Key)
		s += ": " + fmt.Sprint(pair.Value) + ", "
	}

	s += "\n]"
	return s
}

// A data structure to hold a key/value pair.
type Pair struct {
  Key rune
  Value float64
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

// A function to turn a map into a PairList, then sort and return it. 
func sortMapByValue(m map[rune]float64) PairList {
   p := make(PairList, len(m))
   i := 0
   for k, v := range m {
      p[i] = Pair{k, v}
	  i++
   }
   sort.Sort(p)
   return p
}

// for counts

type prettyintmap map[rune]int

func (p prettyintmap) String() string {
	p2 := make(prettymap)
	for k,v := range p {
		p2[k] = float64(v)
	}
	return p2.String()
}
