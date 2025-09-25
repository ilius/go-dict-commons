package search_utils

import "testing"

func TestScoreWordMatch(t *testing.T) {
	test := func(terms []string, query string, score uint8) {
		actualScore := ScoreWordMatch(terms, query)
		if actualScore != score {
			t.Errorf("terms=%#v, query=%#v, score=%v, actualScore=%v", terms, query, score, actualScore)
			t.Fail()
		}
	}
	test([]string{"test"}, "test", 200)
	test([]string{"test"}, "test 123", 189)
	test([]string{"test"}, "112 test", 188)
	test([]string{"test"}, "test 123 x", 178)
	test([]string{"test"}, "x test abc", 177)
	test([]string{"test"}, "test 1 2 3 4 5", 145)
	test([]string{"test"}, "x", 0)
	test([]string{"test"}, "hello", 0)
	test([]string{"test"}, " ", 0)
	test([]string{"test"}, "hello world", 0)

	test([]string{"Test"}, "test", 200)
	test([]string{"TEST"}, "test 123", 189)
	test([]string{"tESt"}, "112 test", 188)

	test([]string{"test"}, "TESt", 200)
	test([]string{"test"}, "Test 123", 189)
	test([]string{"test"}, "112 tesT", 188)

	test([]string{"test word"}, "test word", 200)
	test([]string{"test word"}, "test", 199)
	test([]string{"test word"}, "word", 198)
	test([]string{"test word"}, "test 123", 188)
	test([]string{"test word"}, "word 123", 187)
	test([]string{"test word"}, "112 test", 187)
	test([]string{"test word"}, "x test abc", 176)
	test([]string{"test word"}, "test word 1 2 3 4 5", 145)
	test([]string{"test word"}, "hello", 0)
	test([]string{"test word"}, " ", 0)
	test([]string{"test word"}, "hello world", 0)

	test([]string{"x", "test word"}, "test word", 199)
	test([]string{"x", "y", "test word"}, "test word", 198)
}
