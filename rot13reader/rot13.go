/* Taken from the Go Tour (Exercise: rot13Reader)
 *
 * A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.
 *
 * For example, the gzip.NewReader function takes an io.Reader (a stream of compressed data) and
 * returns a *gzip.Reader that also implements io.Reader (a stream of the decompressed data).
 *
 * Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the
 * stream by applying the rot13 substitution cipher to all alphabetical characters.
 *
 * The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method.
 *
 */

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rotr rot13Reader) Read(b []byte) (int, error) {
	n, e := rotr.r.Read(b)
	for i := 0; i < n; i++ {
		switch {
		case 'a' <= b[i] && b[i] <= 'z':
			b[i] = 'a'+(b[i]-'a'+13)%26
		case 'A' <= b[i] && b[i] <= 'Z':
			b[i] = 'A'+(b[i]-'A'+13)%26
		}
	}
	return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
