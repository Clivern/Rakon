// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Execute executes a command
func Execute(command []string) (string, error) {
	out, err := exec.Command(command[0], command[1:]...).Output()

	if err != nil {
		return "", fmt.Errorf(
			"Error while executing command %s: %s",
			strings.Join(command, " "),
			err.Error(),
		)
	}

	log.Debug(fmt.Sprintf(
		"Execute command %s",
		strings.Join(command, " "),
	))

	return string(out[:]), nil
}
