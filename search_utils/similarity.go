package search_utils

import (
	"github.com/ilius/go-dict-commons/levenshtein"
)

// Make sure you don't use the same buff in multiple goroutines
func Similarity(
	r1 []rune,
	r2 []rune,
	buff []uint16,
	subtract uint8,
) uint8 {
	if len(r1) > len(r2) {
		r1, r2 = r2, r1
	}
	// now len(r1) <= len(r2)
	n := len(r2)
	if len(r1) < (n - n/3) {
		// this optimization assumes we want to ignore below %66 similarity
		return 0
	}
	score := uint8(200 * (n - levenshtein.ComputeDistance(r1, r2, buff)) / n)
	if score <= subtract {
		return 0
	}
	return score - subtract
}
