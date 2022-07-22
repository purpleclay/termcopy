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

package termcopy_test

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/purpleclay/termcopy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSupported(t *testing.T) {
	t.Run("XTERM", func(t *testing.T) {
		setEnv(t, "TERM", "xterm-256color")
		assert.True(t, termcopy.Supported())
	})

	t.Run("Screen", func(t *testing.T) {
		setEnv(t, "TERM", "screen-256color")
		assert.False(t, termcopy.Supported())
	})

	t.Run("TMUX", func(t *testing.T) {
		setEnv(t, "TERM", "xterm-256color")
		setEnv(t, "TMUX", "/private/tmp/tmux-502/default,33442,0")
		assert.False(t, termcopy.Supported())
	})

	t.Run("AppleTerminal", func(t *testing.T) {
		setEnv(t, "TERM", "xterm-256color")
		setEnv(t, "TERM_PROGRAM", "Apple_Terminal")
		assert.False(t, termcopy.Supported())
	})

	t.Run("NoTERM", func(t *testing.T) {
		setEnv(t, "TERM", "")
		assert.False(t, termcopy.Supported())
	})
}

func setEnv(t *testing.T, key, value string) {
	t.Helper()

	curEnv := os.Getenv(key)
	os.Setenv(key, value)

	t.Cleanup(func() {
		os.Setenv(key, curEnv)
	})
}

func TestStream(t *testing.T) {
	out := captureStdout(t, func() {
		termcopy.Stream(strings.NewReader("this is a test"))
	})

	assert.Equal(t, "\033]52;c;dGhpcyBpcyBhIHRlc3Q=\a", out)
}

func TestBytes(t *testing.T) {
	out := captureStdout(t, func() {
		termcopy.Bytes([]byte("this is a test"))
	})

	assert.Equal(t, "\033]52;c;dGhpcyBpcyBhIHRlc3Q=\a", out)
}

func TestString(t *testing.T) {
	out := captureStdout(t, func() {
		termcopy.String("this is a test")
	})

	assert.Equal(t, "\033]52;c;dGhpcyBpcyBhIHRlc3Q=\a", out)
}

func captureStdout(t *testing.T, capture func()) string {
	t.Helper()

	// Replace stdout with a temporary file for capturing output
	f, err := ioutil.TempFile(t.TempDir(), "test")
	require.NoError(t, err)

	stdout := os.Stdout
	os.Stdout = f

	t.Cleanup(func() {
		os.Stdout = stdout
	})

	capture()

	cb, err := ioutil.ReadFile(f.Name())
	require.NoError(t, err)

	return string(cb)
}
