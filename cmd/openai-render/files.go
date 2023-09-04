package main

import (
	"regexp"
)

// outputFiles is a list of code blocks extracted as files
type outputFiles []*outputFile

// outputFile is a code block extracted as a file with a relative path and content
type outputFile struct {
	path    string
	content string
}

// Filenames just returns the filenames of the output files as []string
func (o outputFiles) Filenames() []string {
	res := []string{}
	for _, file := range o {
		res = append(res, file.path)
	}
	return res
}

// filesFromChatCompletion extracts files from a ChatCompletionResponse code blocks.
// If lang is specified, it will be used to determine which code blocks to extract as
// files. If lang is empty, all code blocks will be extracted to files
func filesFromChatCompletion(resp string, lang string) outputFiles {
	filesGroup := outputFiles{}
	// you can check this regex at
	// https: //regex101.com/r/Rs5m3T/1
	codeBlockRegex := regexp.MustCompile(`\x60([^\x60]+)\x60:\W\x60\x60\x60go([^\x60]+)\x60\x60\x60`)
	for _, match := range codeBlockRegex.FindAllStringSubmatch(resp, -1) {
		fPath := match[1]
		fContent := match[2]
		filesGroup = append(filesGroup, &outputFile{
			path:    fPath,
			content: fContent,
		})
	}
	return filesGroup
}
