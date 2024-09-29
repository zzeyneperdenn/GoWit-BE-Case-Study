package appbase

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	LogLevel      string `env:"LOG_LEVEL" env-default:"debug"`
	ServerAddress string `env:"SERVER_ADDRESS" env-default:"0.0.0.0:8080"`
	ServerTimeout int64  `env:"SERVER_TIMEOUT" env-default:"120"`

	// Application Database
	DatabaseName     string `env:"DATABASE_NAME" env-required:"true"`
	DatabasePassword string `env:"DATABASE_PASSWORD" env-required:"true"`
	DatabasePort     string `env:"DATABASE_PORT" env-default:"5432"`
	DatabaseUser     string `env:"DATABASE_USERNAME" env-required:"true"`
	DatabaseHost     string `env:"DATABASE_HOST" env-required:"true"`
}

type goEnv struct {
	GoMod string `json:"GOMOD"`
}

func (c *Config) HTTPTimeoutDuration() time.Duration {
	return time.Duration(c.ServerTimeout) * time.Second
}

func LoadConfig() (*Config, error) {
	c := new(Config)

	err := loadConfig(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func loadConfig(cfg interface{}) error {
	envFilePath, envFileExists := initLocal()

	var err error

	if envFileExists {
		err = cleanenv.ReadConfig(envFilePath, cfg)
	} else {
		err = cleanenv.ReadEnv(cfg)
	}

	if err != nil {
		return err
	}

	return nil
}

func initLocal() (string, bool) {
	if !gitExists() {
		return "", false
	}

	modRoot := getModuleRoot()
	envFilePath := fmt.Sprintf("%s/.env", modRoot)

	if envFileExists(envFilePath) {
		return envFilePath, true
	}

	return "", false
}

func envFileExists(envFilePath string) bool {
	_, err := os.Stat(envFilePath)

	return err == nil
}

func gitExists() bool {
	_, err := exec.LookPath("git")

	return err == nil
}

func getModuleRoot() string {
	goEnvRaw, err := exec.Command("go", "env", "-json").Output()
	if err != nil {
		panic(fmt.Errorf("go env command failed: %w", err))
	}

	env := new(goEnv)

	err = json.Unmarshal(goEnvRaw, env)
	if err != nil {
		panic(fmt.Errorf("go mod unmarshalling failed: %w", err))
	}

	return strings.TrimSuffix(env.GoMod, "/go.mod")
}
