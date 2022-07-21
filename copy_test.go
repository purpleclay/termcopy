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
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSupported(t *testing.T) {
	setEnv(t, termEnv, "xterm-256color")

	assert.True(t, Supported())
}

func setEnv(t *testing.T, key, value string) {
	t.Helper()
	os.Setenv(key, value)

	t.Cleanup(func() {
		os.Unsetenv(key)
	})
}

func TestCopy(t *testing.T) {
	out = tempFile(t)

	Copy(strings.NewReader("this is a test"))

	cb, err := ioutil.ReadFile(out.Name())
	require.NoError(t, err)

	assert.Equal(t, "\033]52;c;dGhpcyBpcyBhIHRlc3Q=\a", string(cb))
}

func TestCopyBytes(t *testing.T) {
	out = tempFile(t)

	CopyBytes([]byte("this is a test"))

	cb, err := ioutil.ReadFile(out.Name())
	require.NoError(t, err)

	assert.Equal(t, "\033]52;c;dGhpcyBpcyBhIHRlc3Q=\a", string(cb))
}

func TestCopyString(t *testing.T) {
	out = tempFile(t)

	CopyString("this is a test")

	cb, err := ioutil.ReadFile(out.Name())
	require.NoError(t, err)

	assert.Equal(t, "\033]52;c;dGhpcyBpcyBhIHRlc3Q=\a", string(cb))
}

func tempFile(t *testing.T) *os.File {
	t.Helper()

	f, err := ioutil.TempFile(t.TempDir(), "test")
	require.NoError(t, err)

	return f
}
