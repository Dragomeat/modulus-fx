package main

import (
	"demo/internal/errors"
	"demo/internal/framework"
	"demo/internal/user"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	app := fx.New(
		framework.ConfigModule(),
		errors.Module(),
		framework.LoggerModule(),
		framework.HttpModule(),
		framework.GormModule(),
		user.UserPlugin(),
		fx.WithLogger(
			func(logger *zap.Logger) fxevent.Logger {
				return &fxevent.ZapLogger{Logger: logger}
			},
		),
	)

	app.Run()
}
