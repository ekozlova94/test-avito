package utils

import (
	"fmt"
	"regexp"
)

const (
	IP           string = `(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`
	URLSchema    string = `((ftp|tcp|udp|wss?|https?):\/\/)`
	URLUsername  string = `(\S+(:\S*)?@)`
	URLPath      string = `((\/|\?|#)[^\s]*)`
	URLPort      string = `(:(\d{1,5}))`
	URLIP        string = `([1-9]\d?|1\d\d|2[01]\d|22[0-3]|24\d|25[0-5])(\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-5]))`
	URLSubdomain string = `((www\.)|([a-zA-Z0-9]+([-_\.]?[a-zA-Z0-9])*[a-zA-Z0-9]\.[a-zA-Z0-9]+))`
	URL                 = `^` + URLSchema + `?` + URLUsername + `?` + `((` + URLIP + `|(\[` + IP + `\])|(([a-zA-Z0-9]([a-zA-Z0-9-_]+)?[a-zA-Z0-9]([-\.][a-zA-Z0-9]+)*)|(` + URLSubdomain + `?))?(([a-zA-Z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-zA-Z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-zA-Z\x{00a1}-\x{ffff}]{1,}))?))\.?` + URLPort + `?` + URLPath + `?$`
)

var (
	emailRegex = regexp.MustCompile(
		"^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+" +
			"@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?" +
			"(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$",
	)
	rxURL        = regexp.MustCompile(URL)
	ErrIncorrect = fmt.Errorf("link is incorrect")
)
