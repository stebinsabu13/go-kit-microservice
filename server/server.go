package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPServer(r *gin.Engine, endpoints Endpoints) http.Handler {
	// r.Use(gin.Logger())
	r.Use(commonMiddleware())
	r.POST("/user", adaptHandler(httptransport.NewServer(
		endpoints.CreateUser,
		decodeUserReq,
		encodeResponse,
	)))

	r.GET("/user/:id", adaptHandler(httptransport.NewServer(
		endpoints.GetUser,
		decodeEmailReq,
		encodeResponse,
	)))

	return r

}

func commonMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Next()
	}
}

func adaptHandler(server *httptransport.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Invoke the server.ServeHTTP method to handle the request
		server.ServeHTTP(c.Writer, c.Request)
	}
}
