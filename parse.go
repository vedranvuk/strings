// Copyright 2013 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strings

import (
	"strings"
)

// Returns length of "s" by counting UTF-8 sequences.
func Len(s string) int {
	return len(strings.Split(s, ""))
}

// Returns "count" UTF-8 sequences from left side of "s".
func LeftByRune(s string, count int) string {
	l := len(s)
	if count >= l {
		return s
	}

	c := strings.Split(s, "")
	r := ""
	for i := 0; i < count; i++ {
		r = r + c[i]
	}
	return r
}

// Returns "count" UTF-8 sequences from right side of "s".
func RightByRune(s string, count int) string {
	l := len(s)
	if count >= l {
		return s
	}

	c := strings.Split(s, "")
	r := ""
	for i := l - 1 - count; i < l; i++ {
		r = r + c[i]
	}
	return r
}

// Returns "count" UTF-8 sequences from "s" with offset "start".
func MidByRune(s string, start, count int) string {
	str := RightByRune(s, len(s)-start-1)
	return LeftByRune(str, count)
}
