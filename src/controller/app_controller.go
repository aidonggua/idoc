package controller

import (
	"net/http"
	"service"
)

func ListAllApp(rw http.ResponseWriter, req *http.Request) {
	apps := service.QueryAllApp()
	rw.Write(GetSuccessStatus().SetResult(apps).GetJson())
}
