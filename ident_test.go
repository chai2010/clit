// Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package clit

import (
	"testing"
)

func TestIdent(t *testing.T) {
	for _, tt := range testsIdent {
		got := Ident(tt.Pkg, tt.ScopeId...)
		if got != tt.Expected {
			t.Fatalf(
				"expected = %q, got = %q; pkg = %q, scopeId = %q",
				tt.Expected, got, tt.Pkg, tt.ScopeId,
			)
		}
	}
}

var testsIdent = []struct {
	Pkg      string
	ScopeId  []string
	Expected string
}{
	// 普通情况
	{
		Pkg:      "main",
		ScopeId:  []string{"main"},
		Expected: "$main_$main",
	},
	{
		Pkg:      "github.com/chai2010/pbgo",
		ScopeId:  []string{"ProtoFiles"},
		Expected: "$github_com_chai2010_pbgo_$ProtoFiles",
	},
	{
		Pkg:      "path/to/pkg",
		ScopeId:  []string{"main", "count"},
		Expected: "$path_to_pkg_$main_$count",
	},

	// Unicode字符
	{
		Pkg:      "主包",
		ScopeId:  []string{"世界"},
		Expected: "$主包_$世界",
	},

	// 非字母字符
	{
		Pkg:      "主包?",
		ScopeId:  []string{"世界()"},
		Expected: "$主包__$世界__",
	},
}
