package project

import (
	"encoding/json"
	"os"
	"raijin/internal/config"
	"raijin/internal/generator"
	"raijin/internal/shell"
	"slices"
)

func CreateScaffold(name string, args []string) error {
	appDirs := config.GetAppDirs(&name)

	frontendDir := appDirs.FrontendSrc

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

	os.WriteFile(appDirs.EntryFile, []byte(generator.GenerateExampleEntryFile()), config.FileMode)

	return nil
}
