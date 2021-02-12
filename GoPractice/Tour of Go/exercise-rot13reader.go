package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error){
	read, err := rot.r.Read(b)
	
	for i, val := range b {
		b[i] = revRot13(val)
	}
	
	return read, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

func revRot13 (c byte) byte {
	asciiVal := c - 13
	
	//space character
	if c < 65 {
		return c
	}
	
	//uppercase letter
	if 65 >= c && c <= 90 {
		if asciiVal < 65 {
			asciiVal += 26
		}
	} else { //lowercase letter
		if asciiVal < 97 {
			asciiVal += 26
		}
	}
	return asciiVal
}
