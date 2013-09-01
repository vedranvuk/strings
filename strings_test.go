// Copyright 2013 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strings

import (
	// "fmt"
	"math/rand"
	"testing"
	"time"
)

func TestCompare(t *testing.T) {
	if Compare("a", "b") != -1 {
		t.Error("Compare() -1 failed.")
	}
	if Compare("b", "a") != 1 {
		t.Error("Compare() 1 failed.")
	}
	if Compare("c", "c") != 0 {
		t.Error("Compare() 0 failed.")
	}

	if CompareFold("a", "B") != -1 {
		t.Error("CompareFold() -1 failed.")
	}
	if CompareFold("b", "A") != 1 {
		t.Error("CompareFold() 1 failed.")
	}
	if CompareFold("TEST", "test") != 0 {
		t.Error("CompareFold() 0 failed.")
	}
}

func TestRandoms(t *testing.T) {
	/*
		fmt.Println(RandomNum())
		fmt.Println(RandomNums(10))
		fmt.Println(RandomLower())
		fmt.Println(RandomLowers(10))
		fmt.Println(RandomUpper())
		fmt.Println(RandomUppers(10))
		fmt.Println(RandomString(true, true, true, 10))
	*/
}

func TestAlphaNumCompares(t *testing.T) {
	if !IsNumsOnly("1234") {
		t.Error("IsNumsOnly() failed.")
	}
	if IsNumsOnly("ABCD") {
		t.Error("IsNumsOnly() failed.")
	}
	if !IsAlphaLowerOnly("abcd") {
		t.Error("IsAlphaLowerOnly() failed.")
	}
	if IsAlphaLowerOnly("abCD") {
		t.Error("IsAlphaLowerOnly() failed.")
	}
	if !IsAlphaUpperOnly("ABCD") {
		t.Error("IsAlphaUpperOnly() failed.")
	}
	if IsAlphaUpperOnly("ABcd") {
		t.Error("IsAlphaUpperOnly() failed.")
	}
	if !IsAlphaOnly("ABCD") {
		t.Error("IsAlphaOnly() failed.")
	}
	if IsAlphaOnly("AB34") {
		t.Error("IsAlphaOnly() failed.")
	}
}

func TestParse(t *testing.T) {
	teststring := "teststring"

	if LeftByRune(teststring, 4) != "test" {
		t.Error("LeftByRune() failed.")
	}

	if RightByRune(teststring, 5) != "string" {
		t.Error("RightByRune() failed.")
	}

	if MidByRune(teststring, 2, 4) != "stst" {
		t.Error("MidByRune() failed.")
	}
}

func init() {
	rand.Seed(time.Now().Unix())
}
