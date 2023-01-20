package command

import (
	"errors"
	"regexp"
	"strings"
	"unicode"

	"github.com/manifoldco/promptui"
	"github.com/simon-co/fltr-cli/internal/apperr"
)

var (
	ErrClassNameTooShort      = errors.New("Supplied class name is too short")
	ErrClassNameNotPascal     = errors.New("Class name must be pascal")
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

func selectViewname(viewNames any) (int, error) {
	prompt := promptui.Select{
		Label: "Select a default view",
		Items: viewNames,
	}
	i, _, err := prompt.Run()
	if err != nil {
		return -1, apperr.Parse(err)
	}
	return i, nil
}

func selectNavigator(navigatorNames any) (int, error){
prompt := promptui.Select{
		Label: "Select navigator",
		Items: navigatorNames,
	}
	i, _, err := prompt.Run()
	if err != nil {
		return -1, apperr.Parse(err)
	}
	return i, nil
}

func promptClassName() (string, error) {
	prompt := promptui.Prompt{
		Label:    "Please provide a base class name in Pascal case.",
		Validate: validateClassName,
    HideEntered: false,
	}
	className, err := prompt.Run()
	if err != nil {
		return "", apperr.Parse(err)
	}
	return className, nil
}

// validates that the provided input is Pascal case
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
	if strings.ContainsAny(input, "/\\,'?!\"0123456789.") {
		return ErrClassNameNotPascal
	}
	return nil
}

func promptRoute() (string, error) {
	prompt := promptui.Prompt{
		Label:       "What route would you like to use?",
		HideEntered: false,
	}

	route, err := prompt.Run()
	if err != nil {
		return "", apperr.Parse(err)
	}
	return route, nil
}

var ErrInvalidConfirmInput = errors.New("Invaild input: n/no/y/yes accepted.")

// validates that that input is a valid confirm string
func validateConfirm(input string) error {
	regex := regexp.MustCompile(`^([nN])+[oO]{0,1}$|^([yY]+(es)*)$`)
	matches := regex.FindStringSubmatch(input)
	if matches == nil {
		return ErrInvalidConfirmInput
	}
	return nil
}

//parses confirm string to bool
func confirmStrToBool(input string) (bool, error) {
  regex := regexp.MustCompile(`^([nN])+[oO]{0,1}$|^([yY])+(?:es)?$`)
	matches := regex.FindStringSubmatch(input)
	if matches == nil {
		return false,ErrInvalidConfirmInput
  }
  if unicode.ToLower(rune(matches[0][0])) == rune('n'){
    return false, nil
  }
  return true, nil
}

//prompts user for confirmation
func promptConfirm(label string) (bool, error) {
  prompt := promptui.Prompt{
  	Label:       label,
  	Validate:    validateConfirm,
  }
  input, err := prompt.Run()
  if err != nil {
    return false, apperr.Parse(err)
  }
  return confirmStrToBool(input)
}
