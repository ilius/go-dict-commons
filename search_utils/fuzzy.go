package search_utils

import (
	"strings"
)

type ScoreFuzzyArgs struct {
	Query          string
	QueryRunes     []rune
	QueryMainWord  []rune
	QueryWordCount int // number of words in query, excluding * (indicates any word)
	MinWordCount   int
	MainWordIndex  int
}

// ScoreFuzzy ...
// Make sure you don't use the same buff in multiple goroutines
// returns a number between 50 and 200
func ScoreFuzzy(
	terms []string,
	args *ScoreFuzzyArgs,
	buff []uint16,
) uint8 {
	bestScore := uint8(0)
	queryMainWord := args.QueryMainWord
	mainWordIndex := args.MainWordIndex
	for termIndex, termOrig := range terms {
		// subtract: to give better scores for terms that come first
		// after 3rd term, they are all treated the same
		subtract := uint8(3)
		if termIndex < 3 {
			subtract = uint8(termIndex)
		}
		term := strings.ToLower(termOrig)
		if term == args.Query { // exact match, done
			return 200 - subtract
		}
		words := strings.Split(term, " ")
		if len(words) < args.MinWordCount { // term has too few words, skip
			continue
		}
		score := Similarity(args.QueryRunes, []rune(term), buff, subtract)
		if score > bestScore {
			bestScore = score
			if score >= 180 { // term is good enough without looking at its words
				continue
			}
		}
		bestWordScore := uint8(0)
		for wordI, word := range words {
			wordScore := Similarity(queryMainWord, []rune(word), buff, subtract)
			if wordScore < 50 {
				continue
			}
			if wordI == mainWordIndex {
				wordScore -= 1
				// just to differentiate between word match and whole term match
			} else {
				// matching word's position is not the same as query, %5 penalty
				wordScore -= wordScore / 20
			}
			if wordScore > bestWordScore {
				bestWordScore = wordScore
			}
		}
		if bestWordScore < 50 { // don't waste any more time on this term
			continue
		}
		if args.QueryWordCount > 1 {
			// To make sure word scores don't precede small misspells on two word query
			// for example when you type "gold lef", term "gold leaf" should come
			// before term "gold"
			bestWordScore = uint8(float64(bestWordScore)*0.87) - uint8(len(words))
		}
		if bestWordScore > bestScore {
			bestScore = bestWordScore
		}
	}
	return bestScore
}
