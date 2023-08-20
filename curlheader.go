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

	re := regexp.MustCompile(`-H\s+'([^:]+):\s+([^']+)'`)
	matched := re.FindAllStringSubmatch(str, -1)

	var currentHeaders http.Header = http.Header{}

	for _, header := range matched {
		currentHeaders.Add(header[1], header[2])
	}

	return currentHeaders, nil
}
