package command

import (
	"errors"
	"strings"
	"unicode"

	"github.com/manifoldco/promptui"
	"github.com/simon-co/fltr-cli/internal/apperr"
)

var (
  ErrClassNameTooShort = errors.New("Supplied class name is too short")
  ErrClassNameNotPascal = errors.New("Class name must be pascal")
  ErrClassNameContainsSpace = errors.New("Class name contains space")
)

func selectComponent() ComponentType {
	prompt := promptui.Select{
		Label: "Select Component To Add",
		Items: ComponentTypes{}.ToMap().ToList(),
	}

	_, res, err := prompt.Run()
	if err != nil {
		panic(err)
	}
	return ComponentType(res)
}

func promptClassName() (string, error) {
  prompt := promptui.Prompt{
    Label: "Please provide a base class name in Pascal case.",
    Validate: validateClassName,
    HideEntered: true,
  }
  className, err := prompt.Run()
  if err != nil {
    return "", apperr.Parse(err)
  }
  return className, nil
}

//validates that the provided input is Pascal case
func validateClassName(input string) error {
	if len(input) <= 1 {
		return ErrClassNameTooShort
	}
	if unicode.IsLower(rune(input[0])) {
		return ErrClassNameNotPascal
	}
	if strings.Contains(input, " ") {
		return ErrClassNameContainsSpace
	}
	return nil
}

func promptRoute() (string, error){
  prompt := promptui. Prompt{
    Label: "What route would you like to use?",
    HideEntered: true,
  }

  route, err := prompt.Run()
  if err != nil {
    return "", apperr.Parse(err)
  }
  return route, nil
}
