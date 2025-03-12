// Package levenshtein is a Go implementation to calculate Levenshtein Distance.
//
// Copyright (C) 2024 Saeed Rasooli
// Copyright (c) 2015 Agniva De Sarker

// Based on https://github.com/agnivade/levenshtein/blob/master/levenshtein.go

// Implementation taken from
// https://gist.github.com/andrei-m/982927#gistcomment-1931258
package levenshtein

// ComputeDistance computes the levenshtein distance between the two
// strings passed as an argument. The return value is the levenshtein distance
//
// Works on runes (Unicode code points) but does not normalize
// the input strings. See https://blog.golang.org/normalization
// and the golang.org/x/text/unicode/norm package.
// THE CALLER MUST MAKE SURE THAT: len(s1) <= len(s2) <= 0xffff
// buff can be nil, but it's best to have an array longer than
// all given strings to avoid re-allocation of memory.
// Make sure you don't use the same buff in multiple goroutines
func ComputeDistance(s1 []rune, s2 []rune, buff []uint16) uint16 {
	lenS1 := len(s1)
	lenS2 := len(s2)

	// Init the row.
	var x []uint16
	if lenS1 >= len(buff) {
		x = make([]uint16, lenS1+1)
	} else {
		x = buff[:lenS1+1]
	}

	xn := uint16(lenS1) + 1
	for i := uint16(0); i < xn; i++ {
		x[i] = i
	}

	// make a dummy bounds check to prevent the 2 bounds check down below.
	// The one inside the loop is particularly costly.
	_ = x[lenS1]
	// fill in the rest
	lenS2_i16 := uint16(lenS2)
	prev := uint16(0)
	for i := uint16(0); i < lenS2_i16; i++ {
		prev = i + 1
		for j := 0; j < lenS1; j++ {
			if s2[i] != s1[j] {
				x[j], prev = prev, min3(x[j], prev, x[j+1])+1
			} else { // match
				x[j], prev = prev, x[j]
			}
		}
		x[lenS1] = prev
	}
	return x[lenS1]
}

func min3(a uint16, b uint16, c uint16) uint16 {
	if a > b {
		a = b
	}
	if a > c {
		return c
	}
	return a
}
