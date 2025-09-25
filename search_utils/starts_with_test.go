package search_utils

import "testing"

func TestScoreStartsWith(t *testing.T) {
	test := func(terms []string, query string, score uint8) {
		actualScore := ScoreStartsWith(terms, query)
		if actualScore != score {
			t.Errorf("terms=%#v, query=%#v, score=%v, actualScore=%v", terms, query, score, actualScore)
			t.Fail()
		}
	}
	test([]string{"test"}, "test", 200)
	test([]string{"test"}, "test 123", 0)
	test([]string{"test"}, "112 test", 0)

	test([]string{"Test"}, "test", 200)
	test([]string{"TEST"}, "test 123", 0)
	test([]string{"tESt"}, "112 test", 0)
	test([]string{"test"}, "TESt", 0)
	test([]string{"test"}, "Test 123", 0)
	test([]string{"test"}, "112 tesT", 0)

	test([]string{"test word"}, "test word", 200)
	test([]string{"test word"}, "test", 195)
	test([]string{"test word"}, "word", 0)
	test([]string{"test word"}, "test 123", 0)

	test([]string{"x", "test word"}, "test word", 199)
	test([]string{"x", "test word"}, "test", 194)
	test([]string{"x", "y", "test word"}, "test", 193)

	test([]string{"x", "test word", "testing"}, "test word", 199)
	test([]string{"x", "test word", "testing"}, "test", 195)
	test([]string{"x", "y", "test word", "testing"}, "test", 194)
}
