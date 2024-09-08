package search_utils

import (
	"testing"
)

func TestSimilarity1(t *testing.T) {
	buff := make([]uint16, 100)
	test := func(s1 string, s2 string, score uint8) {
		actualScore := Similarity([]rune(s1), []rune(s2), buff, 0)
		if actualScore != score {
			t.Fatalf(
				"s1=%#v, s2=%#v, score=%v, actualScore=%v",
				s1, s2, score, actualScore,
			)
		}
	}
	test("signature", "signature", 200)
	test("signature", "sigature", 177)
	test("signature", "signatory", 155)
	// test("", "", )
}
