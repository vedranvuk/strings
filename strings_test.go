// Copyright 2013 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strings

import (
	//"fmt"
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

	if RightByRune(teststring, 6) != "string" {
		t.Error("RightByRune() failed.")
	}

	if MidByRune(teststring, 2, 4) != "stst" {
		t.Error("MidByRune() failed.")
	}

	if Len("ÖʘΏѠ") != 4 {
		t.Error("Len() failed.")
	}
}

func TestIndexes(t *testing.T) {
	a := Indexes("a11b11c11d11e11", "11")
	if len(a) != 5 {
		t.Error("Indexes() failed.")
	}
	if a[0] != 1 {
		t.Error("Indexes() failed.")
	}
	if a[1] != 4 {
		t.Error("Indexes() failed.")
	}
	if a[2] != 7 {
		t.Error("Indexes() failed.")
	}
	if a[3] != 10 {
		t.Error("Indexes() failed.")
	}
	if a[4] != 13 {
		t.Error("Indexes() failed.")
	}

	a = Indexes("abcdefghijk", "wat?")
	if len(a) > 0 {
		t.Error("Indexes() failed.")
	}
}

func TestContainsOther(t *testing.T) {
	if !ContainsOther("abcd1", "abcd") {
		t.Error("ContainsOther() failed.")
	}
	if ContainsOther("abcd", "abcd1") {
		t.Error("ContainsOther() failed.")
	}
}

func TestFetch(t *testing.T) {
	v := "SomethingToSplit"
	if FetchLeft(v, "To") != "Something" {
		t.Error("FetchLeft() failed.")
	}
	if FetchRight(v, "To") != "Split" {
		t.Error("FetchRight() failed.")
	}
	if FetchLeftFold(v, "to") != "Something" {
		t.Error("FetchLeftFold() failed.")
	}
	if FetchRightFold(v, "to") != "Split" {
		t.Error("FetchRightFold() failed.")
	}

}

func init() {
	rand.Seed(time.Now().Unix())
}
