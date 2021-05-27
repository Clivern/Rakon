// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Verbose var
var Verbose bool

var rootCmd = &cobra.Command{
	Use: "bunglon",
	Short: `Never Remember a Linux Command Tool for SysAdmins

If you have any suggestions, bug reports, or annoyances please report
them to our issue tracker at <https://github.com/clivern/bunglon/issues>`,
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(
		&Verbose,
		"verbose",
		"v",
		false,
		"verbose output",
	)
}

// Execute runs cmd tool
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
