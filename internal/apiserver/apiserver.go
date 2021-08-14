package apiserver

import (
	"github.com/daz2yy/go-base/internal/apiserver/config"
	"github.com/daz2yy/go-base/internal/apiserver/options"
	"github.com/daz2yy/go-base/pkg/app"
	"github.com/daz2yy/go-base/pkg/log"
)

const commandDesc = `api server`

func NewApp(basename string) *app.App {
	opts := options.NewOptions()
	application := app.NewApp("api server",
		basename,
		app.WithOptions(opts),
		app.WithDescription(commandDesc),
		app.WithDefaultValidArgs(),
		app.WithRunFunc(run(opts)),
	)

	return application
}

func run(opts *options.Options) app.RunFunc {
	return func(basename string) error {
		log.Init(opts.Log)
		defer log.Flush()

		cfg, err := config.CreateConfigFromOptions(opts)
		if err != nil {
			return err
		}

		return Run(cfg)
	}
}
