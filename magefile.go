//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"

	_ "github.com/joho/godotenv/autoload"
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
	cmd := exec.Command("go", "build", "-o", "MyApp", ".")
	return cmd.Run()
}

// A custom install step if you need your bin someplace other than go/bin
func Install() error {
	mg.Deps(Build)
	fmt.Println("Installing...")
	return os.Rename("./MyApp", "/usr/bin/MyApp")
}

// Manage your deps, or running package managers.
func InstallDeps() error {
	fmt.Println("Installing Deps...")
	cmd := exec.Command("go", "get", "github.com/stretchr/piglatin")
	return cmd.Run()
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll("MyApp")
}

// MySQL マイグレーションコマンド
type Migrate mg.Namespace

func (Migrate) Up() error {
	return execMigrate("up")
}

func (Migrate) Down() error {
	return execMigrate("down")
}

func (Migrate) Force(version string) error {
	return execMigrate("force", version)
}

func execMigrate(args ...string) error {
	args = append([]string{
		"-source", "file://infrastructure/mysql/migrations",
		"-database", fmt.Sprintf(
			"mysql://%s:%s@tcp(%s:%s)/%s",
			os.Getenv("MYSQL_ROOT_USER"),
			os.Getenv("MYSQL_ROOT_PASSWORD"),
			os.Getenv("MYSQL_HOST"),
			os.Getenv("MYSQL_PORT"),
			os.Getenv("MYSQL_DATABASE"),
		),
	}, args...)
	if err := sh.RunV("migrate", args...); err != nil {
		return err
	}
	return nil
}
