package utils

import (
	"net"
	"strings"
)

func IsEmailValid(email string) bool {
	if len(email) < 3 && len(email) > 254 {
		return false
	}
	if !emailRegex.MatchString(email) {
		return false
	}
	parts := strings.Split(email, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return false
	}
	return true
}

func IsURLValid(link string) (string, error) {
	if !strings.HasPrefix(link, "https://www.avito.ru") && !strings.HasPrefix(link, "https://m.avito.ru") {
		return "", ErrIncorrect
	}
	adNumber, err := getAdNumber(link)
	if err != nil {
		return "", err
	}
	if !rxURL.MatchString(link) {
		return "", ErrIncorrect
	}
	return adNumber, nil
}

func getAdNumber(link string) (string, error) {
	i := len(link) - 1
	for ; i > 0; i-- {
		if link[i] >= '0' && link[i] <= '9' {
			continue
		}
		if link[i] != '_' {
			return "", ErrIncorrect
		}
		break
	}
	return link[i+1:], nil
}
