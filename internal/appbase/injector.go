package appbase

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/samber/do"
	"github.com/zzeyneperdenn/GoWit-BE-Case-Study/internal/repository"
	"github.com/zzeyneperdenn/GoWit-BE-Case-Study/internal/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewInjector(serviceName string, cfg *Config) *do.Injector {
	injector := do.New()

	do.Provide(injector, func(i *do.Injector) (*zerolog.Logger, error) {
		logLevel, err := zerolog.ParseLevel(cfg.LogLevel)
		if err != nil {
			return nil, err
		}

		logger := zerolog.New(os.Stdout).
			Level(logLevel).
			With().
			Str("service_name", serviceName).
			Timestamp().
			Logger()

		return &logger, nil
	})
	do.Provide(injector, func(i *do.Injector) (*gorm.DB, error) {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
			cfg.DatabaseHost,
			cfg.DatabaseUser,
			cfg.DatabasePassword,
			cfg.DatabaseName,
			cfg.DatabasePort,
			"disable",
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}

		return db, nil
	})
	do.Provide(injector, func(i *do.Injector) (repository.Repository, error) {
		return repository.New(do.MustInvoke[*gorm.DB](i)), nil
	})
	do.Provide(injector, func(i *do.Injector) (services.ITicketsService, error) {
		return services.NewTicketsService(do.MustInvoke[repository.Repository](i)), nil
	})

	return injector
}
