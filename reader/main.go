/*
Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
*/
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(buff []byte) (int, error) {
	n := 0
	for i := range buff {
		buff[i] = 'A'
		n++
	}
	return n, nil
}

func main() {
	reader.Validate(MyReader{})
}
