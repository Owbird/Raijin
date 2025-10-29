package config

import (
	"os"
	"path/filepath"
)

type AppConfig struct {
	Name string `json:"name"`
}

type AppDirs struct {
	Wd           string
	Frontend     string
	RaijinConfig string
	EntryFile    string
	ActionsDir  string
}

const (
	FileMode = 0755
)

var (
	ViteCmd       = []string{"pnpm", "create", "vite", "."}
	InstallCmd    = []string{"pnpm", "install"}
	FrontedDevCmd = []string{"pnpm", "dev"}
)

func GetAppDirs(path *string) AppDirs {
	fullPath := ""

	if path == nil || *path == "" {
		fullPath, _ = os.Getwd()
	} else {
		fullPath = *path
	}

	frontendDir := filepath.Join(fullPath, "frontend")
	raijinConfig := filepath.Join(fullPath, "raijin.json")
	entryFile := filepath.Join(fullPath, "main.go")
	actionsDir := filepath.Join(frontendDir, "src", "actions")

	return AppDirs{
		Wd:           fullPath,
		Frontend:     frontendDir,
		RaijinConfig: raijinConfig,
		EntryFile:    entryFile,
		ActionsDir:  actionsDir,
	}
}
