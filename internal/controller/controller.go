package controller

import "bot-faq-qq/internal/store"

type Controller interface {
}

var _ Controller = (*controller)(nil)

type controller struct {
	store store.Factory
}

// NewController ...
func NewController(factory store.Factory) Controller {
	return &controller{
		store: factory,
	}
}
