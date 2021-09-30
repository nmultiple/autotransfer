package filter

import "strings"

type TextFilter struct {
	filterFunc []func(line []string) []string
}

func New() *TextFilter {
	result := new(TextFilter)
	result.filterFunc = append(result.filterFunc, filterHeader)
	result.filterFunc = append(result.filterFunc, filterQuote)
	result.filterFunc = append(result.filterFunc, filterFooter)
	result.filterFunc = append(result.filterFunc, filterUselessEmptyLine)

	return result
}

func (f *TextFilter) FilterText(text string) string {
	lines := strings.Split(text, "\r\n")
	for _, filter := range f.filterFunc {
		lines = filter(lines)
	}
	return strings.Join(lines, "\n")
}
