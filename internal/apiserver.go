package apiserver

import (
	conf "bot-faq-qq/config"
	"bot-faq-qq/internal/app"
	"bot-faq-qq/internal/option"
)

func NewAPIServer() *app.App {
	opts := option.NewOptions()
	application := app.NewApp("Bus API Server",
		conf.ProjectName,
		app.WithOptions(opts),
		app.WithRunFunc(run(opts)),
	)

	return application
}

func run(opts *option.Options) app.RunFunc {
	return func(basename string) error {
		return nil
	}
}
