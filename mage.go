//go:build mage

package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/magefile/mage/sh"
)

// Install the binaries.
func Build() error {
	return sh.Run(
		"go",
		"install",
		"./...")
}

// Performs a go mod tidy check
func TidyCheck() error {
	if err := sh.RunV("go", "mod", "tidy"); err != nil {
		return err
	}
	return sh.RunV("git", "diff", "--exit-code", "--", "go.mod", "go.sum")
}

// Performs a lint analysis on the Go code
func Lint() error {
	return sh.RunV("golangci-lint", "run")
}

// Performs a Vulnerability check on the Go code
func Vuln() error {
	return sh.RunV("govulncheck", "./...")
}

// Run the tests.
func Test() error {
	return sh.RunV(
		"go",
		"test",
		"./...")
}

// Cleans the binaries.
func Clean() error {
	return sh.Run("go", "clean", "-i", "./...")
}

// Updates the module dependencies.
func Update() error {
	return sh.Run("go", "get", "-u", "./...")
}

// Starts the Go documentation server.
func Doc() error {
	return sh.Run("godoc", "-http=:6060")
}

// Install all tool dependencies
func ToolInstall() error {
	tools, err := findTools()
	if err != nil {
		return err
	}

	for _, t := range tools {
		if err := sh.Run("go", "install", t); err != nil {
			return err
		}
	}

	return nil
}

func findTools() ([]string, error) {
	f, err := os.Open("tools.go")
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var tools []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " \t")
		if strings.HasPrefix(line, "_") {
			tokens := strings.Split(line, " ")
			tools = append(tools, strings.Trim(tokens[1], "\""))
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return tools, nil
}
