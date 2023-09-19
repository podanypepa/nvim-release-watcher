package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func getLocalVersion() (string, error) {
	out, err := exec.Command("nvim", "-v").Output()
	if err != nil {
		return "", err
	}

	d := strings.Split(string(out), "\n")
	if len(d) < 1 {
		return "", fmt.Errorf("bad nvim version string: %s", string(out))
	}

	return d[0], nil
}
