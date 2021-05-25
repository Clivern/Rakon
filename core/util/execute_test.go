// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// +build unit

package util

import (
	"testing"

	"github.com/franela/goblin"
)

// TestExecute test cases
func TestExecute(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#ExecuteFunc", func() {
		g.It("It should satisfy all provided test cases", func() {
			var tests = []struct {
				input      []string
				wantResult string
			}{
				{[]string{"/bin/echo", "a"}, "a\n"},
				{[]string{"/bin/echo", "b"}, "b\n"},
				{[]string{"/bin/echo", "c"}, "c\n"},
				{[]string{"/bin/echo", "d"}, "d\n"},
				{[]string{"/bin/echo", "f"}, "f\n"},
			}

			for _, tt := range tests {
				res, err := Execute(tt.input)
				g.Assert(res).Equal(tt.wantResult)
				g.Assert(err).Equal(nil)
			}
		})
	})
}
