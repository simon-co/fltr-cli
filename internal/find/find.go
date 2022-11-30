package find

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/simon-co/fltr-cli/internal/apperr"
)

var (
  ErrNotFound = errors.New("can't find entity of supplied name")
)

// searches the current working directory for an entity matching the supllied entityName.
// If no match is found searches back through each directory on the working directory path.
// Once on match is found returns a full path to the entity.
// if no entity is found returns an find.NotFound
func EntityPathBack(entityName string) (string, error) {
	currentPath, _ := os.Getwd()
	for {
		searchPath := filepath.Join(currentPath, entityName)
		_, err := os.Stat(searchPath)
		if err != nil {
			newPath := filepath.Dir(currentPath)
			if newPath == currentPath {
				return "", apperr.Parse(ErrNotFound)
			}
			currentPath = newPath
			continue
		}
		return searchPath, nil
	}
}

func ProjectDir() (string, error) {
  pubspecYamlPath, err := EntityPathBack("pubspec.yaml")  
  if err != nil {
    return "", apperr.Parse(err)
  }
  return filepath.Dir(pubspecYamlPath), nil
}

//searches for a pubspec.yaml using find.EntityPathBack.  If this file is found
//returns a path to the fltr views directory in relation to this path.
//returns find.ErrNotFound if views directory nos found
func ViewsDirPath() (string, error) {
	projectPath, err := ProjectDir()
	if err != nil {
		return "", apperr.Parse(err)
	}
	viewsPath := filepath.Join(projectPath, "lib", "src", "views")
	_, err = os.Stat(viewsPath)
	if err != nil {
		return "", apperr.Parse(ErrNotFound)
	}
	return viewsPath, nil
}

func DialogsDirPath() (string, error) {
	projectPath, err := ProjectDir()
	if err != nil {
		return "", apperr.Parse(err)
	}
	dialogsPath := filepath.Join(projectPath, "lib", "src", "dialogs")
	_, err = os.Stat(dialogsPath)
	if err != nil {
		return "", apperr.Parse(ErrNotFound)
	}
	return dialogsPath, nil
}

func ProjectName() (string, error){
	projectPath, err := ProjectDir()
	if err != nil {
		return "", apperr.Parse(err)
	}
  return filepath.Base(projectPath), nil
}
