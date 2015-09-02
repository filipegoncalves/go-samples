/* Code based on the post "Go Slices: usage and internals", from the
 * Go blog: http://blog.golang.org/go-slices-usage-and-internals
 *
 * [...]
 * To fix this problem one can copy the interesting data to a new slice before returning it:
 *
 * func CopyDigits(filename string) []byte {
 *     b, _ := ioutil.ReadFile(filename)
 *     b = digitRegexp.Find(b)
 *     c := make([]byte, len(b))
 *     copy(c, b)
 *     return c
 * }
 *
 * A more concise version of this function could be constructed by using append. This is left as an
 * exercise for the reader.
 *
 */

package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

var digitRegexp = regexp.MustCompile("[0-9]+")

/* Note: this version not only makes it smaller (it uses return append([]byte{}, b...) instead of
 *       make+copy), but it also deals with error handling, which is important, but often ignored.
 */

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
