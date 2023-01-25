package files

import (
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/simon-co/fltr-cli/internal/apperr"
)

type File struct {
	OutputString       string //string to be output to the outfile
	OutputFilename     string
	OutputFilePath     string
	outputPathParts    []string
	EmbeddedFileReader *FileReader
	Replacements       map[string]string
	Mu                 sync.Mutex
}

func (self *File) loadTemplateFileOutput() error {
	if err := self.EmbeddedFileReader.Read(); err != nil {
		return err
	}
	self.OutputString = self.EmbeddedFileReader.Data.String()
	return nil
}

func (self *File) addRepalcementsToOutput() {
	var wg sync.WaitGroup
	for t, r := range self.Replacements {
		wg.Add(1)
		go func(target string, replace string) {
			defer wg.Done()
			regex := regexp.MustCompile(target)
			self.Mu.Lock()
			defer self.Mu.Unlock()
			self.OutputString = regex.ReplaceAllString(self.OutputString, replace)
		}(t, r)
	}
	wg.Wait()
}

func (self *File) setOutputPath(projectPath string) {
	var sb strings.Builder
	sb.WriteString(projectPath)
	for _, val := range self.outputPathParts {
		sb.WriteRune(os.PathSeparator)
		sb.WriteString(val)
	}
	sb.WriteRune(os.PathSeparator)
	sb.WriteString(self.OutputFilename)
	self.OutputFilePath = sb.String()
}

func (self *File) Instantiate(projectPath string) error {
	if err := self.loadTemplateFileOutput(); err != nil {
		return err
	}
	self.setOutputPath(projectPath)
	self.addRepalcementsToOutput()
	if err := os.WriteFile(self.OutputFilePath, []byte(self.OutputString), 0644); err != nil {
		return apperr.Parse(err)
	}
	return nil
}

func Main(projectName string) *File {
	file := &File{
		OutputString:       "",
		OutputFilename:     "main.dart",
		OutputFilePath:     "",
		outputPathParts:    []string{"lib"},
		EmbeddedFileReader: FileReader{}.New(EmbeddedFsPaths.Main),
		Replacements: map[string]string{
			"PROJECT_NAME": projectName,
		},
		Mu: sync.Mutex{},
	}
	return file
}

func App(projectName string) *File {
	return &File{
		OutputString:       "",
		OutputFilename:     "app.dart",
		OutputFilePath:     "",
		outputPathParts:    []string{"lib", "src", "app"},
		EmbeddedFileReader: FileReader{}.New(EmbeddedFsPaths.App),
		Replacements: map[string]string{
			"PROJECT_NAME": projectName,
		},
		Mu: sync.Mutex{},
	}
}

func AppConfig(projectName string) *File {
	return &File{
		OutputString:       "",
		OutputFilename:     "app_config.dart",
		OutputFilePath:     "",
		outputPathParts:    []string{"lib", "src", "app"},
		EmbeddedFileReader: FileReader{}.New(EmbeddedFsPaths.AppConfig),
		Replacements: map[string]string{
			"PROJECT_NAME": projectName,
		},
		Mu: sync.Mutex{},
	}
}

func AppError(projectName string) *File {
	return &File{
		OutputString:       "",
		OutputFilename:     "app_error.dart",
		OutputFilePath:     "",
		outputPathParts:    []string{"lib", "src", "app"},
		EmbeddedFileReader: FileReader{}.New(EmbeddedFsPaths.AppError),
		Replacements: map[string]string{
			"PROJECT_NAME": projectName,
		},
		Mu: sync.Mutex{},
	}
}

func AppErrorG() *File {
	return &File{
		OutputString:       "",
		OutputFilename:     "app_error.g.dart",
		OutputFilePath:     "",
		outputPathParts:    []string{"lib", "src", "app"},
		EmbeddedFileReader: FileReader{}.New(EmbeddedFsPaths.AppErrorG),
		Replacements:       map[string]string{},
		Mu:                 sync.Mutex{},
	}
}

func AppResult(projectName string) *File {
	return &File{
		OutputString:       "",
		OutputFilename:     "app_result.dart",
		OutputFilePath:     "",
		outputPathParts:    []string{"lib", "src", "app"},
		EmbeddedFileReader: FileReader{}.New(EmbeddedFsPaths.AppResult),
		Replacements: map[string]string{
			"PROJECT_NAME": projectName,
		},
		Mu: sync.Mutex{},
	}
}

func AppTheme() *File {
	return &File{
		OutputString:       "",
		OutputFilename:     "app_theme.dart",
		OutputFilePath:     "",
		outputPathParts:    []string{"lib", "src", "app"},
		EmbeddedFileReader: FileReader{}.New(EmbeddedFsPaths.AppTheme),
		Replacements:       map[string]string{},
		Mu:                 sync.Mutex{},
	}
}

