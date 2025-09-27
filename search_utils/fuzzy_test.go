package search_utils

import (
	"strings"
	"testing"
)

type FuzzyTestCase1 struct {
	Query         string
	Term          string
	Score         uint8
	QueryMainWord int
}

func TestScoreFuzzy1(t *testing.T) {
	for _, tc := range []FuzzyTestCase1{
		{Query: "signature", Term: "signature", Score: 200},
		{Query: "signature", Term: "signature tune", Score: 171},
		{Query: "signature", Term: "key signature", Score: 163},
		{Query: "gold lef", Term: "gold leaf", Score: 177},
		{Query: "gold lef", Term: "gold", Score: 172},
		{Query: "gold lef", Term: "gold beetle", Score: 171},
		{Query: "gold lef", Term: "green gold", Score: 163},
	} {
		query := tc.Query
		term := tc.Term
		buff := make([]uint16, 500)
		args := &ScoreFuzzyArgs{
			Query:          query,
			QueryRunes:     []rune(query),
			QueryMainWord:  []rune(strings.Split(query, " ")[tc.QueryMainWord]),
			QueryWordCount: 2,
			MinWordCount:   1,
		}

		score := ScoreFuzzy([]string{term}, args, buff)
		if score != tc.Score {
			t.Errorf("TestScoreFuzzy: score=%v, query=%#v, term=%#v", score, query, term)
			t.Fail()
		}
		// t.Logf("TestScoreFuzzy: score=%v, query=%#v, term=%#v", score, query, term)
	}
}
