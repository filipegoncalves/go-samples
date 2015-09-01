
package main

import (
	"unicode"
	"fmt"
)

type path []rune

func (p path) ToUpper() {
	for i, c := range p {
		p[i] = unicode.ToUpper(c)
	}
}

func main() {
	paths := []path{path("/usr/bin/tso"), path("/home/filipe/gonçalves"),
		path("/집/사용자/test/파일"), path("/こんにちは/世界/パス"),
		path("/ェ/エ/"), path("/home/josé/niño/ßẞeta/läöür/ẞẞ/€€—")}
	for _, s := range paths {
		p := s;
		fmt.Printf("%s\n", string(p))
		p.ToUpper()
		fmt.Printf("%s\n\n", string(p))
	}
}
