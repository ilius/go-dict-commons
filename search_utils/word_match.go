package search_utils

import (
	"strings"
)

// ScoreWordMatch returns a score based on exact word match between query words and term words
func ScoreWordMatch(terms []string, query string) uint8 {
	query = strings.ToLower(query)
	queryWords := strings.Split(query, " ")
	if len(queryWords) == 1 {
		// optmization, should not change the result
		return scoreWordMatchSingleWordQuery(terms, query)
	}
	queryWordMap := map[string]int{}
	for wordIndex, word := range queryWords {
		queryWordMap[word] = wordIndex
	}
	bestScore := 0
	for termI, termOrig := range terms {
		term := strings.ToLower(termOrig)
		termWords := strings.Split(term, " ")
		termWordMap := map[string]bool{}
		score := -termI
		for wordIndex, word := range termWords {
			termWordMap[word] = true
			queryWordIndex, ok := queryWordMap[word]
			if !ok {
				score--
				continue
			}
			if queryWordIndex > wordIndex {
				score += 10 - (queryWordIndex - wordIndex)
			} else {
				score += 10 - (wordIndex - queryWordIndex)
			}
		}
		for word := range queryWordMap {
			if !termWordMap[word] {
				score--
			}
		}
		if score < 0 {
			continue
		}
		if score > bestScore {
			bestScore = score
		}
	}
	if bestScore <= 0 {
		return 0
	}
	outputScore := 200 - 10*len(queryWordMap) + bestScore
	if outputScore > 200 {
		return 200
	}
	return uint8(outputScore)
}

// query is one word
func scoreWordMatchSingleWordQuery(terms []string, query string) uint8 {
	bestScore := 0
	for termI, termOrig := range terms {
		term := strings.ToLower(termOrig)
		score := -termI
		for wordIndex, word := range strings.Split(term, " ") {
			if word != query {
				score--
				continue
			}
			score += 10 - wordIndex
		}
		if score < 0 {
			continue
		}
		if score > bestScore {
			bestScore = score
		}
	}
	if bestScore <= 0 {
		return 0
	}
	outputScore := 190 + bestScore
	if outputScore > 200 {
		return 200
	}
	return uint8(outputScore)
}
