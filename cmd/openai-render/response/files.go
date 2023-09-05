package response

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

// OutputFiles is a list of code blocks extracted as files
type OutputFiles []*OutputFile

// OutputFile is a code block extracted as a file with a relative path and content
type OutputFile struct {
	Path    string
	Content string
}

// Filenames just returns the filenames of the output files as []string
func (o OutputFiles) Filenames() []string {
	res := []string{}
	for _, file := range o {
		res = append(res, file.Path)
	}
	return res
}

// Write writes the output files to the filesystem, creating the folders if required
func (o OutputFiles) Write(baseDir string) error {
	for _, file := range o {
		path := filepath.Join(baseDir, file.Path)
		dir := filepath.Dir(path)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("error creating dir %v : %v", dir, err)
		}
		if err := os.WriteFile(path, []byte(file.Content), 0644); err != nil {
			return fmt.Errorf("error writing file %v : %v", path, err)
		}
	}
	return nil
}

// FilesFromChat extracts files from a ChatCompletionResponse code blocks.
// If lang is specified, it will be used to determine which code blocks to extract as
// files. If lang is empty, all code blocks will be extracted to files
func FilesFromChat(resp string, lang string) OutputFiles {
	filesGroup := OutputFiles{}
	for _, regex := range knownRegexForCodeBlocks() {
		for _, match := range regex.FindAllStringSubmatch(resp, -1) {
			fPath := match[1]
			fLang := match[2]
			fContent := match[3]
			if lang == "" || fLang == lang {
				filesGroup = append(filesGroup, &OutputFile{
					Path:    fPath,
					Content: fContent,
				})
			}
		}
	}

	return filesGroup
}

// knownRegexForCodeBlocks returns a list of regex that can be used for extracting
// code blocks from a string (a response from the OpenAI API). The responses sometimes
// use one format or another, so we need to try different regexes to extract the code
// The regex must return 3 capturing groups:
// 1. The relative path of the file
// 2. The language of the code block
// 3. The content of the code block
// Additionally, not 2 regex should match the same code block
func knownRegexForCodeBlocks() []*regexp.Regexp {
	// you can check this regex at https: //regex101.com/r/Rs5m3T/1
	longFormat := regexp.MustCompile(`\x60([^\x60]+)\x60:\W\x60\x60\x60(\w+)([^\x60]+)\x60\x60\x60`)
	// you can check this regex at https://regex101.com/r/ndDhvK/2
	shortFormat := regexp.MustCompile(`\W+[^:]+\.[^\n]+:\W+\x60\x60\x60go\W+//\W([^\.]+.([^\n]+))([^\x60]+)\x60\x60\x60`)

	return []*regexp.Regexp{
		longFormat,
		shortFormat,
	}
}
