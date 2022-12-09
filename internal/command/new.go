package command

import (
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/simon-co/fltr-cli/internal/apperr"
	"github.com/simon-co/fltr-cli/internal/files"
)

var newCmd = &Command{
	Name:    "new",
	Aliases: []string{"n"},
	runner:  newCmdRunner{},
}

func init() {
	newCmd.RegisterCommand()
}

type newCmdRunner struct{}

func (self newCmdRunner) Run(output io.Writer, workingDir string, args []string) error {
	blueprint, err := ProjectBlueprint{}.New(workingDir)
	if err != nil {
		return apperr.Parse(err)
	}
	if err := blueprint.Instantiate(output); err != nil {
		os.RemoveAll(blueprint.ProjectPath)
		return apperr.Parse(err)
	}
	return nil
}

type ProjectBlueprint struct {
	ProjectName     string
	ProjectPath     string
	Files           files.FileList
	Directories     []string
	Dependencies    []string
	DevDependencies []string
}

func (self ProjectBlueprint) New(workingDir string) (*ProjectBlueprint, error) {
	prompt := promptui.Prompt{
		Label: "Project Name",
	}
	projectName, err := prompt.Run()
	if err != nil {
		return nil, apperr.Parse(err)
	}
	projectPath := filepath.Join(workingDir, projectName)
	return &ProjectBlueprint{
		ProjectName: projectName,
		ProjectPath: projectPath,
		Files: files.FileList{
			files.Main(projectName),
			files.App(projectName),
			files.AppConfig(projectName),
			files.AppTheme(),
			files.AppCalltrace(),
			files.AppResult(projectName),
			files.AppError(projectName),
			files.AppErrorG(),
			files.SettingsModel(projectName),
			files.SettingsModelG(),
			files.IsarService(projectName),
			files.Router(projectName),
      files.HomeNavigator(projectName),
      files.SplashNavigator(projectName),
			files.SplashView(projectName),
			files.SplashCtlr(),
			files.StartView(projectName),
			files.StartCtlr(),
			files.SettingsDialogView(projectName),
			files.SettingsDialogCtlr(),
		},
		Directories: []string{"app", "dialogs", "models", "routing", "util", "services",
      "routing" + string(os.PathSeparator) + "route_navigators", "views" + string(os.PathSeparator) + "splash", 
      "views" + string(os.PathSeparator) + "start", "dialogs" + string(os.PathSeparator) + "app_settings",
		},
		Dependencies:    []string{"equatable", "go_router", "flutter_bloc", "json_annotation", "path_provider", "shared_preferences", "reactive_forms", "yaml", "isar", "isar_flutter_libs"},
		DevDependencies: []string{"build_runner", "json_serializable", "isar_generator"},
	}, nil
}

func (self *ProjectBlueprint) Instantiate(output io.Writer) error {
	if err := self.initFlutterProject(output); err != nil {
		return apperr.Parse(err)
	}
	self.createDirectories()
	if err := self.Files.InstantiateAll(self.ProjectName); err != nil {
		return apperr.Parse(err)
	}
	if err := self.addDependencies(output); err != nil {
		return apperr.Parse(err)
	}
	if err := self.buildRunner(output); err != nil {
		return apperr.Parse(err)
	}
	return nil
}

func (self *ProjectBlueprint) initFlutterProject(output io.Writer) error {
	cmd := exec.Command("flutter", "create", self.ProjectName)
	cmd.Stdout = output
	cmd.Stderr = output
	if err := cmd.Run(); err != nil {
		return apperr.Parse(err)
	}
	return nil
}

func (self *ProjectBlueprint) createDirectories() {
	srcPath := filepath.Join(self.ProjectPath, "lib", "src")
	os.MkdirAll(srcPath, 0751)
	for _, val := range self.Directories {
		os.MkdirAll(srcPath+string(os.PathSeparator)+val, 0751)
	}
}

func (self *ProjectBlueprint) addDependencies(output io.Writer) error {
	for _, val := range self.Dependencies {
		cmd := exec.Command("flutter", "pub", "add", val)
		cmd.Dir = self.ProjectPath
		cmd.Stdout = output
		cmd.Stderr = output
		if err := cmd.Run(); err != nil {
			return apperr.Parse(err)
		}
	}
	for _, val := range self.DevDependencies {
		cmd := exec.Command("flutter", "pub", "add", "--dev", val)
		cmd.Dir = self.ProjectPath
		cmd.Stdout = output
		cmd.Stderr = output
		if err := cmd.Run(); err != nil {
			return apperr.Parse(err)
		}
	}
	return nil
}

func (self *ProjectBlueprint) buildRunner(output io.Writer) error {
	cmd := exec.Command("flutter", "pub", "run", "build_runner", "build", "--delete-conflicting-outputs")
	cmd.Dir = self.ProjectPath
	cmd.Stdout = output
	cmd.Stderr = output
	if err := cmd.Run(); err != nil {
		return apperr.Parse(err)
	}
	return nil
}
