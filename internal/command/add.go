package command

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/simon-co/fltr-cli/internal/apperr"
	"github.com/simon-co/fltr-cli/internal/files"
	"github.com/simon-co/fltr-cli/internal/find"
)

var addCmd = &Command{
	Name:    "add",
	Aliases: []string{"a"},
	runner:  addCmdRunner{},
}

func init() {
	addCmd.RegisterCommand()
}

type addCmdRunner struct{}

func (self addCmdRunner) Run(output io.Writer, workingDir string, args []string) error {
	selectedComponent := selectComponent()
	blueprint, err := selectedComponent.toBlueprint(output)
	if err != nil {
		return apperr.Parse(err)
	}
	if err := blueprint.AddToProject(output); err != nil {
		return apperr.Parse(err)
	}
	return nil
}

type ComponentBlueprint interface {
	New(io.Writer) (ComponentBlueprint, error)
	AddToProject(io.Writer) error
}

type ViewBlueprint struct {
	ViewDirPath string
	Files       files.FileList
}

func (self ViewBlueprint) New(output io.Writer) (ComponentBlueprint, error) {
	projectName, err := find.ProjectName()
	if err != nil {
		return nil, apperr.Parse(err)
	}

	var classname Classname
	classname, err = classname.fromPrompt()

	dirname := classname.toSnakeCase()
	viewFilename := "v_" + dirname + ".dart"
	ctlrFilename := "c_" + dirname + ".dart"

	viewsDirPath, err := find.ViewsDirPath()
	if err != nil {
		return nil, apperr.Parse(err)
	}

	self.ViewDirPath = viewsDirPath + string(os.PathSeparator) + dirname

	route, err := promptRoute()
	if err != nil {
		return nil, apperr.Parse(err)
	}

	self.Files = files.FileList{
		files.View(projectName, dirname, viewFilename, string(classname), ctlrFilename, route),
		files.ViewCtlr(dirname, ctlrFilename, string(classname), viewFilename),
	}

	return self, nil
}

func (self ViewBlueprint) AddToProject(output io.Writer) error {
	projectPath, err := find.ProjectDir()
	if err != nil {
		return apperr.Parse(err)
	}
	os.MkdirAll(self.ViewDirPath, 0751)
	if err := self.Files.InstantiateAll(projectPath); err != nil {
		return apperr.Parse(err)
	}
	return nil
}

type DialogBlueprint struct {
	DialogsDirPath string
	Files          files.FileList
}

func (self DialogBlueprint) New(output io.Writer) (ComponentBlueprint, error) {
	projectName, err := find.ProjectName()
	if err != nil {
		return nil, apperr.Parse(err)
	}

	var classname Classname
	classname, err = classname.fromPrompt()

	dirname := classname.toSnakeCase()
	dialogFilename := "d_" + dirname + ".dart"
	ctlrFilename := "c_" + dirname + ".dart"

	dialogsDirPath, err := find.DialogsDirPath()
	if err != nil {
		return nil, apperr.Parse(err)
	}

	self.DialogsDirPath = dialogsDirPath + string(os.PathSeparator) + dirname

	self.Files = files.FileList{
		files.Dialog(projectName, dirname, dialogFilename, string(classname), ctlrFilename),
		files.DialogCtlr(dirname, ctlrFilename, string(classname), dialogFilename),
	}

	return self, nil
}

func (self DialogBlueprint) AddToProject(output io.Writer) error {
	projectPath, err := find.ProjectDir()
	if err != nil {
		return apperr.Parse(err)
	}
	os.MkdirAll(self.DialogsDirPath, 0751)
	if err := self.Files.InstantiateAll(projectPath); err != nil {
		return apperr.Parse(err)
	}
	return nil
}

type NavigatorBlueprint struct {
	NavigatorDirPath      string
	Files                 files.FileList
	routerOutputString    string
	navigatorImportString string
}

