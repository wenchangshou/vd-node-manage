package http

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func configDeviceRoutes(){
	router.GET("/device", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	})
	router.POST("/device/:id/exhibition", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	})
	router.POST("/device", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	})
}