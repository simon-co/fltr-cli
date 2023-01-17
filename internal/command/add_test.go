package command

import (
	"errors"
	"fmt"
	"path/filepath"
	"sync"
	"testing"

	"github.com/simon-co/fltr-cli/internal/apperr"
)

func TestComponentTypesToMap(t *testing.T) {
	m := ComponentTypes{}.ToMap()

	for k := range m {
		fmt.Println(k)
	}
}

func TestComponentMapToList(t *testing.T) {
	j := ComponentTypes{}.ToMap().ToList()

	for _, v := range j {
		fmt.Println(v)
	}
}

func TestViewFilenameFromDir(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		error    error
		expected viewFilename
	}
	testCases := []testCase{
		{
			name:     "success",
			input:    "test_this",
			error:    nil,
			expected: "v_test_this.dart",
		},
	}
	var wg sync.WaitGroup
	for _, tc := range testCases {
		wg.Add(1)
		go func(tc testCase) {
			defer wg.Done()
			t.Run(tc.name, func(t *testing.T) {
				result := viewFilename("").fromDirName(tc.input)
				if result != tc.expected {
					t.Errorf("expected: %s received: %s\n", tc.expected, result)
				}
			})
		}(tc)
	}
	wg.Wait()
}

func TestViewNameToClassName(t *testing.T) {
	type testCase struct {
		name     string
		input    viewFilename
		error    error
		expected Classname
	}
	testCases := []testCase{
		{
			name:     "long success",
			input:    "v_test_filename_to_class.dart",
			error:    nil,
			expected: "TestFilenameToClassView",
		},
		{
			name:     "short",
			input:    "v_test.dart",
			error:    nil,
			expected: "TestView",
		},
		{
			name:     "error filetype",
			input:    "v_test.js",
			error:    apperr.Parse(errors.New("invalid view filename")),
			expected: "",
		},
		{
			name:     "error file name",
			input:    "vfilename",
			error:    apperr.Parse(errors.New("invalid view filename")),
			expected: "",
		},
	}
	var wg sync.WaitGroup
	for _, tc := range testCases {
		wg.Add(1)
		go func(tc testCase) {
			defer wg.Done()
			t.Run(tc.name, func(t *testing.T) {
				result, err := tc.input.toClassName()
				if err != nil {
					if tc.error == nil {
						t.Errorf("%s unexpected error\nexpected: nil\nreceived: %s\n", tc.name, err)
						return
					}
					if ok := errors.Is(err, tc.error); !ok {
						t.Errorf("%s unexpected error\nexpected: %s\nreceived: %s\n", tc.name, tc.error, err)
						return
					}
				}
				if result != tc.expected {
					t.Errorf("%s unexpected result\nexpected: %s\nreceived: %s\n", tc.name, tc.expected, result)
				}
			})
		}(tc)
	}
	wg.Wait()
}

func TestValidateConfirm(t *testing.T) {
	type testCase struct {
		name  string
		input string
		error error
	}
	testCases := []testCase{
		{
			name:  "Failure invalid start char",
			input: "eyes",
			error: ErrInvalidConfirmInput,
		},
		{
			name:  "Failure ye",
			input: "ye",
			error: ErrInvalidConfirmInput,
		},
		{
			name:  "Success n",
			input: "n",
			error: nil,
		},
		{
			name:  "Success no",
			input: "no",
			error: nil,
		},
		{
			name:  "Success y",
			input: "y",
			error: nil,
		},
		{
			name:  "Success yes",
			input: "yes",
			error: nil,
		},
	}
	var wg sync.WaitGroup
	for _, tc := range testCases {
		wg.Add(1)
		go func(tc testCase) {
			defer wg.Done()
			t.Run(tc.name, func(t *testing.T) {
				err := validateConfirm(tc.input)
				if err != nil {
					if tc.error == nil {
						t.Errorf("%s Unexpected Error: \nExpected: nil\nReceived: %s\n", tc.name, err)
					}
					if !errors.Is(err, tc.error) {
						t.Errorf("%s Unexpected Error: \nExpected: %s\nReceived: %s\n", tc.name, tc.error, err)
					}
				}
			})
		}(tc)
	}
	wg.Wait()
}

func TestConfirmStrToBool(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected bool
		error    error
	}
	testCases := []testCase{
		{
			name:     "Failure invalid start char",
			input:    "eyes",
			expected: false,
			error:    ErrInvalidConfirmInput,
		},
		{
			name:     "Failure ye",
			input:    "ye",
			expected: false,
			error:    ErrInvalidConfirmInput,
		},
		{
			name:     "Success n",
			input:    "n",
			expected: false,
			error:    nil,
		},
		{
			name:     "Success no",
			input:    "no",
			expected: false,
			error:    nil,
		},
		{
			name:     "Success y",
			input:    "y",
			expected: true,
			error:    nil,
		},
		{
			name:     "Success yes",
			input:    "yes",
			expected: true,
			error:    nil,
		},
	}
	var wg sync.WaitGroup
	for _, tc := range testCases {
		wg.Add(1)
		go func(tc testCase) {
			defer wg.Done()
			t.Run(tc.name, func(t *testing.T) {
				result, err := confirmStrToBool(tc.input)
				if err != nil {
					if tc.error == nil {
						t.Errorf("%s Unexpected Error: \nExpected: nil\nReceived: %s\n", tc.name, err)
					}
					if !errors.Is(err, tc.error) {
						t.Errorf("%s Unexpected Error: \nExpected: %s\nReceived: %s\n", tc.name, tc.error, err)
					}
				}
				if result != tc.expected {
					t.Errorf("%s Unexpected Result:\nExpected: %t\nReceived: %t", tc.name, tc.expected, result)
				}
			})
		}(tc)
	}
	wg.Wait()
}

func TestNavigatorFileFromPath(t *testing.T) {
	type testCase struct {
		name     string
		input    string
    expected *navigatorFile
		error    error
	}
	testCases := []testCase{{
		name:  "success",
		input: filepath.Join("test", "this", "here", "n_this_test.dart"),
    expected: &navigatorFile{
			fileName:  "n_this_test.dart",
			className: "ThisTestRouteNavigator",
			path:      filepath.Join("test", "this", "here", "n_this_test.dart"),
		},
		error: nil,
	},
   
  }
  var wg sync.WaitGroup
  for _, tc := range testCases {
    wg.Add(1)
    go func(tc testCase){
      defer wg.Done()
      t.Run(tc.name, func(t *testing.T) {
        nf, err := navigatorFile{}.fromPathname(tc.input)
        if err != nil {
          if tc.error == nil {
            t.Errorf("%s Unexpected Error\nExpected: nil\nReceived: %s", tc.name, err)
            return
          } 
          if !errors.Is(err, tc.error){
            t.Errorf("%s Unexpected Error\nExpected: %s\nReceived: %s", tc.name, tc.error, err)
            return
          }
          return
        }
        if *nf != *tc.expected {
          t.Errorf("%s Unexpected Result\nExpected: %+v\nReceived: %+v", tc.name, tc.expected, nf)
          return
        }
      })
    }(tc)
    wg.Wait()
  }
}
