// Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package clit

import (
	"strings"
	"unicode"
)

func Ident(pkg string, scopeId ...string) string {
	var newPkgPath = "$"
	for _, c := range ([]rune)(pkg) {
		if unicode.IsLetter(c) || unicode.IsNumber(c) {
			newPkgPath += string(c)
		} else {
			newPkgPath += "_"
		}
	}

	var newIdPath []string
	for i, id := range scopeId {
		newIdPath = append(newIdPath, "$")
		for _, c := range ([]rune)(id) {
			if unicode.IsLetter(c) || unicode.IsNumber(c) {
				newIdPath[i] += string(c)
			} else {
				newIdPath[i] += "_"
			}
		}
	}

	return newPkgPath + "_" + strings.Join(newIdPath, "_")
}
