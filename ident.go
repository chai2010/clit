// Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.10

package clit

import (
	"strings"
	"unicode"
)

func Ident(pkg string, scopeId ...string) string {
	var sb strings.Builder

	sb.WriteRune('$')
	for _, c := range pkg {
		if unicode.IsLetter(c) || unicode.IsNumber(c) {
			sb.WriteRune(c)
		} else {
			sb.WriteRune('_')
		}
	}

	for _, id := range scopeId {
		sb.WriteRune('_')
		sb.WriteRune('$')
		for _, c := range id {
			if unicode.IsLetter(c) || unicode.IsNumber(c) {
				sb.WriteRune(c)
			} else {
				sb.WriteRune('_')
			}
		}
	}

	return sb.String()
}
