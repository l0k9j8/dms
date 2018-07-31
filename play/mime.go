// +build ignore

package main

import (
	"flag"
	"fmt"

	"../../dms"
)

func main() {
	flag.Parse()
	for _, arg := range flag.Args() {
		fmt.Println(dms.MimeTypeByPath(arg))
	}
}
