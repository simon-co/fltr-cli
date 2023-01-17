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

// searches for a pubspec.yaml using find.EntityPathBack.  If this file is found
// returns a path to the fltr views directory in relation to this path.
// returns find.ErrNotFound if views directory nos found
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

func RouteNavigatorsPath() (string, error) {
	projectPath, err := ProjectDir()
	if err != nil {
		return "", apperr.Parse(err)
	}
	navigationpath := filepath.Join(projectPath, "lib", "src", "routing", "route_navigators")
	_, err = os.Stat(navigationpath)
	if err != nil {
		return "", apperr.Parse(ErrNotFound)
	}
	return navigationpath, nil
}

func AllRouteNavigatorPaths() ([]string, error) {
	navigators := []string{}
	navigatorsPath, err := RouteNavigatorsPath()
	if err != nil {
		return nil, apperr.Parse(err)
	}
	err = filepath.WalkDir(navigatorsPath, func(path string, entry os.DirEntry, err error) error {
		if !entry.IsDir() {
			navigators = append(navigators, navigatorsPath + string(os.PathSeparator) + entry.Name())
		}
    return nil
	})
	return navigators, nil
}

func ProjectName() (string, error) {
	projectPath, err := ProjectDir()
	if err != nil {
		return "", apperr.Parse(err)
	}
	return filepath.Base(projectPath), nil
}
func RouterPath() (string, error) {
	projectPath, err := ProjectDir()
	if err != nil {
		return "", apperr.Parse(err)
	}
	routerPath := filepath.Join(projectPath, "lib", "src", "routing", "router.dart")
	return routerPath, nil
}

func AllViewDirNames() ([]string, error) {
	viewsDirPath, err := ViewsDirPath()
	viewFileNames := []string{}
	if err != nil {
		return viewFileNames, apperr.Parse(err)
	}
	err = filepath.WalkDir(viewsDirPath, func(path string, entry os.DirEntry, err error) error {
		if entry.IsDir() {
			if entry.Name() == "views" {
				return nil
			} else {
				viewFileNames = append(viewFileNames, entry.Name())
				return filepath.SkipDir
			}
		}
		return nil
	})
	return viewFileNames, nil
}
