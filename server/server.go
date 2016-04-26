package server

import (
	"github.com/gorilla/mux"
	"github.com/vjscjp/sampleutil/controllers"
	"github.com/vjscjp/sampleutil/controllers/app"
	"github.com/vjscjp/sampleutil/controllers/host"
)

const (
	App       = "/app/{id}"
	HostPorts = "/hostport"
	HostPort  = "/hostport/{id}/{port}"
)

func InitRoutes(muxRouter *mux.Router) {
	controllers.SetEnv()
	muxRouter.HandleFunc(App, appcontroller.GetApp).Methods("GET")
	muxRouter.HandleFunc(HostPorts, hostcontroller.GetHostPorts).Methods("GET")
	muxRouter.HandleFunc(HostPort, hostcontroller.GetHostPort).Methods("GET")
}
