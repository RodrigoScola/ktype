package formatting

import (
	"strings"
)

func FormatAll(str string) string {

	if strings.Contains(str, "—") {
		str = strings.ReplaceAll(str, "—", "-")
	}
	if strings.Contains(str, "’") {
		str = strings.ReplaceAll(str, "’", "'")
	}
	if strings.Contains(str, "“") {
		str = strings.ReplaceAll(str, "“", "\"")
	}
	if strings.Contains(str, "è") {
		str = strings.ReplaceAll(str, "è", "e")
	}
	if strings.Contains(str, "…") {
		str = strings.ReplaceAll(str, "…", "_")
	}
	if strings.Contains(str, "”") {
		str = strings.ReplaceAll(str, "”", "\"")
	}
	if strings.Contains(str, "é") {
		str = strings.ReplaceAll(str, "è", "e")

	}
	if strings.Contains(str, "é") {
		str = strings.ReplaceAll(str, "é", "e")
	}
	if strings.Contains(str, "‘") {
		str = strings.ReplaceAll(str, "‘", "\"")
	}
	if strings.Contains(str, "•") {
		str = strings.ReplaceAll(str, "•", "*")
	}
	if strings.Contains(str, "î") {
		str = strings.ReplaceAll(str, "î", "i")
	}
	str = strings.ToLower(str)
	return str
}

func FormatChar(char *byte) {

}
