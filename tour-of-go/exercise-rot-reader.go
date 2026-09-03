package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13Byte(b[i])
	}
	return n, err
}

func rot13Byte(c byte) byte {
	switch {
	case 'a' <= c && c <= 'z':
		return 'a' + (c-'a'+13)%26
	case 'A' <= c && c <= 'Z':
		return 'A' + (c-'A'+13)%26
	default:
		return c
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
