package appbase

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/samber/do"
)

type AppBase struct {
	Config      *Config
	ServiceName string
	Injector    *do.Injector
}

func New(options ...func(*AppBase)) *AppBase {
	appBase := &AppBase{}
	for _, o := range options {
		o(appBase)
	}

	return appBase
}

func Init() func(*AppBase) {
	return func(appBase *AppBase) {
		cfg, err := LoadConfig()
		if err != nil {
			panic(err)
		}

		appBase.Config = cfg
	}
}

func (a *AppBase) Shutdown() {
	err := a.Injector.Shutdown()
	if err != nil {
		log.Panic().Err(err).Msg("injector's shutdown failed")
	}
}

func WithDependencyInjector() func(*AppBase) {
	return func(appBase *AppBase) {
		appBase.Injector = NewInjector(appBase.ServiceName, appBase.Config)
	}
}

func WithLogger() func(*AppBase) {
	return func(appBase *AppBase) {
		log.Logger = *do.MustInvoke[*zerolog.Logger](appBase.Injector)
	}
}
