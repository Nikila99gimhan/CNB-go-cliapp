package main

import (
	"errors"
	"os"
	"os/exec"
)

type Builder struct {}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) SelectForLanguage(language string) string {
	switch language {
	case "Java":
		return "paketobuildpacks/builder-jammy-base"
	case "Python":
		return "gcr.io/buildpacks/builder:v1"
	case "Go":
		return "paketobuildpacks/builder-jammy-base"
	case "Node.js":
		return "paketobuildpacks/builder-jammy-buildpackless-static"
	case ".NET":
		return "gcr.io/buildpacks/builder:v1"
	default:
		return ""
	}
}

func (b *Builder) Build(builder, imageName, path string) error {
	cmd := exec.Command("pack", "build", imageName, "--builder", builder, "--path", path, "--pull-policy", "if-not-present")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return errors.New("Error during build process: " + err.Error())
	}
	return nil
}
