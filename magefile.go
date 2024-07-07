//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
	"github.com/magefile/mage/sh"
)

func run(cmd string, args ...string) {
	if err := sh.RunV(cmd, args...); err != nil {
		panic(err)
	}
}

func runWithEnv(env map[string]string, cmd string, args ...string) {
	if err := sh.RunWithV(env, cmd, args...); err != nil {
		panic(err)
	}
}

// パッケージインストールコマンド
func Install() {
	run("go", "install", "-tags", "mysql", "github.com/golang-migrate/migrate/v4/cmd/migrate@v4.17.1")
	run("go", "install", "github.com/volatiletech/sqlboiler/v4@v4.16.2")
	run("go", "install", "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@v4.16.2")
	run("go", "mod", "tidy")
}

// MySQL マイグレーションコマンド
type Migrate mg.Namespace

func (Migrate) Up() {
	mg.Deps(Install)
	migrate("up")
}

func (Migrate) Down() {
	mg.Deps(Install)
	migrate("down")
}

func (Migrate) Force(version string) {
	mg.Deps(Install)
	migrate("force", version)
}

func migrate(args ...string) {
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
	run("migrate", args...)
}

// ファイル・コードの自動生成コマンド
type Generate mg.Namespace

func (Generate) Sqlboiler() {
	env := map[string]string{
		"MYSQL_USER":    os.Getenv("MYSQL_ROOT_USER"),
		"MYSQL_PASS":    os.Getenv("MYSQL_ROOT_PASSWORD"),
		"MYSQL_HOST":    os.Getenv("MYSQL_HOST"),
		"MYSQL_PORT":    os.Getenv("MYSQL_PORT"),
		"MYSQL_DBNAME":  os.Getenv("MYSQL_DATABASE"),
		"MYSQL_SSLMODE": "false",
	}
	runWithEnv(env, "sqlboiler", "mysql", "-o", "infrastructure/mysql/tables", "-p", "tables", "--no-tests", "--wipe")
}
