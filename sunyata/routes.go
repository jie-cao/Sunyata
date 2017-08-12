package main
import "sunyata/core/route"
import "sunyata/controller"

func RegisterRoutes(r *route.Router) {
	var mainController = &controller.MainController{}
	r.GET("/TestJson", mainController.TestJson)
	r.GET("/TestParams/:id", mainController.TestParam)
}