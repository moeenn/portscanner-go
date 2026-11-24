package ip

import "net"

func IsValid(s string) bool {
	return net.ParseIP(s) != nil
}
