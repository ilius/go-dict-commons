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
		{Query: "abstracted", Term: "abstracted", Score: 200},
		{Query: "abstracted", Term: "abstract", Score: 120},
		{Query: "abstracted", Term: "abstrac", Score: 80},

		{Query: "signature", Term: "signature", Score: 200},
		{Query: "signature", Term: "signatur", Score: 154},
		{Query: "signature", Term: "signature tune", Score: 142},
		{Query: "signature", Term: "key signature", Score: 123}, // 126 in v0.6.0
		{Query: "signature", Term: "signatory", Score: 110},
		{Query: "signature", Term: "signatu", Score: 110},
		{Query: "signature", Term: "signat", Score: 66},
		{Query: "signature", Term: "signa", Score: 0},
		{Query: "signature", Term: "sign", Score: 0},
		{Query: "gold leaf", Term: "gold lef", Score: 154},
		{Query: "gold lef", Term: "gold leaf", Score: 154},
		{Query: "gold lef", Term: "gold", Score: 143}, // 144 in v0.6.0
		{Query: "gold lef", Term: "gold beetle", Score: 142},
		{Query: "gold lef", Term: "green gold", Score: 123}, // 126 in v0.6.0
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
			t.Errorf(
				"TestScoreFuzzy: score=%v != %v, query=%#v, term=%#v",
				score,
				tc.Score,
				query,
				term,
			)
			t.Fail()
		}
		// t.Logf("TestScoreFuzzy: score=%v, query=%#v, term=%#v", score, query, term)
	}
}
