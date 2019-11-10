<<<<<<< HEAD
// Copyright 2013 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package strings adds additional string utility functions.
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
	if strings.ToLower(a) < strings.ToLower(b) {
		return -1
	}
	if strings.ToLower(a) > strings.ToLower(b) {
		return 1
	}
	return 0
}

// Return everything from left of "s" up to "sep".
func FetchLeft(s, sep string) string {
	v := strings.Split(s, sep)
	if len(v) > 1 {
		return v[0]
	}
	return ""
}

// Return everything from left of "s" up to "sep". Case-insensitive.
func FetchLeftFold(s, sep string) string {
	i := strings.Index(strings.ToLower(s), strings.ToLower(sep))
	return s[0:i]
}

// Return everything from "sep" up to end of "s".
func FetchRight(s, sep string) string {
	v := strings.Split(s, sep)
	if len(v) > 1 {
		return v[1]
	}
	return ""
}

// Return everything from "sep" up to end of "s". Case-insensitive.
func FetchRightFold(s, sep string) string {
	i := strings.Index(strings.ToLower(s), strings.ToLower(sep))
	return s[i+len(sep) : len(s)]
}

// HasPrefixFold tests whether the string "s" begins with "prefix".
// Case-insensitive.
func HasPrefixFold(s, prefix string) bool {
	return strings.HasPrefix(strings.ToLower(s), strings.ToLower(prefix))
}

// HasSuffixFold tests whether the string "s" ends with "suffix"
// Case-insensitive.
func HasSuffixFold(s, suffix string) bool {
	return strings.HasSuffix(strings.ToLower(s), strings.ToLower(suffix))
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

// Indexes returns a slice of all indexes of "sep" starting byte positions in
// "s", or an empty slice if none are present in "s".
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

// Indexes returns a slice of all indexes of "sep" starting byte positions in
// "s", or an empty slice if none are present in "s". Case-insensitive.
func IndexesFold(s, sep string) []int {
	return Indexes(strings.ToLower(s), strings.ToLower(sep))
}

// Matches "text" against "wildcard". Case insensitive. * and ? supported.
func MatchesWildcard(text, wildcard string) bool {
	s := RunesAsStrings(text)
	if len(s) < 1 {
		return false
	}

	w := RunesAsStrings(wildcard)
	if len(w) < 1 {
		return false
	}

	is := 0
	iw := 0
	sw := 0
	ss := -1

	for is < len(s) && w[iw] != "*" {
		if w[iw] != "?" && !strings.EqualFold(s[is], w[iw]) {
			return false
		}
		is++
		iw++
	}

	for is < len(s) {
		if w[iw] == "*" {
			iw++
			if iw >= len(w) {
				return true
			}
			sw = iw
			ss = is
		} else {
			if strings.EqualFold(s[is], w[iw]) || w[iw] == "?" {
				is++
				iw++
			} else {
				is = ss
				ss++
				iw = sw
			}
		}
	}

	for iw < len(w) && w[iw] == "*" {
		iw++
	}

	return iw == len(w)
}
=======
// Copyright 2013 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package strings adds additional string utility functions.
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
	if strings.ToLower(a) < strings.ToLower(b) {
		return -1
	}
	if strings.ToLower(a) > strings.ToLower(b) {
		return 1
	}
	return 0
}

// Return everything from left of "s" up to "sep".
func FetchLeft(s, sep string) string {
	v := strings.Split(s, sep)
	if len(v) > 1 {
		return v[0]
	}
	return ""
}

// Return everything from left of "s" up to "sep". Case-insensitive.
func FetchLeftFold(s, sep string) string {
	i := strings.Index(strings.ToLower(s), strings.ToLower(sep))
	return s[0:i]
}

// Return everything from "sep" up to end of "s".
func FetchRight(s, sep string) string {
	v := strings.Split(s, sep)
	if len(v) > 1 {
		return v[1]
	}
	return ""
}

// Return everything from "sep" up to end of "s". Case-insensitive.
func FetchRightFold(s, sep string) string {
	i := strings.Index(strings.ToLower(s), strings.ToLower(sep))
	return s[i+len(sep) : len(s)]
}

// HasPrefixFold tests whether the string "s" begins with "prefix".
// Case-insensitive.
func HasPrefixFold(s, prefix string) bool {
	return strings.HasPrefix(strings.ToLower(s), strings.ToLower(prefix))
}

// HasSuffixFold tests whether the string "s" ends with "suffix"
// Case-insensitive.
func HasSuffixFold(s, suffix string) bool {
	return strings.HasSuffix(strings.ToLower(s), strings.ToLower(suffix))
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

// Indexes returns a slice of all indexes of "sep" starting byte positions in
// "s", or an empty slice if none are present in "s".
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

// Indexes returns a slice of all indexes of "sep" starting byte positions in
// "s", or an empty slice if none are present in "s". Case-insensitive.
func IndexesFold(s, sep string) []int {
	return Indexes(strings.ToLower(s), strings.ToLower(sep))
}

// Matches "text" against "pattern". Case insensitive. * and ? supported.
func MatchesWildcard(text, pattern string) bool {
	if text == "" || pattern == "" {
		return false
	}

	t, w := []rune{}, []rune{}
	for _, v := range text {
		t = append(t, v)
	}
	for _, v := range pattern {
		w = append(w, v)
	}

	it := 0
	iw := 0
	for it < len(t) && iw < len(w) {
		if w[iw] == '*' {
			break
		}
		if w[iw] != '?' && !strings.EqualFold(string(t[it]), string(w[iw])) {
			return false
		}
		it++
		iw++
	}

	sw := 0
	st := -1
	for it < len(t) && iw < len(w) {
		if w[iw] == '*' {
			iw++
			if iw >= len(w) {
				return true
			}
			sw = iw
			st = it
		} else {
			if w[iw] == '?' || strings.EqualFold(string(t[it]), string(w[iw])) {
				it++
				iw++
			} else {
				it = st
				st++
				iw = sw
			}
		}
	}

	for iw < len(w) && w[iw] == '*' {
		iw++
	}

	return iw == len(w)
}
>>>>>>> a07ceb1c23755f4a2207feb8cb265376cb7dfa6b
