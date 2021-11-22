package http

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func configCommonRoutes() {
	router.GET("/health", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("ok"))
	})
	router.GET("/version", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte(""))
	})
	router.POST("/file", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	})

}
