package command

import (
	"errors"
	"fmt"
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
			expected: "TestFilenameToClass",
		},
		{
			name:     "short",
			input:    "v_test.dart",
			error:    nil,
			expected: "Test",
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
