/*
A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.

For example, the gzip.NewReader function takes an io.Reader (a stream of compressed data) and
returns a *gzip.Reader that also implements io.Reader (a stream of the decompressed data).

Implement a rot13Reader that implements io.Reader and reads from an io.Reader,
modifying the stream by applying the rot13(https://en.wikipedia.org/wiki/ROT13) substitution
cipher to all alphabetical characters.

The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method.
*/
package main

import (
	"fmt"
	"errors"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader *rot13Reader) Read(out []byte) (int, error) {
	i_out := 0
	buffer := make([]byte, 8)
	for {
		n, err := reader.r.Read(buffer);
		if err == io.EOF {
			break
		}

		for i := 0; i < n; i++ {
			if len(out) < i_out + 1 {
				return i_out, errors.New(fmt.Sprintf(
					"output buffer too small: %d < %d", len(out), i_out + 1))
			}
			b := buffer[i]
			if b >= 'a' && b <= 'z' {
				out[i_out] = ((b - 'a') + 13) % 26 + 'a'
			} else if b >= 'A' && b <= 'Z' {
				out[i_out] = ((b - 'A') + 13) % 26 + 'A'
			} else {
				out[i_out] = b
			}
			i_out++
		}
	}
	return i_out, io.EOF  // needs to EOF to signal the end
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
