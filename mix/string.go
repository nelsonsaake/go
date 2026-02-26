package mix

import (
	"bytes"
	"fmt"
	"text/template"
)

func String(text string, data any) (string, error) {

	die := func(f string, a ...any) (string, error) {
		return "", fmt.Errorf(f, a...)
	}

	t := template.New("main").Funcs(funcs())

	t, err := t.Parse(text)
	if err != nil {
		return die("error parsing text: %v", err)
	}

	var (
		res string
		buf = bytes.Buffer{}
	)

	err = t.Execute(&buf, data)
	if err != nil {
		return die("error executing template: %v", err)
	}

	res = buf.String()
	return res, nil
}
