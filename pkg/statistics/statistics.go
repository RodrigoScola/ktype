package statistics

import (
	"strings"
	"time"
)

func GetWords(str string) int {
	return strings.Count(str, " ")
}

// calculateWPM calculates the words per minute.
// It assumes each time.Time in the slice represents a timestamp for a typed word.

func CalculateWPM(startTime time.Time, endTime time.Time, wordCount int) float64 {
	duration := endTime.Sub(startTime).Seconds()

	wpm := float64(wordCount) * (60 / duration) / 5
	return wpm

}
