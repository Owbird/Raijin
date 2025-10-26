package project

import (
	"encoding/json"
	"os"
	"raijin/internal/config"
	"raijin/internal/shell"
	"slices"
)

func CreateScaffold(name string, args []string) error {
	appDirs := config.GetAppDirs(&name)

	frontendDir := appDirs.Frontend

	os.MkdirAll(frontendDir, config.FileMode)

	appConfig := config.AppConfig{Name: name}

	raijinJson, err := json.Marshal(appConfig)
	if err != nil {
		return err
	}

	os.WriteFile(appDirs.RaijinConfig, raijinJson, config.FileMode)

	_, err = shell.Run(shell.ShellCmd{
		Cmd:  "npx",
		Dir:  frontendDir,
		Args: slices.Concat(config.ViteCmd, args),
	})
	if err != nil {
		return err
	}

	_, err = shell.Run(shell.ShellCmd{
		Cmd:  "npx",
		Dir:  frontendDir,
		Args: config.InstallCmd,
	})
	if err != nil {
		return err
	}

	os.WriteFile(appDirs.EntryFile, []byte(`package main

import "raijin/pkg/app"

type AuthActions struct{}

func NewAuthActions() *AuthActions {
	return &AuthActions{}
}

func (aa *AuthActions) Login() { println("Logging in") }

func main() {
	a := app.NewApp()

	a.Bind(NewAuthActions())

	a.Run()
}`), config.FileMode)

	return nil
}
