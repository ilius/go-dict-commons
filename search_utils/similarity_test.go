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
	test("", "hello", 0)
	test("hello", "", 0)
	test("hello", "hello", 200)
	test("ab", "aa", 0)
	test("ab", "ba", 0)
	test("ab", "aaa", 0)
	test("bbb", "a", 0)
	test("kitten", "sitting", 0)
	test("distance", "difference", 0)
	test("levenshtein", "frankenstein", 0)
	test("resume and cafe", "resumes and cafes", 176)
	test("a very long string that is meant to exceed", "another very long string that is meant to exceed", 175)
	test("signature", "sigature", 177)
	test("resumé and café", "resumés and cafés", 176)
	test("resume and cafe", "resumé and café", 173)
	test("Hafþór Júlíus Björnsson", "Hafþor Julius Bjornsson", 165)

	// Only 2 characters are less in the 2nd string
	test("།་གམ་འས་པ་་མ།", "།་གམའས་པ་་མ", 177)

	test("gold leaf", "gold lef", 177)
	// test("", "", 0)
}
