package web

type ActionResponse struct {
	ContentString string
	ContentType   string
	ContentHeader map[string]string
}

func View(viewName string, data interface{}) ActionResponse {
	return ActionResponse{}
}
