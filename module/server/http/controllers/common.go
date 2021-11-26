package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//func configCommonRoutes() {
//	http2.router.GET("/health", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
//		w.Write([]byte("ok"))
//	})
//	http2.router.GET("/version", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
//		w.Write([]byte(""))
//	})
//}

func Health(c *gin.Context){
	c.String(http.StatusOK,"OK")
}
func Version (c *gin.Context){
	c.String(http.StatusOK,"v1.0")
}