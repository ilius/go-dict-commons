package search_utils

import (
	"github.com/ilius/go-dict-commons/levenshtein"
)

// Similarity calculates a similarity score between 0 and 200
// Except any score between 1 and 131 (<%66) will become zero as optimization
// SimilaritySlow func does not have this limitation
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
	if len(r1) < n-n/3 {
		// this optimization assumes we want to ignore below %66 similarity
		return 0
	}
	n16 := uint16(n)
	dist := levenshtein.ComputeDistance(r1, r2, buff)
	if dist > n16/3 {
		// this optimization assumes we want to ignore below %66 similarity
		return 0
	}
	score := uint8(200 * (n16 - dist) / n16)
	if score <= subtract {
		return 0
	}
	return score - subtract
}

// SimilaritySlow calculates a similarity score between 0 and 200
// Make sure you don't use the same buff in multiple goroutines
func SimilaritySlow(
	r1 []rune,
	r2 []rune,
	buff []uint16,
	subtract uint8,
) uint8 {
	if len(r1) > len(r2) {
		r1, r2 = r2, r1
	}
	// now len(r1) <= len(r2)
	n16 := uint16(len(r2))
	score := uint8(200 * (n16 - levenshtein.ComputeDistance(r1, r2, buff)) / n16)
	if score <= subtract {
		return 0
	}
	return score - subtract
}
