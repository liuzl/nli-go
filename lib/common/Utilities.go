package common

import (
	"io/ioutil"
	"os"
	"path"
	"runtime"
)

func IntArrayContains(haystack []int, needle int) bool {
	for _, value := range haystack {
		if needle == value {
			return true
		}
	}
	return false
}

func IntArrayDeduplicate(array []int) []int {
	newArray := []int{}
	previous := map[int]bool{}

	for _, value := range array {
		_, found := previous[value]
		if !found {
			previous[value] = true
			newArray = append(newArray, value)
		}
	}

	return newArray
}

func StringArrayDeduplicate(array []string) []string {
	newArray := []string{}
	uniques := map[string]bool{}

	for _, value := range array {
		_, found := uniques[value]
		if !found {
			uniques[value] = true
			newArray = append(newArray, value)
		}
	}

	return newArray
}

// Returns the directory of the file calling this function
// https://coderwall.com/p/_fmbug/go-get-path-to-current-file
func Dir() string {

	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

func ReadFile(path string) (string, error) {

	source := ""

	bytes, err := ioutil.ReadFile(path)
	if err == nil {
		source = string(bytes)
	}

	return source, err
}

// If path is an absolute path, returns path
// Otherwise, adds path to baseDir to create an absolute path
func AbsolutePath(baseDir string, path string) string {

	absolutePath := path

	if len(path) > 0 && path[0] != os.PathSeparator {
		absolutePath = baseDir + string(os.PathSeparator) + path
	}

	return absolutePath
}
