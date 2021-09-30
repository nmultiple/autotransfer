package filter

func filterFooter(lines []string) []string {
	endIndex := -1

	for i, line := range lines {
		if line == "-- " || line == "○●-------------------------------------●○" {
			endIndex = i
			break
		}
	}

	return lines[:endIndex]
}
