// Copyright 2013 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strings

import (
	"strings"
)

// Returns number of runes in the string "s".
func RuneCount(s string) (n int) {
	for _ = range s {
		n++
	}
	return
}

// Returns a slice of runes which form string "s".
func Runes(s string) []rune {
	result := []rune{}
	for _, r := range s {
		result = append(result, r)
	}
	return result
}

// Returns "count" runes from left side of string "s" as a string.
// If "count" is more than num of chars in "s" all is returned w/o error.
func LeftByRune(s string, count int) string {
	if count <= 0 || len(s) == 0 {
		return ""
	}
	r := Runes(s)
	if count > len(r) {
		count = len(r)
	}
	return string(r[0:count])
}

// Returns "count" runes from right side of string "s" as a string.
// If "count" is more than num of chars in "s" all is returned w/o error.
func RightByRune(s string, count int) string {
	if count <= 0 || len(s) == 0 {
		return ""
	}
	r := Runes(s)
	if count > len(r) {
		count = len(r)
	}
	return string(r[len(r)-count : len(r)])
}

// Returns "count" runes from left "offset" side of string "s" as a string.
// If "count" is more than num of chars in "s" all is returned w/o error.
func MidByRune(s string, start, count int) string {
	if count <= 0 || len(s) == 0 || start < 0 {
		return ""
	}
	r := Runes(s)
	if start >= len(r) {
		return ""
	}
	return string(r[start : start+count])
}

// Checks if string "s" contains characters not in "set".
func ContainsOther(s, set string) bool {
	for _, c := range strings.Split(s, "") {
		if !strings.Contains(set, c) {
			return true
		}
	}
	return false
}
