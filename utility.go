package webframe

import (
	"strings"
)

func BuildString(elements ...string) string {
	if len(elements) == 0 {
		return ""
	}
	if len(elements) == 1 {
		return elements[0]
	}

	var buffer strings.Builder
	for _, element := range elements {
		buffer.WriteString(element)
	}
	return buffer.String()
}
