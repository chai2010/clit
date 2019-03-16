// Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package clit

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

func String(s string) string {
	var sb strings.Builder
	for _, c := range s {
		if c > 127 {
			buf := make([]byte, 5)
			n := utf8.EncodeRune(buf, c)
			for i := 0; i < n; i++ {
				sb.WriteString(`\x`)
				sb.WriteString(strconv.FormatInt(int64(buf[i]), 16))
			}
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
			sb.WriteRune(c)
		}
	}

	return sb.String()
}
