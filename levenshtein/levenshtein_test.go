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
	return int(ComputeDistance([]rune(a), []rune(b), buff))
}

func TestSanity(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want int
	}{
		{
			a:    "",
			b:    "hello",
			want: 5,
		},
		{
			a:    "hello",
			b:    "",
			want: 5,
		},
		{
			a:    "hello",
			b:    "hello",
			want: 0,
		},
		{
			a:    "ab",
			b:    "aa",
			want: 1,
		},
		{
			a:    "ab",
			b:    "ba",
			want: 2,
		},
		{
			a:    "ab",
			b:    "aaa",
			want: 2,
		},
		{
			a:    "bbb",
			b:    "a",
			want: 3,
		},
		{
			a:    "kitten",
			b:    "sitting",
			want: 3,
		},
		{
			a:    "distance",
			b:    "difference",
			want: 5,
		},
		{
			a:    "levenshtein",
			b:    "frankenstein",
			want: 6,
		},
		{
			a:    "resume and cafe",
			b:    "resumes and cafes",
			want: 2,
		},
		{
			a:    "a very long string that is meant to exceed",
			b:    "another very long string that is meant to exceed",
			want: 6,
		},
	}
	for i, d := range tests {
		n := computeDistanceHL(d.a, d.b)
		if n != d.want {
			t.Errorf("Test[%d]: ComputeDistance(%q,%q) returned %v, want %v",
				i, d.a, d.b, n, d.want)
		}
	}
}

func TestUnicode(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want int
	}{
		// Testing acutes and umlauts
		{
			a:    "resumé and café",
			b:    "resumés and cafés",
			want: 2,
		},
		{
			a:    "resume and cafe",
			b:    "resumé and café",
			want: 2,
		},
		{
			a:    "Hafþór Júlíus Björnsson",
			b:    "Hafþor Julius Bjornsson",
			want: 4,
		},
		// Only 2 characters are less in the 2nd string
		{
			a:    "།་གམ་འས་པ་་མ།",
			b:    "།་གམའས་པ་་མ",
			want: 2,
		},
	}
	for i, d := range tests {
		n := computeDistanceHL(d.a, d.b)
		if n != d.want {
			t.Errorf("Test[%d]: ComputeDistance(%q,%q) returned %v, want %v",
				i, d.a, d.b, n, d.want)
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
