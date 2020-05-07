package routes

import "github.com/gin-gonic/gin"

type Routes struct{}

func Register(r *gin.Engine) {
	routes := &Routes{}
	r.GET("/listen", routes.Subscribe)
}
