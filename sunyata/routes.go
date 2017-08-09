package main

import (
	"net/http"
	"sunyata/core/web"
)

func MapRouteHandler(template string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(template, handler)
}

func MapRouteString(template string, route string) {
}

func MapRoute(template string, route web.Route) {
}
