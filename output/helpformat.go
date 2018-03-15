package output

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/logrusorgru/aurora"
)

//
// formatHelp formats a Help structure
//
func formatHelp(v Help) string {
	lenArguments, lenOptions := v.GetParamLen()
	arguments := aurora.Colorize("Arguments:", aurora.BrownFg).String() + "\n"
	options := aurora.Colorize("Options:", aurora.BrownFg).String() + "\n"

	msg := aurora.Colorize("Usage:", aurora.BrownFg).String() + "\n"
	msg += "  " + v.Command()

	if lenOptions > 0 {
		msg += " [options]"
	}

	for _, p := range v.Params() {
		if p.Flag {
			options += formatFlagLines(p)

			continue
		}

		if p.Required {
			msg += " <" + p.Name + ">"
		} else {
			msg += " [" + p.Name + "]"
		}

		arguments += formatArgumentLines(p)
		arguments += formatExample(p)
	}

	msg += "\n"

	if lenArguments > 0 {
		msg += "\n" + arguments
	}

	if lenOptions > 0 {
		msg += "\n" + options
	}

	msg += "\n" + aurora.Colorize("Help:", aurora.BrownFg).String() + "\n"

	helpText := ""
	reader := bufio.NewReader(strings.NewReader(v.Help()))

	for {
		l, err := reader.ReadString('\n')

		if strings.HasPrefix(l, "   ") {
			l = aurora.Colorize(l, aurora.GreenFg).String()
		}

		helpText += "  " + formatText(l)

		if err != nil {
			break
		}
	}

	msg += helpText

	return msg
}

//
// formatExample formats an example of a HelpParam
//
func formatExample(p HelpParam) string {
	var ret string

	if p.Example != "" {
		ret = fmt.Sprintf(
			"   %20s %s\n",
			"",
			aurora.Colorize("[example:"+p.Example+"]", aurora.BrownFg).String(),
		)
	}

	return ret
}

//
// formatDefaultValue formats a default value of a HelpParam
//
func formatDefaultValue(p HelpParam) string {
	var ret string

	if p.Flag && p.DefaultValue != "" {
		ret = " " + aurora.Colorize("[default:"+p.DefaultValue+"]", aurora.BrownFg).
			String()
	}

	return ret
}

//
// formatText ...
//
func formatText(text string) string {
	for _, t := range boldRegex.FindAllString(text, -1) {
		text = strings.Replace(
			text,
			t,
			aurora.Colorize(strings.Trim(t, "*"), aurora.BoldFm).String(),
			-1,
		)
	}

	return text
}

func formatOptions(p HelpParam) (shortUsage string, longUsage string) {
	shortUsage = "-" + p.Name

	if p.LongName != "" {
		longUsage = "--" + p.LongName

		if p.Name != "" {
			longUsage = ", " + longUsage
		}
	}

	return
}

//
// formatFlagLines ...
//
func formatArgumentLines(p HelpParam) string {
	var ret string

	if p.Flag {
		return ret
	}

	argumentsLineNum := 0
	reader := bufio.NewReader(strings.NewReader(p.Description))

	for {
		l, err := reader.ReadString('\n')
		l = formatText(strings.Trim(l, "\n"))

		if argumentsLineNum == 0 {
			ret += fmt.Sprintf(
				"  %-30s %s\n",
				aurora.Colorize(p.Name, aurora.GreenFg).String(),
				l,
			)
		} else {
			ret += fmt.Sprintf(
				"   %-20s %s\n",
				"",
				l,
			)
		}

		if err != nil {
			ret = strings.TrimRight(ret, "\n") + formatDefaultValue(p) + "\n"
			break
		}

		argumentsLineNum++
	}

	return ret
}

func formatFlagLines(p HelpParam) string {
	var ret string

	if !p.Flag {
		return ret
	}

	argumentsLineNum := 0
	reader := bufio.NewReader(strings.NewReader(p.Description))

	for {
		l, err := reader.ReadString('\n')
		l = formatText(strings.Trim(l, "\n"))

		if argumentsLineNum == 0 {
			shortUsage, longUsage := formatOptions(p)

			ret += fmt.Sprintf(
				"  %-2s%-28s %s\n",
				aurora.Colorize(shortUsage, aurora.GreenFg).String(),
				aurora.Colorize(longUsage, aurora.GreenFg).String(),
				l,
			)
		} else {
			ret += fmt.Sprintf(
				"  %-20s %s\n",
				"",
				l,
			)
		}

		if err != nil {
			break
		}

		argumentsLineNum++
	}

	return ret
}