func (self NavigatorBlueprint) New(output io.Writer) (ComponentBlueprint, error) {
	projectName, err := find.ProjectName()
	if err != nil {
		return nil, apperr.Parse(err)
	}

	var navClassname Classname
	navClassname, err = navClassname.fromPrompt()
	navClassname = navClassname + "Navigator"

	dirname := navClassname.toSnakeCase()
	navigatorFilename := "n_" + dirname + ".dart"

	route, err := promptRoute()
	if err != nil {
		return nil, apperr.Parse(err)
	}

	viewDirNames, err := find.AllViewDirNames()
	if err != nil {
		return nil, apperr.Parse(err)
	}

	var wg sync.WaitGroup
	viewFilenames := make([]viewFilename, len(viewDirNames))
	viewClassNames := make([]Classname, len(viewDirNames))
	for i, dirName := range viewDirNames {
		wg.Add(1)
		go func(i int, name string) {
			defer wg.Done()
			vn := viewFilename("").fromDirName(name)
			cn, err := vn.toClassName()
			if err != nil {
				return
			}
			viewFilenames[i] = vn
			viewClassNames[i] = cn
		}(i, dirName)
	}
	wg.Wait()
	if len(viewClassNames) <= 0 {
		return nil, apperr.Parse(errors.New("no views found"))
	}
	i, err := selectViewname(viewClassNames)
	if err != nil {
		return nil, apperr.Parse(err)
	}
	viewClassname := viewClassNames[i]
	navigatorDirPath, err := find.RouteNavigatorsPath()
	if err != nil {
		return nil, apperr.Parse(err)
	}

	viewPathname := viewDirNames[i] + string(os.PathSeparator) + "v_" + viewDirNames[i] + ".dart"

	self.NavigatorDirPath = navigatorDirPath + string(os.PathSeparator) + dirname

	self.Files = files.FileList{
		files.RouteNavigator(projectName, dirname, navigatorFilename, string(navClassname), route, viewPathname, string(viewClassname)),
	}

	var sb strings.Builder
	sb.WriteString("\tGoRoute(\n")
	sb.WriteString(fmt.Sprintf("\t\t\tname: %s.route,\n", navClassname))
	sb.WriteString(fmt.Sprintf("\t\t\tpath: %s.route,\n", navClassname))
	sb.WriteString(fmt.Sprintf("\t\t\tbuilder: (context, state) => %s(state.queryParams)),\n", navClassname))
	sb.WriteString("]);")
	self.routerOutputString = sb.String()
  self.navigatorImportString = filepath.Join(projectName, "src", "routing", "route_navigators", navigatorFilename)
  self.navigatorImportString = fmt.Sprintf("import 'package:%s';\n", self.navigatorImportString)

	return self, nil
}

func (self NavigatorBlueprint) AddToProject(output io.Writer) error {
	projectPath, err := find.ProjectDir()
	if err != nil {
		return apperr.Parse(err)
	}
	routerPath, err := find.RouterPath()
	if err != nil {
		return apperr.Parse(err)
	}
	content, err := os.ReadFile(routerPath)
	if err != nil {
		return apperr.Parse(err)
	}
	regex := regexp.MustCompile(`]\);`)
	content = []byte(regex.ReplaceAllString(string(content), self.routerOutputString))
  content = []byte(self.navigatorImportString + string(content))
	err = os.WriteFile(routerPath, content, 0644)
	if err != nil {
		return apperr.Parse(err)
	}
	if err := self.Files.InstantiateAll(projectPath); err != nil {
		return apperr.Parse(err)
	}
	return nil
}

type viewFilename string

func (self viewFilename) fromDirName(dirName string) viewFilename {
	self = "v_" + viewFilename(dirName) + ".dart"
	return self
}

func (self viewFilename) toClassName() (Classname, error) {
	split := strings.Split(string(self), ".dart")
	if len(split) != 2 {
		return "", apperr.Parse(errors.New("invalid view filename"))
	}
	split = strings.Split(split[0], "_")
	if len(split) <= 1 {
		return "", apperr.Parse(errors.New("invalid view filename"))
	}
	split = split[1:]
	for i, word := range split {
		runes := []rune(word)
		r := runes[0]
		upper := []rune(strings.ToUpper(string(r)))
		capitalised := append(upper, runes[1:]...)
		split[i] = string(capitalised)
	}
	classname := strings.Join(split, "")
	classname = classname + "View"
	return Classname(classname), nil
}
