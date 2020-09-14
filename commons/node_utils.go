package commons

import (
	"fmt"
	"os"
	"strings"
)

func GetHostname() (string, error) {

	nodeName, err := os.Hostname()
	if err != nil {
		return "", fmt.Errorf("couldn't determine hostname: %v", err)
	}

	// Trim whitespaces first to avoid getting an empty hostname
	// For linux, the hostname is read from file /proc/sys/kernel/hostname directly
	nodeName = strings.TrimSpace(nodeName)
	if len(nodeName) == 0 {
		return "", fmt.Errorf("empty hostname is invalid")
	}
	return strings.ToLower(nodeName), nil
}
