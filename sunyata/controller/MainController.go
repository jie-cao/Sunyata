package controller

import (
	"net/http"
	"sunyata/core/response"
	"sunyata/core/route"
)

type TestModel struct {
	Name   string
	Gender string
}

type MainController struct {
}

func (m *MainController) Index(w http.ResponseWriter, r *http.Request, ps route.Params) {
	response.Json(w, "TestActition")
}

func (m *MainController) TestView(w http.ResponseWriter, r *http.Request, ps route.Params) {
	response.View(w, "Test", &TestModel{"AAA", "M"})
}

func (m *MainController) TestJson(w http.ResponseWriter, r *http.Request, ps route.Params) {
	response.Json(w, &TestModel{"AAA", "M"})
}

func (m *MainController) TestParam(w http.ResponseWriter, r *http.Request, ps route.Params) {
	response.Json(w, &TestModel{"id", ps.ByName("id")})
}
