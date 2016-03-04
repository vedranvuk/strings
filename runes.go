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
func Runes(s string) (result []rune) {
	for _, r := range s {
		result = append(result, r)
	}
	return
}

// Explodes "s" into an array of runes in a string array.
func RunesAsStrings(s string) (result []string) {
	for _, r := range s {
		result = append(result, string(r))
	}
	return
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

// Limits "text" to "max" length which is expressed in bytes.
// Limiting is done at UTF-8 unicode code points and resulting string is ensured
// to be less than or equal to "max" bytes in length.
func LenLimitByRune(text string, max int) string {
	if len(text) > max {
		u := strings.Split(text, "")
		i := 0
		s := ""
		for i < len(u) {
			if len(s)+len(u[i]) > max {
				break
			}
			s = s + u[i]
			i++
		}
		return s
	}
	return text
}

// Split "msg" to array of string of "max" length which is expressed in bytes.
// Splitting is done at UTF-8 unicode cod points and resulting strings are ensured
// to be less than or equal to "max" bytes in length.
func LenSplitByRune(text string, max int) (out []string) {
	if len(text) > max {
		u := strings.Split(text, "")
		l := len(u)
		i := 0
		for i < l {
			s := ""
			for i < l {
				if len(s)+len(u[i]) > max {
					break
				} else {
					s = s + u[i]
					i++
				}
			}
			out = append(out, s)
		}
	} else {
		out = append(out, text)
	}
	return
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
