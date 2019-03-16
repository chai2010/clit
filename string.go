// Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package clit

import (
	"strconv"
	"strings"
	"unicode"
)

func String(s string) string {
	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		c := s[i]

		if c > 127 {
			sb.WriteString(`\x`)
			sb.WriteString(strconv.FormatInt(int64(c), 16))
			continue
		}

		switch c {
		case 0:
			sb.WriteString(`\0`)
		case '?':
			sb.WriteString(`\?`)

		case '\a':
			sb.WriteString(`\a`)
		case '\b':
			sb.WriteString(`\b`)
		case '\f':
			sb.WriteString(`\f`)
		case '\n':
			sb.WriteString(`\n`)
		case '\r':
			sb.WriteString(`\r`)
		case '\t':
			sb.WriteString(`\t`)
		case '\v':
			sb.WriteString(`\v`)
		case '\\':
			sb.WriteString(`\\`)
		case '\'':
			sb.WriteString(`\'`)
		case '"':
			sb.WriteString(`\"`)

		default:
			switch {
			case unicode.IsPrint(rune(c)):
				sb.WriteByte(c)
			default:
				sb.WriteString(`\x`)
				sb.WriteString(strconv.FormatInt(int64(c), 16))
			}
		}
	}

	return sb.String()
}
