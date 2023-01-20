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
	ViewDirPath   string
	ViewDirName   string
	viewFilename  string
	viewClassname string
	Files         files.FileList
	AddToNav      bool
	Navigator     *navigatorFile
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
	self.viewFilename = viewFilename
	self.viewClassname = string(classname) + "View"

	viewsDirPath, err := find.ViewsDirPath()
	if err != nil {
		return nil, apperr.Parse(err)
	}
	self.ViewDirName = dirname
	self.ViewDirPath = viewsDirPath + string(os.PathSeparator) + dirname

	route, err := promptRoute()
	if err != nil {
		return nil, apperr.Parse(err)
	}

	//add to a navigator
	ok, err := promptConfirm("Would you like to add the view to an existing navigator?")
	if err != nil {
		return nil, apperr.Parse(err)
	}
	if ok {
		nfns, err := find.AllRouteNavigatorPaths()
		if err != nil {
			return nil, apperr.Parse(err)
		}
		navigators := []*navigatorFile{}
		var wg sync.WaitGroup
		var mu sync.Mutex
		for _, nfn := range nfns {
			wg.Add(1)
			go func(nfn string) {
				defer wg.Done()
				nf, err := navigatorFile{}.fromPathname(nfn)
				if err == nil {
					mu.Lock()
					defer mu.Unlock()
					navigators = append(navigators, nf)
				}
			}(nfn)
		}
		wg.Wait()
		if len(navigators) > 0 {
			ncn := make([]string, 0, len(navigators))
			for _, cn := range navigators {
				ncn = append(ncn, cn.className)
			}
			i, err := selectNavigator(ncn)
			if err != nil {
				return nil, apperr.Parse(err)
			}
			self.Navigator = navigators[i]
			self.AddToNav = true
		}
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
	if self.AddToNav {
		if err := self.addToNav(); err != nil {
			return apperr.Parse(err)
		}
	}
	os.MkdirAll(self.ViewDirPath, 0751)
	if err := self.Files.InstantiateAll(projectPath); err != nil {
		return apperr.Parse(err)
	}
	return nil
}

var viewNavFuncStr = `
  static AppError? toVIEW_CLASSNAME() {
    final navState = navKey.currentState;
    if (navState != null) {
      navState.pushReplacementNamed(VIEW_CLASSNAME.route);
    } else {
      return AppError(AppErrorCode.e500, "navState is null");
    }
    return null;
  }

  static AppError? pushVIEW_CLASSNAME() {
    final navState = navKey.currentState;
    if (navState != null) {
      navState.pushNamed(VIEW_CLASSNAME.route);
    } else {
      return AppError(AppErrorCode.e500, "navState is null");
    }
    return null;
  }
  static AppError? pop() {`

func (self ViewBlueprint) addToNav() error {
	projectName, err := find.ProjectName()
	if err != nil {
		return apperr.Parse(err)
	}
	navContent, err := os.ReadFile(self.Navigator.path)
	if err != nil {
		return apperr.Parse(err)
	}

	importStr := fmt.Sprintf("import 'package:%s/src/views/%s/%s';\n", projectName, self.ViewDirName, self.viewFilename)
	navContent = append([]byte(importStr), navContent...)
  
  navContent = self.addNavFuncs(self.viewClassname, navContent)

	replaceStr := "switch (settings.name) {"
	var sb strings.Builder
	sb.WriteString(replaceStr)
	sb.WriteString(fmt.Sprintf("\n\t\t\t\t\t\tcase %s.route:\n", self.viewClassname))
	sb.WriteString(fmt.Sprintf("\t\t\t\t\t\t\tview = %s(params);\n", self.viewClassname))
	sb.WriteString("\t\t\t\t\t\t\tbreak;")
	navContent = []byte(strings.Replace(string(navContent), replaceStr, sb.String(), 1))

	if err := os.WriteFile(self.Navigator.path, navContent, 0751); err != nil {
		return apperr.Parse(err)
	}
	return nil
}

// adds nav funcs to the suppliend navigator content and returns updated content
func (self ViewBlueprint) addNavFuncs(viewClassname string, navContent []byte) []byte {
  funcsStr := strings.ReplaceAll(viewNavFuncStr,"VIEW_CLASSNAME", viewClassname)
  navContent = []byte(strings.ReplaceAll(string(navContent),"static AppError? pop() {", funcsStr))
	return navContent
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
	dirname := navClassname.toSnakeCase()
	navClassname = navClassname + "Route" + "Navigator"

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

type navigatorFile struct {
	fileName  string
	className string
	path      string
}

var ErrInvalidNavFilePath = errors.New("supplied path is not a valid nav filename")

// returns the filename and classname of the associated file path
func (self navigatorFile) fromPathname(path string) (*navigatorFile, error) {
	self.path = path
	filename := filepath.Base(path)
	if filename[0] != byte('n') {
		return nil, apperr.Parse(ErrInvalidNavFilePath)
	}
	self.fileName = filename
	split := strings.Split(filename, ".dart")
	if len(split) != 2 {
		return nil, apperr.Parse(ErrInvalidNavFilePath)
	}
	split = strings.Split(split[0], "_")
	if len(split) <= 1 {
		return nil, apperr.Parse(ErrInvalidNavFilePath)
	}
	split = split[1:]
	for i, word := range split {
		runes := []rune(word)
		r := runes[0]
		upper := []rune(strings.ToUpper(string(r)))
		capitalised := append(upper, runes[1:]...)
		split[i] = string(capitalised)
	}
	self.className = strings.Join(split, "") + "RouteNavigator"
	return &self, nil
}

func (self *navigatorFile) addViewToNavigator(viewFilename string, viewClassname string) error {
	return nil
}
