package filter

import "regexp"

var unneededHeaderRegexp = regexp.MustCompile("^(全学)?学生団体[ 　]?各位$")

func filterHeader(lines []string) []string {
	if len(lines) < 2 {
		return lines
	}

	startIndex := 0

	for _, line := range lines {
		if len(line) == 0 {
			startIndex++
		}
		break
	}

	if unneededHeaderRegexp.MatchString(lines[startIndex]) {
		startIndex++
	}

	return lines[startIndex:]
}
