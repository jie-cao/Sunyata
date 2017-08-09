package controller

import "sunyata/core/web"

type MainController struct {
}

func (m *MainController) Index() string {
	return "Hello Action"
}

func (m *MainController) Test() web.ActionResponse {
	return web.View("Test", nil)
}
