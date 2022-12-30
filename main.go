package main

import (
	jb "jenkinsServer/JenkinsBuild"
	"jenkinsServer/Utill"

	"github.com/gin-gonic/gin"
)

func init() {
	Utill.InitConfig()
}

func main() {
	r := gin.Default()
	r.GET("/debBuild",jb.Build)
	r.Run(":9000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}