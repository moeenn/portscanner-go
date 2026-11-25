package host

import (
	"errors"
	"net"
	"net/url"
	"strings"
)

func isValidIP(s string) bool {
	return net.ParseIP(s) != nil
}

func isValidUrl(s string) (string, bool) {
	if !strings.HasPrefix(s, "http") {
		s = "https://" + s
	}

	u, err := url.Parse(s)
	if err != nil {
		return "", false
	}

	host := u.Hostname()
	return host, host != ""
}

func Validate(host *string) (string, error) {
	if isValidIP(*host) {
		return *host, nil
	}

	h, ok := isValidUrl(*host)
	if !ok {
		return "", errors.New("not a valid ip address or url")
	}

	return h, nil
}
