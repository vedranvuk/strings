// Copyright 2013 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strings

import (
	"strings"
	"testing"
)

func TestRuneFunctions(t *testing.T) {
	in := "teststring"

	if LeftByRune(in, 4) != "test" {
		t.Error("LeftByRune() failed.")
	}

	if RightByRune(in, 6) != "string" {
		t.Error("RightByRune() failed.")
	}

	if MidByRune(in, 2, 4) != "stst" {
		t.Error("MidByRune() failed.")
	}

	if !ContainsOther("abcde", "abc") {
		t.Error("ContainsOther() failed.")
	}

	in = strings.Repeat("A", 1024)

	a := LenLimitByRune(in, 500)
	if len(a) != 500 {
		t.Error("LenLimitByRune() failed.")
	}

	b := LenSplitByRune(in, 384)
	if len(b) != 3 {
		t.Error("LenSplitByRune() failed.")
	} else {
		if len(b[0]) != 384 || len(b[1]) != 384 || len(b[2]) != 256 {
			t.Error("LenSplitByRune() failed.")
		}
	}

}

func TestStringFunctions(t *testing.T) {
	in := "teststring"

	if FetchLeft(in, "str") != "test" {
		t.Error("FetchLeft() failed.")
	}

	if FetchRight(in, "est") != "string" {
		t.Error("FetchRight() failed.")
	}

	if FetchLeftFold(in, "STR") != "test" {
		t.Error("FetchLeftFold() failed.")
	}

	if FetchRightFold(in, "EST") != "string" {
		t.Error("FetchRightFold() failed.")
	}

	if !HasPrefixFold(in, "TEST") {
		t.Error("HasPrefixFold() failed.")
	}

	if !HasSuffixFold(in, "STRING") {
		t.Error("HasSuffixFold() failed.")
	}

	i := Indexes("a.b.c.d.e", ".")
	if len(i) != 4 {
		t.Error("Indexes() failed.")
	}
	if i[0] != 1 && i[1] != 3 && i[2] != 5 && i[3] != 7 {
		t.Error("Indexes() failed.")
	}

	j := IndexesFold("1a2a3a4a5", "A")
	if len(j) != 4 {
		t.Error("Indexes() failed.")
	}
	if j[0] != 1 && j[1] != 3 && j[2] != 5 && j[3] != 7 {
		t.Error("Indexes() failed.")
	}

	if !MatchesWildcard("Dickson", "?ic*n") {
		t.Error("MatchesWildcard() failed")
	}
	if !MatchesWildcard("Sensible", "se?s*le") {
		t.Error("MatchesWildcard() failed")
	}
	if MatchesWildcard("ThisIsWrong", "?hi*IswrongO") {
		t.Error("MatchesWildcard() failed")
	}
}

func TestAsciiFunctions(t *testing.T) {
	if !IsNumsOnly("12345") {
		t.Error("IsNumsOnly() failed.")
	}
	if !IsAlphaLowerOnly("abcde") {
		t.Error("IsAlphaLowerOnly() failed.")
	}
	if !IsAlphaUpperOnly("ABCDE") {
		t.Error("IsAlphaUpperOnly() failed.")
	}
	if !IsAlphaOnly("abcdeABCDE") {
		t.Error("IsAlphaOnly() failed.")
	}
	if !IsAlphaNumsOnly("12345abcdeABCDE") {
		t.Error("IsAlphaNumsOnly() failed.")
	}
}
