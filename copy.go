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

package termcopy

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	termEnv        = "TERM"
	termProgramEnv = "TERM_PROGRAM"
	tmuxEnv        = "TMUX"
	xterm          = "xterm"
)

// Supported identifies whether the current terminal supports copying
// to the system clipboard
func Supported() bool {
	var termType string
	if termType = os.Getenv(termEnv); termType == "" {
		return false
	}

	// TMUX isn't supported
	if tmuxTerm := os.Getenv(tmuxEnv); tmuxTerm != "" {
		return false
	}

	// Terminal.app isn't supported
	if termProg := os.Getenv(termProgramEnv); termProg == "Apple_Terminal" {
		return false
	}

	return strings.HasPrefix(termType, xterm)
}

// Stream the contents of the reader to the system clipboard. If the current
// terminal is not supported the operating system code (OSC) will be
// ignored and the existing system clipboard will remain unmodified
func Stream(in io.Reader) {
	buf := bufio.NewWriter(os.Stdout)
	fmt.Fprint(buf, "\033]52;c;")

	enc := base64.NewEncoder(base64.StdEncoding, buf)
	io.Copy(enc, in)
	enc.Close()
	fmt.Fprint(buf, "\a")

	buf.Flush()
}

// Bytes copies the contents of a byte array to the system clipboard. If the
// current terminal is not supported the operating system code (OSC) will be
// ignored and the existing system clipboard will remain unmodified
func Bytes(in []byte) {
	b64 := base64.StdEncoding.EncodeToString(in)

	buf := bufio.NewWriter(os.Stdout)
	fmt.Fprint(buf, "\033]52;c;")
	fmt.Fprint(buf, b64)
	fmt.Fprint(buf, "\a")

	buf.Flush()
}

// String copies the contents of a string to the system clipboard. If the
// current terminal is not supported the operating system code (OSC) will
// be ignored and the existing system clipboard will remain unmodified
func String(in string) {
	Bytes([]byte(in))
}
