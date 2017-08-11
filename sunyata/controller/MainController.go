package controller

import (
	"net/http"
	"sunyata/core/action"
)

type TestModel struct {
	Name   string
	Gender string
}

type MainController struct {
}

func (m *MainController) Index() string {
	return "Hello Action"
}

func (m *MainController) TestView(r *http.Request) action.Response {
	return action.View("Test", nil)
}

func (m *MainController) TestJson(r *http.Request) action.Response {
	return action.Json(&TestModel{"AAA", "M"})
}
