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
func ScoreFuzzy(
	terms []string,
	args *ScoreFuzzyArgs,
	buff []uint16,
) uint8 {
	bestScore := uint8(0)
	for termI, termOrig := range terms {
		subtract := uint8(3)
		if termI < 3 {
			subtract = uint8(termI)
		}
		term := strings.ToLower(termOrig)
		if term == args.Query {
			return 200 - subtract
		}
		words := strings.Split(term, " ")
		if len(words) < args.MinWordCount {
			continue
		}
		score := Similarity(args.QueryRunes, []rune(term), buff, subtract)
		if score > bestScore {
			bestScore = score
			if score >= 180 {
				continue
			}
		}
		if len(words) > 1 {
			bestWordScore := uint8(0)
			for wordI, word := range words {
				wordScore := Similarity(args.QueryMainWord, []rune(word), buff, subtract)
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
				continue
			}
			if args.QueryWordCount > 1 {
				bestWordScore = bestWordScore>>1 + bestWordScore/7
			}
			if bestWordScore > bestScore {
				bestScore = bestWordScore
			}
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
