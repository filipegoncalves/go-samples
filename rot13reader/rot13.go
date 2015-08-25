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

/* A note about error handling here:
 *
 * According to the documentation for the Reader interface, callers of the Read() method
 * should always process the n > 0 number of bytes returned before processing a (possible)
 * error status. This is the correct behavior because if there was an error, we still want
 * to rot13 the part of the stream that was successfully read (and if there was an error before
 * anything could be read at all, n will be 0 so we won't process anything).
 *
 * Here's the relevant quote from the documentation regarding the Read() method of the 
 * io.Reader interface (see http://golang.org/pkg/io/#Reader):
 *
 * When Read encounters an error or end-of-file condition after successfully reading n > 0 bytes, it
 * returns the number of bytes read. It may return the (non-nil) error from the same call or return
 * the error (and n == 0) from a subsequent call. An instance of this general case is that a Reader
 * returning a non-zero number of bytes at the end of the input stream may return either err == EOF
 * or err == nil. The next Read should return 0, EOF.
 *
 * Callers should always process the n > 0 bytes returned before considering the error err. Doing
 * so correctly handles I/O errors that happen after reading some bytes and also both of the allowed
 * EOF behaviors.
 *
 */

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
