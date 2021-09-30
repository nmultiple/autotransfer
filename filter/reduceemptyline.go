package filter

func filterUselessEmptyLine(lines []string) []string {
	startIndex := 0

	for _, line := range lines {
		if len(line) == 0 {
			startIndex++
		}
		break
	}

	endIndex := 0
	for i, line := range lines {
		if len(line) != 0 {
			endIndex = i
		}
	}

	return lines[startIndex:endIndex + 1]
}
