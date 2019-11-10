// Copyright 2013 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strings

import (
	"strings"
)

const (
	Nums       = "0123456789"
	AlphaUpper = "ABCDEFGHIJKLMNOPQRSTUVXYZ"
	AlphaLower = "abcdefghijklmnopqrstuvxyz"
	Alpha      = AlphaUpper + AlphaLower
	AlphaNums  = Nums + Alpha
)

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
