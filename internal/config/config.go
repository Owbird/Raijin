package config

type AppConfig struct {
	Name string `json:"name"`
}

const (
	FileMode = 0755
)

var (
	ViteCmd    = []string{"pnpm", "create", "vite", "."}
	InstallCmd = []string{"pnpm", "install"}
)
