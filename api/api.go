//api/api
package api

import (

    "pdf-APP/api/router"
    "github.com/gin-gonic/gin"
)

func InitServer() {
    r1 := gin.New()
    r1.Use(gin.Logger(), gin.Recovery())

    v1 := r1.Group("/api/v1")
    {
        router.UserRouter(v1)
    }

    r1.Run(":8001")
}
