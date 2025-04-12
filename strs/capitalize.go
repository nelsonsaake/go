package strs

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var Capitalize = cases.Title(language.AmericanEnglish, cases.NoLower).String
