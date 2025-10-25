package project

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"raijin/internal/config"
	"raijin/internal/shell"
	"slices"
)

func CreateScaffold(name string, args []string) error {
	mainDir := name

	frontendDir := filepath.Join(mainDir, "frontend")
	backendDir := filepath.Join(mainDir, "backend")
	raijinJsonPath := filepath.Join(mainDir, "raijin.json")

	os.MkdirAll(frontendDir, config.FileMode)
	os.MkdirAll(backendDir, config.FileMode)

	appConfig := config.AppConfig{Name: name}

	raijinJson, err := json.Marshal(appConfig)
	if err != nil {
		return err
	}

	os.WriteFile(raijinJsonPath, raijinJson, config.FileMode)

	createOutput, err := shell.Run(shell.ShellCmd{
		Cmd:  "npx",
		Dir:  frontendDir,
		Args: slices.Concat(config.ViteCmd, args),
	})
	if err != nil {
		return err
	}

	log.Println(string(createOutput))

	installOutput, err := shell.Run(shell.ShellCmd{
		Cmd:  "npx",
		Dir:  frontendDir,
		Args: config.InstallCmd,
	})
	if err != nil {
		return err
	}

	log.Println(string(installOutput))

	return nil
}
