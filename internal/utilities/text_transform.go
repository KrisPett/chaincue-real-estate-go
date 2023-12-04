package utilities

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

func FormatTitleCaseString(input string) string {
	replaced := strings.Replace(input, "_", " ", -1)

	formatted := cases.Title(language.English).String(replaced)

	return formatted
}
