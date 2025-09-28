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

// stretchScore converts 100..200 range to 0..200
func stretchScore(score uint8) uint8 {
	if score < 100 {
		return 0
	}
	return score - (200 - score)
}

// ScoreFuzzy returns fuzzy score between query and term list
// Make sure you don't use the same buff in multiple goroutines
// Returns a number in 0..200 range, but 80 is a good minimum score to use
// which is equivalent of %70 similarity in single-word query and term
// Because we spread out %50 .. %100 similarity to scores of 100..200
// for example "abstracted" is %70 similar to "abstrac" query, but gets score of 80
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
			return 200 - subtract*2
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
		if bestWordScore < 100 { // don't waste any more time on this term
			continue
		}
		if args.QueryWordCount > 1 {
			// To make sure word scores don't precede small misspells on two word query
			// for example when you type "gold lef", term "gold leaf" should come
			// before term "gold"
			bestWordScore = bestWordScore - 25 - uint8(len(words))
		}
		if bestWordScore > bestScore {
			bestScore = bestWordScore
		}
	}
	if bestScore <= 100 {
		return 0
	}
	return stretchScore(bestScore)
}
