package client

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App `yaml:"app"`
		// HTTP `yaml:"http"`
		Log `yaml:"logger"`
		// PG   `yaml:"postgres"`
	}

	App struct {
		Name    string `yaml:"name"    env:"APP_NAME"`
		Version string `yaml:"version" env:"APP_VERSION"`
	}

	// HTTP struct {
	// 	Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	// }

	Log struct {
		Level string `yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// PG struct {
	// 	PoolMax int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
	// 	URL     string `env-required:"true"                 env:"PG_URL"`
	// }
)

var (
	currentConfig Config    //nolint:gochecknoglobals // pattern singleton
	once          sync.Once //nolint:gochecknoglobals // pattern singleton
)

// LoadConfig returns app config.
func LoadConfig() Config {
	var err error

	once.Do(func() {
		err = cleanenv.ReadConfig("./config/client/config.yml", &currentConfig)
		if err != nil {
			log.Panicln("LoadConfig - %w", err)
		}

		err = cleanenv.ReadEnv(&currentConfig)
		if err != nil {
			log.Panicln("LoadConfig - %w", err)
		}
	})

	return currentConfig
}
