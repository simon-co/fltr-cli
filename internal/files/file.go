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

type FileList []*File

func (self FileList) InstantiateAll(projectPath string) error {
	for _, val := range self {
		if err := val.Instantiate(projectPath); err != nil {
			return apperr.Parse(err)
		}
	}
	return nil
}
