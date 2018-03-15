package output

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/logrusorgru/aurora"
)

var boldRegex = regexp.MustCompile(`\*\S[^\*]*\S\*`)

//
// Sprintln ...
//
func Sprintln(data ...interface{}) string {
	var ret string
	var msg string

	for i := 0; i < len(data); i++ {
		msg = ""

		switch v := data[i].(type) {
		case error:
			msg = aurora.Colorize(
				strings.TrimSpace(v.Error()),
				aurora.RedBg,
			).String()

		case Success:
			msg = aurora.Colorize(
				strings.TrimSpace(v.Success()),
				aurora.GreenBg,
			).String()

		case Help:
			msg = formatHelp(v)

		case string:
			msg = v
		}

		ret += "\n" + msg
	}

	return ret[1:]
}

//
// Println ...
//
func Println(data ...interface{}) {
	fmt.Print(Sprintln(data...))
}
