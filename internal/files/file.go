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

func AppTheme(projectName string) *File {
	return &File{
		OutputString:       "",
		OutputFilename:     "app_theme.dart",
		OutputFilePath:     "",
		outputPathParts:    []string{"lib", "src", "app"},
		EmbeddedFileReader: FileReader{}.New(EmbeddedFsPaths.AppCalltrace),
		Replacements: map[string]string{
			"PROJECT_NAME": projectName,
		},
		Mu: sync.Mutex{},
	}
}

func RootNav(projectName string) *File {
	return &File{
		OutputString:       "",
		OutputFilename:     "n_root.dart",
		OutputFilePath:     "",
		outputPathParts:    []string{"lib", "src", "navigation"},
		EmbeddedFileReader: FileReader{}.New(EmbeddedFsPaths.RootNav),
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

func SplashCtlr() *File {
	return &File{
		OutputString:       "",
		OutputFilename:     "c_splash.dart",
		OutputFilePath:     "",
		outputPathParts:    []string{"lib", "src", "views", "splash"},
		EmbeddedFileReader: FileReader{}.New(EmbeddedFsPaths.SplashView),
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

func StartCtlr() *File {
	return &File{
		OutputString:       "",
		OutputFilename:     "c_start.dart",
		OutputFilePath:     "",
		outputPathParts:    []string{"lib", "src", "views", "start"},
		EmbeddedFileReader: FileReader{}.New(EmbeddedFsPaths.StartView),
		Replacements:       map[string]string{},
		Mu:                 sync.Mutex{},
	}
}

func View(projectName string, dirname string, filename string, classname string, controllerFilename string, root string) *File {
	return &File{
		OutputString:       "",
		OutputFilename:     filename,
		OutputFilePath:     "",
		outputPathParts:    []string{"lib", "src", "views", dirname},
		EmbeddedFileReader: FileReader{}.New(EmbeddedFsPaths.ViewView),
		Replacements: map[string]string{
			"PROJECT_NAME": projectName,
      "CONTROLLER_FILE_NAME": controllerFilename,
      "CLASS_NAME": classname
		},
		Mu: sync.Mutex{},
	}
}

func ViewCtlr(dirname string, filename string, classname string, viewFilename string) *File {
	return &File{
		OutputString:       "",
		OutputFilename:     filename,
		OutputFilePath:     "",
		outputPathParts:    []string{"lib", "src", "views", dirname},
		EmbeddedFileReader: FileReader{}.New(EmbeddedFsPaths.ViewCtlr),
		Replacements:       map[string]string{
      "VIEW_FILENAME": viewFilename,
      "CLASS_NAME": classname,
    },
		Mu:                 sync.Mutex{},
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
