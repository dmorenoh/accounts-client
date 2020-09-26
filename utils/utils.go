package utils

import (
	"log"
	"os/exec"
	"strings"
)

func NewUUID() string {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return strings.ToLower(strings.TrimSuffix(string(out), "\n"))
}
