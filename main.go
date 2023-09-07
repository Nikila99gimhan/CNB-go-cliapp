package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	Green = "\033[32m"
	Bold  = "\033[1m"
	Reset = "\033[0m"
)

func main() {
	printLogo()

	if len(os.Args) < 2 {
		fmt.Println("Usage: <repository_url>")
		return
	}

	repoManager := NewRepoManager()
	languageDetector := NewLanguageDetector()
	builder := NewBuilder()
	// reporter := NewReporter()

	repoURL := os.Args[1]
	repoName := repoManager.GetRepoName(repoURL)

	if !repoManager.RepoExists(repoName) {
		printGreenBold("\nCloning repository...\n")
		err := repoManager.CloneRepo(repoURL)
		if err != nil {
			fmt.Println("Error while cloning:", err)
			return
		}
	} else {
		printGreenBold("\nRepository already exists. Using the existing one...\n")
	}

	language := languageDetector.Detect(repoName)
	printGreenBold("\nDetected language:", language)

	printGreenBold("\nSelecting builder for", language, "...\n")
	selectedBuilder := builder.SelectForLanguage(language)
	if selectedBuilder == "" {
		fmt.Println("Unsupported language or framework:", language)
		return
	}

	fmt.Print("\nEnter the name for your image: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	imageName := scanner.Text()
	if imageName == "" {
		fmt.Println("Image name cannot be empty!")
		return
	}

	printGreenBold("\nBuilding and containerizing application...\n")
	err := builder.Build(selectedBuilder, imageName, repoName)
	if err != nil {
		fmt.Println("Error during build:", err)
		return
	}

	printGreenBold("\nImage", imageName, "built successfully!")
	// reporter.Generate(language, repoName)
}

func printLogo() {
	logo := `
  _____ ______       ___    ___ ________  ___       ___  ________  ________  ________   
|\   _ \  _   \    |\  \  /  /|\   ____\|\  \     |\  \|\   __  \|\   __  \|\   __  \  
\ \  \\\__\ \  \   \ \  \/  / | \  \___|\ \  \    \ \  \ \  \|\  \ \  \|\  \ \  \|\  \ 
 \ \  \\|__| \  \   \ \    / / \ \  \    \ \  \    \ \  \ \   __  \ \   ____\ \   ____\
  \ \  \    \ \  \   \/  /  /   \ \  \____\ \  \____\ \  \ \  \ \  \ \  \___|\ \  \___|
   \ \__\    \ \__\__/  / /      \ \_______\ \_______\ \__\ \__\ \__\ \__\    \ \__\   
    \|__|     \|__|\___/ /        \|_______|\|_______|\|__|\|__|\|__|\|__|     \|__|   
                  \|___|/                                                            
`
	fmt.Println(Green + Bold + logo + Reset)
}

func printGreenBold(values ...interface{}) {
	fmt.Println(Green + Bold + fmt.Sprint(values...) + Reset)
}
