package utils

import (
	"net"
)

func HostExists(hostname string) bool {
	_, err := net.LookupHost(hostname)
	return err == nil
}
