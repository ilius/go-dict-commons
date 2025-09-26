package search_utils

import "strings"

type ScoreFuzzyArgs struct {
	Query          string
	QueryRunes     []rune
	QueryMainWord  []rune
	QueryWordCount int
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
		if len(words) == 1 { // no need to look at the words
			continue
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
				// matching word's position is not the same as query, %10 penalty
				wordScore -= wordScore / 10
			}
			if wordScore > bestWordScore {
				bestWordScore = wordScore
			}
		}
		if bestWordScore < 50 { // don't waste any more time on this term
			continue
		}
		if args.QueryWordCount > 1 {
			// Fastest way to make an int become about 2/3 of its current value
			// in this case about %64 because 1/2 + 1/7 ~= 0.642
			// Tip: Python prioritizes + over >> so you need to add ()
			bestWordScore = bestWordScore>>1 + bestWordScore/7
		}
		if bestWordScore > bestScore {
			bestScore = bestWordScore
		}
	}
	return bestScore
}

// ScoreFuzzySingle ...
// Make sure you don't use the same buff in multiple goroutines
func ScoreFuzzySingle(
	termOrig string,
	args *ScoreFuzzyArgs,
	buff []uint16,
) uint8 {
	term := strings.ToLower(termOrig)
	words := strings.Split(term, " ")
	if len(words) < args.MinWordCount {
		return 0
	}
	score := Similarity(args.QueryRunes, []rune(term), buff, 0)
	if score >= 180 {
		return score
	}
	if len(words) < 1 {
		return score
	}
	bestWordScore := uint8(0)
	for wordI, word := range words {
		wordScore := Similarity(args.QueryMainWord, []rune(word), buff, 0)
		if wordScore < 50 {
			continue
		}
		if wordI == args.MainWordIndex {
			wordScore -= 1
		} else {
			wordScore -= wordScore / 10
		}
		if wordScore > bestWordScore {
			bestWordScore = wordScore
		}
	}
	if bestWordScore < 50 {
		return 0
	}
	if args.QueryWordCount > 1 {
		bestWordScore = bestWordScore>>1 + bestWordScore/7
	}
	return bestWordScore
}
