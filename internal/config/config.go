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
	FrontendSrc  string
	FrontendDist string
	RaijinConfig string
	EntryFile    string
	ActionsDir   string
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

	frontendSrcDir := filepath.Join(fullPath, "frontend")
	frontendDistDir := filepath.Join(frontendSrcDir, "dist")
	raijinConfig := filepath.Join(fullPath, "raijin.json")
	entryFile := filepath.Join(fullPath, "main.go")
	actionsDir := filepath.Join(frontendSrcDir, "src", "actions")

	return AppDirs{
		Wd:           fullPath,
		FrontendSrc:  frontendSrcDir,
		FrontendDist: frontendDistDir,
		RaijinConfig: raijinConfig,
		EntryFile:    entryFile,
		ActionsDir:   actionsDir,
	}
}