func AppCalltrace() *File {
	return &File{
		OutputString:       "",
		OutputFilename:     "app_calltrace.dart",
		OutputFilePath:     "",
		outputPathParts:    []string{"lib", "src", "app"},
		EmbeddedFileReader: FileReader{}.New(EmbeddedFsPaths.AppCalltrace),
		Replacements:       map[string]string{},
		Mu:                 sync.Mutex{},
	}
}

func SettingsModel(projectName string) *File {
	return &File{
		OutputString:       "",
		OutputFilename:     "m_settings.dart",
		OutputFilePath:     "",
		outputPathParts:    []string{"lib", "src", "models"},
		EmbeddedFileReader: FileReader{}.New(EmbeddedFsPaths.SettingsModel),
		Replacements: map[string]string{
			"PROJECT_NAME": projectName,
		},
		Mu: sync.Mutex{},
	}
}

func SettingsModelG() *File {
	return &File{
		OutputString:       "",
		OutputFilename:     "m_settings.g.dart",
		OutputFilePath:     "",
		outputPathParts:    []string{"lib", "src", "models"},
		EmbeddedFileReader: FileReader{}.New(EmbeddedFsPaths.SettingsModelG),
		Replacements:       map[string]string{},
		Mu:                 sync.Mutex{},
	}
}

func IsarService(projectName string) *File {
	return &File{
		OutputString:       "",
		OutputFilename:     "s_isar.dart",
		OutputFilePath:     "",
		outputPathParts:    []string{"lib", "src", "services"},
		EmbeddedFileReader: FileReader{}.New(EmbeddedFsPaths.IsarService),
		Replacements: map[string]string{
			"PROJECT_NAME": projectName,
		},
		Mu: sync.Mutex{},
	}
}

func Router(projectName string) *File {
	return &File{
		OutputString:       "",
		OutputFilename:     "router.dart",
		OutputFilePath:     "",
		outputPathParts:    []string{"lib", "src", "routing"},
		EmbeddedFileReader: FileReader{}.New(EmbeddedFsPaths.Router),
		Replacements: map[string]string{
			"PROJECT_NAME": projectName,
		},
		Mu: sync.Mutex{},
	}
}

func SplashNavigator(projectName string) *File {
	return &File{
		OutputString:       "",
		OutputFilename:     "n_splash.dart",
		OutputFilePath:     "",
		outputPathParts:    []string{"lib", "src", "routing", "route_navigators"},
		EmbeddedFileReader: FileReader{}.New(EmbeddedFsPaths.SplashNav),
		Replacements: map[string]string{
			"PROJECT_NAME": projectName,
		},
		Mu: sync.Mutex{},
	}
}

func HomeNavigator(projectName string) *File {
	return &File{
		OutputString:       "",
		OutputFilename:     "n_home.dart",
		OutputFilePath:     "",
		outputPathParts:    []string{"lib", "src", "routing", "route_navigators"},
		EmbeddedFileReader: FileReader{}.New(EmbeddedFsPaths.HomeNav),
		Replacements: map[string]string{
			"PROJECT_NAME": projectName,
		},
		Mu: sync.Mutex{},
	}
}

func SplashView(projectName string) *File {
	return &File{
		OutputString:       "",
		OutputFilename:     "v_splash.dart",
		OutputFilePath:     "",
		outputPathParts:    []string{"lib", "src", "views", "splash"},
		EmbeddedFileReader: FileReader{}.New(EmbeddedFsPaths.SplashView),
		Replacements: map[string]string{
			"PROJECT_NAME": projectName,
		},
		Mu: sync.Mutex{},
	}
}

func SplashCtrl() *File {
	return &File{
		OutputString:       "",
		OutputFilename:     "c_splash.dart",
		OutputFilePath:     "",
		outputPathParts:    []string{"lib", "src", "views", "splash"},
		EmbeddedFileReader: FileReader{}.New(EmbeddedFsPaths.SplashCtrl),
		Replacements:       map[string]string{},
		Mu:                 sync.Mutex{},
	}
}

func StartView(projectName string) *File {
	return &File{
		OutputString:       "",
		OutputFilename:     "v_start.dart",
		OutputFilePath:     "",
		outputPathParts:    []string{"lib", "src", "views", "start"},
		EmbeddedFileReader: FileReader{}.New(EmbeddedFsPaths.StartView),
		Replacements: map[string]string{
			"PROJECT_NAME": projectName,
		},
		Mu: sync.Mutex{},
	}
}

