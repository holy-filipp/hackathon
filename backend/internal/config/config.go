package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

// Структуры конфига

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	PDFStorage string `yaml:"pdf_storage" env-required:"true"`
	Postgresql `yaml:"postgresql"`
	HTTPServer `yaml:"http_server"`
	Parser     `yaml:"parser"`
}

type Postgresql struct {
	Host     string `yaml:"host" env-required:"true"`
	Port     int    `yaml:"port" env-required:"true"`
	Username string `yaml:"username" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
	Database string `yaml:"database" env-required:"true"`
	SSLMode  string `yaml:"ssl_mode" env-default:"prefer"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Parser struct {
	VerticalThreshold   float64 `yaml:"vertical_threshold" required:"true"`
	HorizontalThreshold float64 `yaml:"horizontal_threshold" required:"true"`
}

func MustLoad() *Config { // Загрузка конфига
	// configPath := os.Getenv("CONFIG_PATH")
	configPath := "config/local.yaml"
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil { // Читаем конфиг, юзаю библиотеку cleanenv
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
