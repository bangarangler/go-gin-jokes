// +build mage

package main

import (
	"fmt"
	"os"
	// "os/exec"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
	"github.com/magefile/mage/sh"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

// A build step that requires additional params, or platform specific steps for example
func Build() error {
	mg.Deps(InstallDeps)
	fmt.Println("Building...")
	// cmd := exec.Command("go", "build", "-o", "MyApp", ".")
	return sh.Run("go", "build", "-o", "go-gin-jokes")
	// return cmd.Run()
}

func Start() error {
	mg.Deps(Build)
	fmt.Println("Running go-gin-jokes...")
	// cmd := exec.Command("go", "build", "-o", "MyApp", ".")
	return sh.Run("./go-gin-jokes")
}

// A custom install step if you need your bin someplace other than go/bin
func Install() error {
	mg.Deps(Build)
	fmt.Println("Installing...")
	return os.Rename("./go-gin-jokes", "/usr/bin/go-gin-jokes")
}

// Manage your deps, or running package managers.
func InstallDeps() error {
	fmt.Println("Installing Deps...")
	// cmd := exec.Command("go", "get", "github.com/stretchr/piglatin")
	// return cmd.Run()
	return sh.Run("go", "mod", "tidy")
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll("./go-gin-jokes")
}
