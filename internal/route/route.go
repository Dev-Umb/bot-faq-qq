package route

import (
	"bot-faq-qq/internal/store"
	"github.com/gin-gonic/gin"
)

func ApiServer(g *gin.Engine, sro store.DataStore) *gin.Engine {
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	return g
}
