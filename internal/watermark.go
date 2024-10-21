package internal

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ricochhet/modmanager/pkg/logger"
)

func DrawWatermark(text []string) {
	result := []string{}

	longestLength := 0

	for _, textItem := range text {
		length := textLength(textItem)
		if length > longestLength {
			longestLength = length
		}
	}

	line := strings.Repeat("-", longestLength)
	result = append(result, fmt.Sprintf("┌─%s─┐", line))

	for _, textItem := range text {
		spacingSize := longestLength - textLength(textItem)
		spacingText := textItem + strings.Repeat(" ", spacingSize)
		result = append(result, fmt.Sprintf("│ %s │", spacingText))
	}

	result = append(result, fmt.Sprintf("└─%s─┘", line))

	for _, textItem := range result {
		logger.SharedLogger.Info(textItem)
	}
}

func textLength(s string) int {
	re := regexp.MustCompile(`[\p{Han}\p{Katakana}\p{Hiragana}\p{Hangul}]`)
	processedString := re.ReplaceAllString(s, "ab")

	return len(processedString)
}
