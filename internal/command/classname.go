package command

import (
	"strings"
	"unicode"

	"github.com/manifoldco/promptui"
	"github.com/simon-co/fltr-cli/internal/apperr"
)

type Classname string

func (self Classname) fromPrompt() (Classname, error) {
	prompt := promptui.Prompt{
		Label:       "Please provide a base class name in Pascal case.",
		Validate:    validateClassName,
		HideEntered: true,
	}
	classname, err := prompt.Run()
	if err != nil {
		return "", apperr.Parse(err)
	}
	return Classname(classname), nil
}

func (self Classname) toSnakeCase() string{
var sb strings.Builder
	sb.Grow(len(self))
	for n, char := range self {
		if sb.Len() == 0 {
			if unicode.IsLetter(char) {
				sb.WriteString(strings.ToLower(string(char)))
			} else if !unicode.IsSpace(char) {
				sb.WriteRune(char)
			}
		} else {
			if unicode.IsSpace(char) {
				if !unicode.IsSpace(rune(self[n-1])) && !unicode.IsSpace(rune(self[n+1])) {
					sb.WriteString("_")
				}
			} else if unicode.IsUpper(char) {
				if !unicode.IsSpace(rune(self[n-1])) {
					sb.WriteString("_")
				}
				sb.WriteString(strings.ToLower(string(char)))
			} else {
				sb.WriteRune(char)
			}
		}
	}
	return sb.String()
}
