package http

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func configFileRoutes(){
	router.GET("/file", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	})
	router.POST("/file", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	})
}