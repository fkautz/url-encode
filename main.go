package main

import (
	"io"
	"net/url"
	"os"
)

func main() {
	// read from stdin
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	// urlencode
	encoded := url.QueryEscape(string(b))

	// write urlencoded version out
	os.Stdout.WriteString(encoded)
}
