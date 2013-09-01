// Copyright 2013 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strings

import (
	"math/rand"
)

const Nums = "0123456789"
const AlphaUpper = "ABCDEFGHIJKLMNOPQRSTUVXYZ"
const AlphaLower = "abcdefghijklmnopqrstuvxyz"
const Alpha = AlphaUpper + AlphaLower
const AlphaNums = Nums + Alpha

// Returns a string containing a random number.
func RandomNum() string {
	return string(Nums[rand.Intn(len(Nums))])
}

// Returns a string of random numbers of length.
func RandomNums(length int) string {
	if length < 1 {
		return ""
	}
	r := ""
	for i := 0; i < length; i++ {
		r = r + RandomNum()
	}
	return r
}

// Returns a string containing a random uppercase letter.
func RandomUpper() string {
	return string(AlphaUpper[rand.Intn(len(AlphaUpper))])
}

// Returns a string of random uppercase letters of length.
func RandomUppers(length int) string {
	if length < 1 {
		return ""
	}
	r := ""
	for i := 0; i < length; i++ {
		r = r + RandomUpper()
	}
	return r
}

// Returns a string containing a random lowercase letter.
func RandomLower() string {
	return string(AlphaLower[rand.Intn(len(AlphaLower))])
}

// Returns a string of random lowercase letters of length.
func RandomLowers(length int) string {
	if length < 1 {
		return ""
	}
	r := ""
	for i := 0; i < length; i++ {
		r = r + RandomLower()
	}
	return r
}

// Returns a random string of length.
// If lo includes lowercase letters.
// If up includes uppercase letters.
// If num includes numbers.
func RandomString(lo, up, nums bool, length int) string {
	if length < 1 {
		return ""
	}
	f := []func() string{}
	if lo {
		f = append(f, RandomLower)
	}
	if up {
		f = append(f, RandomUpper)
	}
	if nums {
		f = append(f, RandomNum)
	}
	r := ""
	for i := 0; i < length; i++ {
		r = r + f[rand.Intn(len(f))]()
	}
	return r
}
