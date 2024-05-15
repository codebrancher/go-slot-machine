package display

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

func (cd *ConsoleDisplay) printCentered(text string, totalWidth int) {
	visibleLength := len([]rune(text))

	totalPadding := totalWidth - visibleLength
	if totalPadding < 0 {
		totalPadding = 0
	}

	leftPadding := totalPadding / 2
	rightPadding := totalPadding - leftPadding

	paddedText := fmt.Sprintf("|%s%s%s|", strings.Repeat(" ", leftPadding), text, strings.Repeat(" ", rightPadding))
	fmt.Println(paddedText)
}

func (cd *ConsoleDisplay) printLeftAligned(label string, totalWidth int) {

	fullMessage := fmt.Sprintf(" %s", label)
	if len(fullMessage) > totalWidth {
		fullMessage = fullMessage[:totalWidth]
	}
	padding := totalWidth - len(fullMessage)
	fmt.Printf("|%s%s|\n", fullMessage, strings.Repeat(" ", padding))
}

func (cd *ConsoleDisplay) printWithMiddlePadding(leftValue, rightValue string, frameWidth int) {
	contentWidth := frameWidth - 2

	leftLength := utf8.RuneCountInString(leftValue)
	rightLength := utf8.RuneCountInString(rightValue)
	totalLength := leftLength + rightLength

	if totalLength > contentWidth {
		// If total length exceeds content width, truncate the longer part or both
		excess := totalLength - contentWidth
		if leftLength > rightLength {
			leftValue = truncateString(leftValue, max(0, leftLength-excess))
		} else {
			rightValue = truncateString(rightValue, max(0, rightLength-excess))
		}
		// Recalculate lengths after truncation
		leftLength = utf8.RuneCountInString(leftValue)
		rightLength = utf8.RuneCountInString(rightValue)
		totalLength = leftLength + rightLength
	}

	padding := contentWidth - totalLength // Calculate the padding needed in the middle
	fmt.Printf("|%s%s%s|\n", leftValue, strings.Repeat(" ", padding), rightValue)
}

func (cd *ConsoleDisplay) printBlankLine() {
	frameWidth := 32               // Includes the border "|"
	contentWidth := frameWidth - 2 // Width available for content between the borders
	fmt.Printf("|%s|\n", strings.Repeat(" ", contentWidth))
}

func (cd *ConsoleDisplay) printIntermediaryLine() {
	frameWidth := 32               // Includes the border "|"
	contentWidth := frameWidth - 2 // Width available for content between the borders
	fmt.Printf("|%s|\n", strings.Repeat("-", contentWidth))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func formatDuration(d time.Duration) string {
	hours := d / time.Hour
	d %= time.Hour
	minutes := d / time.Minute
	d %= time.Minute
	seconds := d / time.Second
	d %= time.Second
	milliseconds := d / time.Millisecond
	return fmt.Sprintf("%02d:%02d:%02d:%02d", hours, minutes, seconds, milliseconds)
}

// Helper to truncate string to the specified number of runes
func truncateString(s string, num int) string {
	if num >= utf8.RuneCountInString(s) {
		return s
	}
	return string([]rune(s)[:num])
}