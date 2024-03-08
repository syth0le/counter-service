package counter_service

import (
	"fmt"
	"log"

	xlogger "github.com/syth0le/gopnik/logger"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/spf13/pflag"

	"github.com/syth0le/counter-service/cmd/counter-service/application"
	"github.com/syth0le/counter-service/cmd/counter-service/configuration"
)

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalf("failed to create config: %v", err)
	}

	if err = cfg.Validate(); err != nil {
		log.Fatalf("config validation failed: %v", err)
	}

	logger, err := xlogger.ConstructLogger(cfg.Logger)
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}

	app := application.New(cfg, logger)
	if err = app.Run(); err != nil {
		logger.Sugar().Fatalf("application stopped with error: %v", err)
	} else {
		logger.Info("application stopped")
	}
}

func loadConfig() (*configuration.Config, error) {
	cfg := configuration.NewDefaultConfig()

	configPath := pflag.StringP("config", "c", "", "config path")
	pflag.Parse()

	if err := cleanenv.ReadConfig(*configPath, cfg); err != nil {
		return nil, fmt.Errorf("cannot load config: %w", err)
	}
	return cfg, nil
}
