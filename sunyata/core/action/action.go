package action

type Response struct {
	ContentString string
	ContentType   string
	ContentHeader map[string]string
}

func View(viewName string, data interface{}) Response {
	return Response{}
}

func Json(data interface{}) Response {
	return Response{}
}
