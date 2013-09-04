// Copyright 2013 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strings

import (
	"strings"
)

// Compare strings a and b, return -1 if a is lower, 1 if greater, 0 if equal.
// Case sensitive.
func Compare(a, b string) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

// Compare strings a and b, return -1 if a is lower, 1 if greater, 0 if equal.
// Case in-sensitive.
func CompareFold(a, b string) int {
	if strings.ToUpper(a) < strings.ToUpper(b) {
		return -1
	}
	if strings.ToUpper(a) > strings.ToUpper(b) {
		return 1
	}
	return 0
}

// HasPrefixFold tests whether the string s begins with prefix without case sensitivity.
func HasPrefixFold(s, prefix string) bool {
	return strings.HasPrefix(strings.ToLower(s), strings.ToLower(prefix))
}

// HasSuffixFold tests whether the string s ends with suffix without case sensitivity.
func HasSuffixFold(s, suffix string) bool {
	return strings.HasSuffix(strings.ToLower(s), strings.ToLower(suffix))
}

// Checks if "s" consists exclusively of numeric characters.
func IsNumsOnly(s string) bool {
	for _, c := range s {
		if !strings.Contains(Nums, string(c)) {
			return false
		}
	}
	return true
}

// Checks if "s" consists exclusively of lowercase alpha characters.
func IsAlphaLowerOnly(s string) bool {
	for _, c := range s {
		if !strings.Contains(AlphaLower, string(c)) {
			return false
		}
	}
	return true
}

// Checks if "s" consists exclusively of uppercase alpha characters.
func IsAlphaUpperOnly(s string) bool {
	for _, c := range s {
		if !strings.Contains(AlphaUpper, string(c)) {
			return false
		}
	}
	return true
}

// Checks if "s" consists exclusively of alpha characters.
func IsAlphaOnly(s string) bool {
	for _, c := range s {
		if !strings.Contains(Alpha, string(c)) {
			return false
		}
	}
	return true
}

// Checks if "s" consists exclusively of alphanumeric characters.
func IsAlphaNumsOnly(s string) bool {
	for _, c := range s {
		if !strings.Contains(AlphaNums, string(c)) {
			return false
		}
	}
	return true
}

// primeRK is the prime base used in Rabin-Karp algorithm.
const primeRK = 16777619

// hashstr returns the hash and the appropriate multiplicative
// factor for use in Rabin-Karp algorithm.
func hashstr(sep string) (uint32, uint32) {
	hash := uint32(0)
	for i := 0; i < len(sep); i++ {
		hash = hash*primeRK + uint32(sep[i])

	}
	var pow, sq uint32 = 1, primeRK
	for i := len(sep); i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return hash, pow
}

// Indexes returns a slice of indexes of sep in s, or an empty slice if none
// are present in s.
func Indexes(s, sep string) (r []int) {
	n := len(sep)
	switch {
	case n == 0:
		return
	case n == 1:
		c := sep[0]
		// special case worth making fast
		for i := 0; i < len(s); i++ {
			if s[i] == c {
				r = append(r, i)
			}
		}
		return
	case n == len(s):
		if sep == s {
			r = append(r, 0)
			return
		}
	case n > len(s):
		return
	}
	// Hash sep.
	hashsep, pow := hashstr(sep)
	var h uint32
	for i := 0; i < n; i++ {
		h = h*primeRK + uint32(s[i])
	}
	if h == hashsep && s[:n] == sep {
		r = append(r, 0)
		return
	}
	for i := n; i < len(s); {
		h *= primeRK
		h += uint32(s[i])
		h -= pow * uint32(s[i-n])
		i++
		if h == hashsep && s[i-n:i] == sep {
			r = append(r, i-n)
		}
	}
	return r
}
