// Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"

	"github.com/chai2010/clit"
)

func main() {
	fmt.Println(clit.Ident("main", "main"))
	fmt.Println(clit.Ident("path/to/pkg", "main", "count"))
	fmt.Println(clit.Ident("github.com/chai2010/pbgo", "HttpRule"))
	fmt.Println(clit.Ident("主包", "主函数"))

	fmt.Println(clit.String("abc\n124你好'\""))

	// Output:
	// $main_$main
	// $path_to_pkg_$main_$count
	// $github_com_chai2010_pbgo_$HttpRule
	// $主包_$主函数
	// abc\n124\xe4\xbd\xa0\xe5\xa5\xbd\'\"
}
