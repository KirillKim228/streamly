// Package config загружает настройки приложения через cleanenv.
package config

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const envFile = ".env"

// Config — настройки приложения.
type Config struct {
	Addr              string        `env:"ADDR" env-default:":8080" env-description:"адрес сервера"`
	MediaDir          string        `env:"MEDIA_DIR" env-default:"media" env-description:"папка с видео и аудио"`
	ReadHeaderTimeout time.Duration `env:"READ_HEADER_TIMEOUT" env-default:"5s" env-description:"таймаут чтения заголовков"`
	ShutdownTimeout   time.Duration `env:"SHUTDOWN_TIMEOUT" env-default:"5s" env-description:"ожидание запросов при остановке"`
}

// Load загружает настройки из файла .env (если он есть) и переменных окружения.
// Переменные окружения главнее файла, значения по умолчанию заданы тегами env-default.
func Load() (Config, error) {
	var cfg Config

	_, err := os.Stat(envFile)
	if errors.Is(err, fs.ErrNotExist) {
		if err := cleanenv.ReadEnv(&cfg); err != nil {
			return Config{}, fmt.Errorf("чтение переменных окружения: %w", err)
		}
		return cfg, nil
	}
	if err != nil {
		return Config{}, fmt.Errorf("проверка %s: %w", envFile, err)
	}

	if err := cleanenv.ReadConfig(envFile, &cfg); err != nil {
		return Config{}, fmt.Errorf("чтение %s: %w", envFile, err)
	}
	return cfg, nil
}
