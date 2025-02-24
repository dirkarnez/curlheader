package curlheader

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
)

func GetCurlHeader(file string) (http.Header, error) {
	if len(file) < 1 {
		return nil, fmt.Errorf("no file")
	}

	dat, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	str := string(dat)

	matchedNonCookieHeader := regexp.MustCompile(`-H\s+'([^:]+):\s+([^']+)'`).FindAllStringSubmatch(str, -1)
	matchedCookieHeader := regexp.MustCompile(`-b\s+'([^']+)'`).FindAllStringSubmatch(str, -1)

	var currentHeaders http.Header = http.Header{}

	for _, header := range matchedNonCookieHeader {
		currentHeaders.Set(header[1], header[2])
	}

	for _, header := range matchedCookieHeader {
		currentHeaders.Set("cookie", header[1])
	}

	return currentHeaders, nil
}
