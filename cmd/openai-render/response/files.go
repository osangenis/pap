package response

import (
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

// FilesFromChat extracts files from a ChatCompletionResponse code blocks.
// If lang is specified, it will be used to determine which code blocks to extract as
// files. If lang is empty, all code blocks will be extracted to files
func FilesFromChat(resp string, lang string) OutputFiles {
	filesGroup := OutputFiles{}
	// you can check this regex at
	// https: //regex101.com/r/Rs5m3T/1
	codeBlockRegex := regexp.MustCompile(`\x60([^\x60]+)\x60:\W\x60\x60\x60(\w+)([^\x60]+)\x60\x60\x60`)
	for _, match := range codeBlockRegex.FindAllStringSubmatch(resp, -1) {
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
	return filesGroup
}
