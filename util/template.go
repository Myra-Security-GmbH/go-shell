package util

import text "text/template"

//
// NewTextTemplate ...
//
func NewTextTemplate(content string) *text.Template {
	ret, err := text.New("").Parse(content)

	if err != nil {
		panic(err)
	}

	return ret
}