func StartCtrl() *File {
	return &File{
		OutputString:       "",
		OutputFilename:     "c_start.dart",
		OutputFilePath:     "",
		outputPathParts:    []string{"lib", "src", "views", "start"},
		EmbeddedFileReader: FileReader{}.New(EmbeddedFsPaths.StartCtrl),
		Replacements:       map[string]string{},
		Mu:                 sync.Mutex{},
	}
}

func View(projectName string, dirname string, filename string, classname string, controllerFilename string, route string) *File {
	return &File{
		OutputString:       "",
		OutputFilename:     filename,
		OutputFilePath:     "",
		outputPathParts:    []string{"lib", "src", "views", dirname},
		EmbeddedFileReader: FileReader{}.New(EmbeddedFsPaths.ViewView),
		Replacements: map[string]string{
			"PROJECT_NAME":         projectName,
			"CONTROLLER_FILE_NAME": controllerFilename,
			"CLASS_NAME":           classname,
			"ROUTE":                route,
		},
		Mu: sync.Mutex{},
	}
}

func ViewCtrl(dirname string, filename string, classname string, viewFilename string) *File {
	return &File{
		OutputString:       "",
		OutputFilename:     filename,
		OutputFilePath:     "",
		outputPathParts:    []string{"lib", "src", "views", dirname},
		EmbeddedFileReader: FileReader{}.New(EmbeddedFsPaths.ViewCtrl),
		Replacements: map[string]string{
			"VIEW_FILENAME": viewFilename,
			"CLASS_NAME":    classname,
		},
		Mu: sync.Mutex{},
	}
}

func SettingsDialogView(projectName string) *File {
	return &File{
		OutputString:       "",
		OutputFilename:     "d_app_settings.dart",
		OutputFilePath:     "",
		outputPathParts:    []string{"lib", "src", "dialogs", "app_settings"},
		EmbeddedFileReader: FileReader{}.New(EmbeddedFsPaths.SettingsDialogView),
		Replacements: map[string]string{
			"PROJECT_NAME": projectName,
		},
		Mu: sync.Mutex{},
	}
}

func SettingsDialogCtrl() *File {
	return &File{
		OutputString:       "",
		OutputFilename:     "c_app_settings.dart",
		OutputFilePath:     "",
		outputPathParts:    []string{"lib", "src", "dialogs", "app_settings"},
		EmbeddedFileReader: FileReader{}.New(EmbeddedFsPaths.SettingsDialogCtrl),
		Replacements:       map[string]string{},
		Mu:                 sync.Mutex{},
	}
}

func Dialog(projectName string, dirname string, filename string, classname string, controllerFilename string) *File {
	return &File{
		OutputString:       "",
		OutputFilename:     filename,
		OutputFilePath:     "",
		outputPathParts:    []string{"lib", "src", "dialogs", dirname},
		EmbeddedFileReader: FileReader{}.New(EmbeddedFsPaths.DialogView),
		Replacements: map[string]string{
			"PROJECT_NAME":        projectName,
			"CONTROLLER_FILENAME": controllerFilename,
			"CLASS_NAME":          classname,
		},
		Mu: sync.Mutex{},
	}
}

func DialogCtrl(dirname string, filename string, classname string, dialogFilename string) *File {
	return &File{
		OutputString:       "",
		OutputFilename:     filename,
		OutputFilePath:     "",
		outputPathParts:    []string{"lib", "src", "dialogs", dirname},
		EmbeddedFileReader: FileReader{}.New(EmbeddedFsPaths.DialogCtrl),
		Replacements: map[string]string{
			"DIALOG_FILENAME": dialogFilename,
			"CLASS_NAME":      classname,
		},
		Mu: sync.Mutex{},
	}
}

func RouteNavigator(projectName, dirname, filename, classname, route, viewPathName, viewClassName string) *File {
	return &File{
		OutputString:       "",
		OutputFilename:     filename,
		OutputFilePath:     "",
		outputPathParts:    []string{"lib", "src", "routing", "route_navigators"},
		EmbeddedFileReader: FileReader{}.New(EmbeddedFsPaths.RouteNav),
		Replacements: map[string]string{
			"PROJECT_NAME":    projectName,
			"VIEW_PATH":       viewPathName,
			"NAV_CLASSNAME":   classname,
			"ROUTE":           route,
			"VIEW_CLASS_NAME": viewClassName,
		},
		Mu: sync.Mutex{},
	}
}

type FileList []*File

func (self FileList) InstantiateAll(projectPath string) error {
	for _, val := range self {
		if err := val.Instantiate(projectPath); err != nil {
			return apperr.Parse(err)
		}
	}
	return nil
}
