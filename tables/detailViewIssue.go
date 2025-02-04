package table

import (
	"fmt"
	"strings"

	"github.com/Kaya-Sem/commandtrein/api"
	"github.com/Kaya-Sem/commandtrein/internal/util"
)

func getDetailedIssueInfo(i api.Issue) string {
	output := ""

	output += fmt.Sprintf("\x1b[1m%s\x1b[0m\n", i.Title)
	output += util.ConvertEpochStringToDate(i.Timestamp)

	if i.Type == "planned" {
		output += " (planned)"
	}

	output += "\n\n"

	// Properly wrap description to 80 chars per line
	wrapped := wrapText(i.Description, 80)
	output += wrapped + "\n\n"

	// All links are present in object, but due to wrapping and hyperlink issues, not included yet.
	output += fmt.Sprintf("%s:\n%s\n", "Info", "https://www.belgiantrain.be/nl/travel-info/current/ongoing-disturbances-and-works")

	return output
}

// Helper function to wrap text to a given width
func wrapText(text string, limit int) string {
	var result strings.Builder
	var line strings.Builder
	words := strings.Fields(text)

	for _, word := range words {
		if line.Len()+len(word) > limit {
			result.WriteString(line.String() + "\n")
			line.Reset()
		}
		if line.Len() > 0 {
			line.WriteString(" ")
		}
		line.WriteString(word)
	}

	result.WriteString(line.String())
	return result.String()
}
