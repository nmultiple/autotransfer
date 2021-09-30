package filter

import (
	"regexp"
)

var quoteBeginRegexp = regexp.MustCompile("^On .+ wrote:$")

func filterQuote(lines []string) []string {
	var result []string
	isInQuote := false

	for _, line := range lines {
		if isInQuote {
			if len(line) > 0 && line[0] == '>' {
				continue
			}
			isInQuote = false
		}

		if quoteBeginRegexp.MatchString(line) {
			isInQuote = true
			continue
		}

		result = append(result, line)
	}

	return result
}
