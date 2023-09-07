package main

import (
	"os"
)

type LanguageDetector struct{}

func NewLanguageDetector() *LanguageDetector {
	return &LanguageDetector{}
}

func (ld *LanguageDetector) Detect(directory string) string {
	// Check for existence of certain files or directories
	if fileExists(directory + "/pom.xml") {
		return "Java"
	}
	if fileExists(directory + "/requirements.txt") {
		return "Python"
	}
	if fileExists(directory + "/go.mod") {
		return "Go"
	}
	if fileExists(directory + "/package.json") {
		return "Node.js"
	}
	if fileExists(directory + "/.csproj") {
		return ".NET"
	}
	// Extend with other detections as necessary
	return "Unknown"
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
