package command

import (
	"io"
	"os"

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
	ViewDirPath  string
  Files files.FileList
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

type DialogBlueprint struct{
  DialogsDirPath string
  Files files.FileList
}

func (self DialogBlueprint) New(output io.Writer) (ComponentBlueprint, error) {
  projectName, err := find.ProjectName()
  if err != nil {
    return nil, apperr.Parse(err)
  }

  var classname Classname
  classname, err =classname.fromPrompt()

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

func (self DialogBlueprint) AddToProject(output io.Writer) error{
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
