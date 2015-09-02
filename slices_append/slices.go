package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

var digitRegexp = regexp.MustCompile("[0-9]+")

func FindDigits(filename string) ([]byte, error) {
	b, e := ioutil.ReadFile(filename)

	if e != nil {
		return b, e
	}

	b = digitRegexp.Find(b)

	return append([]byte{}, b...), nil
}

func main() {
	d, e := FindDigits("sample.in")

	if (e != nil) {
		fmt.Printf("Error: %v\n", e)
	} else {
		fmt.Printf("Digits found: %s\n", d)
	}
}
