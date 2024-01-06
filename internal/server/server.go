package server

import (
	"bot-faq-qq/internal/option"
	"bot-faq-qq/internal/route"
	"bot-faq-qq/internal/store"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Web *gin.Engine
}

func NewServer(opt *option.Options) (*Server, error) {
	sro := store.DataStore{}
	g := route.ApiServer(opt.WebOptions.Engine, sro)
	return &Server{
		Web: g,
	}, nil
}
