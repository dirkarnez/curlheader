package curlheader

func GetCurlHeader() {
ex, err := os.Executable()
	checkErr(err)
	exPath := filepath.Dir(ex)

dat, err := os.ReadFile(filepath.Join(exPath, "seed.txt"))
	checkErr(err)
	str := string(dat)

	re := regexp.MustCompile(`-H\s+'([^:]+):\s+([^']+)'`)
	matched := re.FindAllStringSubmatch(str, -1)


	var currentHeaders map[string][]string = map[string][]string{}

	for _, header := range matched {
		currentHeaders[header[1]] = []string{header[2]}
	}
}
