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
