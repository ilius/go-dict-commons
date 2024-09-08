package levenshtein

import (
	"testing"
	"unicode/utf8"
)

var buff = make([]uint16, 100)

func computeDistanceHL(a string, b string) int {
	if len(a) == 0 {
		return utf8.RuneCountInString(b)
	}

	if len(b) == 0 {
		return utf8.RuneCountInString(a)
	}

	if a == b {
		return 0
	}
	ar := []rune(a)
	br := []rune(b)
	if len(ar) > len(br) {
		ar, br = br, ar
	}
	return int(ComputeDistance(ar, br, buff))
}

func TestSanity(t *testing.T) {
	tests := []struct {
		a string
		b string
		n int
	}{
		{
			a: "",
			b: "hello",
			n: 5,
		},
		{
			a: "hello",
			b: "",
			n: 5,
		},
		{
			a: "hello",
			b: "hello",
			n: 0,
		},
		{
			a: "ab",
			b: "aa",
			n: 1,
		},
		{
			a: "ab",
			b: "ba",
			n: 2,
		},
		{
			a: "ab",
			b: "aaa",
			n: 2,
		},
		{
			a: "bbb",
			b: "a",
			n: 3,
		},
		{
			a: "kitten",
			b: "sitting",
			n: 3,
		},
		{
			a: "distance",
			b: "difference",
			n: 5,
		},
		{
			a: "levenshtein",
			b: "frankenstein",
			n: 6,
		},
		{
			a: "resume and cafe",
			b: "resumes and cafes",
			n: 2,
		},
		{
			a: "a very long string that is meant to exceed",
			b: "another very long string that is meant to exceed",
			n: 6,
		},
		{
			a: "signature",
			b: "sigature",
			n: 1,
		},
	}
	for index, test := range tests {
		actual := computeDistanceHL(test.a, test.b)
		if actual != test.n {
			t.Errorf(
				"Test[%d]: ComputeDistance(%q,%q) returned %v, want %v",
				index,
				test.a,
				test.b,
				actual,
				test.n,
			)
		}
	}
}

func TestUnicode(t *testing.T) {
	tests := []struct {
		a string
		b string
		n int
	}{
		// Testing acutes and umlauts
		{
			a: "resumé and café",
			b: "resumés and cafés",
			n: 2,
		},
		{
			a: "resume and cafe",
			b: "resumé and café",
			n: 2,
		},
		{
			a: "Hafþór Júlíus Björnsson",
			b: "Hafþor Julius Bjornsson",
			n: 4,
		},
		// Only 2 characters are less in the 2nd string
		{
			a: "།་གམ་འས་པ་་མ།",
			b: "།་གམའས་པ་་མ",
			n: 2,
		},
	}
	for index, test := range tests {
		actual := computeDistanceHL(test.a, test.b)
		if actual != test.n {
			t.Errorf(
				"Test[%d]: ComputeDistance(%q,%q) returned %v, want %v",
				index,
				test.a,
				test.b,
				actual,
				test.n,
			)
		}
	}
}

var sink int

func BenchmarkSimple(b *testing.B) {
	tests := []struct {
		a    string
		b    string
		name string
	}{
		// ASCII
		{
			a:    "levenshtein",
			b:    "frankenstein",
			name: "ASCII",
		},
		// Testing acutes and umlauts
		{
			a:    "resumé and café",
			b:    "resumés and cafés",
			name: "French",
		},
		{
			a:    "Hafþór Júlíus Björnsson",
			b:    "Hafþor Julius Bjornsson",
			name: "Nordic",
		},
		{
			a:    "a very long string that is meant to exceed",
			b:    "another very long string that is meant to exceed",
			name: "long string",
		},
		// Only 2 characters are less in the 2nd string
		{
			a:    "།་གམ་འས་པ་་མ།",
			b:    "།་གམའས་པ་་མ",
			name: "Tibetan",
		},
	}
	tmp := 0
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for range b.N {
				tmp = computeDistanceHL(test.a, test.b)
			}
		})
	}
	sink = tmp
}
