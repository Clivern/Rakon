// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var distroCmd = &cobra.Command{
	Use:   "distro",
	Short: "Print the Linux Distribution",
	Run: func(cmd *cobra.Command, args []string) {
		if Verbose {
			log.SetLevel(log.DebugLevel)
		}

		log.Debug("Distribution command got called.")

		fmt.Println(
			fmt.Sprintf(
				`OS: %s`,
				strings.Title(runtime.GOOS),
			),
		)
	},
}

func init() {
	rootCmd.AddCommand(distroCmd)
}
