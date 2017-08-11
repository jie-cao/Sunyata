package action

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	ContentString string
	ContentType   string
	ContentHeader map[string]string
}

func View(viewName string, data interface{}) Response {
	return Response{}
}

func Json(data interface{}) Response {
	var response = Response{}
	jsonResponse, _ := json.Marshal(data)
	fmt.Println(string(jsonResponse))
	response.ContentString = string(jsonResponse)
	fmt.Println(response.ContentString)
	return response
}
