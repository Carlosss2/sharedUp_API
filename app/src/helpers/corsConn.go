package helpers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func InitCORS(r *gin.Engine) {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"*",},
		AllowedMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		},
		AllowedHeaders: []string{
			"Content-Type", "Authorization", "X-Requested-With", "Accept",
		},
		AllowCredentials: true,
		MaxAge:           int((12 * time.Hour).Seconds()),
	})

	// Gin usa http.Handler internamente
	r.Use(func(ctx *gin.Context) {
		c.HandlerFunc(ctx.Writer, ctx.Request)
		ctx.Next()
	})
}
