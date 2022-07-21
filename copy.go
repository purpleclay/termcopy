/*
Copyright (c) 2022 Purple Clay

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package clipboard

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	termEnv = "TERM"
	xterm   = "xterm"

	// OCS52 supports copying a maximum of 100000 bytes to the system clipboard
	maxBuffer = 100000
)

//
var out = os.Stdout

// Supported ...
func Supported() bool {
	var termType string
	if termType = os.Getenv(termEnv); termType == "" {
		return false
	}

	return strings.HasPrefix(termType, xterm)
}

// Copy ...
func Copy(in io.Reader) {
	buf := bufio.NewWriterSize(out, maxBuffer)
	fmt.Fprint(buf, "\033]52;c;")

	enc := base64.NewEncoder(base64.StdEncoding, buf)
	io.Copy(enc, io.LimitReader(in, 99992))
	enc.Close()
	fmt.Fprint(buf, "\a")

	// Flush buffer to terminal
	buf.Flush()
}

// CopyBytes ...
func CopyBytes(in []byte) {
	// TODO: identify the length, and truncate to the limit as needed

	b64 := base64.StdEncoding.EncodeToString(in)

	buf := bufio.NewWriterSize(out, maxBuffer)
	fmt.Fprint(buf, "\033]52;c;")
	fmt.Fprint(buf, b64)
	fmt.Fprint(buf, "\a")

	// Flush buffer to terminal
	buf.Flush()
}

// CopyString ...
func CopyString(in string) {
	CopyBytes([]byte(in))
}
