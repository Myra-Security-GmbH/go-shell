package command

import (
	"errors"
	"strings"
	"text/scanner"
)

var scan *scanner.Scanner

func init() {
	scan = &scanner.Scanner{
		Mode: scanner.ScanIdents | scanner.ScanStrings | scanner.ScanInts,
	}
}

//
// ParseEditorCommand ...
//
func ParseEditorCommand(editorCmd string) (Command, error) {
	var flags = make(map[string]bool)
	var params = []string{}
	var ret Command
	var err error

	sc := scan.Init(strings.NewReader(editorCmd))

	var token rune
	var arg string

	flagMode := false
	longFlag := false
	quote := false

	token = sc.Scan()
	if token == scanner.Ident {
		ret, err = FindCommand(sc.TokenText())

		if err != nil {
			return nil, err
		}
	} else if token == scanner.EOF {
		return nil, nil
	} else {
		return nil, errors.New("Missing command")
	}

	for token != scanner.EOF {
		switch {
		case (token == '-' && !flagMode):
			flagMode = true

		case (token == '-' && flagMode):
			longFlag = true
			arg = ""

		case (longFlag && flagMode && token > 0):
			arg += string(token)

			if sc.Peek() == ' ' || sc.Peek() == scanner.EOF {
				flags[strings.ToLower(arg)] = true
				arg = ""
				flagMode = false
				longFlag = false
			}

		case (flagMode && token > 0):
			flags[string(token)] = true

			if sc.Peek() == ' ' || sc.Peek() == scanner.EOF {
				flagMode = false
				longFlag = false
			}

		case (token == '"' && !quote):
			quote = true

		case token > 0:
			for ((!quote && token != ' ') || (quote && token != '"')) && token != scanner.EOF {
				arg += string(token)
				token = sc.Next()
			}

			if quote && token != '"' {
				return nil, errors.New("Missing closing doublequote")
			}

			if arg != "" {
				params = append(params, arg)
				arg = ""
				quote = false
			}
		}

		token = sc.Next()
	}

	ret.SetFlags(flags)
	ret.SetParams(params)

	return ret, nil
}
