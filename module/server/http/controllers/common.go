package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
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

func Health(c *gin.Context) {
	//1.创建子span
	span, _ := opentracing.StartSpanFromContext(c, "span_foo3")
	defer func() {
		//4.接口调用完，在tag中设置request和reply
		span.SetTag("request", c.Request)
		span.Finish()
	}()
	c.String(http.StatusOK, "OK")
}
func Version(c *gin.Context) {
	c.String(http.StatusOK, "v1.0")
}
