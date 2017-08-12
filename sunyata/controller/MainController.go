package controller

import (
	"net/http"
	"sunyata/core/action"
	"sunyata/core/route"
)

type TestModel struct {
	Name   string
	Gender string
}

type MainController struct {
}

func (m *MainController) Index(w http.ResponseWriter, r *http.Request, ps route.Params) {
	action.Json(w, "TestActition")
}

func (m *MainController) TestView(w http.ResponseWriter, r *http.Request, ps route.Params){
	action.View("Test", nil)
}

func (m *MainController) TestJson(w http.ResponseWriter, r *http.Request, ps route.Params){
	action.Json(w, &TestModel{"AAA", "M"})
}

func (m *MainController) TestParam(w http.ResponseWriter, r *http.Request, ps route.Params){
	action.Json(w, &TestModel{"id", ps.ByName("id")})
}
